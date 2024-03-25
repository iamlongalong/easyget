package easyget

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// read string from stdin
type StdGetter struct {
	filters []IFilter
}

func NewStdGetter(filters ...IFilter) *StdGetter {
	return &StdGetter{
		filters: filters,
	}
}

func (sg *StdGetter) Get(key string) (string, bool) {
	// 如果在终端中运行, 打印提示信息
	if isTerminal() {
		fmt.Printf("please input the value of %s:\t", key)
	}

	reader := bufio.NewReader(os.Stdin)
	val, err := reader.ReadString('\n')
	if err != nil {
		l.Errorf("read from std getter fail: %s", err)
	}

	val = strings.TrimSuffix(val, "\n") // 去掉换行符

	// 遍历过滤器并对输入进行处理
	for _, filter := range sg.filters {
		valBytes := []byte(val)
		val = string(filter.Filt(valBytes))
	}

	return val, true
}
