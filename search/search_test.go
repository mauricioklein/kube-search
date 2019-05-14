package search

import (
	"bytes"
	"errors"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

type fakeRunner struct {
	stdout *bytes.Buffer
	stderr *bytes.Buffer
	err    error
}

func (fr fakeRunner) run(namespace string) (*bytes.Buffer, *bytes.Buffer, error) {
	return fr.stdout, fr.stderr, fr.err
}

func TestSearch_Success(t *testing.T) {
	content, _ := ioutil.ReadFile("../kubectl_test_outputs/resource.txt")
	buffer := bytes.NewBuffer(content)

	s := New("pod.spec.containers.livenessProbe", "livenessProbe")
	s.setRunner(fakeRunner{stdout: buffer, stderr: nil, err: nil})

	matches, err := s.Run()

	assert.Nil(t, err)
	assert.Equal(t, []Match{{Namespace: "pod.spec.containers.livenessProbe"}}, matches)
}

func TestSearch_Fail(t *testing.T) {
	errorMsg := "kubectl failed"

	stderr := bytes.NewBufferString("kubectl failed")

	s := New("pod.spec.containers.livenessProbe", "livenessProbe")
	s.setRunner(fakeRunner{
		stdout: nil,
		stderr: bytes.NewBufferString(errorMsg),
		err:    errors.New(""),
	})

	matches, err := s.Run()

	assert.Error(t, err)
	assert.EqualError(t, err, stderr.String())
	assert.Empty(t, matches)
}
