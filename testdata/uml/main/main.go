package main

import (
	uml "github.com/maobuji/go-package-plantuml/testdata/uml"
	sub "github.com/maobuji/go-package-plantuml/testdata/uml/sub"
	sub2 "github.com/maobuji/go-package-plantuml/testdata/uml/sub2"
	"fmt"
)


func main() {
	fmt.Println(uml.SA{})
	fmt.Println(sub.SA{})
	fmt.Println(sub2.Sub2A{})
}

