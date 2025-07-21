/*
Copyright © 2025 Nevin Manoj Nevin-Manoj@ibm.com
*/
package pkg

import (
	"fmt"
	"net/url"
	"os/exec"
	"strings"
)

func LoginUsingCredsDefault(user string, pwd string, server string) error {
	fullServer := "https://api." + server + ".cp.fyre.ibm.com:6443"
	return LoginUsingCreds(user, pwd, fullServer)
}

func LoginUsingCreds(user string, pwd string, server string) error {
	cmd := exec.Command("oc", "login", server, "--username", user, "--password", pwd, "--insecure-skip-tls-verify=true")
	output, err := cmd.CombinedOutput()
	fmt.Println(string(output))
	if err != nil {
		return err
	}
	cmd1 := exec.Command("oc", "project")
	output1, err := cmd1.CombinedOutput()
	if err != nil {
		return err
	}
	fmt.Println(string(output1))
	return nil
}

func SaveCurrentOC() error {

	cmd0 := exec.Command("oc", "project")
	output0, err := cmd0.CombinedOutput()
	if err != nil {
		return err
	}
	fmt.Println(string(output0))

	cmd := exec.Command("oc", "whoami")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	user := strings.TrimSpace(strings.ReplaceAll(string(output), ":", ""))

	var pwd string
	fmt.Print("Enter ocp password: ")
	fmt.Scanln(&pwd)

	cmd1 := exec.Command("oc", "whoami", "--show-server")
	output1, err := cmd1.CombinedOutput()
	if err != nil {
		return err
	}
	server, err := extractSubdomain(string(output1))
	if err != nil {
		return err
	}

	err = SaveCredentials(user, pwd, server)
	if err != nil {
		return err
	}
	fmt.Printf("Saved current oc login creds, use 'kbox login %s' to oc login in future", server)
	return nil
}

func SaveOcByPrompt() error {

	var user string
	fmt.Print("Enter ocp user(eg: kubeadmin ): ")
	fmt.Scanln(&user)

	var pwd string
	fmt.Print("Enter ocp password(eg: A1B2C-a1b2c-A1B2C-A1B2C ): ")
	fmt.Scanln(&pwd)

	var server string
	fmt.Print("Enter ocp cluster name (eg: wsa-190 )full server url not required: ")
	fmt.Scanln(&server)

	err := SaveCredentials(user, pwd, server)
	if err != nil {
		return err
	}
	fmt.Printf("Saved oc login creds, use 'kbox login %s' to oc login in future", server)
	return nil

}

func extractSubdomain(rawURL string) (string, error) {
	rawURL = strings.TrimSpace(rawURL)
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}

	host := parsedURL.Hostname()
	parts := strings.Split(host, ".")
	if len(parts) < 4 {
		return "", fmt.Errorf("unexpected URL format")
	}

	// Assuming the format is api.xxxxx.cp.fyre.ibm.com
	return parts[1], nil
}

func IsLoggedIn() bool {
	cmd := exec.Command("oc", "whoami")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return false
	}
	trimmed := strings.TrimSpace(string(output))
	return trimmed != ""
}
