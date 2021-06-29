package simple_set

type Iterator struct {
	C    <-chan interface{}
	stop chan struct{}
}

func (i *Iterator) Stop() {
	defer func() {
		recover()
	}()
	close(i.stop)

	//排空channel
	for range i.C{
	}
}

func newIterator() (*Iterator, chan <- interface{}, <-chan struct{}) {
	itemChan := make(chan interface{})
	stopChan := make(chan struct{})
	return &Iterator{
		C:    itemChan,
		stop: stopChan,
	}, itemChan, stopChan
}
