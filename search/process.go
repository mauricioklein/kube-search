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
	matches := make([]Match, 0)

	if strings.HasPrefix(node.Value, "FIELD:") {
		// FIELD match
		matches = append(matches, newMatch(namespace, matchingScore(nodeValue, resource)))
	} else if strings.HasPrefix(node.Value, "RESOURCE:") {
		// RESOURCE match
		matches = append(matches, newMatch(namespace, matchingScore(nodeValue, resource)))
	} else if strings.HasPrefix(node.Value, "FIELDS:") {
		matches = append(matches, processFieldsTree(node, namespace, resource)...)
	}

	return matches
}

func processFieldsTree(node *spacetree.Node, namespace string, resource string) []Match {
	matches := make([]Match, 0)

	for _, child := range node.Children {
		childValue := sanitizeLine(child.Value)
		ns := strings.Join([]string{namespace, childValue}, ".")

		matches = append(matches, newMatch(ns, matchingScore(childValue, resource)))

		// Recursive call
		matches = append(matches, processFieldsTree(child, ns, resource)...)
	}

	return matches
}
