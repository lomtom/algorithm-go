package leetcode

import "testing"

func Test(t *testing.T) {
	t.Log(entityParser("&apos;"))
	t.Log(entityParser("leetcode.com&frasl;problemset&frasl;all"))
	t.Log(entityParser("&amp; is an HTML entity but &ambassador; is not."))
	t.Log(entityParser("and I quote: &quot;...&quot;"))
	t.Log(entityParser("Stay home! Practice on Leetcode :)"))
	t.Log(entityParser("x &gt; y &amp;&amp; x &lt; y is always false"))
}
