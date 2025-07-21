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

// saveCmd represents the save command
var saveCmd = &cobra.Command{
	Use:   "save",
	Short: "Save oc creds",
	Long: "Saves current logged in oc creds or prompt for creds if not logged in" +
		"\nExample:\n  kbox save",
	Run: save,
}

func init() {
	rootCmd.AddCommand(saveCmd)
}

func save(cmd *cobra.Command, args []string) {

	if len(args) > 0 {
		fmt.Println("Error: Expected 0 arguments")
		os.Exit(1)
	}

	loggedIN := pkg.IsLoggedIn()

	if loggedIN {
		err := pkg.SaveCurrentOC()
		if err != nil {
			println("Error while saving current oc creds:", err)
			os.Exit(1)
		}
	} else {
		println("Not Logged into a server")
		err := pkg.SaveOcByPrompt()
		if err != nil {
			println("Error while saving:", err)
			os.Exit(1)
		}
	}
	os.Exit(0)
}
