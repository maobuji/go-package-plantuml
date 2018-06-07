package main

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/jessevdk/go-flags"
	"github.com/maobuji/go-package-plantuml/codeanalysis"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func main() {

	log.SetLevel(log.InfoLevel)

	var opts struct {
		CodeDir    string   `long:"codedir" description:"Code directory to scan" required:"true"`
		GopathDir  string   `long:"gopath" description:"GOPATH directory"`
		OutputFile string   `long:"outputfile" description:"The result of the analysis is saved in this file"`
		IgnoreDirs []string `long:"ignoredir" description:"Need to be excluded, no need to scan and parse"`
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

	if opts.CodeDir == "" {
		panic("Code directory cannot be empty")
		os.Exit(1)
	}

	if opts.GopathDir == "" {
		opts.GopathDir = os.Getenv("GOPATH")
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

	for _, dir := range opts.IgnoreDirs {
		if !strings.HasPrefix(dir, opts.CodeDir) {
			panic(fmt.Sprintf("Need to be excluded %s, must be subdirectory of GOPATH directory %s ", dir, opts.CodeDir))
			os.Exit(1)
		}
	}

	config := codeanalysis.Config{
		CodeDir:    opts.CodeDir,
		GopathDir:  opts.GopathDir,
		VendorDir:  path.Join(opts.CodeDir, "vendor"),
		IgnoreDirs: opts.IgnoreDirs,
	}

	result := codeanalysis.AnalysisCode(config)

	result.OutputToFile(opts.OutputFile)

}

func getCurrentDirectory(tempFile string) string {
	dir, err := filepath.Abs(filepath.Dir(tempFile))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}
