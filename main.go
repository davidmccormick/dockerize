package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

type sliceVar []string

type Context struct {
}

func (c *Context) Env() map[string]string {
	env := make(map[string]string)
	for _, i := range os.Environ() {
		sep := strings.Index(i, "=")
		env[i[0:sep]] = i[sep+1:]
	}
	return env
}

var (
	buildVersion string
	version      bool

)

func (s *sliceVar) Set(value string) error {
	*s = append(*s, value)
	return nil
}

func (s *sliceVar) String() string {
	return strings.Join(*s, ",")
}

func main() {

	flag.BoolVar(&version, "version", false, "show version")
	//flag.Var(&templatesFlag, "template", "Template (/template:/dest). Can be passed multiple times")

	flag.Parse()

	if version {
		fmt.Println(buildVersion)
		return
	}

	if flag.NArg() == 0 {
		log.Fatalln("you must enter a list of templates and files to write")
	}

	for _, t := range flag.Args() {
		parts := strings.Split(t, ":")
		if len(parts) != 2 {
			log.Fatalf("bad template argument: %s. expected \"/template:/dest\"", t)
		}
		generateFile(parts[0], parts[1])
	}
}
