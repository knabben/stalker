/* Copyright © 2024 Amim Knabben */
package cmd

import (
	"fmt"
	"github.com/knabben/stalker/pkg/tui"
	"log"

	"github.com/spf13/cobra"
)

// abstractCmd represents the abstract command
var abstractCmd = &cobra.Command{
	Use:   "abstract",
	Short: "Summarize the board status and present the flake or failing ones",
	RunE:  RunAbstract,
}

func init() {
	rootCmd.AddCommand(abstractCmd)
}

func RunAbstract(cmd *cobra.Command, args []string) error {
	for _, dashboard := range testBoards {
		summary, err := tg.FetchSummary(dashboard)
		if err != nil {
			log.Fatal("error fetching the summary", err)
		}

		var counter = 0
		for tab, dashboard := range *summary.Dashboards {
			if hasStatus(dashboard.OverallStatus, brokenStatus) {
				table, err := tg.FetchTable(dashboard.DashboardName, tab)
				if err != nil {
					_ = fmt.Errorf("error fetching table : %s", err)
					continue
				}
				issue := tui.NewDashboardIssue(summary.URL, tab, dashboard, table)
				if err = issue.RenderVisual(counter); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
