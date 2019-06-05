package cmd

import (
	"fmt"
	"os"
	"sort"

	"github.com/mauricioklein/kube-search/search"
	"github.com/spf13/cobra"
)

var cfg config

// RootCmd define the Cobra command for the root
// (i.e. calling kube-search with no command)
var RootCmd = &cobra.Command{
	Use:     "kube-search",
	Short:   "Fuzzy search K8s fields path by namespace and resource name",
	Version: search.Version,
	Run: func(cmd *cobra.Command, args []string) {
		doSearch(cfg)
	},
}

func init() {
	RootCmd.Flags().StringVarP(&cfg.namespace, "namespace", "n", "", "the Kubectl namespace")
	RootCmd.MarkFlagRequired("namespace")

	RootCmd.Flags().StringVarP(&cfg.resource, "resource", "r", "", "the Kubectl resource")
	RootCmd.MarkFlagRequired("resource")

	RootCmd.Flags().Uint16VarP(&cfg.nRecords, "count", "c", 1, fmt.Sprintf("number of results returned by %s", RootCmd.Use))

	RootCmd.Flags().BoolVar(&cfg.printScore, "show-score", false, "print the matching score along with the matches")
}

func doSearch(cfg config) {
	s := search.New(cfg.namespace, cfg.resource)

	matches, err := s.Run()
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

	// Sort the matches by descending matching score
	sort.Sort(search.ByMatchingScore(matches))

	// Get the top n records
	matches = matches[:cfg.nRecords]

	printer := cfg.printer()
	for _, match := range matches {
		printer.print(match)
	}
}
