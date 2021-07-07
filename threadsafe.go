package simple_set

import "sync"

type threadSafeSet struct {
	s threadUnsafeSet
	sync.RWMutex
}

func newThreadSafeSet() threadSafeSet {
	return threadSafeSet{s: newThreadUnsafeSet()}
}

func (s *threadSafeSet) Add(i interface{}) bool {
	defer func() {
		s.Unlock()
	}()
	s.Lock()
	return s.s.Add(i)
}

func (s *threadSafeSet) Clear() {
	defer func() {
		s.Unlock()
	}()
	s.Lock()
	s.s.Clear()
}

func (s *threadSafeSet) Contains(i ...interface{}) bool{
	defer func() {
		s.RUnlock()
	}()
	s.RLock()
	return s.s.Contains(i...)
}

func (s *threadSafeSet) Equal(other Set) bool {
	o := other.(*threadSafeSet)

	s.RLock()
	o.RLock()
	defer func() {
		s.RUnlock()
		o.RUnlock()
	}()

	return s.s.Equal(&o.s)
}

func (s *threadSafeSet) Len() int {
	defer func() {
		s.RUnlock()
	}()
	s.RLock()
	return s.s.Len()
}

func (s *threadSafeSet) Each(f func(interface{}) bool) {
	defer func() {
		s.RUnlock()
	}()
	s.RLock()
	for item := range s.s{
		if f(item){
			return
		}
	}
}

func (s *threadSafeSet) Intersect(other Set) Set {
	o := other.(*threadSafeSet)

	s.RLock()
	o.RLock()
	unsafeIntersection := s.s.Intersect(&o.s).(*threadUnsafeSet)

	ret := &threadSafeSet{s: *unsafeIntersection}
	s.RUnlock()
	o.RUnlock()
	return ret
}

func (s *threadSafeSet) Iterator() *Iterator {
	iterator, itemCh, stopCh := newIterator()
	go func() {
		defer func() {
			s.RUnlock()
		}()
		s.RLock()
	L:
		for item := range s.s{
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

func (s *threadSafeSet) Remove(i interface{}) {
	defer func() {
		s.Unlock()
	}()
	s.Lock()
	delete(s.s, i)
}

func (s *threadSafeSet) String() string {
	defer func() {
		s.RUnlock()
	}()
	s.RLock()
	return s.s.String()
}

func (s *threadSafeSet) Pop() interface{} {
	defer func() {
		s.Unlock()
	}()
	s.Lock()
	return s.s.Pop()
}

func (s *threadSafeSet) ToSlice() []interface{}{
	defer func() {
		s.RUnlock()
	}()
	s.RLock()
	return s.s.ToSlice()
}
