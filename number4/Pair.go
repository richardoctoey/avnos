package number4

import (
    "fmt"
)

type Element struct {
    FirstNum int
    SecondNum int
}

var Pairs []Element

func PairSum(arr []int, sum int){
    Pairs = []Element{}
    for _, firstelm := range arr {
        for _, secondelm := range arr {
            if firstelm+secondelm == sum {
                Pairs = append(Pairs, Element{FirstNum:firstelm, SecondNum:secondelm})
            }
        }
    }
}

func Pair() {
    Pairs = make([]Element, 0)
    PairSum([]int{1,2,3}, 5)
    fmt.Println(Pairs)
    PairSum([]int{1,2,3}, 10)
    fmt.Println(Pairs)
}