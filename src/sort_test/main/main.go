package main

import (
	"fmt"
	"math"
	"sort"
)

type bucket struct {
	upperBound float64
	count float64
}

func main() {
	a := []bucket{
		{upperBound: 0.05, count: 199881},
		{upperBound: 0.1, count: 212210},
		{upperBound: 0.2, count: 215395},
		{upperBound: 0.4, count: 319435},
		{upperBound: 0.8, count: 419576},
		{upperBound: 1.6, count: 469593},
		{upperBound: math.Inf(1), count: 519593}, // 因为是无穷大,所以这里可以表示记录总数
	}
	//fmt.Println(len(a)-1)


	b := sort.Search(len(a)-1, func(i int) bool { fmt.Println(i); return a[i].count >= 389694.75 })
	fmt.Println(a[4])
	fmt.Println(b)
}