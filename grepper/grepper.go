package grepper

import (
	"fmt"
	"io/ioutil"

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

func (gp *Grepper) SearchVal(val string) ([]string, error) {
	return nil, fmt.Errorf("Can't find '%s' in %s", val, gp.filename)
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
