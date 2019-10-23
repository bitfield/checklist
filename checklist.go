package checklist

import (
	"io/ioutil"
	"strings"
)

type checklist struct {
	list []string
}

func Load(path string) (checklist, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return checklist{}, err
	}
	data := string(bytes)
	list := strings.Split(data, "\n")
	return checklist{list}, nil
}

func (c checklist) Len() int {
	return len(c.list)
}
