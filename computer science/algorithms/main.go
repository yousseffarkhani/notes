package main

import (
	"fmt"
)

type Node struct {
	Value int
	left  *Node
	right *Node
}

func main() {
	nodes := read()
	fmt.Println(nodes)
}

func printNode(n Node) {
	fmt.Print("Value: ", n.Value)
	if n.left != nil {
		fmt.Print(" Left: ", n.left.Value)
	}
	if n.right != nil {
		fmt.Print(" Right: ", n.right.Value)
	}
	fmt.Println()
}

func read() []Node {
	var N int
	fmt.Scanf("%d", &N)
	var nodes []Node = make([]Node, N)

	for i := 0; i < N; i++ {
		var value, indexLeft, indexRight int
		fmt.Scanf("%d %d %d", &value, &indexLeft, &indexRight)
		nodes[i].Value = value
		if indexLeft >= 0 {
			nodes[i].left = &nodes[indexLeft]
		}
		if indexRight >= 0 {
			nodes[i].right = &nodes[indexRight]
		}
	}
	return nodes
}

/*
type Stack struct {
	slice []int
}

func (s *Stack) String() string {
	return fmt.Sprint(s.slice)
}

func (s *Stack) push(num int) {
	s.slice = append(s.slice, num)
}

func (s *Stack) pop() int {
	if len(s.slice) == 0 {
		return -1
	}
	s.slice = s.slice[:len(s.slice)-1]
	return s.peek()
}

func (s *Stack) peek() int {
	return s.slice[len(s.slice)-1]
}

func binarySearch(numList []int, valueToFind, low, high int) int {
	if low > high {
		return -1
	}
	mid := low + (high-low)/2
	middleVal := numList[mid]
	if middleVal == valueToFind {
		return mid
	}
	if middleVal > valueToFind {
		return binarySearch(numList, valueToFind, low, mid-1)
	}
	return binarySearch(numList, valueToFind, mid+1, high)
}

func search(input []int, value int) int {
	for index, val := range input {
		if val == value {
			return index
		}
	}
	return 0
}

func search2(input []int, value int) bool {
	middleIndex := int(math.Ceil(float64(len(input) / 2)))
	if input[middleIndex] == value {
		return true
	}
	if len(input) == 1 {
		return false
	}
	if input[middleIndex] < value {
		return search2(input[:middleIndex], value)
	}
	return search2(input[middleIndex:], value)
} */

// func merge(inputA, inputB []string) []string {
// 	var result []string
// 	if len(inputA) == 1 {
// 		return inputA
// 	}
// 	if len(inputB) == 1 {
// 		return inputB
// 	}
// 	if inputA[0] == inputB[0] {
// 		result = append(result, inputA[0])
// 	}
// 	if inputA[0] > inputB[0] {
// 		result = append(result, inputB[0])
// 	}
// 	result = append(result, inputA[0])
// 	return result
// }
/*
func mergeSort(A []int) []int {
	if len(A) == 1 {
		return A
	}
	middle := int(math.Floor(float64(len(A)) / float64(2)))
	leftHald := A[:middle]
	rightHalf := A[middle:]
	return merge(mergeSort(leftHald), mergeSort(rightHalf))
}
func merge(A, B []int) []int {
	var result []int
	if len(A) == 0 {
		return B
	}
	if len(B) == 0 {
		return A
	}
	if A[0] > B[0] {
		result = append([]int{B[0]}, merge(A, B[1:])...)
		return result
	} else {
		result = append([]int{A[0]}, merge(A[1:], B)...)
		return result
	}
}
*/

/* func mergeLoop(A, B []int) []int {
	var result []int
	for _, valA := range A {
		for i, valB := range B {
			if valA < valB {
				result = append(result, valA)
				break
			}
			result = append(result, valB)
			if len(B)-1 == i {
				result = append(result, A[i:]...)
			}
		}
	}
	return result
} */

// func Fact(number int) int {
// 	fact := 1
// 	if number == 1 {
// 		return fact
// 	}
// 	return number * Fact(number-1)
// }
/*
func sort(input []int) []int {
	var sorted []int
	for _ = range input {
		max, index := Max(input)
		sorted = append(sorted, max)
		input = append(input[:index], input[index+1:]...)
	}
	return sorted
}

func Max(input []int) (int, int) {
	max := input[0] // 2 instructions : 1 pour rechercher input[0] et l'autre pour l'assigner à max
	index := 0

	for i := 0; i < len(input); i++ {
		if max < input[i] {
			max = input[i] // 1 instructions
			index = i
		}
	}
	return max, index
}

func Repeat(name string, n int) {
	startTime := time.Now()
	for i := 0; i < n; i++ {
		for i := 0; i < n; i++ {
			for i := 0; i < n; i++ {
				for i := 0; i < n; i++ {
					fmt.Println(name)
				}
			}
		}
	}
	endTime := time.Now().Sub(startTime)
	fmt.Println(endTime)
}
*/
