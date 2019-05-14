package search

import (
	"strings"

	"github.com/mauricioklein/go-spacetree/spacetree"
)

// searchOnTree starts the recursive search on the tree
// for matching resources
func searchOnTree(root *spacetree.Node, namespace string, resource string) []Match {
	matches := make([]Match, 0)

	for _, child := range root.Children {
		matches = append(matches, dig(child, namespace, resource)...)
	}

	return matches
}

// dig is the recursive part of the "searchOnTree" method
func dig(node *spacetree.Node, namespace string, resource string) []Match {
	nodeValue := sanitizeLine(node.Value)
	var ns string

	if strings.HasPrefix(node.Value, "FIELDS:") {
		ns = namespace
	} else {
		ns = strings.Join([]string{namespace, nodeValue}, ".")
	}

	matches := make([]Match, 0)

	if strings.HasPrefix(node.Value, "FIELD:") && nodeValue == resource {
		// FIELD match
		matches = append(matches, Match{Namespace: namespace})
	} else if strings.HasPrefix(node.Value, "RESOURCE:") && nodeValue == resource {
		// RESOURCE match
		matches = append(matches, Match{Namespace: namespace})
	} else if nodeValue == resource {
		// Child FIELD match
		matches = append(matches, Match{Namespace: ns})
	}

	// Dig into the children
	for _, child := range node.Children {
		matches = append(matches, dig(child, ns, resource)...)
	}

	return matches
}
