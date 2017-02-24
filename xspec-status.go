//usr/bin/env go run $0 "$@" ; exit

package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Test struct {
	Success string `xml:"successful,attr"`
}

type Scenario struct {
	Label string `xml:"label"`
	Test  *Test  `xml:"test"`
}

func (s *Scenario) String() string {
	return fmt.Sprintf("[Scenario label='%s' success='%v']", s.Label, s.Test.Success)
}

type Report struct {
	Scenarios []*Scenario `xml:"scenario"`
}

func (r *Report) Errors() int {
	errors := 0
	for _, scenario := range r.Scenarios {
		if scenario.Test.Success == "false" {
			errors++
		}
	}
	return errors
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("You must pass Xspec report files on command line\n")
		os.Exit(1)
	}
	files := os.Args[1:]
	errors := 0
	for _, file := range files {
		fmt.Printf("Parsing report file '%s'\n", file)
		source, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Printf("Could not read file '%s': %v\n", file, err)
			os.Exit(2)
		}
		var report Report
		err = xml.Unmarshal(source, &report)
		if err != nil {
			fmt.Printf("Could parse file '%s': %v\n", file, err)
			os.Exit(3)
		}
		errors += report.Errors()
	}
	fmt.Println("Errors:", errors)
	os.Exit(errors)
}
