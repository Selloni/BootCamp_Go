package presnt

type PresentHeap []Present

type Present struct {
	Value int
	Size  int
}

func (pq PresentHeap) Len() int {
	return len(pq)
}

func (pq PresentHeap) Less(i, j int) bool {
	if pq[i].Value > pq[j].Value {
		return true
	}
	if pq[i].Value == pq[j].Value && pq[i].Size < pq[i].Size {
		return true
	}
	return false
}

func (pq PresentHeap) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
