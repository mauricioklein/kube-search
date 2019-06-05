package cmd

import (
	"fmt"

	"github.com/mauricioklein/kube-search/search"
)

type printable interface {
	print(match search.Match)
}

type simplePrinter struct{}

func (sp simplePrinter) print(match search.Match) {
	fmt.Println(match.Namespace)
}

type scorePrinter struct{}

func (sp scorePrinter) print(match search.Match) {
	fmt.Printf("%s (matching score: %f)\n", match.Namespace, match.MatchScore)
}
