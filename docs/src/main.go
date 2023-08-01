package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	dir := "/Users/lomtom/company/项目/go/astroplate/src/content/blog"
	// 读取dir下所有的文件
	files, err := os.ReadDir(dir)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		name := file.Name()
		if !strings.HasSuffix(name, ".md") {
			continue
		}
		name = strings.Replace(name, ".md", "", -1)
		filePath := dir + "/" + name + ".md"
		name = strings.Replace(name, "_readme", "", -1)
		// 读取文件内容
		readFile, err := os.ReadFile(filePath)
		if err != nil {
			panic(err)
		}
		// 写入文件
		readFile = append([]byte(fmt.Sprintf(`---
title: %s
categories:
  - 中等
tags:
  - 队列
---
`, name)), readFile...)
		err = os.WriteFile(filePath, readFile, 0644)
		if err != nil {
			panic(err)
		}
	}
	fmt.Println(len(files))
}
