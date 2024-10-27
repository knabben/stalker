package cmd

import (
	"github.com/knabben/stalker/pkg/testgrid"
	"os"

	"github.com/spf13/cobra"
)

var (
	tg           = testgrid.NewTestGrid()
	testBoards   = []string{"sig-release-master-blocking", "sig-release-master-informing"}
	brokenStatus = []string{testgrid.FAILING_STATUS, testgrid.FLAKY_STATUS}
	rootCmd      = &cobra.Command{
		Use:   "stalker",
		Short: "stalker search for issues and flaky tests on Kubernetes",
		Long:  "stalker search for issues and flaky tests on Kubernetes",
	}
)

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func hasStatus(boardStatus string, statuses []string) bool {
	for _, status := range statuses {
		if boardStatus == status {
			return true
		}
	}
	return false
}
