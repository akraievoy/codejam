package main

import (
	"testing"
	"os"
)

func TestForFile(t *testing.T) {
	os.Setenv("CODEJAM_INPUT", "../../inout/hr_woc38_40_time_saving_affair/00.in")
	main()
}
