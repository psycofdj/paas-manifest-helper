package main

import (
	"fmt"
	"github.com/prometheus/common/version"
	"gopkg.in/alecthomas/kingpin.v2"
	yml "gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

var (
	line  = kingpin.Flag("line", "Cursor line").Default("0").Int()
	col   = kingpin.Flag("col", "Cursor column").Default("0").Int()
	fpath = kingpin.Flag("path", "Path to yml file").String()
	stdin = kingpin.Flag("stdin", "Force to read file content from stdin. path argument is still mandatory").Default("0").Bool()
)

type Manifest struct {
	Consts map[string]interface{} `yaml:"const"`
}

func main() {
	kingpin.Version(version.Print("paas-manifest-helper"))
	kingpin.HelpFlag.Short('h')
	kingpin.Parse()

	var buff []byte
	var err error
	if !*stdin {
		buff, err = ioutil.ReadFile(*fpath)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	} else {
		buff, _ = ioutil.ReadAll(os.Stdin)
	}

	fpath := path.Clean(*fpath)
	if !path.IsAbs(fpath) {
		pwd, _ := os.Getwd()
		fpath = path.Clean(path.Join(pwd, fpath))
	}

	instance := path.Dir(fpath)
	instance_name := path.Base(instance)
	product := path.Dir(path.Dir(instance))
	product_name := path.Base(product)
	root := path.Dir(product)

	manifest := Manifest{
		Consts: map[string]interface{}{
			"root":          root,
			"common":        path.Join(root, "common"),
			"product":       product,
			"product_name":  product_name,
			"instance":      instance,
			"instance_name": instance_name,
		},
	}
	err = yml.Unmarshal(buff, &manifest)
	if err != nil {
		fmt.Printf("unable to parse input file: %s", err)
		os.Exit(1)
	}

	value, err := yml.ValueAtPoint(*line-1, *col, buff)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for key, replace := range manifest.Consts {
		value = strings.Replace(value, "(("+key+"))", fmt.Sprintf("%s", replace), -1)
	}

	fmt.Println(value)
	os.Exit(0)
}
