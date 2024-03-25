package easyget

import "regexp"

type RegFilter struct {
	reg    string
	regexp *regexp.Regexp
}

func NewRegFilter(reg string) *RegFilter {
	regexp := regexp.MustCompile(reg)

	return &RegFilter{
		regexp: regexp,
		reg:    reg,
	}
}

func (rf *RegFilter) Filt(d []byte) []byte {
	return rf.regexp.Find(d)
}
