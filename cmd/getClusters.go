package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/Krucible/krucible-go-client/krucible"
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

		activeClusters := []krucible.Cluster{}
		for _, c := range clusters {
			if c.State != "deprovisioned" {
				activeClusters = append(activeClusters, c)
			}
		}

		if len(activeClusters) == 0 {
			fmt.Println("No active clusters found.")
			return
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetBorder(false)
		table.SetHeader([]string{
			"ID",
			"Name",
			"State",
			"Created",
			"Expires",
		})
		for _, c := range activeClusters {
			table.Append([]string{
				c.ID,
				c.DisplayName,
				c.State,
				c.CreatedAt.Format(time.RFC822),
				c.ExpiresAt.Format(time.RFC822),
			})
		}
		table.Render()
	},
}

func init() {
	getCmd.AddCommand(getClustersCmd)
}
