package main

import "fmt"

// ! CHALLENGE URL: https://www.hackerrank.com/challenges/grading/problem

func main() {
	grades := []int32{4, 73, 67, 38, 33}
	result := gradingStudents(grades)
	for _, r := range result {
		fmt.Printf("%d\n", r)
	}
}

func gradingStudents(grades []int32) []int32 {
	min := int32(38)
	var updatedGrades []int32

	for i, g := range grades {
		if i != 0 {

			if g >= min {
				shouldRound := g%5 >= 3

				if shouldRound {
					diff := 5 - g%5
					updatedGrades = append(updatedGrades, g+diff)
				} else {
					updatedGrades = append(updatedGrades, g)
				}
			} else {
				updatedGrades = append(updatedGrades, g)
			}
		}
	}

	return updatedGrades
}
