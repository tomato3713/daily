package cmd

import (
	"io"
	"os"
	"os/exec"
	"runtime"
	"sort"
)

func runCmd(command string, r io.Reader, w io.Writer, args ...string) error {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", append([]string{"/c", command}, args...)...)
	} else {
		cmd = exec.Command("sh", append([]string{"-c", command}, args...)...)
	}
	cmd.Stderr = os.Stderr
	cmd.Stdout = w
	cmd.Stdin = r
	return cmd.Run()
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func sortByDate(files []string) []string {
	sort.Sort(sort.Reverse(sort.StringSlice(files)))
	return files
}
