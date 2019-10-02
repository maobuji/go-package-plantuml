package a

import (
	"github.com/ahilbig/go-package-plantuml/testdata/uml/sub2"
	"sync"
)

type IA interface {
	Add()
}

type SA struct {
	a int
	b sync.Mutex
	c sub2.Sub2A
	m map[string]sub2.Sub2A
}

func (this *SA) Add() {}

func NewSA() *SA {
	return &SA{}
}
