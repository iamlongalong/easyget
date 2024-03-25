package easyget

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

type JSONGetter struct {
	m map[string]interface{}
}

func NewJSONGetter(m map[string]interface{}) *JSONGetter {
	if m == nil {
		m = map[string]interface{}{}
	}
	return &JSONGetter{
		m: m,
	}
}

func (sg *JSONGetter) Get(key string) (string, bool) {
	v, ok := sg.m[key]
	if ok {
		return toString(v), true
	}

	return "", false
}

func (sg *JSONGetter) Gets() []KVPair {
	res := make([]KVPair, 0, len(sg.m))

	for k, v := range sg.m {
		res = append(res, KVPair{Key: k, Value: toString(v)})
	}

	return res
}

func NewJSONGetterFromHTTP(method string, url string, headers map[string]string, body map[string]interface{}) (*JSONGetter, error) {
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(strings.ToUpper(method), url, bytes.NewBuffer(bodyBytes))
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp != nil {
		defer resp.Body.Close()
	}

	var result JSONGetter
	err = json.NewDecoder(resp.Body).Decode(&result.m)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func NewJSONGetterFromJSONFile(file string, subkeys ...string) *JSONGetter {
	b, err := os.ReadFile(file)
	if err != nil {
		l.Errorf("[JSONGetter] read file fail: %s", err)
	}

	m := map[string]interface{}{}

	err = json.Unmarshal(b, &m)
	if err != nil {
		l.Errorf("[JSONGetter] from json file fail: %s", err)
	}

	for _, key := range subkeys {
		submap, ok := m[key].(map[string]interface{})
		if !ok {
			l.Errorf("[JSONGetter] subkey does not exist or is not an object: %s", key)
			return nil
		}

		m = submap
	}

	return &JSONGetter{
		m: m,
	}
}

func NewJSONGetterFromCmd(cmd string, args []string, envs []string, subkeys ...string) *JSONGetter {
	command := exec.Command(cmd, args...)
	command.Env = envs
	out, err := command.Output()
	if err != nil {
		l.Errorf("[NewJSONGetterFromCmd] cmd output fail: %s", err)
		return &JSONGetter{map[string]interface{}{}}
	}

	var m map[string]interface{}
	err = json.Unmarshal(out, &m)
	if err != nil {
		l.Errorf("[NewJSONGetterFromCmd] cmd output unmarshal fail: %s", err)
		return &JSONGetter{map[string]interface{}{}}
	}

	for _, key := range subkeys {
		submap, ok := m[key].(map[string]interface{})
		if !ok {
			l.Errorf("[JSONGetter] subkey does not exist or is not an object: %s", key)
			return nil
		}

		m = submap
	}

	return &JSONGetter{m: m}
}
