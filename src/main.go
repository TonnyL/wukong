package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {

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
