package test

import (
	"testing"
)

const slug = "6CE719"

// 文档
func TestDocument(t *testing.T) {
	resp := getContent(slug)
	err := GenFile(resp)
	if err != nil {
		panic(err)
		return
	}
}

// 题目
func TestQuestion(t *testing.T) {
	resp := getContent(slug)
	GenQuestionCode(resp)
}
