package main

import "testing"

func TestRepository_DisplayLanguage(t *testing.T) {
	repo := Repository{
		Language: "",
	}

	if len(repo.DisplayLanguage()) == 0 {
		t.Error("failed")
	} else {
		t.Log("pass")
	}

	repo.Language = "Go"

	if repo.DisplayLanguage() != "Go" {
		t.Error("failed")
	} else {
		t.Log("pass")
	}
}

func TestRepository_DisplayDescription(t *testing.T) {
	repo := Repository{
		Description: "",
	}

	if len(repo.DisplayDescription()) == 0 {
		t.Error("failed")
	} else {
		t.Log("pass")
	}

	repo.Description = "The Go programming language"

	if repo.DisplayDescription() != "The Go programming langu\nage" {
		t.Error("failed")
	} else {
		t.Log("pass")
	}
}

func TestDeveloper_DisplayFullName(t *testing.T) {
	dev := Developer{
		Name:     "Li Zhao Tai Lang",
		Username: "tonnyl",
	}

	if dev.DisplayFullName() != "Li Zhao Tai Lang(tonnyl)" {
		t.Error("failed")
	} else {
		t.Log("pass")
	}

	dev.Name = ""
	if dev.DisplayFullName() != "tonnyl" {
		t.Error("failed")
	} else {
		t.Log("pass")
	}
}

func TestDeveloper_DisplayRepoInfo(t *testing.T) {
	dev := Developer{
		Repo: Repo{
			Name:        "Wukong",
			Description: "A command-line tool for browsing GitHub trending repositories and developers written by Go.",
		},
	}

	if dev.DisplayRepoInfo() != "Wukong - A command-line\ntool for browsing GitHub\n trending repositories a\nnd developers written by\n Go." {
		t.Error("failed")
	} else {
		t.Log("pass")
	}

	dev.Repo.Description = ""

	if dev.DisplayRepoInfo() != "Wukong - Description not\n provided" {
		t.Error("failed")
	} else {
		t.Log("pass")
	}
}

func TestLimitStringWithBreakLines(t *testing.T) {
	s := "string"

	if LimitStringWithBreakLines(s, 0) != "s\nt\nr\ni\nn\ng" {
		t.Error("failed")
	} else {
		t.Log("pass")
	}

	if LimitStringWithBreakLines(s, 9) != "string" {
		t.Error("failed")
	} else {
		t.Log("pass")
	}
}
