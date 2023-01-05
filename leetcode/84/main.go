// https://leetcode.com/problems/largest-rectangle-in-histogram/description/

package main

import "fmt"

type Instance struct {
	height int
	count  int
}

func (i *Instance) area() int {
	return i.height * i.count
}

type Stack struct {
	instances []Instance
}

func (s *Stack) isEmpty() bool {
	return len(s.instances) == 0
}

func (s *Stack) push(i Instance) {
	s.instances = append(s.instances, i)
}

func (s *Stack) peek() (Instance, bool) {
	if s.isEmpty() {
		return Instance{}, false
	}

	result := s.instances[len(s.instances)-1]
	return result, true
}

func (s *Stack) pop() (Instance, bool) {
	if s.isEmpty() {
		return Instance{}, false
	}

	result := s.instances[len(s.instances)-1]
	s.instances = s.instances[:len(s.instances)-1]
	return result, true
}

func (s *Stack) modifyTop(i Instance) {
	s.instances[len(s.instances)-1] = i
}

func largestRectangleArea(heights []int) int {
	stack := Stack{
		instances: []Instance{},
	}

	result := 0

	for _, h := range heights {
		top, exists := stack.peek()
		if !exists || (top.height < h) {
			stack.push(Instance{count: 1, height: h})
			continue
		}

		times := 0
		for {
			popped, exists := stack.pop()
			if !exists {
				stack.push(Instance{count: times + 1, height: h})
				break
			}

			if popped.height < h {
				stack.push(popped)
				stack.push(Instance{count: times + 1, height: h})
				break
			}

			if popped.height == h {
				stack.push(Instance{count: times + popped.count + 1, height: h})
				break
			}

			popped.count, times = popped.count+times, popped.count+times
			area := popped.area()
			if area > result {
				result = area
			}
		}
	}

	times := 0
	for {
		popped, exists := stack.pop()
		if !exists {
			break
		}
		popped.count, times = popped.count+times, popped.count+times
		area := popped.area()
		if area > result {
			result = area
		}
		continue
	}

	return result
}

func main() {
	fmt.Println(largestRectangleArea([]int{4, 2, 0, 3, 2, 4, 3, 4}))
}
