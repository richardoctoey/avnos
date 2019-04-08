package number3

type TreeS struct {
	Num int
	Childs []*TreeS
}

var t []*TreeS

func InputTree(curtree *TreeS, parent int, v int) {
	if len(t) == 0 {
		childs := []*TreeS{}
		childs[0] = &TreeS{Num: v}
		d := &TreeS{Num: parent, Childs: childs}
		t = append(t, d)
	}
	for _, elmtree := range t {
		if elmtree.Num == parent {
			elmtree.Childs = append(elmtree.Childs, &TreeS{Num: v})
			return
		}
	}
}

func Tree() {
	t = make([]*TreeS, 0)
	InputTree(0, 1)
}