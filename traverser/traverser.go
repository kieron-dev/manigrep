package traverser

import (
	"fmt"
	"strings"
)

type Traverser struct {
	data          *interface{}
	keyMatcher    func(key string) bool
	valMatcher    func(val string) bool
	foundCallback func(path []string)
}

func New(data *interface{}) (*Traverser, error) {
	tr := Traverser{}
	tr.data = data
	return &tr, nil
}

func (tr *Traverser) dfs(data *interface{}, keyStack []string) {
	if list, ok := (*data).([]map[string]interface{}); ok {
		for i, item := range list {
			newStack := append(keyStack, fmt.Sprintf("[%d]", i))
			var interfaceItem interface{}
			interfaceItem = item
			tr.dfs(&interfaceItem, newStack)
		}
	} else if mp, ok := (*data).(map[string]interface{}); ok {
		for key, val := range mp {
			newStack := append(keyStack, key)
			if tr.keyMatcher(key) {
				tr.foundCallback(newStack)
			}
			tr.dfs(&val, newStack)
		}
	} else {
		valStr := fmt.Sprintf("%v", *data)
		if tr.valMatcher(valStr) {
			tr.foundCallback(keyStack)
		}
	}
}

func noMatch(_ string) bool {
	return false
}

func (tr *Traverser) SearchKey(keySearch string, resChan chan string) {
	tr.keyMatcher = func(key string) bool {
		return strings.Contains(key, keySearch)
	}
	tr.valMatcher = noMatch
	tr.foundCallback = func(path []string) {
		resChan <- strings.Join(path, "/")
	}
	tr.dfs(tr.data, []string{""})
	close(resChan)
}

func (tr *Traverser) SearchVal(valSearch string, resChan chan string) {
	tr.keyMatcher = noMatch
	tr.valMatcher = func(val string) bool {
		return strings.Contains(val, valSearch)
	}
	tr.foundCallback = func(path []string) {
		resChan <- strings.Join(path, "/")
	}
	tr.dfs(tr.data, []string{""})
	close(resChan)
}
