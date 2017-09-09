package main

import (
	"fmt"
	"sort"
)

type Food struct {
	Name  string
	Price int
}

func main() {
	foods := make([]Food, 6)
	foods[0] = Food{Name: "みかん", Price: 100}
	foods[1] = Food{Name: "バナナ", Price: 150}
	foods[2] = Food{Name: "りんご", Price: 100}
	foods[3] = Food{Name: "パイナップル", Price: 120}
	foods[4] = Food{Name: "梨", Price: 100}
	foods[5] = Food{Name: "メロン", Price: 200}

	sort.Slice(foods, func(i, j int) bool {
		return foods[i].Price < foods[j].Price
	})

	result := sort.SliceIsSorted(foods, func(i, j int) bool {
		return foods[i].Price < foods[j].Price
	})

	if result {
		fmt.Println("ソートされてるよ")
	} else {
		fmt.Println("ソートされてないよ")
	}

	for _, food := range foods {
		fmt.Printf("%+v\n", food)
	}
}
