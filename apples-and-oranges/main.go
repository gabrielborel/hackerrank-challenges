package main

import "fmt"

// ! CHALLENGE URL: https://www.hackerrank.com/challenges/apple-and-orange/problem

func main() {
	s := 7
	t := 11
	a := 5
	b := 15
	apples := []int32{-2, 2, 1}
	oranges := []int32{5, -6}

	countApplesAndOranges(int32(s), int32(t), int32(a), int32(b), apples, oranges)
}

func countApplesAndOranges(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) {
	appleCount := 0
	orangeCount := 0

	for i := range apples {
		apples[i] += a
	}
	for i := range oranges {
		oranges[i] += b
	}

	for _, apple := range apples {
		if apple >= s && apple <= t {
			appleCount++
		}
	}
	for _, orange := range oranges {
		if orange >= s && orange <= t {
			orangeCount++
		}
	}

	fmt.Printf("%d\n", appleCount)
	fmt.Printf("%d\n", orangeCount)
}
