package main

import (
	"encoding/json"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

func main() {
	//resp, err := FetchTrendingRepositories("kotlin", "daily")
	//if err != nil {
	//	fmt.Println("error: " + err.Error())
	//
	//	return
	//}
	//
	//ShowTable(resp)

	//resp, err := FetchTrendingDevelopers("go", "daily")
	//if err != nil {
	//	fmt.Println("error: " + err.Error())
	//	return
	//}
	//
	//ShowTableOfDevelopers(resp)
}

// Receive an array of trending repositories.
func FetchTrendingRepositories(language, since string) ([]Repository, error) {
	resp, err := http.Get(fmt.Sprintf("https://github-trending-api.now.sh/repositories?language=%s&since=%s", language, since))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	repositories := make([]Repository, 0)
	jsonErr := json.NewDecoder(resp.Body).Decode(&repositories)
	if jsonErr != nil {
		return nil, jsonErr
	}

	return repositories, nil
}

// Receive an array of trending developers.
func FetchTrendingDevelopers(language, since string) ([]Developer, error) {
	resp, err := http.Get(fmt.Sprintf("https://github-trending-api.now.sh/developers?language=%s&since=%s", language, since))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	developers := make([]Developer, 0)
	jsonErr := json.NewDecoder(resp.Body).Decode(&developers)
	if jsonErr != nil {
		return nil, jsonErr
	}

	return developers, nil
}

// Read popular languages and all languages.
func ReadLanguages() ([]Language, error) {
	jsonFile, err := os.Open("./resources/languages.json")
	if err != nil {
		return nil, err
	}

	defer jsonFile.Close()

	bytes, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	languages := make([]Language, 0)
	jsonErr := json.Unmarshal(bytes, &languages)
	if jsonErr != nil {
		return nil, jsonErr
	}

	return languages, nil
}

func ShowTableOfRepositories(repos []Repository) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Rank", "Name", "Description", "Language", "Stars(Total/Period)", "Url"})
	table.SetRowLine(true)

	for index, repo := range repos {
		table.Append([]string{strconv.Itoa(index + 1), repo.Name, repo.DisplayDescription(), repo.DisplayLanguage(), fmt.Sprintf("%d/%d", repo.Stars, repo.CurrentPeriodStars), repo.Url})
	}

	table.Render()
}

func ShowTableOfDevelopers(developers []Developer) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Rank", "Name", "Repo Name/Description", "Url"})
	table.SetRowLine(true)

	for index, dev := range developers {
		table.Append([]string{strconv.Itoa(index + 1), dev.DisplayFullName(), dev.DisplayRepoInfo(), dev.Url})
	}

	table.Render()
}

func ShowTableOfLanguages(languages []Language) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Id", "Name"})
	table.SetRowLine(true)

	for _, lang := range languages {
		table.Append([]string{lang.UrlParam, lang.Name})
	}

	table.Render()
}
