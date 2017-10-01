package grepper

import (
	"fmt"
	"io/ioutil"

	"github.com/kieron-pivotal/manigrep/traverser"
	yaml "gopkg.in/yaml.v2"
)

type Grepper struct {
	filename string
	data     interface{}
}

func New(filename string) (*Grepper, error) {
	gp := Grepper{filename: filename}
	err := gp.loadFile()
	if err != nil {
		return nil, err
	}
	return &gp, nil
}

func (gp *Grepper) SearchVal(val string, resultCh chan string) {
	tr, _ := traverser.New(&gp.data)
	go tr.SearchKey(val, resultCh)
}

func (gp *Grepper) loadFile() error {
	content, err := ioutil.ReadFile(gp.filename)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(content, &gp.data)
	if err != nil {
		return fmt.Errorf("Error parsing YAML\n%v", err)
	}

	return nil
}
