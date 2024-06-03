package test

import (
	"encoding/json"
	"fmt"
	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/go-resty/resty/v2"
	"log"
	"os"
	"testing"
)

var m = map[string]string{
	"Easy":   "简单",
	"Medium": "中等",
	"Hard":   "困难",
}

func TestQuestionContent(t *testing.T) {
	resp := getContent("letter-case-permutation")
	err := GenFile(resp)
	if err != nil {
		panic(err)
		return
	}
}

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

	converter := md.NewConverter("", true, nil)

	markdown, err := converter.ConvertString(resp.Data.Question.TranslatedContent)
	if err != nil {
		log.Fatal(err)
	}
	resp.Data.Question.TranslatedContent = markdown
	resp.Data.Question.Difficulty = m[resp.Data.Question.Difficulty]
	return resp
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
		} `json:"question"`
	} `json:"data"`
}

func GenFile(questionContent QuestionContent) error {
	fileContent := GenContent(questionContent)
	create, err := os.Create(fmt.Sprintf("../content/leetcode/%s%s.md", questionContent.Data.Question.QuestionFrontendId, questionContent.Data.Question.TranslatedTitle))
	if err != nil {
		return err
	}
	_, err = create.WriteString(fileContent)
	return err
}

func GenContent(questionContent QuestionContent) string {
	tags := ""
	for index := range questionContent.Data.Question.TopicTags {
		tags += `
  - ` + questionContent.Data.Question.TopicTags[index].TranslatedName
	}

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

- 时间复杂度：O(N * M)
- 空间复杂度：O(1)

**执行结果：**

- 执行耗时:1 ms,击败了40.84 的Go用户
- 内存消耗:2.4 MB,击败了28.50 的Go用户
`, questionContent.Data.Question.TranslatedTitle,
		questionContent.Data.Question.Difficulty,
		tags,
		questionContent.Data.Question.TitleSlug,
		questionContent.Data.Question.QuestionFrontendId,
		questionContent.Data.Question.TranslatedContent)
}
