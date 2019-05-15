package search

import (
	"bufio"
	"os"
	"testing"

	"github.com/mauricioklein/go-spacetree/spacetree"
	"github.com/stretchr/testify/assert"
)

type testCase struct {
	name      string
	namespace string
	resource  string
	want      []Match
}

func TestProcess_FieldOutput(t *testing.T) {
	root := loadTree("../kubectl_test_outputs/field.txt")

	cases := []testCase{
		{
			name:      "Matching field",
			namespace: "pod.spec.containers.livenessProbe.exec.command",
			resource:  "command",
			want:      []Match{{Namespace: "pod.spec.containers.livenessProbe.exec.command"}},
		},
		{
			name:      "No match",
			namespace: "pod.spec.containers.livenessProbe.exec",
			resource:  "foobar",
			want:      []Match{},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			matches := searchOnTree(root, c.namespace, c.resource)
			assert.Equal(t, c.want, matches)
		})
	}
}

func TestProcess_ResourceOutput(t *testing.T) {
	root := loadTree("../kubectl_test_outputs/resource.txt")

	cases := []testCase{
		{
			name:      "Matching resource",
			namespace: "pod.spec.containers.livenessProbe",
			resource:  "livenessProbe",
			want:      []Match{{Namespace: "pod.spec.containers.livenessProbe"}},
		},
		{
			name:      "Matching field on first level",
			namespace: "pod.spec.containers.livenessProbe",
			resource:  "exec",
			want:      []Match{{Namespace: "pod.spec.containers.livenessProbe.exec"}},
		},
		{
			name:      "Matching field on deeper level",
			namespace: "pod.spec.containers.livenessProbe",
			resource:  "command",
			want:      []Match{{Namespace: "pod.spec.containers.livenessProbe.exec.command"}},
		},
		{
			name:      "Multiple matches",
			namespace: "pod.spec.containers.livenessProbe",
			resource:  "host",
			want: []Match{
				{Namespace: "pod.spec.containers.livenessProbe.httpGet.host"},
				{Namespace: "pod.spec.containers.livenessProbe.tcpSocket.host"},
			},
		},
		{
			name:      "No match",
			namespace: "pod.spec.containers.livenessProbe",
			resource:  "foobar",
			want:      []Match{},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			matches := searchOnTree(root, c.namespace, c.resource)
			assert.Equal(t, c.want, matches)
		})
	}
}

func loadTree(path string) *spacetree.Node {
	file, _ := os.Open(path)
	buffer := bufio.NewReader(file)
	scanner := bufio.NewScanner(buffer)

	root, _ := spacetree.New(scanner, "   ")

	return root
}
