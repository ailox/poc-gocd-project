package main

import (
	"log"
	"os"
	"path"
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

	for _, pipeline := range pipelines {
		os.Mkdir(path.Join("services", pipeline.Name), 0755)
		filename := path.Join("services", pipeline.Name, ".gocd.yaml")
		f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		tmpl.ExecuteTemplate(f, "build.yaml.tmpl", map[string]interface{}{
			"Pipelines": []*Pipeline{pipeline},
			"Envs":      environments,
		})
		f.Close()
	}

}
