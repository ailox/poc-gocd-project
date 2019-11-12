# poc-gocd-project
gocd configuration-as-code demo

## Typical created pipelines per project

 - [pipeline]-build
 - [pipeline]-deploy-dev
 - [pipeline]-deploy-demo
 - [pipeline]-deploy-prod

## Templated values

For build pipelines:
 - `Pipeline.Name` 
 - `Pipeline.Type` [dockerfile, maven]
 - `Pipeline.Git.URL` 
 - `Pipeline.Git.Branch` 

