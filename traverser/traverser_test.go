package traverser_test

import (
	"log"

	"github.com/kieron-pivotal/manigrep/traverser"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Traverser", func() {

	var data interface{}

	BeforeEach(func() {
		data = []map[string]interface{}{
			map[string]interface{}{"alf": 30},
			map[string]interface{}{"bob": 50},
		}
	})

	It("can match a key value in a list", func() {
		tr, _ := traverser.New(&data)
		resChan := make(chan string)
		go tr.SearchKey("lf", resChan)
		res := []string{}
		for r := range resChan {
			log.Println(r)
			res = append(res, r)
		}
		Expect(len(res)).To(Equal(1))
	})

	It("can match a value within a list", func() {
		tr, _ := traverser.New(&data)
		resChan := make(chan string)
		go tr.SearchVal("5", resChan)
		res := []string{}
		for r := range resChan {
			log.Println(r)
			res = append(res, r)
		}
		Expect(len(res)).To(Equal(1))
	})

})
