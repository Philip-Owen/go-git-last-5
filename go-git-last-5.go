package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// Repo Struct to map repositories from JSON response
type Repo struct {
	RepoName string `json:"name"`
	Language string `json:"language"`
	URL      string `json:"html_url"`
}

func main() {

	if len(os.Args) <= 1 {
		fmt.Println("You must specify a Github username...")
		os.Exit(1)
	}

	user := os.Args[1]
	userURL := fmt.Sprintf("https://api.github.com/users/%s/repos?sort=updated&per_page=5", user)

	res, err := http.Get(userURL)
	if err != nil {
		fmt.Println(err.Error())
	}

	if res.StatusCode != 200 {
		fmt.Printf("User '%s' not found...\n", user)
		os.Exit(1)
	}

	resData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err.Error())
	}

	var RepoList []Repo
	json.Unmarshal(resData, &RepoList)

	fmt.Println("******************************************************************")
	fmt.Println("Last 5 updated repositories for", user)
	fmt.Println("******************************************************************")
	for i := 0; i < len(RepoList); i++ {
		fmt.Println("* Repo:", RepoList[i].RepoName)
		fmt.Println("  - Language:", RepoList[i].Language)
		fmt.Println("  - URL:", RepoList[i].URL)
	}
}
