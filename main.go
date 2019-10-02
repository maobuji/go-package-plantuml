package main

import (
	"fmt"
	"github.com/ahilbig/go-package-plantuml/codeanalysis"
	"github.com/jessevdk/go-flags"
	log "github.com/sirupsen/logrus"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func main() {

	log.SetLevel(log.InfoLevel)

	var opts struct {
		CodeDir          string   `short:"c" long:"codedir" description:"Code directory to scan" required:"true"`
		GopathDir        string   `short: "g" long:"gopath" description:"GOPATH directory"`
		OutputFile       string   `short:"o" long:"outputfile" description:"The result of the analysis is saved in this file"`
		IgnoreDirs       []string `short:"i" long:"ignoredir" description:"Need to be excluded, no need to scan and parse"`
		IgnoreImplements []string `long:"ii" description:"Implementation that needs to be excluded"`
		//Svg bool `long:"svg" description:"/Output svg format"`
	}

	if len(os.Args) == 1 {
		fmt.Println("Use examples\n" +
			os.Args[0] + " --codedir /appdev/gopath/src/github.com/contiv/netplugin --gopath /appdev/gopath --outputfile  result.md")
		os.Exit(1)
	}

	_, err := flags.ParseArgs(&opts, os.Args)

	if err != nil {
		os.Exit(1)
	}

	curPath, _ := filepath.Abs(filepath.Dir("."))
	curPath = strings.Replace(curPath, "\\", "/", -1)

	GOPATH := os.Getenv("GOPATH")
	GOPATH = strings.Replace(GOPATH, "\\", "/", -1)

	if opts.CodeDir == "" {
		opts.CodeDir = curPath
	} else {
		if !filepath.IsAbs(opts.CodeDir) {
			opts.CodeDir = filepath.Join(curPath, opts.CodeDir)
		}

	}

	if opts.GopathDir == "" {
		opts.GopathDir = GOPATH
		if opts.GopathDir == "" {
			panic("GOPATH directory cannot be empty")
			os.Exit(1)
		}
	}
	if opts.OutputFile == "" {
		fmt.Println("Output file is set to use puml.txt as output file")
		opts.OutputFile = "puml.txt"
	}
	opts.OutputFile, _ = filepath.Abs(opts.OutputFile)

	currentPath := getCurrentDirectory(opts.OutputFile)
	createErr := os.MkdirAll(currentPath, 0777)
	if err != nil {
		fmt.Printf("%s", createErr)
		panic("GOPATH directory cannot be empty")
		os.Exit(1)
	}

	if !strings.HasPrefix(opts.CodeDir, opts.GopathDir) {
		panic(fmt.Sprintf("Code directory %s, must be subdirectory of GOPATH directory %s ", opts.CodeDir, opts.GopathDir))
		os.Exit(1)
	}

	for i, dir := range opts.IgnoreDirs {
		if !filepath.IsAbs(dir) {
			dir = filepath.Join(opts.CodeDir, dir)
			opts.IgnoreDirs[i] = dir
		}

		if !strings.HasPrefix(dir, opts.CodeDir) {
			panic(fmt.Sprintf("Need to be excluded %s, must be subdirectory of GOPATH directory %s ", dir, opts.CodeDir))
			os.Exit(1)
		}
	}

	config := codeanalysis.Config{
		CodeDir:          opts.CodeDir,
		GopathDir:        opts.GopathDir,
		VendorDir:        path.Join(opts.CodeDir, "vendor"),
		IgnoreDirs:       opts.IgnoreDirs,
		IgnoreImplements: opts.IgnoreImplements,
	}

	result := codeanalysis.AnalysisCode(config)

	result.OutputToFile(opts.OutputFile)

	//dir := filepath.Dir(opts.OutputFile)

	jarPath := GOPATH + "/third"
	cmd := fmt.Sprintf("java -jar %s/plantuml.jar %s -tsvg", jarPath, opts.OutputFile)

	rst := codeanalysis.Exec(cmd)
	fmt.Print(cmd)
	fmt.Print(rst.Err)
	fmt.Print(string(rst.Output))
}

func getCurrentDirectory(tempFile string) string {
	dir, err := filepath.Abs(filepath.Dir(tempFile))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}
