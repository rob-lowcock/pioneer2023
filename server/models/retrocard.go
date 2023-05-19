package models

type Retrocard struct {
	ID     string `jsonapi:"primary,retrocard"`
	Title  string `jsonapi:"attr,title"`
	Column int    `jsonapi:"attr,column"`
	Active bool   `jsonapi:"attr,active"`
}
