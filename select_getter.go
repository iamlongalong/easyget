package easyget

import (
	"fmt"
	"strings"

	"github.com/manifoldco/promptui"
)

type KVPair struct {
	Key   string
	Value string
}

type KVPariGetter interface {
	Gets() []KVPair
}

// read string from stdin
type SelectGetter struct {
	kvsGetter KVPariGetter
}

func NewSelectGetter(kvsGetter KVPariGetter) *SelectGetter {
	return &SelectGetter{
		kvsGetter: kvsGetter,
	}
}

func (sg *SelectGetter) Get(key string) (string, bool) {
	kvs := sg.kvsGetter.Gets()

	items := make([]string, len(kvs))
	for i, kv := range kvs {
		val := ""
		if len(kv.Key) > 0 {
			val = fmt.Sprintf("%s: %s", kv.Key, kv.Value)
		} else {
			val = fmt.Sprintf("%s", kv.Value)
		}

		items[i] = val
	}

	selectPrompt := promptui.Select{
		Label: fmt.Sprintf("Select Value for [%s]", key),
		Items: items,
		// HideSelected: true,
		StartInSearchMode: true,
		Searcher: func(input string, index int) bool {
			return strings.Contains(items[index], input)
		},
	}

	i, _, err := selectPrompt.Run()

	if err != nil {
		l.Errorf("Prompt failed %v", err)
		return "", false
	}

	return kvs[i].Value, true
}
