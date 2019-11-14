package main

import (
	"log"
	"os"
	"strings"
	"text/template"

	"gopkg.in/yaml.v2"
)

type GitRepo struct {
	URL    string `yaml:"url"`
	Branch string `yaml:"branch"`
}

type Pipeline struct {
	Name            string   `yaml:"name"`
	Type            string   `yaml:"type"`
	Group           string   `yaml:"group"`
	DockerImage     string   `yaml:"docker_image"`
	NamespacePrefix string   `yaml:"namespace_prefix"`
	Git             *GitRepo `yaml:"git"`
	HelmGit         *GitRepo `yaml:"helm_git"`
	HelmChartPath   string   `yaml:"helm_chart_path"`
}

func main() {
	funcMap := template.FuncMap{
		"ToUpper": strings.ToUpper,
		"ToLower": strings.ToLower,
		"Title":   strings.Title,
	}

	tmpl := template.Must(template.New("yaml").Funcs(funcMap).ParseGlob("templates/*"))

	environments := []string{"dev", "demo", "prod"}

	pipelines := []*Pipeline{}

	f, err := os.OpenFile("services.yaml", os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	yd := yaml.NewDecoder(f)
	err = yd.Decode(&pipelines)
	if err != nil {
		log.Fatal(err)
	}

	tmpl.ExecuteTemplate(os.Stdout, "build.yaml.tmpl", map[string]interface{}{
		"Pipelines": pipelines,
		"Envs":      environments,
	})
}
