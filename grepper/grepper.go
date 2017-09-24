package grepper

import (
	"fmt"
)

type Grepper struct {
	filename string
}

func New(filename string) (*Grepper, error) {
	// parse file here and store in struct
	return &Grepper{filename: filename}, nil
}

func (grep *Grepper) SearchVal(val string) ([]string, error) {
	return nil, fmt.Errorf("Can't find '%s' in %s", val, grep.filename)
}
