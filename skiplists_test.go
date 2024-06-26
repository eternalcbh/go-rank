package rank

//go test -covermode=count -v -coverprofile=coverage.out -run=. -cpuprofile=rank.p
//go tool cover -html=coverage.out
//go tool pprof rank.p
//go test -v -run=^$ -bench BenchmarkRank -count 10
import (
	"math/rand"
	"testing"
)

func TestSkipLists(t *testing.T) {
	for c := 0; c < 10000; c++ {

		//fmt.Println(c)

		sl1 := newSkipLists(0)

		var nodes []*node
		for i := 0; i < rand.Int()%10000+1000; i++ {
			nodes = append(nodes, &node{
				key:   rand.Int() % 10000,
				value: uint64(i + i),
			})
		}

		for _, v := range nodes {
			sl1.InsertNode(v)
		}

		sl2 := sl1.split()

		if !sl1.checkLink() {
			panic("checkLink1")
		}

		if !sl2.checkLink() {
			panic("checkLink2")
		}

		sl1.merge(sl2)

		if !sl1.checkLink() {
			panic("checkLink3")
		}

	}
}
