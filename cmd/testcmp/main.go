package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"lesiw.io/testcmp"
)

func main() { singlechecker.Main(testcmp.Analyzer) }
