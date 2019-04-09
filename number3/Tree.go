package number3

import (
	"fmt"
)

type TreeS struct {
	Num int
	Childs []*TreeS
}

var t []*TreeS

func NestedSearch(nt  *TreeS, parent int, v int) int {
	if nt.Num == parent {
		nt.Childs = append(nt.Childs, &TreeS{Num: v})
		return 1
	}
	for _, tchild := range nt.Childs {
		return NestedSearch(tchild, parent, v)
	}
	return 0
}

func InputTree(parent int, v int) {
	top := 0
	for _, elmtree := range t {
		if elmtree.Num == parent {
			elmtree.Childs = append(elmtree.Childs, &TreeS{Num: v})
			return
		}
		for _, tchild := range elmtree.Childs {
			top += NestedSearch(tchild, parent, v)
		}
	}
	if top == 0 {
		childs := []*TreeS{}
		childs = append(childs, &TreeS{Num: v})
		t = append(t, &TreeS{Num: parent, Childs: childs})
	}
}

func Tree() {
	t = make([]*TreeS, 0)
	InputTree(0, 1)
	InputTree(0, 2)
	InputTree(3, 4)
	fmt.Println("Number 3:  Total Tree: ", len(t))
}