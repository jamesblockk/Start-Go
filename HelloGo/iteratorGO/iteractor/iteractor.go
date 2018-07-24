package iteractor

import "context"

type Iterator interface {
	Iterator(ctx context.Context) <-chan interface{}
}

type iterator []interface{}

func (i iterator) Iterator(ctx context.Context) <-chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		for _, v := range i {
			select {
			case c <- v:
				break
			case <-ctx.Done():
				return
			}
		}
	}()

	return c
}

func Loop(datas []interface{}) {
	nums := iterator(datas)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	for range nums.Iterator(ctx) {

	}
}
