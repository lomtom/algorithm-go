package test

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

const solutionContent = `package leetcode`

func TestMkdir(t *testing.T) {
	dir := "../../../leetcode"
	number := "2938"
	d := filepath.Join(dir, number)
	if _, err := os.Stat(d); err != nil {
		err = os.Mkdir(d, os.ModePerm)
		if err != nil {
			panic(err)
			return
		}
	}
	solution := filepath.Join(dir, number, "solution.go")
	if _, err := os.Stat(solution); err != nil {
		create, err := os.Create(solution)
		if err != nil {
			panic(err)
			return
		}
		_, err = create.Write([]byte(solutionContent))
		if err != nil {
			panic(err)
			return
		}
	}
	solution = filepath.Join(dir, number, "solution_test.go")
	if _, err := os.Stat(solution); err != nil {
		create, err := os.Create(solution)
		if err != nil {
			panic(err)
			return
		}
		_, err = create.Write([]byte(solutionContent))
		if err != nil {
			panic(err)
			return
		}
	}
	cmd := exec.Command("git", "add", d)
	cmd.Run()
}
