package leetcode

import (
	"testing"
)

// "2x=x"
// "x=x+2"
// "x+5-3+x=6+x-2"
// "2x=x"
// "0x=0"
func Test(t *testing.T) {
	t.Log(solveEquation("-x=1"))
	t.Log(solveEquation("2x=x"))
	t.Log(solveEquation("x=x+2"))
	t.Log(solveEquation("x+5-3+x=6+x-2"))
	t.Log(solveEquation("2x=x"))
	t.Log(solveEquation("0x=0"))
}
