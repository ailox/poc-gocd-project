package main

import (
	"os"
	"text/template"
)

type GitRepo struct {
	URL    string
	Branch string
}

type Pipeline struct {
	Type            string
	Name            string
	Group           string
	DockerImage     string
	NamespacePrefix string
	Git             GitRepo
	HelmGit         GitRepo
	HelmChartPath   string
}

func main() {
	tmpl := template.Must(template.New("yaml").ParseGlob("templates/*"))

	pipeline := &Pipeline{
		Type:            "dockerfile",
		Name:            "web-ui",
		Group:           "wip-paul",
		DockerImage:     "ubirch/ubirch-web-ui",
		NamespacePrefix: "ubirch",
		Git: GitRepo{
			URL:    "git@github.com:ubirch/ubirch-web-ui.git",
			Branch: "master",
		},
		HelmGit: GitRepo{
			URL:    "git@github.com:ubirch/wip-helmcharts.git",
			Branch: "develop",
		},
		HelmChartPath: "web-ui/",
	}

	tmpl.ExecuteTemplate(os.Stdout, "build.yaml.tmpl", map[string]interface{}{"Pipeline": pipeline})
}
