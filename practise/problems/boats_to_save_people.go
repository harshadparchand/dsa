package problems

import (
	"fmt"
	"sort"
)

//https://leetcode.com/problems/boats-to-save-people/
//You are given an array people where people[i] is the weight of the ith person, and an infinite number of boats where each boat can carry a maximum weight of limit.
//Each boat carries at most two people at the same time, provided the sum of the weight of those people is at most limit.

func NumberOfBoats() {
	// Given array of numbers with each number representing weight of a person in KG
	persons := []int{3, 3, 7, 10, 4, 8, 3}
	//Sort the array in ascending order
	sort.Ints(persons)
	// Boat can carry maximum of 10KG weight and a boat can carry at max 2 person. How many boats will be needed
	maxWeight := 10
	leftPointer := 0
	rightPointer := len(persons) - 1
	numberOfBoats := 0
	for leftPointer <= rightPointer {
		remain := maxWeight - persons[rightPointer]
		rightPointer--
		numberOfBoats++
		if leftPointer <= rightPointer && remain >= persons[leftPointer] {
			leftPointer++
		}
	}
	fmt.Println(numberOfBoats)
}
