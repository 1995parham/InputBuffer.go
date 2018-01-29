/*
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 26-01-2018
 * |
 * | File Name:     main.go
 * +===============================================
 */

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/AUTProjects/InputBuffer.go/run"
)

type fakeWriter struct{}

func (*fakeWriter) Write(p []byte) (n int, err error) {
	return len(p), nil
}

func main() {
	var f string
	var r string

	fmt.Print("Simulation file name: ")
	fmt.Scanf("%s", &f)

	var w io.Writer = &fakeWriter{}

	fmt.Print("Results file name (empty if you do not want it): ")
	fmt.Scanf("%s", &r)
	if r != "" {
		r = r + ".results"
		f, err := os.OpenFile(r, os.O_WRONLY|os.O_CREATE, 0755)
		if err == nil {
			w = f
		}
	}

	data, err := ioutil.ReadFile(f)
	if err != nil {
		log.Fatalf("Reading file %s failed with: %s\n", f, err)
		os.Exit(1)
	}

	if err := run.Run(data, w); err != nil {
		log.Fatalf("Parsing file %s failed with: %s\n", f, err)
	}
}
