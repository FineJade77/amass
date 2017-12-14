package pool

import (
	"testing"
)

type Test struct{}

func (t *Test) Close() error {
	return nil
}

func newTest() (Objecter, error) {
	return &Test{}, nil
}

func TestPool(t *testing.T) {
	p, err := NewPool(newTest, 5)
	if err != nil {
		t.Fatal(err)
	}

	if length := p.Len(); length != 5 {
		t.Fatalf("the length of pool expect %d, but be %d", 5)
	}

	obj, err := p.Get()
	if err != nil {
		t.Fatal(err)
	}

	if length := p.Len(); length != 4 {
		t.Fatalf("the length of pool expect %d, but be %d", 4)
	}

	err = p.Put(obj)
	if err != nil {
		t.Fatal(err)
	}

	if length := p.Len(); length != 5 {
		t.Fatalf("the length of pool expect %d, but be %d", 5)
	}

	p.Empty()

	obj, err = p.Get()
	if err != nil {
		t.Fatal(err)
	}
}
