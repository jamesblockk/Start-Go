package queue

import "testing"

type intQueue interface {
	Push(interface{})
	Pop() (interface{}, bool)
	Len() int
	Cap() int
}

func testqueue(t *testing.T, q intQueue) {
	for j := 0; j < 100; j++ {
		if q.Len() != 0 {
			t.Fatal("expected no elements")
		} else if _, ok := q.Pop(); ok {
			t.Fatal("expected no elements")
		}

		for i := 0; i < j; i++ {
			q.Push(i)
		}

		for i := 0; i < j; i++ {
			if x, ok := q.Pop(); !ok {
				t.Fatal("expected an element")
			} else if x != i {
				t.Fatalf("expected %d got %d", i, x)
			}
		}
	}

	a := 0
	r := 0
	for j := 0; j < 100; j++ {
		for i := 0; i < 4; i++ {
			q.Push(a)
			a++
		}

		for i := 0; i < 2; i++ {
			if x, ok := q.Pop(); !ok {
				t.Fatal("expected an element")
			} else if x != r {
				t.Fatalf("expected %d got %d", r, x)
			}
			r++
		}
	}

	if q.Len() != 200 {
		t.Fatalf("expected 200 elements have %d", q.Len())
	}
}

func TestSqueue(t *testing.T) {
	testqueue(t, New())
}

func benchmarkAdd(b *testing.B, q intQueue) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		q.Push(i)
	}
}

func BenchmarkSliceAdd(b *testing.B) {
	benchmarkAdd(b, New())
}

func benchmarkRemove(b *testing.B, q intQueue) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		q.Push(i)

		if q.Len() > 10 {
			q.Pop()
		}
	}
}

func BenchmarkSliceRemove(b *testing.B) {
	benchmarkRemove(b, New())
}
