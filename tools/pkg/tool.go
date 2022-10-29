package pkg

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/exp/constraints"
)

type ChainItem struct {
	File  string
	Line  int
	Chain string
	Step  string
	Desc  string
}

func (c *ChainItem) Show() string {
	ps := strings.Split(c.File, "/")
	lastFile := ps[len(ps)-1]
	return fmt.Sprintf("%s %v %s %s", lastFile, c.Line, c.Step, c.Desc)
}

func ParseChain(input string) (*ChainItem, error) {
	var file string
	var line int
	var chain string
	var step string
	var desc string
	var err error

	m := strings.SplitN(input, ":", 2)
	file = m[0]
	left := m[1]
	m = strings.SplitN(left, ":", 2)
	line, err = strconv.Atoi(m[0])
	if err != nil {
		return nil, err
	}
	left = m[1]
	chain = ""
	left, _, err = take(left, "wg-chain")
	if err != nil {
		return nil, err
	}
	left, meta, err := take(left, ":")
	meta = meta[:len(meta)-1]
	if err != nil {
		return nil, err
	}
	chain = strings.Fields(meta)[0]
	step = strings.Fields(meta)[1]
	desc = strings.Trim(left, " ")
	return &ChainItem{
		File:  file,
		Line:  line,
		Chain: chain,
		Step:  step,
		Desc:  desc,
	}, nil
}

func take(input string, m string) (left string, out string, err error) {
	index := strings.Index(input, m)
	if index == -1 {
		return input, "", fmt.Errorf("could not find %v", m)
	}
	return input[index+len(m):], input[:index+len(m)], nil
}

type Chain struct {
	items []ChainItem
}

func NewChain(cs []ChainItem) (*Chain, error) {
	c := &Chain{items: cs}
	err := c.sort()
	if err != nil {
		return nil, err
	}
	return c, nil
}

func readNum(in string) (left string, num int, err error) {
	left = in
	pos := 1
	if in[0] == '-' {
		pos = -1
		left = left[1:]
	}
	if in[0] == '+' {
		left = left[1:]
	}
	index := strings.IndexFunc(left, func(c rune) bool {
		return !(c >= '0' && c <= '9')
	})
	end := len(left)
	if index != -1 {
		end = index
	}
	numStr := left[:end]
	num, err = strconv.Atoi(numStr)
	if err != nil {
		return "", 0, err
	}
	return left[end:], num * pos, nil
}

func parseStep(in string) ([]int, error) {
	out := []int{}
	left := in
	var num int
	var err error
	for {
		left, num, err = readNum(left)
		if err != nil {
			return nil, err
		}
		out = append(out, num)
		if left == "" {
			return out, nil
		}
	}
}

func min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func is_left_first(left []int, right []int) bool {
	end := min(len(left), len(right))
	for i := 0; i < end; i++ {
		lv := left[i]
		rv := right[i]
		if lv == rv {
			continue
		}
		return lv < rv
	}
	if len(left) == len(right) {
		return true
	}
	if len(left) > len(right) {
		val := left[len(right)]
		// 1-1 1 true
		// 1+1 1 false
		return val < 0
	}
	if len(left) < len(right) {
		// 1 1-1 false
		// 1 1+1 true
		val := right[len(left)]
		return val > 0
	}
	return len(left) > len(right)
}

func (c *Chain) sort() error {
	steps := map[string][]int{}
	for _, c := range c.items {
		s, err := parseStep(c.Step)
		if err != nil {
			return err
		}
		steps[c.Step] = s
	}

	sort.SliceStable(c.items, func(i, j int) bool {
		left := steps[c.items[i].Step]
		right := steps[c.items[j].Step]
		ret := is_left_first(left, right)
		// fmt.Printf("%v %v %v %v %v \n", left, right, ret, c.items[i].Step, c.items[j].Step)
		return ret
	})
	return nil
}
func (c *Chain) GetItems() []ChainItem {
	return c.items
}
