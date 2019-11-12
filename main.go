package main

import (
	"os"
	"strings"
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
	funcMap := template.FuncMap{
		"ToUpper": strings.ToUpper,
		"ToLower": strings.ToLower,
		"Title":   strings.Title,
	}

	tmpl := template.Must(template.New("yaml").Funcs(funcMap).ParseGlob("templates/*"))

	environments := []string{"dev", "demo", "prod"}

	webui := &Pipeline{
		Type:            "dockerfile",
		Name:            "web-ui",
		Group:           "wip-paul",
		DockerImage:     "ubirch/ubirch-web-ui",
		NamespacePrefix: "ubirch",
		Git: GitRepo{
			URL:    "git@github.com:ubirch/ubirch-web-ui.git",
			Branch: "dev",
		},
		HelmGit: GitRepo{
			URL:    "git@github.com:ubirch/ubirch-web-ui.git",
			Branch: "dev",
		},
		HelmChartPath: "helm-charts/web-ui/",
	}
	webuiRest := &Pipeline{
		Type:            "maven",
		Name:            "web-ui-rest",
		Group:           "wip-paul",
		DockerImage:     "ubirch/web-admin-api-server",
		NamespacePrefix: "ubirch",
		Git: GitRepo{
			URL:    "git@github.com:ubirch/ubirch-web-ui-rest.git",
			Branch: "master",
		},
		HelmGit: GitRepo{
			URL:    "git@github.com:ubirch/ubirch-web-ui-rest.git",
			Branch: "master",
		},
		HelmChartPath: "helm-charts/webui-api/",
	}
	vizEnabler := &Pipeline{
		Type:            "maven",
		Name:            "viz-enabler",
		Group:           "wip-paul",
		DockerImage:     "ubirch/viz-enabler-server",
		NamespacePrefix: "ubirch",
		Git: GitRepo{
			URL:    "git@github.com:ubirch/viz-enabler.git",
			Branch: "master",
		},
		HelmGit: GitRepo{
			URL:    "git@github.com:ubirch/viz-enabler.git",
			Branch: "master",
		},
		HelmChartPath: "helm-charts/viz-enabler/",
	}
	discoveryServiceKafka := &Pipeline{
		Type:            "maven",
		Name:            "discovery-service-kafka",
		Group:           "wip-paul",
		DockerImage:     "ubirch/discovery-service-kafka",
		NamespacePrefix: "ubirch",
		Git: GitRepo{
			URL:    "git@github.com:ubirch/ubirch-discovery-service.git",
			Branch: "master",
		},
		HelmGit: GitRepo{
			URL:    "git@github.com:ubirch/ubirch-discovery-service.git",
			Branch: "master",
		},
		HelmChartPath: "helm-charts/discovery-service-kafka/",
	}

	pipelines := []*Pipeline{webui, webuiRest, vizEnabler, discoveryServiceKafka}

	tmpl.ExecuteTemplate(os.Stdout, "build.yaml.tmpl", map[string]interface{}{
		"Pipelines": pipelines,
		"Envs":      environments,
	})
}
