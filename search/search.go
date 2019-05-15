package search

import (
	"bufio"
	"errors"

	"github.com/mauricioklein/go-spacetree/spacetree"
)

const (
	kubeIndentationSymbol = "   " // 3 spaces
)

// Search defines the parameters used to search for
// occurences in the "Kubectl explain" tree
type Search struct {
	Namespace string
	Resource  string
	Runner    runnable
}

// Match defines a match on the "Kubectl explain" tree
type Match struct {
	Namespace string
}

// New returns a new instance of Search
func New(ns string, res string) Search {
	return Search{
		Namespace: ns,
		Resource:  res,
		Runner:    cmdRunner{},
	}
}

func (s *Search) setRunner(r runnable) {
	s.Runner = r
}

// Run executes the "Kubectl explain" command
// and look for matches in the tree
func (s *Search) Run() ([]Match, error) {
	// Run the Kubectl command
	stdout, stderr, err := s.Runner.run(s.Namespace)
	if err != nil {
		return nil, errors.New(stderr.String())
	}

	// Parse the "Kubectl explain" tree
	scanner := bufio.NewScanner(stdout)
	root, err := spacetree.New(scanner, kubeIndentationSymbol)

	if err != nil {
		return nil, err
	}

	// Calculate the matches
	matches := searchOnTree(root, s.Namespace, s.Resource)

	return matches, nil
}
