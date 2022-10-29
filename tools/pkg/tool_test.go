package pkg

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPasrse(t *testing.T) {
	raw := "/nginx-1.19.9/src/core/ngx_cycle.c:617:    // wg-chain listen 1-1: 开始监听端口"
	ret, err := ParseChain(raw)
	assert.NoError(t, err)
	fmt.Printf("ret %v err %v", ret, err)
	assert.Equal(t, *ret, ChainItem{
		File:  "/nginx-1.19.9/src/core/ngx_cycle.c",
		Line:  617,
		Chain: "listen",
		Step:  "1-1",
		Desc:  "开始监听端口",
	})
}

func TestReadNum(t *testing.T) {
	tests := []struct {
		in      string
		expnum  int
		expleft string
	}{
		{
			in:      "1",
			expnum:  1,
			expleft: "",
		},
		{
			in:      "-1",
			expnum:  -1,
			expleft: "",
		},
		{
			in:      "+1",
			expnum:  1,
			expleft: "",
		},
		{
			in:      "1+1",
			expnum:  1,
			expleft: "+1",
		},
		{
			in:      "1-1",
			expnum:  1,
			expleft: "-1",
		},
	}
	for _, tt := range tests {
		t.Run("test-readnum-"+tt.in, func(t *testing.T) {
			left, num, err := readNum(tt.in)
			assert.NoError(t, err)
			assert.Equal(t, left, tt.expleft)
			assert.Equal(t, num, tt.expnum)
		})
	}
}

func TestParseStep(t *testing.T) {
	out, err := parseStep("1")
	assert.NoError(t, err)
	assert.Equal(t, out, []int{1})
	out, err = parseStep("1-1")
	assert.NoError(t, err)
	assert.Equal(t, out, []int{1, -1})
	out, err = parseStep("1+1")
	assert.NoError(t, err)
	assert.Equal(t, out, []int{1, 1})
}

func TestComp(t *testing.T) {
	assert.Equal(t, is_left_first([]int{1}, []int{1, -1}), false)
	assert.Equal(t, is_left_first([]int{1, -1, -1}, []int{1, -1}), true)
	assert.Equal(t, is_left_first([]int{1, -2}, []int{1, -3}), false)
	assert.Equal(t, is_left_first([]int{1, 1}, []int{1}), false)
	assert.Equal(t, is_left_first([]int{1, 1}, []int{2}), false)
}
