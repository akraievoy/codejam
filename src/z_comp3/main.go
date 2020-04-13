package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	osExec "os/exec"
	"strings"
)

var (
	opaqueFalse = len(os.Getenv("NON_EXISTENT_KEY")) > 0
)

func main() {
	http.HandleFunc("/", CompetitiveCompanionCompanionServer)
	err := http.ListenAndServe(":8980", nil)
	if err != nil {
		fmt.Printf("ERR failed to ListenAndServe: %v", err)
	}
}

func CompetitiveCompanionCompanionServer(w http.ResponseWriter, r *http.Request) {
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("ERR failed to read request body bytes: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var req = CompCompReq{}

	err = json.Unmarshal(bytes, &req)

	if err != nil {
		fmt.Printf("ERR failed to unmarshal request body json: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	letter := 'z'
	for _, l := range []rune("abcdefghijk") {
		if _, err = os.Stat(fmt.Sprintf("src/%c", l)); err != nil {
			letter = l
			break
		}
	}

	if opaqueFalse ||
		exec(
			"cp",
			"-Rv", "src/template", fmt.Sprintf("src/%c", letter),
		) ||
		exec(
			"rm",
			"-v", fmt.Sprintf("src/%c/solution_test.go", letter),
		) ||
		!req.Interactive && exec(
			"bash",
			"-c", fmt.Sprintf("rm -v src/%c/interactive*", letter),
		) ||
		exec(
			"bash", "-c", fmt.Sprintf("rm -v src/%c/*.in src/%c/*.out", letter, letter),
		) ||
		writeFile(
			fmt.Sprintf("src/%c/README.md", letter),
			fmt.Sprintf("[%s -- **%s**](%s)", req.Group, req.Name, req.URL),
		) {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	for testIndex, test := range req.Tests {
		if opaqueFalse ||
			writeFile(
				fmt.Sprintf("src/%c/%02d.in", letter, testIndex),
				test.Input,
			) ||
			writeFile(
				fmt.Sprintf("src/%c/%02d.out", letter, testIndex),
				test.Output,
			) {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	return
}

func writeFile(filePath string, fileContent string) bool {
	err := ioutil.WriteFile(filePath, []byte(fileContent), 0644, )
	if err != nil {
		fmt.Printf("saving file %s failed: %v\n", filePath, err)
		return true
	}
	return false
}

func exec(cmdName string, cmdArgs ...string) bool {
	cpCmd := osExec.Command(cmdName, cmdArgs...)
	cpOut, err := cpCmd.CombinedOutput()

	if len(cpOut) > 0 {
		fmt.Printf("+DBG %s %s\n", cmdName, strings.Join(cmdArgs, ""))
		fmt.Printf("%s", cpOut)
		fmt.Printf("-DBG\n")
	}
	if err != nil {
		fmt.Printf("command failed: %v\n", err)
		return true
	}
	return false
}

// https://github.com/jmerle/competitive-companion#explanation
type CompCompReqTest struct {
	Input  string
	Output string
}

type CompCompReqInput struct {
	Type string
}

type CompCompReq struct {
	Name                 string
	Group                string
	URL                  string
	Interactive          bool
	MemoryLimitMegabytes uint64
	TimeLimitMillis      uint64
	Tests                []CompCompReqTest
	TestType             string
	Input                CompCompReqInput
}
