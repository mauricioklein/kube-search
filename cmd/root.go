package cmd

import (
	"fmt"
	"os"

	"github.com/mauricioklein/kube-search/search"
	"github.com/spf13/cobra"
)

// RootCmd define the Cobra command for the root
// (i.e. calling kube-search with no command)
var RootCmd = &cobra.Command{
	Use:     "kube-search",
	Version: search.Version,
	Run: func(cmd *cobra.Command, args []string) {
		namespace := cmd.Flag("namespace").Value.String()
		resource := cmd.Flag("resource").Value.String()

		doSearch(namespace, resource)
	},
}

func init() {
	RootCmd.Flags().StringP("namespace", "n", "", "the Kubectl namespace")
	RootCmd.MarkFlagRequired("namespace")

	RootCmd.Flags().StringP("resource", "r", "", "the Kubectl resource")
	RootCmd.MarkFlagRequired("resource")
}

func doSearch(namespace, resource string) {
	s := search.New(namespace, resource)

	matches, err := s.Run()
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

	for _, match := range matches {
		fmt.Println(match.Namespace)
	}
}
