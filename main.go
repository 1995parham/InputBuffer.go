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

	"github.com/1995parham/InputBuffer.go/run"
)

func main() {
	fmt.Print("Simulation file name (without .yml extension): ")

	var f string
	if _, err := fmt.Scanf("%s", &f); err != nil {
		log.Fatal(err)
	}

	f += ".yml"

	fmt.Print("Results file name (empty if you do not want it): ")

	var r string
	if _, err := fmt.Scanf("%s", &r); err != nil {
		log.Fatal(err)
	}

	var w io.Writer = ioutil.Discard

	if r != "" {
		r += ".results"

		f, err := os.OpenFile(r, os.O_WRONLY|os.O_CREATE, 0755)
		if err == nil {
			w = f
		}
	}

	data, err := ioutil.ReadFile(f)
	if err != nil {
		log.Fatalf("Reading file %s failed with: %s\n", f, err)
	}

	if err := run.Run(data, w); err != nil {
		log.Fatalf("Parsing file %s failed with: %s\n", f, err)
	}
}
