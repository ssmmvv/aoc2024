package problems

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func setupStdoutMock() (*os.File, *os.File, chan string) {
	old_stdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	mock_stdout := make(chan string)
	go func() {
		var buff bytes.Buffer
		io.Copy(&buff, r)
		mock_stdout <- buff.String()
	}()
	return old_stdout, w, mock_stdout
}

func cleanupStdoutMock(old_stdout *os.File, w *os.File) {
	w.Close()
	os.Stdout = old_stdout
}

const (
	expect_distance   = "distance: 1830467"
	expect_similarity = "simialrity: 26674158"
)

func TestProblem1(t *testing.T) {
	old_stdout, w, mock_stdout := setupStdoutMock()
	Problem1("../inputs/1")
	cleanupStdoutMock(old_stdout, w)
	stdout_str := <-mock_stdout

	if !strings.Contains(stdout_str, expect_distance) {
		t.Errorf("Incorrect distance answer, expected %s\n", expect_distance)
	}

	if !strings.Contains(stdout_str, expect_similarity) {
		t.Errorf("Incorrect similarity, expected %s\n", expect_similarity)
	}

	close(mock_stdout)
}
