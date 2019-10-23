package checklist

import (
	"errors"
	"io/ioutil"
	"strings"
)

type checklist struct {
	list     []string
	nextItem int
}

type item struct {
	text string
}

func Load(path string) (checklist, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return checklist{}, err
	}
	data := string(bytes)
	list := strings.Split(data, "\n")
	return checklist{list, 0}, nil
}

func (c checklist) Len() int {
	return len(c.list)
}

func (c *checklist) NextItem() (item, error) {

	if c.nextItem >= len(c.list) {
		return item{}, errors.New("no more items")
	}

	line := c.list[c.nextItem]
	c.nextItem++
	return item{line}, nil
}
