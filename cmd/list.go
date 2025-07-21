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

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List saved cluster names",
	Long: "Used to list all saved cluster names" +
		"\nExample:\n  kbox list",
	Run: list,
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func list(cmd *cobra.Command, args []string) {
	names, err := pkg.GetServerNames()
	if err != nil {
		fmt.Println("Error while fetching names:", err)
		os.Exit(1)
	}

	for _, name := range names {
		fmt.Printf("%s\n", name)
	}

	fmt.Println("\nTotal count: ", len(names))
	os.Exit(0)
}
