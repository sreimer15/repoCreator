package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	r "./lib"
)

func main() {
	username := getUserName()
	fmt.Printf("%T", username)
	fmt.Printf("%s", username)
	// quickTest(username)

	fmt.Println("The user's username is" + username)
	baseDirectory := getWorkingDirectory()

	repo := r.CreateDefaultRepo(baseDirectory)

	data, err := json.Marshal(repo)
	if err != nil {
		return
	}
	// need to get an auth token, put it on the header and it shoudl work
	jsonstr := string(data)
	fmt.Println(jsonstr)
	// getToken(username)
	msg := createRepo(username, jsonstr)
	fmt.Println(msg)
	url := "https://github.com/" + username + "/" + baseDirectory
	setUpRemoteOrigin(url)
	pushToMaster(url)
	fmt.Println("Everything is done visit your repo at: " + url)

}

func getUserName() string {
	var (
		cmdOut []byte
		err    error
	)
	cmdName := "git"
	cmdArgs := []string{"config", "user.name"}

	if cmdOut, err = exec.Command(cmdName, cmdArgs...).Output(); err != nil {
		fmt.Fprintln(os.Stderr, "There was an error running git config user.name command: ", err)
		os.Exit(1)
	}

	username := string(cmdOut)
	strings.Trim(username, "\n")

	fmt.Printf("%d %q\n", len(username), username)
	t := strings.TrimSpace(username)
	fmt.Printf("%d %q\n", len(t), t)
	return t
}

func getWorkingDirectory() string {

	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	baseDirectory := filepath.Base(pwd)

	return baseDirectory
}

func quickTest(username string) {
	var (
		cmdOut []byte
		err    error
	)
	fmt.Printf("%T", username)
	fmt.Printf("%s", username)

	cmdName := "curl"
	// for some reason adding serimer15 as a variable adds an extra new line that ruins everything
	cmdArgs := []string{"-i", "-u", username, "-d", "{\"scopes\": [\"repo\", \"user\"], \"note\": \"getting-started\"}", "https://api.github.com/authorizations"}
	fmt.Println(cmdArgs)
	if cmdOut, err = exec.Command(cmdName, cmdArgs...).Output(); err != nil {

		fmt.Fprintln(os.Stderr, "There was an error running git remote add origin : ", err)
		os.Exit(1)
	}
	fmt.Println("Succesfully added something origin at ", string(cmdOut))
}

func setUpRemoteOrigin(url string) {
	cmdName := "git"
	cmdArgs := []string{"remote", "add", "origin", url}

	if err := exec.Command(cmdName, cmdArgs...).Run(); err != nil {
		fmt.Fprintln(os.Stderr, "There was an error running git remote add origin : ", err)
		os.Exit(1)
	}
	fmt.Println("Succesfully set up remote origin at " + url)
}

func pushToMaster(url string) {
	cmdName := "git"
	cmdArgs := []string{"push", "origin", "master"}

	if err := exec.Command(cmdName, cmdArgs...).Run(); err != nil {
		fmt.Fprintln(os.Stderr, "There was an error running git push origin master : ", err)
		os.Exit(1)
	}
	fmt.Println("Succesfully pushed to master at " + url)
}

func getToken(username string) string {

	var (
		cmdOut []byte
		err    error
	)
	authURL := "https://api.github.com/authorizations"
	cmdName := "curl"
	cmdArgs := []string{"-i", "-u", username, "-d", "{\"scopes\": [\"repo\", \"user\"], \"note\": \"getting-started\"}", authURL}

	if cmdOut, err = exec.Command(cmdName, cmdArgs...).Output(); err != nil {
		fmt.Fprintln(os.Stderr, "There was an error running curl: ", err)
		os.Exit(1)
	}

	fmt.Println(string(cmdOut))
	// seems like it is json?

	cmdOutType := fmt.Sprintf("%T", cmdOut)
	fmt.Println(cmdOutType)
	return string(cmdOut)

}

func createRepo(username string, jsonstr string) string {
	apiURL := "https://api.github.com/user/repos"
	var (
		cmdOut []byte
		err    error
	)
	cmdName := "curl"
	cmdArgs := []string{"-u", username, apiURL, "-d", jsonstr}
	fmt.Printf("%s", cmdArgs)

	if cmdOut, err = exec.Command(cmdName, cmdArgs...).Output(); err != nil {
		fmt.Fprintln(os.Stderr, "There was an error running curl: ", err)
		os.Exit(1)
	}
	fmt.Println("The token name is " + username)
	fmt.Println("The api url is " + apiURL)
	fmt.Println("The jsonstr is" + jsonstr)

	return string(cmdOut)

}
