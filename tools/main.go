package main

import (
	"bufio"
	"fmt"
	"os"

	. "woodgear.com/chain/pkg"
)

func main() {
	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
	raws := []string{}
	for scanner.Scan() {
		raw := scanner.Text()
		raws = append(raws, raw)
	}
	cs := []ChainItem{}
	for _, raw := range raws {
		chain, err := ParseChain(raw)
		if err != nil {
			panic(err)
		}
		cs = append(cs, *chain)
	}
	c, err := NewChain(cs)
	if err != nil {
		panic(err)
	}
	for _, c := range c.GetItems() {
		fmt.Println(c.Show())
	}
}
