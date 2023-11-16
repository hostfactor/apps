package main

import (
  "fmt"
  "strings"
)

type GithubAction struct {
  Name string                     `json:"name,omitempty" yaml:"name,omitempty"`
  On   *GithubActionTrigger       `json:"on,omitempty" yaml:"on,omitempty"`
  Jobs map[string]GithubActionJob `json:"jobs,omitempty" yaml:"jobs,omitempty"`
  Env  map[string]string          `json:"env,omitempty" yaml:"env,omitempty"`
}

type GithubActionJob struct {
  Name   string                `json:"name,omitempty" yaml:"name,omitempty"`
  Uses   string                `json:"uses,omitempty" yaml:"uses,omitempty"`
  With   map[string]string     `json:"with,omitempty" yaml:"with,omitempty"`
  RunsOn string                `json:"runs-on,omitempty" yaml:"runs-on,omitempty"`
  Steps  []GithubActionJobStep `json:"steps,omitempty" yaml:"steps,omitempty"`
}

type GithubActionJobStep struct {
  Uses string            `json:"uses,omitempty" yaml:"uses,omitempty"`
  Run  string            `json:"run,omitempty" yaml:"run,omitempty"`
  Env  map[string]string `json:"env,omitempty" yaml:"env,omitempty"`
  Name string            `json:"name,omitempty" yaml:"name,omitempty"`
  With map[string]string `json:"with,omitempty" yaml:"with,omitempty"`
  Id   string            `json:"id,omitempty" yaml:"id,omitempty"`
}

type GithubActionTrigger struct {
  Schedule []GithubActionScheduleCron `json:"schedule,omitempty" yaml:"schedule,omitempty"`
  Push     *GithubActionPush          `json:"push,omitempty" yaml:"push,omitempty"`
}

type GithubActionPush struct {
  Branches []string `json:"branches,omitempty" yaml:"branches,omitempty"`
  Paths    []string `json:"paths,omitempty" yaml:"paths,omitempty"`
}

type GithubActionScheduleCron struct {
  Cron string `json:"cron,omitempty" yaml:"cron,omitempty"`
}

var CheckoutGithubAction = GithubActionJobStep{
  Uses: "actions/checkout@v4",
  Name: "Checkout",
}

type BuildImageGithubActionSpec struct {
  Name             string
  File             string
  Context          string
  Tags             []string
  ImageDescription string
  BuildArgs        map[string]string
}

func BuildImageGithubAction(spec BuildImageGithubActionSpec) GithubActionJobStep {
  out := GithubActionJobStep{
    Uses: "./.github/actions/build_image",
    Name: "Build and deploy image",
    With: map[string]string{},
  }

  out.With["image_description"] = spec.ImageDescription
  out.With["name"] = spec.Name
  out.With["file"] = spec.File

  if spec.Context != "" {
    out.With["context"] = spec.Context
  }

  if len(spec.Tags) > 0 {
    out.With["tags"] = strings.Join(spec.Tags, "\n")
  }

  if len(spec.BuildArgs) > 0 {
    ba := make([]string, 0, len(spec.BuildArgs))
    for k, v := range spec.BuildArgs {
      ba = append(ba, fmt.Sprintf("%s=%s", k, v))
    }
    out.With["build_args"] = strings.Join(ba, "\n")
  }

  return out
}

type GetLatestGithubReleaseActionOutputs struct {
  TagName string
}

func GetLatestGithubReleaseAction(repo string) (GithubActionJobStep, GetLatestGithubReleaseActionOutputs) {
  out := GithubActionJobStep{
    Uses: "./.github/actions/get_github_release",
    Name: "Get Github release",
    Id:   "github_release",
    With: map[string]string{
      "repo": repo,
    },
  }
  return out, GetLatestGithubReleaseActionOutputs{TagName: fmt.Sprintf("steps.%s.outputs.tag_name", out.Id)}
}
