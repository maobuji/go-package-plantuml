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
		GopathDir        string   `short:"g" long:"gopath" description:"GOPATH directory"`
		OutputFile       string   `short:"o" long:"outputfile" description:"The result of the analysis is saved in this file"`
		OutputReplaceTag string   `short:"r" long:"replacetag" description:"Tags for marking section containing the analysis output"`
		IgnoreDirs       []string `short:"i" long:"ignoredir" description:"Need to be excluded, no need to scan and parse"`
		IgnoreImplements []string `long:"ii" description:"Implementation that needs to be excluded"`
		IgnoreTypeAlias  []string `long:"it" description:"Type to be excluded, will be scanned and parsed but not put into UML"`
		IncludeTypeAlias []string `short:"t" long:"type" description:"Type to be included - if set all other will be ignored"`
		IgnoreFilenames  []string `long:"ignorefile" description:"Filename (without .go suffix) to be excluded, no need to scan and parse"`
		IncludeFilenames []string `short:"f" long:"file" description:"Filename (without .go suffix) to be included - if set all other will be ignored"`
		//Svg bool `long:"svg" description:"/Output svg format"`
	}

	if len(os.Args) == 1 {
		fmt.Println("Use examples\n" +
			os.Args[0] + " --codedir /appdev/gopath/src/github.com/contiv/netplugin --type Contract --type Subscription" +
			"--gopath /appdev/gopath -f myFile1 -f myFile2 -r classDiagramContainingOnlyClassesFromFile1And2 --outputfile result.md")

		fmt.Println()
		fmt.Println(`If --replacetag/-r is specified the UML will be framed by comments containing this tag.\n
Subsequent calls of the analysis will replace just this section if present instead of overwriting the file.\n
This can be used for automatically generating UML snippets during CI/CD Pipelines for a project and\n
embedding them into markdown containing other contents e.g. readme.md.`)
		os.Exit(1)
	}

	_, err := flags.ParseArgs(&opts, os.Args)

	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	curPath, _ := filepath.Abs(filepath.Dir("."))

	opts.CodeDir = filepath.FromSlash(opts.CodeDir)
	opts.GopathDir = filepath.FromSlash(opts.GopathDir)

	GOPATH := os.Getenv("GOPATH")

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
		dir = filepath.FromSlash(dir)
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
		IgnoreTypeAlias:  opts.IgnoreTypeAlias,
		IncludeTypeAlias: opts.IncludeTypeAlias,
		IgnoreFilenames:  opts.IgnoreFilenames,
		IncludeFilenames: opts.IncludeFilenames,
		OutputReplaceTag: opts.OutputReplaceTag,
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
