package main

import (
	uml "github.com/ahilbig/go-package-plantuml/testdata/uml"
	sub "github.com/ahilbig/go-package-plantuml/testdata/uml/sub"
	sub2 "github.com/ahilbig/go-package-plantuml/testdata/uml/sub2"
	"fmt"
)


func main() {
	fmt.Println(uml.SA{})
	fmt.Println(sub.SA{})
	fmt.Println(sub2.Sub2A{})
}

