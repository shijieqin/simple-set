package simple_set

type Set interface {
	// 添加元素到Set内
	Add(item interface{}) bool
	// 清空Set
	Clear()
	// 判断Set是否包含给定的item
	Contains(items ...interface{}) bool
	// 判断两个set是否相同
	Equal(s Set) bool
	// 返回Set的元素个数
	Len() int
	// 遍历Set，直到函数返回true
	Each(func(interface{}) bool)
	// 迭代
	Iterator() *Iterator
	//移除一个元素
	Remove(i interface{})

	String() string

	// 删除并返回
	Pop() interface{}

	//返回包含所有元素的slice
	ToSlice() []interface{}
}

func NewSet(s ...interface{}) Set {
	set := newThreadUnsafeSet()
	for _, item := range s {
		set.Add(item)
	}
	return &set
}

func NewThreadSafeSet(s ...interface{}) Set {
	set := newThreadSafeSet()
	for _, item := range s {
		set.Add(item)
	}
	return &set
}