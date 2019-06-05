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
			name:      "Exact match",
			namespace: "pod.spec.containers.livenessProbe.exec.command",
			resource:  "command",
			want: []Match{
				{
					Namespace:  "pod.spec.containers.livenessProbe.exec.command",
					MatchScore: matchingScore("command", "command"),
				},
			},
		},
		{
			name:      "Partial match",
			namespace: "pod.spec.containers.livenessProbe.exec.command",
			resource:  "kowmand",
			want: []Match{
				{
					Namespace:  "pod.spec.containers.livenessProbe.exec.command",
					MatchScore: matchingScore("command", "kowmand"),
				},
			},
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
			name:      "Exact match on resource name",
			namespace: "pod.spec.containers.livenessProbe.exec",
			resource:  "exec",
			want: []Match{
				{
					Namespace:  "pod.spec.containers.livenessProbe.exec",
					MatchScore: matchingScore("exec", "exec"),
				},
				{
					Namespace:  "pod.spec.containers.livenessProbe.exec.command",
					MatchScore: matchingScore("exec", "command"),
				},
			},
		},
		{
			name:      "Partial match on resource name",
			namespace: "pod.spec.containers.livenessProbe.exec",
			resource:  "ezek",
			want: []Match{
				{
					Namespace:  "pod.spec.containers.livenessProbe.exec",
					MatchScore: matchingScore("exec", "ezek"),
				},
				{
					Namespace:  "pod.spec.containers.livenessProbe.exec.command",
					MatchScore: matchingScore("exec", "command"),
				},
			},
		},
		{
			name:      "Exact match on field name",
			namespace: "pod.spec.containers.livenessProbe.exec",
			resource:  "command",
			want: []Match{
				{
					Namespace:  "pod.spec.containers.livenessProbe.exec",
					MatchScore: matchingScore("exec", "command"),
				},
				{
					Namespace:  "pod.spec.containers.livenessProbe.exec.command",
					MatchScore: matchingScore("command", "command"),
				},
			},
		},
		{
			name:      "Partial match on field name",
			namespace: "pod.spec.containers.livenessProbe.exec",
			resource:  "cumand",
			want: []Match{
				{
					Namespace:  "pod.spec.containers.livenessProbe.exec",
					MatchScore: matchingScore("exec", "cumand"),
				},
				{
					Namespace:  "pod.spec.containers.livenessProbe.exec.command",
					MatchScore: matchingScore("command", "cumand"),
				},
			},
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
