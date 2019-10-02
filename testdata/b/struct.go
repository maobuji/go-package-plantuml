package b

import (
	sub2 "github.com/ahilbig/go-package-plantuml/testdata/b/sub"
	"github.com/ahilbig/go-package-plantuml/testdata/b/suba"
	a "sync"
)

type SB struct {
}

func (this SB) Add(a sub2.SubSA, locker a.Locker, b B, subsa1 suba.SubSa1) {}
