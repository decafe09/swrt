package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"sort"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/spec"
)

var version = "1.0.0"

type route struct {
	Path    string
	Methods []string
}

type routes []route

func (r routes) Len() int {
	return len(r)
}

func (r routes) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func (r routes) Less(i, j int) bool {
	return r[i].Path < r[j].Path
}

func main() {

	var showVersion bool
	var showHelp bool
	var importPath string

	flag.BoolVar(&showVersion, "v", false, "show version")
	flag.BoolVar(&showVersion, "-version", false, "show version")
	flag.BoolVar(&showHelp, "h", false, "show help")
	flag.BoolVar(&showHelp, "-help", false, "show help")
	flag.StringVar(&importPath, "f", "swagger.json", "import path your swagger.json")

	flag.Parse()
	if showHelp {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
		return
	}

	if showVersion {
		fmt.Println("version:", version)
		return
	}

	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("failed to get working direcroty")
		return
	}

	src := filepath.Join(cwd, importPath)

	sw, err := loadSpec(src)
	if err != nil {
		fmt.Println(err)
		return
	}

	var rs routes = make([]route, len(sw.Paths.Paths))
	var i int
	for path, pathItem := range sw.Paths.Paths {
		r := route{Path: path}
		if pathItem.Get != nil {
			r.Methods = append(r.Methods, "GET")
		}
		if pathItem.Put != nil {
			r.Methods = append(r.Methods, "PUT")
		}
		if pathItem.Post != nil {
			r.Methods = append(r.Methods, "POST")
		}
		if pathItem.Delete != nil {
			r.Methods = append(r.Methods, "DELETE")
		}
		if pathItem.Options != nil {
			r.Methods = append(r.Methods, "OPTIONS")
		}
		if pathItem.Head != nil {
			r.Methods = append(r.Methods, "HEAD")
		}
		if pathItem.Patch != nil {
			r.Methods = append(r.Methods, "PATCH")
		}
		rs[i] = r
		i++
	}

	sort.Sort(rs)

	for _, s := range sw.Schemes {
		fmt.Printf("%s://%s\n", s, path.Join(sw.Host, sw.BasePath))
	}
	for _, r := range rs {
		for _, m := range r.Methods {
			fmt.Printf("%6s\t%s\n", m, r.Path)
		}
	}

}

func loadSpec(src string) (*spec.Swagger, error) {
	fi, err := os.Stat(src)
	if err != nil {
		return nil, fmt.Errorf("not exists a file %q", src)
	}

	if fi.IsDir() {
		return nil, fmt.Errorf("expected %q to be a file not a directory", src)
	}

	sp, err := loads.Spec(src)
	if err != nil {
		return nil, err
	}

	return sp.Spec(), nil
}
