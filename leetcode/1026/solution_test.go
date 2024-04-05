package leetcode

import "testing"

func Test(t *testing.T) {
	t.Log(maxAncestorDiff(construct([]any{1, nil, 2, nil, nil, nil, 0, nil, nil, nil, nil, nil, nil, 3})))
	t.Log(maxAncestorDiff1(construct([]any{1, nil, 2, nil, nil, nil, 0, nil, nil, nil, nil, nil, nil, 3})))
	t.Log(maxAncestorDiff1(construct([]any{2, nil, 0, nil, nil, 1})))
}
