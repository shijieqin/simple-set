package simple_set

import "testing"

func TestThreadUnsafeSet_Add(t *testing.T) {
	set := NewSet()
	set.Add("aaaa")
	set.Add("bbb")
	t.Log(set.String())
}
