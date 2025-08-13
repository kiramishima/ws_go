package main

import (
	"fmt"
)

/*
 * Complete the 'gradingStudents' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts INTEGER_ARRAY grades as parameter.
 */

func gradingStudents(grades []int32) []int32 {

	for idx, grade := range grades {
		if grade >= 38 && grade%5 >= 3 {
			for grade%5 != 0 {
				grade++
			}
			grades[idx] = grade
		}
	}

	return grades
}

func main() {
	grades := []int32{73, 67, 38, 33}
	fmt.Println(gradingStudents(grades)) // [33 40 67 75]
}
