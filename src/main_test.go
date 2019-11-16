package main

import (
	"bytes"
	"log"
	"os"
	"strings"
	"testing"
)

func TestShowTableOfLanguages(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stderr)
	}()

	lang := Language{
		UrlParam: "go",
		Name:     "Go",
	}

	ShowTableOfLanguages([]Language{lang})

	s := buf.String()

	if !strings.Contains(s, "ID") || !strings.Contains(s, "Name") || !strings.Contains(s, "go") || !strings.Contains(s, "Go") {
		t.Error("failed")
	} else {
		t.Log("pass")
	}
}

func TestShowTableOfRepositories(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stderr)
	}()

	repo := Repository{
		Author:             "golang",
		Name:               "Go",
		Stars:              1000,
		CurrentPeriodStars: 256,
	}

	ShowTableOfRepositories([]Repository{repo})

	s := buf.String()

	if !strings.Contains(s, "golang") || !strings.Contains(s, "Go") || !strings.Contains(s, "1000") || !strings.Contains(s, "256") {
		t.Error("failed")
	} else {
		t.Log("pass")
	}
}

func TestShowTableOfDevelopers(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stderr)
	}()

	dev := Developer{
		Username: "google",
		Name:     "Google",
		Type:     "organization",
	}

	ShowTableOfDevelopers([]Developer{dev})

	s := buf.String()

	if !strings.Contains(s, "google") || !strings.Contains(s, "Google") || !strings.Contains(s, "organization") {
		t.Error("failed")
	} else {
		t.Log("pass")
	}
}
