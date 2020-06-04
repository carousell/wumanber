package wumanber

import (
	"reflect"
	"testing"
)

func TestInit(t *testing.T) {
	var patterns []string = []string{"你好", "世界"}
	var wm WuManber
	err := wm.Init(patterns)
	if err != nil {
		t.Error("init wumanber error")
	}
	if 3 != wm.Block {
		t.Error("block size is not 3")
	}
	if 1003 != wm.TableSize {
		t.Error("table size if not 2")
	}

	if 1003 != len(wm.ShiftTable) {
		t.Error("shift table size if not 2")
	}
	if 1003 != len(wm.HashTable) {
		t.Error("hash table size if not 2")
	}
}

var (
	cases = []struct {
		keywords                []string
		text                    string
		expectedMatches         []string
		expectedMatchPostitions []int
	}{
		{
			[]string{"brown", "jump"},
			"the quick brown fox jumps over the lazy dog.",
			[]string{"brown", "jump"},
			[]int{10, 20},
		},
		{
			[]string{"dog."},
			"the quick brown fox jumps over the lazy dog.",
			[]string{"dog."},
			[]int{40},
		},
		{
			[]string{"brown fox"},
			"the quick brown fox jumps over the lazy dog.",
			[]string{"brown fox"},
			[]int{10},
		},
		{
			[]string{"brown", "dog", "jump"},
			"The Quick Brown Fox Jumps Over The Lazy Dog.",
			[]string{},
			[]int{},
		},
		{
			// Text contains unicode. (Code points may mess with
			// certain implementations and cause the keyword to match)
			[]string{"na", "🌟n"},
			"🌟nly‼️😄 Quick Snatch",
			[]string{"🌟n", "na"},
			[]int{0, 25},
		},
		{
			// Matches duplicate
			[]string{"one"},
			"one one one one two",
			[]string{"one", "one", "one", "one"},
			[]int{0, 4, 8, 12},
		},
		{
			// Matches are ordered in order of appearance
			[]string{"two", "one"},
			"one two three",
			[]string{"one", "two"},
			[]int{0, 4},
		},
		{
			[]string{"a", "b", "c", "#1 Delta Executive Homes Aniban", "("},
			"aa a c cccc d ddddd #1 Delta Executive Homes Aniban",
			[]string{"a", "a", "a", "c", "c", "c", "c", "c", "#1 Delta Executive Homes Aniban", "a", "c", "b", "a"},
			[]int{0, 1, 3, 5, 7, 8, 9, 10, 20, 27, 32, 48, 49},
		},
		{
			[]string{"你好", "世界", "abc"},
			"北京你好，世界很大啊",
			[]string{"你好", "世界"},
			[]int{6, 15},
		},
	}
)

func TestSearch(t *testing.T) {
	for _, c := range cases {
		wm := WuManber{}
		wm.Init(c.keywords)
		matches, matchPositions := wm.Search(c.text)
		if len(matches) == 0 && len(c.expectedMatches) == 0 {
			continue
		}
		isEqual := reflect.DeepEqual(matches, c.expectedMatches)
		if !isEqual {
			t.Fatalf(`matcher.Match("testBucket", %+v) => using keywords=%+v, got %+v; want %+v`, c.text, c.keywords, matches, c.expectedMatches)
		}

		isEqual = reflect.DeepEqual(matchPositions, c.expectedMatchPostitions)
		if !isEqual {
			t.Fatalf(`matcher.Match("testBucket", %+v) => using keywords=%+v, got %+v; want %+v`, c.text, c.keywords, matchPositions, c.expectedMatchPostitions)
		}
	}
}
