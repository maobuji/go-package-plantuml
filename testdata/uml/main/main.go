package main

import (
	uml "github.com/wwwido/go-package-plantuml/testdata/uml"
	sub "github.com/wwwido/go-package-plantuml/testdata/uml/sub"
	sub2 "github.com/wwwido/go-package-plantuml/testdata/uml/sub2"
	"fmt"
)


func main() {
	fmt.Println(uml.SA{})
	fmt.Println(sub.SA{})
	fmt.Println(sub2.Sub2A{})
}

