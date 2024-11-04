package easyget

import (
	"fmt"
	"io"
	"os"
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
	config    *SelectConfig
}

type SelectConfig struct {
	Prompt string
	Stdin  io.ReadCloser
	Stdout io.WriteCloser
}

func NewSelectGetter(kvsGetter KVPariGetter, config *SelectConfig) *SelectGetter {
	if config == nil {
		config = &SelectConfig{}
	}

	if config.Stdin == nil {
		config.Stdin = os.Stdin
	}
	if config.Stdout == nil {
		config.Stdout = os.Stdout
	}

	return &SelectGetter{
		kvsGetter: kvsGetter,
		config:    config,
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

	prompt := sg.config.Prompt
	if len(prompt) == 0 {
		prompt = fmt.Sprintf("Select Value for [%s]", key)
	}

	selectPrompt := promptui.Select{
		Label: prompt,
		Items: items,
		// HideSelected: true,
		StartInSearchMode: true,
		Searcher: func(input string, index int) bool {
			return strings.Contains(strings.ToLower(items[index]), strings.ToLower(input))
		},
		Stdin:  sg.config.Stdin,
		Stdout: sg.config.Stdout,
	}

	i, _, err := selectPrompt.Run()

	if err != nil {
		l.Errorf("Prompt failed %v", err)
		return "", false
	}

	return kvs[i].Value, true
}
