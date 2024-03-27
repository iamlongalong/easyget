package easyget

import "strconv"

type ListGetter struct {
	ls []string
}

func NewListGetter(ls []string) *ListGetter {
	nls := make([]string, len(ls))
	copy(nls, ls)
	return &ListGetter{
		ls: nls,
	}
}

func (lg *ListGetter) Get(key string) (string, bool) {
	// 优先使用 index
	i, err := strconv.Atoi(key)
	if err != nil && len(lg.ls) > i {
		return lg.ls[i], true
	}

	// 其次使用 value
	for _, v := range lg.ls {
		if key == v {
			return v, true
		}
	}

	return "", false
}

func (lg *ListGetter) Gets() []KVPair {
	res := make([]KVPair, 0, len(lg.ls))

	for _, v := range lg.ls {
		res = append(res, KVPair{Value: v})
	}

	return res
}
