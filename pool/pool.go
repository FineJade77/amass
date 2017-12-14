package pool

import (
	"errors"
)

type Objecter interface {
	Close() error
}

type Factory func() (Objecter, error)

type Pool struct {
	objs    chan Objecter
	factory Factory
	size    int
}

func NewPool(factory Factory, size int) (*Pool, error) {
	if size <= 0 {
		return nil, errors.New("invalid capacity settings")
	}

	pool := &Pool{
		objs:    make(chan Objecter, size),
		factory: factory,
		size:    size,
	}

	for i := 0; i < size; i++ {
		obj, err := pool.factory()
		if err != nil {
			return nil, err
		}

		pool.objs <- obj
	}

	return pool, nil
}

func (p *Pool) Get() (Objecter, error) {
	select {
	case obj := <-p.objs:
		return obj, nil
	default:
		return p.factory()
	}
}

func (p *Pool) Put(obj Objecter) error {
	select {
	case p.objs <- obj:
		return nil
	default:
		// pool is full
		return obj.Close()
	}
}

func (p *Pool) Len() int {
	return len(p.objs)
}

func (p *Pool) Empty() {
	for {
		select {
		case obj := <-p.objs:
			obj.Close()
		default:
			return
		}
	}
}
