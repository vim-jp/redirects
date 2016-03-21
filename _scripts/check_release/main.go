package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"sync"
	"text/template"

	"github.com/koron/go-github"

	"gopkg.in/yaml.v2"
)

var pageTmpl = template.Must(template.New("page").Parse(`---
title: {{.Title}}
redirect_to:
  - {{.RedirectURL}}
---
`))

const (
	dataFile = "_data/redirects.yml"
)

type redirect struct {
	Title string `yaml:"title"`
	Path  string `yaml:"path"`

	GithubRelease *githubRelease `yaml:"github_release,omitempty"`
}

type githubRelease struct {
	Owner       string `yaml:"owner"`
	Repo        string `yaml:"repo"`
	NamePattern string `yaml:"name_pattern"`
}

type tmplData struct {
	Title       string
	RedirectURL string
}

func loadData(name string) ([]redirect, error) {
	b, err := ioutil.ReadFile(name)
	if err != nil {
		return nil, err
	}
	var v []redirect
	err = yaml.Unmarshal(b, &v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func fetchRedirect(d redirect) (*github.Asset, error) {
	if d.GithubRelease == nil {
		return nil, nil
	}
	r, err := github.Latest(d.GithubRelease.Owner, d.GithubRelease.Repo)
	if err != nil {
		return nil, err
	}
	rx, err := regexp.Compile(d.GithubRelease.NamePattern)
	if err != nil {
		return nil, err
	}
	for _, v := range r.Assets {
		if rx.MatchString(v.Name) {
			return &v, nil
		}
	}
	return nil, nil
}

func updateRedirect(d redirect, a *github.Asset) error {
	if a.State != "uploaded" {
		return fmt.Errorf("not uploaded yet: %s", d.Path)
	}
	name := d.Path + ".html"
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer f.Close()
	err = pageTmpl.Execute(f, tmplData{
		Title:       d.Title,
		RedirectURL: a.DownloadURL,
	})
	if err != nil {
		return nil
	}
	return nil
}

func processRedirect(d redirect) {
	a, err := fetchRedirect(d)
	if err != nil {
		log.Printf("fetch failed for %s: %s", d.Path, err)
		return
	}
	if a == nil {
		return
	}
	err = updateRedirect(d, a)
	if err != nil {
		log.Printf("update failed for %s: %s", d.Path, err)
		return
	}
}

func main() {
	targets, err := loadData(dataFile)
	if err != nil {
		log.Fatal(err)
	}
	var wg sync.WaitGroup
	for _, t := range targets {
		wg.Add(1)
		go func(d redirect) {
			processRedirect(d)
			wg.Done()
		}(t)
	}
	wg.Wait()
}
