package catsort

// Split returns chunks of same category runes that the string s is divided into.
func Split(s string) []string {
	var e Splitter
	Parse(s, &e)
	return e.List
}

// Splitter is Eval for Split.
type Splitter struct {
	List []string
}

var _ Eval = (*Splitter)(nil)

func (eval *Splitter) Number(s string) { eval.List = append(eval.List, s) }
func (eval *Splitter) Letter(s string) { eval.List = append(eval.List, s) }
func (eval *Splitter) Symbol(s string) { eval.List = append(eval.List, s) }
