package main

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/wwwido/go-package-plantuml/codeanalysis"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	log.SetLevel(log.InfoLevel)

	//config := codeanalysis.Config{
	//	CodeDir: "/appdev/go-demo/src/github.com/wwwido/go-package-plantuml/testdata/a",
	//	GopathDir : "/appdev/go-demo",
	//}
	//
	//result := codeanalysis.AnalysisCode(config)
	//
	//result.OutputToFile("/tmp/uml.txt")
	GOPATH := os.Getenv("GOPATH")
	GOPATH = strings.Replace(GOPATH, "\\", "/", -1)
	file := "."
	basePath,_ := filepath.Abs(filepath.Dir(file))
	basePath = strings.Replace(basePath, "\\", "/", -1)
	config := codeanalysis.Config{
		CodeDir: GOPATH + "/src/github.com/davyxu/cellnet",
		GopathDir : GOPATH,
	}

	result := codeanalysis.AnalysisCode(config)

	result.OutputToFile("./uml.txt")

	rst := codeanalysis.Exec("java -jar plantuml.jar ./uml.txt -tsvg")
	fmt.Print(string(rst.Output))
	fmt.Print(rst.Err)

	//bytes, _ := ioutil.ReadFile("./uml.txt")

	//fmt.Println(string(bytes))

	// java -jar /app/plantuml.jar  /tmp/uml.txt -tsvg && open2 /tmp/uml.svg

}
