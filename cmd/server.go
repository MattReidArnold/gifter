package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "start server",
	Run:   serverRun,
}

func init() {
	rootCmd.AddCommand(serverCmd)
}

func serverRun(cmd *cobra.Command, args []string) {
	fmt.Println("Starting server...")
}
