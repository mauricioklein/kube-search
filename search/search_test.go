package search

import (
	"bytes"
	"errors"
	"io/ioutil"
	"sort"
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

	s := New("pod.spec.containers.livenessProbe.exec", "command")
	s.setRunner(fakeRunner{stdout: buffer, stderr: nil, err: nil})

	matches, err := s.Run()

	assert.Nil(t, err)
	assert.Equal(t, []Match{
		{
			Namespace:  "pod.spec.containers.livenessProbe.exec",
			MatchScore: matchingScore("exec", "command"),
		},
		{
			Namespace:  "pod.spec.containers.livenessProbe.exec.command",
			MatchScore: matchingScore("command", "command"),
		},
	}, matches)
}

func TestSearch_Fail(t *testing.T) {
	errorMsg := "kubectl failed"

	stderr := bytes.NewBufferString("kubectl failed")

	s := New("pod.spec.containers.livenessProbe.exec", "command")
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

func TestSortByMatchingScore(t *testing.T) {
	matches := []Match{
		{Namespace: "Match 0", MatchScore: 1.0},
		{Namespace: "Match 1", MatchScore: 0.8},
		{Namespace: "Match 2", MatchScore: 0.8},
	}

	cases := []struct {
		name  string
		given []Match
		want  []Match
	}{
		{
			name:  "With different scores in ascending order",
			given: []Match{matches[1], matches[0]},
			want:  []Match{matches[0], matches[1]},
		},
		{
			name:  "With different scores in descending order",
			given: []Match{matches[0], matches[1]},
			want:  []Match{matches[0], matches[1]},
		},
		{
			name:  "With same scores",
			given: []Match{matches[1], matches[2]},
			want:  []Match{matches[1], matches[2]},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			sort.Sort(ByMatchingScore(c.given))
			assert.Equal(t, c.want, c.given)
		})
	}
}
