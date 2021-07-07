package simple_set

import (
	"fmt"
	"strings"
)

type threadUnsafeSet map[interface{}]struct{}

func newThreadUnsafeSet() threadUnsafeSet {
	return make(threadUnsafeSet)
}

func (s *threadUnsafeSet) Add(i interface{}) bool {
	if _, found := (*s)[i]; found {
		return false
	}
	(*s)[i] = struct{}{}
	return true
}

func (s *threadUnsafeSet) Clear() {
	*s = newThreadUnsafeSet()
}

func (s *threadUnsafeSet) Contains(i ...interface{}) bool {
	for _, item := range i {
		if _, found := (*s)[item]; !found {
			return false
		}
	}
	return true
}

func (s *threadUnsafeSet) Equal(set Set) bool {
	_ = set.(*threadUnsafeSet)
	if s.Len() != set.Len() {
		return false
	}
	for key:= range *s {
		if !set.Contains(key) {
			return false
		}
	}
	return true
}

func (s *threadUnsafeSet) Len() int {
	return len(*s)
}

func (s *threadUnsafeSet) Each(f func(interface{}) bool){
	for item:= range *s{
		if f(item) {
			return
		}
	}
}

func (s *threadUnsafeSet) Intersect(other Set) Set {
	o := other.(*threadUnsafeSet)
	intersection := newThreadUnsafeSet()
	if s.Len() < other.Len() {
		for elem := range *s {
			if other.Contains(elem) {
				intersection.Add(elem)
			}
		}
	} else {
		for elem := range *o {
			if s.Contains(elem) {
				intersection.Add(elem)
			}
		}
	}
	return &intersection
}

func (s *threadUnsafeSet) Iterator() *Iterator {
	iterator, itemCh, stopCh := newIterator()

	go func() {
	L:
		for item := range *s {
			select {
			case <-stopCh:
				break L
			case itemCh <- item:
			}
		}
		close(itemCh)
	}()
	return iterator
}

func (s *threadUnsafeSet) Remove(i interface{}){
	delete(*s, i)
}

func (s *threadUnsafeSet) String() string {
	items := make([]string, 0, len(*s))

	for item := range *s {
		items = append(items, fmt.Sprintf("%v", item))
	}

	return fmt.Sprintf("Set{%s}", strings.Join(items, ","))
}

func (s *threadUnsafeSet) Pop() interface{}{
	if s.Len() == 0 {
		return nil
	}
	item := (*s)[0]
	delete(*s, item)
	return item
}

func (s *threadUnsafeSet) ToSlice() []interface{}{
	keys := make([]interface{}, 0, s.Len())
	for key := range *s {
		keys = append(keys, key)
	}
	return keys
}