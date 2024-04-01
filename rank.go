package rank

const (
	realRankCount int = 1000
	maxItemCount  int = 1000
)

type rankItemBlock struct {
	items    []node
	nextFree int
}
