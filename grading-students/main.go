package main

import "fmt"

/*
	! CHALLENGE URL: https://www.hackerrank.com/challenges/grading/problem

 	! HackerLand University has the following grading policy:

	! Every student receives a grade in the inclusive range from 0 to 100.
	! Any grade less than 40 is a failing grade.
	! The university uses the following rounding rules:

	! If the difference between the grade and the next multiple of 5 is less than 3, round the grade up to the next multiple of 5.
	! If the grade is less than 38, no rounding occurs, and the result is the original grade.
	! Write a function to round the grades of a list of students according to the university's grading policy.
*/

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
