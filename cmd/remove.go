/*
Copyright © 2025 Nevin Manoj Nevin-Manoj@ibm.com
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/nevin-manoj1/kbox/pkg"
	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "To remove saved oc creds",
	Long: "Used to remove saved oc creds using server name as input" +
		"\nExample:\n  kbox remove my-cluster my-clsuter-2",
	Run: remove,
}

func init() {
	rootCmd.AddCommand(removeCmd)
}

func remove(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		fmt.Println("Atleast 1 server name required")
		os.Exit(1)
	}
	for _, server := range args {
		pkg.RemoveServer(server)
	}
	os.Exit(0)
}
