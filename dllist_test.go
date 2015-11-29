package dllist_test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/RIscRIpt/dllist"
)

type key int

func (i key) Less(j dllist.Noder) bool {
	return i < j.(key)
}

func (i key) Equals(j dllist.Noder) bool {
	return i == j.(key)
}

func BenchmarkSort1M(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	l := dllist.New()
	b.Log("Filling ...")
	for i := 0; i < 1000000; i++ {
		l.PushBack(key(rand.Int()))
	}
	b.Log("Sorting ...")
	b.ResetTimer()
	l.Sort()
	b.StopTimer()
	b.Log("Checking ...")
	for iter, node := dllist.NewIterator(l.First(), l.Last()); node != iter.Last(); node = iter.Next() {
		if node.Next().Data().Less(node.Data()) {
			b.FailNow()
		}
	}
}
