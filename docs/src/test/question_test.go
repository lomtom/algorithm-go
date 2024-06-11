package test

import (
	"testing"
)

const slug = "maximum-number-of-balloons"

func TestDocument(t *testing.T) {
	resp := getContent(slug)
	err := GenFile(resp)
	if err != nil {
		panic(err)
		return
	}
}

func TestQuestion(t *testing.T) {
	resp := getContent(slug)
	GenQuestionCode(resp)
}
