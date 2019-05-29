package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"

	"github.com/softlandia/xlib"
)

//CollExclude2 - from "i n x y z" makes "x y z"
func CollExclude2(fn string) error {
	if !xlib.FileExists(os.Args[1]) {
		return fmt.Errorf("file '%s' not found", fn)
	}
	if filepath.Ext(fn) == "xyz" {
		return fmt.Errorf("change ext of file: '" + filepath.Ext(fn) + "' from 'xyz' to other")
	}
	iFile, err := os.Open(os.Args[1])
	if err != nil {
		return fmt.Errorf("error on open file: '%s'", fn)
	}
	defer iFile.Close()
	scanner := bufio.NewScanner(iFile)
	var (
		s string
		i,
		n int
		x,
		y,
		z float64
	)
	oFile, err := os.Create(os.Args[1] + ".xyz")
	defer oFile.Close()
	if err != nil {
		fmt.Printf("output file not open")
		os.Exit(1)
	}

	for scanner.Scan() {
		s = scanner.Text()
		fmt.Sscanf(s, "%d %d %f %f %f", &i, &n, &x, &y, &z)
		fmt.Fprintf(oFile, "%f %f %f \n", x, y, -z)
	}
	return nil
}

//CollExclude1 - from "x y z n" makes "x y z" add -999 -999 -999 when n changed
func CollExclude1(fn string) error {
	if !xlib.FileExists(fn) {
		return fmt.Errorf("file '%s' not found", fn)
	}
	if filepath.Ext(fn) == "xyz" {
		return fmt.Errorf("change ext of file: '" + filepath.Ext(fn) + "' from 'xyz' to other")
	}
	iFile, err := os.Open(fn)
	defer iFile.Close()
	if err != nil {
		return fmt.Errorf("error on open file: '%s'", fn)
	}

	fn = xlib.ChangeFileExt(fn, ".xyz")
	oFile, err := os.Create(fn)
	defer oFile.Close()
	if err != nil {
		return fmt.Errorf("output file '%s' not open to write", fn)
	}

	var (
		s string
		x,
		y,
		z float64
	)

	scanner := bufio.NewScanner(iFile)
	i := 0
	j := 0
	for scanner.Scan() {
		s = scanner.Text()
		fmt.Sscanf(s, "%f %f %f %d", &x, &y, &z, &i)
		if (i != j) && (j > 0) {
			oFile.WriteString("-999 -999 -999\n")
		}
		fmt.Fprintf(oFile, "%f %f %f \n", x, y, z)
		j = i
	}
	return nil
}
