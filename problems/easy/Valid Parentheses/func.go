package Valid_Parentheses

import "strings"

func isValid(s string) bool {

	if len(s) == 0 {
		return true
	}
	if len(s)%2 == 1 {
		return false
	}

	var stck Stack

	for _, br := range s {

		if strings.Contains("({[", string(br)) {
			stck.push(br)
		} else {
			upperbr := stck.upper()
			if br == ')' && upperbr == '(' {
				stck.pop()
			} else if br == ']' && upperbr == '[' {
				stck.pop()
			} else if br == '}' && upperbr == '{' {
				stck.pop()
			} else {
				return false
			}

		}
	}

	return stck.empty()

}

type Stack struct {
	bracket []rune
	n       int
}

func (s *Stack) push(r rune) {
	s.n++
	s.bracket = append(s.bracket, r)
}

func (s *Stack) pop() {
	s.bracket = s.bracket[:s.n-1]
	s.n -= 1
}

func (s *Stack) upper() rune {
	if s.n != 0 {
		return s.bracket[s.n-1]
	} else {
		return 0
	}
}
func (s *Stack) empty() bool {
	return len(s.bracket) == 0
}
