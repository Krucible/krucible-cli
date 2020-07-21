package cmd

import (
	"os"
	"time"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

// getClustersCmd represents the get clusters command
var getClustersCmd = &cobra.Command{
	Use:   "clusters",
	Short: "Get all clusters within the account",
	Run: func(cmd *cobra.Command, args []string) {
		client := getClientOrDie()
		clusters, err := client.GetClusters()
		if err != nil {
			panic(err)
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetBorder(false)
		for _, c := range clusters {
			table.SetHeader([]string{
				"ID",
				"Name",
				"State",
				"Created",
				"Expires",
			})
			if c.State == "running" {
				table.Append([]string{
					c.ID,
					c.DisplayName,
					c.State,
					c.CreatedAt.Format(time.RFC822),
					c.ExpiresAt.Format(time.RFC822),
				})
			}
		}
		table.Render()
	},
}

func init() {
	getCmd.AddCommand(getClustersCmd)
}
