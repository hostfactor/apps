package main

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v3"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

var manifestDir fs.FS

func GenerateGithubActions(manifestFile, targetFolder string) error {
	dir, manifestFilename := filepath.Split(manifestFile)
	if manifestDir == nil {
		manifestDir = os.DirFS(filepath.Dir(dir))
	}

	b, err := fs.ReadFile(manifestDir, manifestFilename)
	if err != nil {
		return err
	}

	mani := new(Manifest)
	err = yaml.Unmarshal(b, mani)
	if err != nil {
		return errors.Join(fmt.Errorf("failed to read manifest %s", manifestFile), err)
	}

	for _, v := range mani.Features {
		act, err := FeatureToGithubAction(mani.Name, manifestFile, &v)
		if err != nil {
			return errors.Join(fmt.Errorf("failed to generate action for %s", v.Name), err)
		}

		fileName := fmt.Sprintf("%s_%s.yml", mani.Name, v.Name)
		fp := filepath.Join(targetFolder, fileName)
		err = writeGithubActionFile(fp, &v, act)
		if err != nil {
			return err
		}
	}

	return nil
}

func writeGithubActionFile(fp string, feat *Feature, act *GithubAction) error {
	f, err := os.OpenFile(fp, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return errors.Join(fmt.Errorf("failed to write file %s", fp), err)
	}
	defer f.Close()
	enc := yaml.NewEncoder(f)
	enc.SetIndent(2)
	fmt.Println("Writing Github action for feature", feat.Name, "to", fp)
	err = enc.Encode(act)
	if err != nil {
		return errors.Join(fmt.Errorf("failed to write file %s", fp), err)
	}
	return nil
}

type Manifest struct {
	Name        string    `yaml:"name,omitempty" json:"name,omitempty"`
	Author      string    `yaml:"author,omitempty" json:"author,omitempty"`
	Description string    `yaml:"description,omitempty" json:"description,omitempty"`
	Features    []Feature `yaml:"features,omitempty" json:"features,omitempty"`
}

type Feature struct {
	Name      string            `yaml:"name,omitempty" json:"name,omitempty"`
	Tag       string            `yaml:"tag,omitempty" json:"tag,omitempty"`
	Context   string            `yaml:"context,omitempty" json:"context,omitempty"`
	Watch     *FeatureWatch     `yaml:"watch,omitempty" json:"watch,omitempty"`
	BuildArgs map[string]string `yaml:"build_args,omitempty" json:"build_args,omitempty"`
}

type FeatureWatch struct {
	GithubRelease *GithubReleaseFeatureWatch `yaml:"github_release,omitempty" json:"github_release,omitempty"`
}

type GithubReleaseFeatureWatch struct {
	Repo string `yaml:"repo,omitempty" json:"repo,omitempty"`
}

func FeatureToGithubAction(appName, manifestFilePath string, f *Feature) (*GithubAction, error) {
	dir := filepath.Dir(manifestFilePath)
	jobName := fmt.Sprintf("build_and_deploy_%s_%s", appName, f.Name)
	act := &GithubAction{
		Name: f.Name,
		On: &GithubActionTrigger{
			Push: &GithubActionPush{
				Branches: []string{
					"feat/mani_build",
				},
				Paths: []string{
					filepath.Join(dir, f.Context, "**"),
					manifestFilePath,
				},
			},
		},
		Jobs: map[string]GithubActionJob{
			jobName: {
				Uses: "./.github/workflows/deploy_app.yml",
				With: map[string]string{
					"name":    appName,
					"file":    "Dockerfile",
					"context": filepath.Join(dir, f.Context),
					"tags":    f.Tag,
				},
			},
		},
		Env: map[string]string{},
	}

	buildArgs := make([]string, 0, len(f.BuildArgs))
	for k, v := range f.BuildArgs {
		buildArgs = append(buildArgs, fmt.Sprintf("%s=%s", k, v))
	}
	if len(buildArgs) > 0 {
		act.Jobs[jobName].With["build_args"] = strings.Join(buildArgs, "\n")
	}

	if f.Watch != nil {
		act.On.Schedule = []GithubActionScheduleCron{
			{
				Cron: "0 */12 * * *",
			},
		}
		if f.Watch.GithubRelease != nil {
			act.Jobs[jobName].With["get_github_release_tag_name"] = f.Watch.GithubRelease.Repo
		}
	}

	return act, nil
}
