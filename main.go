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
	"io/ioutil"
	"log"
	"os"

	"github.com/AUTProjects/InputBuffer.go/run"
)

func main() {
	var f string
	var r string

	fmt.Print("Simulation file name: ")
	fmt.Scanf("%s", &f)

	fmt.Print("Results file name (empty if you do not want it): ")
	fmt.Scanf("%s", &r)
	if r != "" {
		r = r + ".results"
	}

	data, err := ioutil.ReadFile(f)
	if err != nil {
		log.Fatalf("Reading file %s failed with: %s\n", f, err)
		os.Exit(1)
	}

	if err := run.Run(data); err != nil {
		log.Fatalf("Parsing file %s failed with: %s\n", f, err)
	}
}
