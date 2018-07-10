package main

import (
	"testing"
	"os"
)

func TestForFile(t *testing.T) {
	os.Setenv("CODEJAM_INPUT", "../../inout/hr_woc38_60_cyclical_queries/00.in")
	main()
}
