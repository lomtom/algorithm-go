package test

import (
	"encoding/json"
	"fmt"
	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/go-resty/resty/v2"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

var m = map[string]string{
	"Easy":   "简单",
	"Medium": "中等",
	"Hard":   "困难",
}

const solutionContent = `package leetcode

%s
`

const solutionTestContent = `package leetcode

import "testing"

func Test(t *testing.T) {
	t.Log()
}
`

// https://cattiek.site/2019/03/03/Leetcode%E7%88%AC%E8%99%AB%E5%AE%9E%E8%B7%B5/#%E9%A2%98%E7%9B%AE%E8%AF%A6%E6%83%85-%E5%85%8D%E7%99%BB%E9%99%86%E5%8F%AF%E8%8E%B7%E5%8F%96
func getContent(slug string) QuestionContent {
	// 请求的 URL
	url := "https://leetcode.cn/graphql/"

	// 设置请求头
	headers := map[string]string{
		"Accept-Language": "zh-CN,zh;q=0.9",
		"Content-Type":    "application/json",
	}

	body := map[string]interface{}{
		"query": `query questionTranslations($titleSlug: String!) {
  question(titleSlug: $titleSlug) {
    translatedTitle
    translatedContent
    questionFrontendId
    difficulty
    titleSlug
    topicTags{
      translatedName
    }
    codeSnippets {
      lang
      langSlug
      code
      __typename
    }
  }
}`,
		"variables": map[string]string{
			"titleSlug": slug,
		},
		"operationName": "questionTranslations",
	}
	var resp QuestionContent

	resty.New().SetJSONMarshaler(json.Marshal).SetJSONUnmarshaler(json.Unmarshal).R().
		SetHeaders(headers).
		SetBody(body).
		SetResult(&resp).
		Post(url)

	if IsHTMLFragment(resp.Data.Question.TranslatedContent) {
		converter := md.NewConverter("", true, nil)

		markdown, err := converter.ConvertString(resp.Data.Question.TranslatedContent)
		if err != nil {
			log.Fatal(err)
		}
		resp.Data.Question.TranslatedContent = markdown
	}
	resp.Data.Question.Difficulty = m[resp.Data.Question.Difficulty]
	resp.Data.Question.QuestionFrontendId = strings.Replace(resp.Data.Question.QuestionFrontendId, " ", "", -1)
	return resp
}

// IsHTMLFragment 判断字符串是否为HTML片段
func IsHTMLFragment(input string) bool {
	return strings.Contains(input, "<p>") && strings.Contains(input, "</p>")
}

type QuestionContent struct {
	Data struct {
		Question struct {
			TranslatedTitle    string `json:"translatedTitle"`
			TranslatedContent  string `json:"translatedContent"`
			QuestionFrontendId string `json:"questionFrontendId"`
			Difficulty         string `json:"difficulty"`
			TitleSlug          string `json:"titleSlug"`
			TopicTags          []struct {
				Name           string `json:"name"`
				Slug           string `json:"slug"`
				TranslatedName string `json:"translatedName"`
			} `json:"topicTags"`
			CodeSnippets []struct {
				Lang     string `json:"lang"`
				LangSlug string `json:"langSlug"`
				Code     string `json:"code"`
				Typename string `json:"__typename"`
			} `json:"codeSnippets"`
		} `json:"question"`
	} `json:"data"`
}

func (q QuestionContent) getDocumentPath() string {
	return fmt.Sprintf("../content/leetcode/%s%s.md", q.Data.Question.QuestionFrontendId, strings.Replace(q.Data.Question.TranslatedTitle, " ", "", -1))
}

func (q QuestionContent) getQuestionPath() string {
	dir := "../../../leetcode"
	number := q.Data.Question.QuestionFrontendId
	return filepath.Join(dir, number)
}

func GenFile(questionContent QuestionContent) error {
	if questionContent.Data.Question.QuestionFrontendId == "" {
		panic("题目ID为空")
		return nil
	}
	fileContent := GenContent(questionContent)
	f := questionContent.getDocumentPath()
	if _, err := os.Stat(f); err == nil {
		return nil
	}
	create, err := os.Create(f)
	if err != nil {
		return err
	}
	_, err = create.WriteString(fileContent)
	cmd := exec.Command("git", "add", f)
	cmd.Run()
	return err
}

func GenContent(questionContent QuestionContent) string {
	tags := ""
	for index := range questionContent.Data.Question.TopicTags {
		tags += `
  - ` + questionContent.Data.Question.TopicTags[index].TranslatedName
	}
	questionContent.RewriteContent()
	return fmt.Sprintf(
		`---
title: %s
categories:
  - %s
tags: %s
slug: %s
number: %s
---

## 题目描述：

%s

---
## 解题分析及思路：

### 方法：方法

**思路：**


**复杂度：**

- 时间复杂度：O(n)
- 空间复杂度：O(1)

**执行结果：**

- 执行耗时:0 ms,击败了100.00 的Go用户
- 内存消耗:2.4 MB,击败了99.83 的Go用户
`, questionContent.Data.Question.TranslatedTitle,
		questionContent.Data.Question.Difficulty,
		tags,
		questionContent.Data.Question.TitleSlug,
		questionContent.Data.Question.QuestionFrontendId,
		WithImg(questionContent.Data.Question.TranslatedContent, fmt.Sprintf("%s%s", questionContent.Data.Question.QuestionFrontendId, questionContent.Data.Question.TranslatedTitle)))
}

func GenQuestionCode(resp QuestionContent) {
	if resp.Data.Question.QuestionFrontendId == "" {
		panic("题目ID为空")
		return
	}
	d := resp.getQuestionPath()
	if _, err := os.Stat(d); err != nil {
		err = os.Mkdir(d, os.ModePerm)
		if err != nil {
			panic(err)
			return
		}
	}
	solution := filepath.Join(d, "solution.go")
	if _, err := os.Stat(solution); err != nil {
		create, err := os.Create(solution)
		if err != nil {
			panic(err)
			return
		}
		var code string
		for index := range resp.Data.Question.CodeSnippets {
			if resp.Data.Question.CodeSnippets[index].LangSlug == "golang" {
				code = resp.Data.Question.CodeSnippets[index].Code
			}
		}
		_, err = create.Write([]byte(fmt.Sprintf(solutionContent, code)))
		if err != nil {
			panic(err)
			return
		}
	}
	solution = filepath.Join(d, "solution_test.go")
	if _, err := os.Stat(solution); err != nil {
		create, err := os.Create(solution)
		if err != nil {
			panic(err)
			return
		}
		_, err = create.Write([]byte(solutionTestContent))
		if err != nil {
			panic(err)
			return
		}
	}
	cmd := exec.Command("git", "add", d)
	cmd.Run()
}

func WithImg(content, title string) string {

	// Regular expression to find image URLs in Markdown content
	imgRegex := regexp.MustCompile(`!\[.*?\]\((http.*?)\)`)

	// Find all image URLs
	matches := imgRegex.FindAllStringSubmatch(content, -1)
	if len(matches) == 0 {
		return content
	}

	title = strings.Replace(title, " ", "", -1)

	// Define the local directory to save images
	localDir := "../../public/img/leetcode/" + title
	replaceDir := "/img/leetcode/" + title
	err := os.MkdirAll(localDir, os.ModePerm)
	if err != nil {
		fmt.Println("Error creating directory:", err)
		return content
	}
	for _, match := range matches {
		if len(match) > 1 {
			imgURL := match[1]
			// Create a local file to save the image
			fileName := filepath.Base(imgURL)
			localPath := filepath.Join(localDir, fileName)
			if _, err := os.Stat(localPath); err == nil {
				continue
			}

			// Download image
			resp, err := http.Get(imgURL)
			if err != nil {
				fmt.Println("Error downloading image:", err)
				continue
			}
			defer resp.Body.Close()

			out, err := os.Create(localPath)
			if err != nil {
				fmt.Println("Error creating file:", err)
				continue
			}
			defer out.Close()

			// Copy the image to the local file
			_, err = io.Copy(out, resp.Body)
			if err != nil {
				fmt.Println("Error saving image:", err)
				continue
			}

			replacePath := filepath.Join(replaceDir, fileName)
			cmd := exec.Command("git", "add", localPath)
			err = cmd.Run()
			if err != nil {
				fmt.Println(err)
				return ""
			}
			// Replace the remote URL with the local path in the content
			content = strings.Replace(content, imgURL, replacePath, -1)
		}
	}

	return content
}

func (q *QuestionContent) RewriteContent() {
	if strings.Contains(q.Data.Question.TranslatedContent, "```") {
		return
	}
	fmt.Println("没有找到代码块，已自动调整，请检查")
	var (
		preIndex  int = -1
		lastIndex int = -1
	)
	lines := strings.Split(q.Data.Question.TranslatedContent, "\n")
	for index := 0; index < len(lines); index++ {
		if strings.Contains(lines[index], "**示例") || strings.Contains(lines[index], "**提示：**") {
			if preIndex == -1 {
				for index+1 < len(lines) && (len(lines[index+1]) == 0 || strings.HasPrefix(lines[index+1], "![")) {
					index++
				}
				preIndex = index
				continue
			}
			lastIndex = index
		}
		if preIndex != -1 && lastIndex != -1 {
			for i := preIndex + 1; i < lastIndex; i++ {
				lines[i] = strings.ReplaceAll(lines[i], "**", "")
				if strings.HasPrefix(lines[i], ">") {
					lines[i] = strings.Replace(lines[i], ">", "", 1)
				}
			}
			lines[preIndex] = lines[preIndex] + "\n```"
			lines[lastIndex] = "```\n" + lines[lastIndex]
			for index+1 < len(lines) && (len(lines[index+1]) == 0 || strings.HasPrefix(lines[index+1], "![")) {
				index++
			}
			preIndex = index
			lastIndex = -1
		}
	}
	q.Data.Question.TranslatedContent = strings.Join(lines, "\n")
}
