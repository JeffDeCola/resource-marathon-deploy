# resource-marathon-deploy

[![Go Report Card](https://goreportcard.com/badge/github.com/JeffDeCola/resource-marathon-deploy)](https://goreportcard.com/report/github.com/JeffDeCola/resource-marathon-deploy)
[![GoDoc](https://godoc.org/github.com/JeffDeCola/resource-marathon-deploy?status.svg)](https://godoc.org/github.com/JeffDeCola/resource-marathon-deploy)
[![Maintainability](https://api.codeclimate.com/v1/badges/ed7ccb9bd7f39ccbfe23/maintainability)](https://codeclimate.com/github/JeffDeCola/resource-marathon-deploy/maintainability)
[![Issue Count](https://codeclimate.com/github/JeffDeCola/resource-marathon-deploy/badges/issue_count.svg)](https://codeclimate.com/github/JeffDeCola/resource-marathon-deploy/issues)
[![License](http://img.shields.io/:license-mit-blue.svg)](http://jeffdecola.mit-license.org)

`resource-marathon-deploy` is a Concourse resource type that deploys an APP to
Marathon using a .json file.

[resource-marathon-deploy Docker Image](https://hub.docker.com/r/jeffdecola/resource-marathon-deploy)
on DockerHub.

[resource-marathon-deploy GitHub Webpage](https://jeffdecola.github.io/resource-marathon-deploy/)

## DEVELOPED

Written in GO, `resource-marathon-deploy` was used as a test to help develop [`marathon-resource`](https://github.com/ckaznocha/marathon-resource).

This resource was built using [_`resource-template`_](https://github.com/JeffDeCola/resource-template).

## SOURCE CONFIGURATION

* `marathonuri`: The uri of marathon.

## BEHAVIOR

### CHECK (a resource version(s))

CHECK will mimic getting the list of versions from a resource.

#### CHECK stdin

```json
{
  "source": {
    "source1": "sourcefoo1",
    "source2": "sourcefoo2"
  },
  "version": {
    "ref": "123 ",
  }
}
```

123 is the current version.

#### CHECK stdout

```json
[
  { "ref": "123" },
  { "ref": "3de" },
  { "ref": "456" }
]
```

456 is the latest version that will be used.

The last number 456 will become the current ref version that will be used by IN.

#### CHECK - go run

```bash
echo '{
"params": {"param1": "Hello Clif","param2": "Nice to meet you"},
"source": {"source1": "sourcefoo1","source2": "sourcefoo2"},
"version":{"ref": "123"}}' |
go run main.go check $PWD
```

### IN (fetch a resource)

IN will mimic fetching a resource and placing a file in the working directory.

#### IN Parameters

* `param1`: Just a placeholder.

* `param2`: Just a placeholder.

#### IN stdin

```json
{
  "source": {
    "source1": "sourcefoo1",
    "source2": "sourcefoo2"
  },
  "params": {
    "param1": "Hello Clif",
    "param2": "Nice to meet you"
  },
  "version": {
    "ref": "456",
  }
```

#### IN stdout

```json
{
  "version":{ "ref": "456" },
  "metadata": [
    { "name": "nameofmonkey", "value": "Larry" },
    { "name": "author","value": "Jeff DeCola" }
  ]
}
```

#### IN file fetched (fetch.json)

The IN will mimic a fetch and place a fake file `fetched.json`
file in the working directory:

#### IN - go run

```bash
echo '{
"params": {"param1": "Hello Clif","param2": "Nice to meet you"},
"source": {"source1": "sourcefoo1","source2": "sourcefoo2"},
"version":{"ref": "777"}}' |
go run main.go in $PWD
```

### OUT (update a resouce)

OUT shall delploy an APP to marathon based on marathon.json file.

Create a marathon .json file.  As an example:

```json
{
    "id": "appname",
    "cpus": 0.1,
    "mem": 16.0,
    "container": {
        "type": "DOCKER",
        "docker": {
            "forcePullImage": true,
            "image": "jeffdecola/hello-go:latest",
            "network": "BRIDGE",
            "portMappings": [{
                "containerPort": 8080,
                "hostPort": 0
            }]
        }
    }
}
```

#### OUT Parameters

* `app_json_path`: the path to your newly created marathon .json file.

#### OUT stdin

```json
{
  "params": {
    "app_json_path": "hello-go/app.json",
  },
  "source": {
    "marathonuri": "http://10.141.141.10:808",
  },
  "versions": {
    "ref": ""
  }
}
```

#### OUT stdout

```json
{
  "version":{ "ref": "(timestamp of when marathon started)" },
  "metadata": [
    { "name": "????????????","value": "????????????" },
  ]
}
```

#### OUT - go run

```bash
echo '{
"params": {"param1": "Hello Jeff","param2": "How are you?"},
"source": {"source1": "sourcefoo1","source2": "sourcefoo2"},
"version":{"ref": ""}}' |
go run main.go out $PWD
```

## PIPELINE EXAMPLE USING PUT

```yaml
jobs:
...
- name: your-job-name
  plan:
    ...
  - put: resource-marathon-deploy
    params: {app_json_path: "hello-go/app.json"}

resource_types:
  ...
- name: marathon-deploy
  type: docker-image
  source:
    marathonuri:http://10.141.141.10:8080
    repository: jeffdecola/resource-marathon-deploy
    tag: latest

resources:
  ...
- name: resource-marathon-deploy
  type: marathon-deploy
  source:
    repository: /username/image-name
    tag: latest
```

GET would look similar.

## TESTED, BUILT & PUSHED TO DOCKERHUB USING CONCOURSE

To automate the creation of the `resource-marathon-deploy` docker image, a concourse pipeline
will,

* Update README.md for resource-marathon-deploy github webpage.
* Unit Test the code.
* Build the docker image `resource-marathon-deploy` and push to DockerHub.

![IMAGE - resource-marathon-deploy concourse ci pipeline - IMAGE](docs/pics/resource-marathon-deploy-pipeline.jpg)

As seen in the pipeline diagram, the _resource-dump-to-dockerhub_ uses
the resource type
[docker-image](https://github.com/concourse/docker-image-resource)
to push a docker image to dockerhub.

`resource-marathon-deploy` also contains a few extra concourse resources:

* A resource (_resource-slack-alert_) uses a [docker image](https://hub.docker.com/r/cfcommunity/slack-notification-resource)
  that will notify slack on your progress.
* A resource (_resource-repo-status_) use a [docker image](https://hub.docker.com/r/dpb587/github-status-resource)
  that will update your git status for that particular commit.
* A resource ([_`resource-template`_](https://github.com/JeffDeCola/resource-template))
  that can be used as a starting point and template for creating other concourse
  ci resources.

For more information on using concourse for continuous integration,
refer to my cheat sheet on [concourse](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/operations-tools/continuous-integration-continuous-deployment/concourse-cheat-sheet).
