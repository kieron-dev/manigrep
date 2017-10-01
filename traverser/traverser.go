package traverser

import "fmt"

import "strings"

type Traverser struct {
	data          *interface{}
	keyStack      []string
	keyMatcher    func(key string) bool
	valMatcher    func(val string) bool
	foundCallback func(path string)
}

func New(data *interface{}) (*Traverser, error) {
	tr := Traverser{}
	tr.data = data
	tr.keyStack = []string{}
	return &tr, nil
}

func (tr *Traverser) dfs(data *interface{}) {
	if list, ok := (*data).([]map[string]interface{}); ok {
		fmt.Println("it's a list", list)
		for i, item := range list {
			tr.keyStack = append(tr.keyStack, fmt.Sprintf("[%d]", i))
			var interfaceItem interface{}
			interfaceItem = item
			tr.dfs(&interfaceItem)
		}
	} else if mp, ok := (*data).(map[string]interface{}); ok {
		fmt.Println("it's a map", mp)
		for key, val := range mp {
			if tr.keyMatcher(key) {
				tr.foundCallback("foo")
			}
			tr.dfs(&val)
		}
	} else {
		fmt.Println("it's not a list or map", *data)
	}
}

func (tr *Traverser) SearchKey(keySearch string, resChan chan string) {
	tr.keyMatcher = func(key string) bool {
		return strings.Contains(key, keySearch)
	}
	tr.valMatcher = func(_ string) bool {
		return false
	}
	tr.foundCallback = func(path string) {
		resChan <- path
	}
	tr.dfs(tr.data)
	fmt.Println("Closing resChan")
	close(resChan)
}
