// https://www.codewars.com/kata/5663f5305102699bad000056/go

package main

import (
	"fmt"
	"math"
)

func MxDifLg(a1 []string, a2 []string) int {
	if len(a1) == 0 || len(a2) == 0 {
		return -1
	}

	// Find the maximum and minimum length of strings inside the a1 array
	var maxLen1 = 0
	var minLen1 = -1
	for _, str := range a1 {
		if len(str) > maxLen1 {
			maxLen1 = len(str)
		}

		if len(str) < minLen1 || minLen1 == -1 {
			minLen1 = len(str)
		}
	}

	// Find the maximum and minimum length of strings inside the a2 array
	var maxLen2 = 0
	var minLen2 = -1
	for _, str := range a2 {
		if len(str) > maxLen2 {
			maxLen2 = len(str)
		}

		if len(str) < minLen2 || minLen2 == -1 {
			minLen2 = len(str)
		}
	}

	var len1 = int(math.Abs(float64(maxLen1 - minLen2)))
	var len2 = int(math.Abs(float64(maxLen2 - minLen1)))
	var len = int(math.Max(float64(len1), float64(len2)))

	return len
}

func main(){
	fmt.Println(MxDifLg([]string{"hoqq", "bbllkw", "oox", "ejjuyyy", "plmiis", "xxxzgpsssa", "xxwwkktt", "znnnnfqknaz", "qqquuhii", "dvvvwz"}, []string{"cccooommaaqqoxii", "gggqaffhhh", "tttoowwwmmww"}))
}