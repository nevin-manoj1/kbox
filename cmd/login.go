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

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Used to login to saved clusters",
	Long: "Used to login or switch to an oc cluster saved previosly using the save command" +
		"\nExample:\n  kbox login my-cluster",
	Run: login,
}

func init() {
	rootCmd.AddCommand(loginCmd)
}

func login(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		fmt.Println("Error: cluster name is required")
		os.Exit(1)
	}

	server := args[0]
	creds, err := pkg.GetCredentials(server)
	if err != nil {
		fmt.Println("Error while retriving creds: ", err)
		os.Exit(1)
	}

	err = pkg.LoginUsingCredsDefault(creds.User, creds.Password, server)
	if err != nil {
		fmt.Println("Error while logging in:", err)
		os.Exit(1)
	}
	os.Exit(0)
}
