package search

import (
	"bytes"
	"os/exec"
)

type runnable interface {
	run(namespace string) (*bytes.Buffer, *bytes.Buffer, error)
}

type cmdRunner struct{}

// runKubeExplain
func (runner cmdRunner) run(namespace string) (*bytes.Buffer, *bytes.Buffer, error) {
	cmd := exec.Command("kubectl", "explain", "--recursive", namespace)

	var stdout, stderr bytes.Buffer

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()

	return &stdout, &stderr, err
}
