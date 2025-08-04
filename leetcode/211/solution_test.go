package leetcode

import "testing"

func Test(t *testing.T) {
	this := Constructor()
	this.AddWord("bad")
	this.AddWord("dad")
	this.AddWord("mad")
	t.Log(this.Search("pad"))
	t.Log(this.Search("bad"))
	t.Log(this.Search(".ad"))
	t.Log(this.Search("b.."))
}

func Test1(t *testing.T) {
	this := Constructor()
	this.AddWord("at")
	this.AddWord("and")
	this.AddWord("an")
	this.AddWord("add")
	t.Log(this.Search("a"))
	t.Log(this.Search(".at"))
	this.AddWord("bat")
	t.Log(this.Search(".at"))
	t.Log(this.Search("an."))
	t.Log(this.Search("a.d."))
	t.Log(this.Search("b."))
	t.Log(this.Search("a.d"))
	t.Log(this.Search("."))
}
