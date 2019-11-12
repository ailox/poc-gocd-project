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

	niomonAuth := &Pipeline{
		Type:            "maven",
		Name:            "niomon-auth",
		Group:           "wip-paul",
		DockerImage:     "ubirch/niomon-auth",
		NamespacePrefix: "ubirch",
		Git: GitRepo{
			URL:    "git@github.com:ubirch/niomon-auth.git",
			Branch: "master",
		},
		HelmGit: GitRepo{
			URL:    "git@github.com:ubirch/niomon-common-files.git",
			Branch: "master",
		},
		HelmChartPath: "helm-charts/auth/",
	}

	niomonDecoder := &Pipeline{
		Type:            "maven",
		Name:            "niomon-decoder",
		Group:           "wip-paul",
		DockerImage:     "ubirch/niomon-decoder",
		NamespacePrefix: "ubirch",
		Git: GitRepo{
			URL:    "git@github.com:ubirch/niomon-decoder.git",
			Branch: "master",
		},
		HelmGit: GitRepo{
			URL:    "git@github.com:ubirch/niomon-common-files.git",
			Branch: "master",
		},
		HelmChartPath: "helm-charts/decoder/",
	}

	niomonEnricher := &Pipeline{
		Type:            "maven",
		Name:            "niomon-enricher",
		Group:           "wip-paul",
		DockerImage:     "ubirch/niomon-enricher",
		NamespacePrefix: "ubirch",
		Git: GitRepo{
			URL:    "git@github.com:ubirch/niomon-enricher.git",
			Branch: "master",
		},
		HelmGit: GitRepo{
			URL:    "git@github.com:ubirch/niomon-common-files.git",
			Branch: "master",
		},
		HelmChartPath: "helm-charts/enricher/",
	}

	niomonHTTP := &Pipeline{
		Type:            "maven",
		Name:            "niomon-http",
		Group:           "wip-paul",
		DockerImage:     "ubirch/niomon-http",
		NamespacePrefix: "ubirch",
		Git: GitRepo{
			URL:    "git@github.com:ubirch/niomon-http.git",
			Branch: "master",
		},
		HelmGit: GitRepo{
			URL:    "git@github.com:ubirch/niomon-common-files.git",
			Branch: "master",
		},
		HelmChartPath: "helm-charts/http/",
	}

	niomonResponder := &Pipeline{
		Type:            "maven",
		Name:            "niomon-responder",
		Group:           "wip-paul",
		DockerImage:     "ubirch/niomon-responder",
		NamespacePrefix: "ubirch",
		Git: GitRepo{
			URL:    "git@github.com:ubirch/niomon-responder.git",
			Branch: "master",
		},
		HelmGit: GitRepo{
			URL:    "git@github.com:ubirch/niomon-common-files.git",
			Branch: "master",
		},
		HelmChartPath: "helm-charts/responder/",
	}

	niomonSigner := &Pipeline{
		Type:            "maven",
		Name:            "niomon-signer",
		Group:           "wip-paul",
		DockerImage:     "ubirch/niomon-signer",
		NamespacePrefix: "ubirch",
		Git: GitRepo{
			URL:    "git@github.com:ubirch/niomon-signer.git",
			Branch: "master",
		},
		HelmGit: GitRepo{
			URL:    "git@github.com:ubirch/niomon-common-files.git",
			Branch: "master",
		},
		HelmChartPath: "helm-charts/signer/",
	}

	niomonVerifier := &Pipeline{
		Type:            "maven",
		Name:            "niomon-verifier",
		Group:           "wip-paul",
		DockerImage:     "ubirch/niomon-verifier",
		NamespacePrefix: "ubirch",
		Git: GitRepo{
			URL:    "git@github.com:ubirch/niomon-verifier.git",
			Branch: "master",
		},
		HelmGit: GitRepo{
			URL:    "git@github.com:ubirch/niomon-common-files.git",
			Branch: "master",
		},
		HelmChartPath: "helm-charts/verifier/",
	}

	pipelines := []*Pipeline{
		webui,
		webuiRest,
		vizEnabler,
		discoveryServiceKafka,
		niomonAuth,
		niomonDecoder,
		niomonEnricher,
		niomonHTTP,
		niomonResponder,
		niomonSigner,
		niomonVerifier,
	}

	tmpl.ExecuteTemplate(os.Stdout, "build.yaml.tmpl", map[string]interface{}{
		"Pipelines": pipelines,
		"Envs":      environments,
	})
}
