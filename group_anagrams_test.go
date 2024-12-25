package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

// 对变位词进行分组

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string, 128)
	for _, str := range strs {
		s := sortCharsInStr(str)
		fmt.Printf("sorted str:%s\n", s)

		var l []string
		if v, present := m[s]; present {
			l = v
		} else {
			l = make([]string, 0, 128)
		}
		l = append(l, str)
		fmt.Printf("k:%s,v:%v\n", s, l)
		m[s] = l
	}
	r := make([][]string, 0, len(m))
	for _, v := range m {
		r = append(r, v)
	}
	return r
}

func sortCharsInStr(s string) string {
	bytes := []byte(s)
	sort.Slice(bytes, func(i, j int) bool {
		return bytes[i] < bytes[j]
	})
	return string(bytes)
}

func TestGroupAnagrams(t *testing.T) {
	cases := []struct {
		title string
		args  []string
		want  [][]string
	}{
		{
			title: "1",
			args:  []string{"eat", "tea", "tan", "ate", "nat", "bat"}, want: [][]string{
			{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"},
		}},
		// {title: "2", args: []string{""}, want: [][]string{{""}}},
		// {title: "3", args: []string{"a"}, want: [][]string{{"a"}}},
	}
	for _, tc := range cases {
		t.Run(tc.title, func(t *testing.T) {
			assert.Equal(t, tc.want, groupAnagrams(tc.args))
		})
	}
}
