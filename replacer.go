// Package replacer provides a drop-in replacement for Go's built-in
// strings.Replacer for cases when it is not quite up to the task. For example:
//
//	r := strings.NewReplacer(
//	  "modifiedbyid", "modified_by_id",
//	  "modifiedby", "modified_by")
//	fmt.Println(r.Replace("modifiedbyid"))
//	// Output may be: modified_by_id
//	//            OR: modified_byid
//
// By contrast, the replacer type returned by replacer.New will always produce
// the same output by preferring the longest match.
package replacer

import _ "embed"

//go:embed version.txt
var Version string

type dfaState struct {
	next           map[rune]*dfaState
	hasReplacement bool
	replacement    string
}

func (s *dfaState) addReplacement(expect []rune, replacement string) {
	if len(expect) == 0 {
		s.hasReplacement = true
		s.replacement = replacement
		return
	}

	if s.next == nil {
		s.next = map[rune]*dfaState{}
	}

	next, ok := s.next[expect[0]]
	if !ok {
		next = &dfaState{}
		s.next[expect[0]] = next
	}

	next.addReplacement(expect[1:], replacement)
}

func (s *dfaState) match(l int, rs []rune) (int, string) {
	if len(rs) > 0 {
		if nextState, hasNext := s.next[rs[0]]; hasNext {
			cmpL, replacement := nextState.match(l+1, rs[1:])
			if cmpL > l {
				return cmpL, replacement
			}
		}
	}

	if s.hasReplacement {
		return l, s.replacement
	}
	return 0, ""
}

type Replacer struct {
	dfa *dfaState
}

func New(replacements ...string) *Replacer {
	root := &dfaState{}
	for i := 0; i < len(replacements); i += 2 {
		if len(replacements[i]) == 0 {
			continue
		}
		root.addReplacement([]rune(replacements[i]), replacements[i+1])
	}
	return &Replacer{root}
}

func (r *Replacer) Replace(s string) string {
	work := []rune(s)

	i := 0
	for i < len(work) {
		l, replacement := r.dfa.match(0, work[i:])
		if l == 0 {
			i++
			continue
		}

		switch {
		case len(replacement) < l:
			// shrink
			delta := l - len(replacement)
			copy(work[i+len(replacement):], work[i+l:])
			work = work[:len(work)-delta]
		case len(replacement) > l:
			// grow
			delta := len(replacement) - l
			if len(work)+delta > cap(work) {
				newWork := make([]rune, len(work)+delta, max(cap(work)*2, (len(work)+delta)*2))
				copy(newWork, work[:i])
				copy(newWork[i+len(replacement):], work[i+l:])
				work = newWork
			} else {
				work = work[:len(work)+delta]
				copy(work[i+len(replacement):], work[i+l:])
			}
		}
		copy(work[i:], []rune(replacement))

		i += len(replacement)
	}

	return string(work)
}
