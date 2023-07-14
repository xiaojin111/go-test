package main

import (
	"log"
	"sort"
)

func main() {
	AAA([]int{5, 6}, 6)
}

type sortArr []int

func (s sortArr) Len() int {
	return len(s)
}

func (s sortArr) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortArr) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func AAA(arr []int, k int) {
	var result []int
	sort.Sort(sortArr(arr))

	temp := 0
	index := 1
	for temp < k {
		ve := -1
		for _, v := range arr {
			if v == index {
				ve = v
				break
			}
		}
		if ve == -1 {
			result = append(result, index)
			temp++

		}
		index++
	}
	sum := 0
	for _, v := range result {
		sum += v
	}
	log.Println(sum)
}
