package test

import (
	"testing"
)

const slug = "maximum-number-of-operations-with-the-same-score-i"

func TestQuestionContent(t *testing.T) {
	resp := getContent(slug)
	err := GenFile(resp)
	if err != nil {
		panic(err)
		return
	}
}

func TestMkdir(t *testing.T) {
	resp := getContent(slug)
	GenQuestionCode(resp)
}
