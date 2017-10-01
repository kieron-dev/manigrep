package traverser_test

import (
	"fmt"

	"github.com/kieron-pivotal/manigrep/traverser"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Traverser", func() {

	It("can iterate through a top-level list", func() {
		var data interface{}
		data = []map[string]interface{}{
			map[string]interface{}{"alf": 30},
			map[string]interface{}{"bob": 50},
		}
		tr, _ := traverser.New(&data)
		resChan := make(chan string)
		go tr.SearchKey("", resChan)
		res := []string{}
		for r := range resChan {
			res = append(res, r)
		}
		fmt.Println("res len", len(res))
		Expect(len(res)).ToNot(Equal(0))
	})

})
