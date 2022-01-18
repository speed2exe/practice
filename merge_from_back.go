package practice

// input: 2 sorted arrays,
// first: [100, 200, 300, -1, -1]
// second: [1, 2]
// desired output: [1, 2, 100, 200, 300]

// Explanation:
// empty space is denoted to be -1
// all valid integer in both arrays needs to be positive, -1 is special number
// to denote empyty space that can be filled by other number
// output: sorted array of both firt and second

// Requirement:
// space complexity: O(1)
// time complexity: O(n + m), where n is len(first) and m is len(second)

// Constraint:
// number of empty(-1) is always the same len of

type mergeBackwardSolver struct {
	src       []int
	tgt       []int
	empty_idx int
	value_idx int
}

// first: [100, 200, 300, -1, -1]
// second: [1, 2]
func MergeBackwards(first, second []int) {
	solver := mergeBackwardSolver{
		src: first,
		tgt: second,
	}

	solver.init()

	for solver.stillHaveItemToMerge() {
		solver.merge()
	}
}

func (s *mergeBackwardSolver) init() {
	s.empty_idx = len(s.src) - 1
	s.value_idx = len(s.src) - 1

	for {
		if s.value_idx <= 0 {
			return
		}
		if s.src[s.value_idx] != -1 {
			break
		}
		s.value_idx--
	}
}

func (s *mergeBackwardSolver) stillHaveItemToMerge() bool {
	if s.empty_idx < 0 {
		return false
	}
	return s.src[s.empty_idx] == -1
}

func (s *mergeBackwardSolver) merge() {
	srcLast, _ := s.peekSrcLast()
	tgtLast := s.popTgtLast()

	for srcLast > tgtLast {
		s.put(srcLast)
		s.src[s.value_idx] = -1
		s.value_idx--

		var ok bool
		srcLast, ok = s.peekSrcLast()
		if !ok {
			break
		}
	}

	s.put(tgtLast)
}

func (s *mergeBackwardSolver) put(val int) {
	s.src[s.empty_idx] = val
	s.empty_idx--
}

func (s *mergeBackwardSolver) popTgtLast() int {
	lastIdx := len(s.tgt) - 1
	last := s.tgt[lastIdx]
	s.tgt = s.tgt[:lastIdx]
	return last
}

func (s *mergeBackwardSolver) peekSrcLast() (int, bool) {
	if s.value_idx < 0 {
		return 0, false
	}
	value := s.src[s.value_idx]
	return value, true
}
