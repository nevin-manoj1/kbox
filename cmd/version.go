/*
Copyright © 2025 Nevin Manoj Nevin-Manoj@ibm.com
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints current version of kbox",
	Long:  "Used to print the version of kbox being used",
	Run:   getVersion,
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

func getVersion(cmd *cobra.Command, args []string) {
	fmt.Printf("kbox version: %s\n", version)
}
