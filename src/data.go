package main

import (
	"fmt"
)

type Repository struct {
	// google
	Author string `json:"author"`
	// gvisor
	Name string `json:"name"`
	// https://github.com/google.png
	Avatar string `json:"avatar"`
	// https://github.com/google/gvisor
	Url string `json:"url"`
	// Container Runtime Sandbox
	Description string `json:"description"`
	// Go
	Language string `json:"language"`
	// #3572A5
	LanguageColor string `json:"languageColor"`
	// 3320
	Stars int32 `json:"stars"`
	// 118
	Forks int32 `json:"forks"`
	// 1624
	CurrentPeriodStars int32     `json:"currentPeriodStars"`
	BuiltBy            []BuiltBy `json:"builtBy"`
}

type BuiltBy struct {
	// https://github.com/viatsko
	Href string `json:"href"`
	// https://avatars0.githubusercontent.com/u/376065
	Avatar string `json:"avatar"`
	// viatsko
	Username string `json:"username"`
}

type Developer struct {
	// google
	Username string `json:"username"`
	// Google
	Name string `json:"name"`
	// organization
	// could be `organization` or `user`.
	Type string `json:"type"`
	// https://github.com/google
	Url string `json:"url"`
	// https://avatars0.githubusercontent.com/u/1342004
	Avatar string `json:"avatar"`
	Repo   Repo   `json:"repo"`
}

type Repo struct {
	// traceur-compiler
	Name string `json:"name"`
	// Traceur is a JavaScript.next-to-JavaScript-of-today compiler
	Description string `json:"description"`
	// https://github.com/google/traceur-compiler
	Url string `json:"url"`
}

type Language struct {
	UrlParam string `json:"urlParam"`
	Name     string `json:"name"`
}

func (r Repository) DisplayLanguage() string {
	if r.Language == "" {
		return "Unknown"
	}

	return r.Language
}

func (r Repository) DisplayDescription() string {
	if r.Description == "" {
		return "Description not provided"
	}

	return LimitStringWithBreakLines(r.Description, 24)
}

func (d Developer) DisplayFullName() string {
	if d.Name != "" {
		return fmt.Sprintf("%s(%s)", d.Name, d.Username)
	}

	return d.Username
}

func (d Developer) DisplayRepoInfo() string {
	if d.Repo.Description == "" {
		d.Repo.Description = "Description not provided"
	}

	return LimitStringWithBreakLines(fmt.Sprintf("%s - %s", d.Repo.Name, d.Repo.Description), 24)
}

func LimitStringWithBreakLines(s string, perLength int) string {
	result, desc, index := "", []rune(s), 0
	for _, s := range desc {
		if index >= perLength {
			index = 0
			result += "\n"
		}

		result += string(s)
		index += len(string(s))
	}

	return result
}
