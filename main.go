package main

import (
	"github.com/wwwido/go-package-plantuml/codeanalysis"
	log "github.com/Sirupsen/logrus"
	"fmt"
	"github.com/jessevdk/go-flags"
	"os"
	"strings"
	"path"
	"path/filepath"
)

func main() {

	log.SetLevel(log.InfoLevel)

	var opts struct {
		CodeDir string `long:"codedir" description:"要扫描的代码目录"`
		GopathDir string `long:"gopath" description:"GOPATH目录"`
		OutputFile string `long:"outputfile" description:"解析结果保存到该文件中"`
		IgnoreDirs []string `long:"ignoredirs" description:"需要排除的目录,不需要扫描和解析"`
		//Svg bool `long:"svg" description:"输出 svg格式"`
	}

	//if len(os.Args) == 1 {
	//	fmt.Println("使用例子\n" +
	//		os.Args[0] + " --codedir /appdev/gopath/src/github.com/contiv/netplugin --gopath /appdev/gopath --outputfile  /tmp/result")
	//	os.Exit(1)
	//}

	_, err := flags.ParseArgs(&opts, os.Args)

	if err != nil {
		os.Exit(1)
	}

	curPath,_ := filepath.Abs(filepath.Dir("."))
	curPath = strings.Replace(curPath, "\\", "/", -1)

	GOPATH := os.Getenv("GOPATH")
	GOPATH = strings.Replace(GOPATH, "\\", "/", -1)

	if opts.CodeDir == "" {
		opts.CodeDir = curPath
	}

	if opts.GopathDir == "" {
		opts.GopathDir = GOPATH
	}

	if ! strings.HasPrefix(opts.CodeDir, opts.GopathDir) {
		panic(fmt.Sprintf("代码目录%s,必须是GOPATH目录%s的子目录", opts.CodeDir, opts.GopathDir))
		os.Exit(1)
	}

	for i, dir := range opts.IgnoreDirs {
		if !filepath.IsAbs(dir) {
			dir = opts.CodeDir + "/" + dir
			opts.IgnoreDirs[i] = dir
		}

		if ! strings.HasPrefix(dir, opts.CodeDir){
			panic(fmt.Sprintf("需要排除的目录%s,必须是代码目录%s的子目录", dir, opts.CodeDir))
			os.Exit(1)
		}
	}

	config := codeanalysis.Config{
		CodeDir: opts.CodeDir,
		GopathDir :opts.GopathDir,
		VendorDir : path.Join(opts.CodeDir, "vendor"),
		IgnoreDirs:opts.IgnoreDirs,
	}

	result := codeanalysis.AnalysisCode(config)

	if opts.OutputFile == "" {
		opts.OutputFile = curPath + "/uml.txt"
	}
	result.OutputToFile(opts.OutputFile)

	//if !opts.Svg {
	//	return
	//}

	//dir := filepath.Dir(opts.OutputFile)

	jarPath := GOPATH + "/third"
	cmd := fmt.Sprintf("java -jar %s/plantuml.jar %s -tsvg", jarPath, opts.OutputFile)

	rst := codeanalysis.Exec(cmd)
	fmt.Print(cmd)
	fmt.Print(rst.Err)
	fmt.Print(string(rst.Output))
}

