package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"github.com/urfave/cli"
	"net/http"
	"os"
	"strconv"
)

func main() {
	lang := ""
	period := "daily"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:     "lang,l",
			Usage:    "language, use list command to see all the available options",
			Required: false,
		},
		cli.StringFlag{
			Name:     "period,p",
			Usage:    "Period, possible values: daily, weekly and monthly",
			Required: false,
		},
	}

	app := cli.NewApp()
	app.Name = "Wukong" // Incredible name!
	app.Usage = "A command-line tool for browsing GitHub trending repositories&developers written by Go."
	app.Version = "0.1.0-alpha01"
	app.HideVersion = true
	app.Copyright = "Wukong is under an MIT license. See the [LICENSE](https://github.com/TonnyL/Wukong/blob/master/LICENSE) for more information."
	app.Commands = []cli.Command{
		{
			Name:        "repo",
			Aliases:     []string{"r", "repositories", "repository"},
			Description: "See the developers that the GitHub community is most excited about.",
			Usage:       "--lang x --period y",
			Flags:       flags,
			Action: func(c *cli.Context) error {
				paramsErr := CheckParams(lang, period)
				if paramsErr != nil {
					return paramsErr
				}

				repos, err := FetchTrendingRepositories(lang, period)
				if err != nil {
					return err
				}

				ShowTableOfRepositories(repos)

				return nil
			},
		},
		{
			Name:        "dev",
			Aliases:     []string{"d", "developers", "developer"},
			Description: "See the repositories that the GitHub community is most excited about.",
			Usage:       "--lang x --period y",
			Flags:       flags,
			Action: func(c *cli.Context) error {
				paramsErr := CheckParams(lang, period)
				if paramsErr != nil {
					return paramsErr
				}

				devs, err := FetchTrendingDevelopers(lang, period)
				if err != nil {
					return err
				}

				ShowTableOfDevelopers(devs)

				return nil
			},
		},
		{
			Name:    "lang",
			Aliases: []string{"l", "languages", "language"},
			Usage:   "List all the available language options",
			Action: func(c *cli.Context) error {
				langs, err := FetchLanguages()
				if err != nil {
					return err
				}

				ShowTableOfLanguages(langs)

				return nil
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		fmt.Print("☹️ command error: " + err.Error())
	}
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

// cache to avoid too many requests.
var cachedLanguages = make([]Language, 0)

// Fetch all languages.
func FetchLanguages() ([]Language, error) {
	if len(cachedLanguages) != 0 {
		return cachedLanguages, nil
	}
	resp, err := http.Get("https://raw.githubusercontent.com/TonnyL/Wukong/master/resources/languages.json")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	jsonErr := json.NewDecoder(resp.Body).Decode(&cachedLanguages)
	if jsonErr != nil {
		return nil, jsonErr
	}

	return cachedLanguages, nil
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

func CheckParams(l, p string) error {
	if p != "daily" && p != "weekly" && p != "monthly" {
		return errors.New("Unknown period value: " + p)
	}

	langs, err := FetchLanguages()
	if err != nil {
		return err
	}

	for _, lang := range langs {
		if lang.UrlParam == l {
			return nil
		}
	}

	return errors.New("Unknown lang value: " + l)
}
