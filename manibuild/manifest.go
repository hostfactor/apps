package main

import (
  "errors"
  "fmt"
  "github.com/goccy/go-yaml"
  "google.golang.org/protobuf/encoding/protojson"
  "io/fs"
  "manibuild/gen/go/manifest"
  "os"
  "path/filepath"
)

var manifestDir fs.FS

var (
  Unmarshaller = protojson.UnmarshalOptions{
    AllowPartial: true,
  }
  Marshaller = protojson.MarshalOptions{
    Multiline:    true,
    Indent:       "  ",
    AllowPartial: true,
  }
)

func LoadAppManifest(f fs.FS, name string) (*manifest.AppManifest, error) {
  b, err := fs.ReadFile(f, name)
  if err != nil {
    return nil, err
  }

  b, err = yaml.YAMLToJSON(b)
  if err != nil {
    return nil, fmt.Errorf("failed to convert from YAML to JSON: %w", err)
  }

  mani := new(manifest.AppManifest)
  err = Unmarshaller.Unmarshal(b, mani)
  if err != nil {
    return nil, errors.Join(fmt.Errorf("failed to read manifest %s", name), err)
  }

  return mani, nil
}

func GenerateGithubActions(manifestFile, targetFolder string) error {
  dir, manifestFilename := filepath.Split(manifestFile)
  if manifestDir == nil {
    manifestDir = os.DirFS(dir)
  }

  mani, err := LoadAppManifest(manifestDir, manifestFilename)
  if err != nil {
    return errors.Join(fmt.Errorf("failed to read manifest %s", manifestFile), err)
  }

  for _, v := range mani.GetBuilds() {
    act, err := FeatureToGithubAction(mani.Name, manifestFile, v)
    if err != nil {
      return errors.Join(fmt.Errorf("failed to generate action for %s", v.Name), err)
    }

    fileName := fmt.Sprintf("%s_%s.yml", mani.Name, v.Name)
    fp := filepath.Join(targetFolder, fileName)
    err = writeGithubActionFile(fp, v, act)
    if err != nil {
      return err
    }
  }

  return nil
}

func writeGithubActionFile(fp string, build *manifest.Build, act *GithubAction) error {
  f, err := os.OpenFile(fp, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
  if err != nil {
    return errors.Join(fmt.Errorf("failed to write file %s", fp), err)
  }
  defer f.Close()

  fmt.Println("Writing Github action for feature", build.Name, "to", fp)
  b, err := yaml.MarshalWithOptions(act, yaml.Indent(2), yaml.IndentSequence(true))
  if err != nil {
    return errors.Join(fmt.Errorf("failed to write file %s", fp), err)
  }

  err = os.WriteFile(fp, b, 0666)
  if err != nil {
    return errors.Join(fmt.Errorf("failed to write file %s", fp), err)
  }

  return nil
}

func FeatureToGithubAction(appName, manifestFilePath string, b *manifest.Build) (*GithubAction, error) {
  dir := filepath.Dir(manifestFilePath)
  jobName := fmt.Sprintf("build_and_deploy_%s_%s", appName, b.Name)

  job := GithubActionJob{
    Name: fmt.Sprintf("Build and deploy %s: %s", appName, b.Name),
    Steps: []GithubActionJobStep{
      CheckoutGithubAction,
    },
    RunsOn: "ubuntu-latest",
  }

  act := &GithubAction{
    Name: fmt.Sprintf("%s-%s", appName, b.Name),
    On: &GithubActionTrigger{
      WorkflowDispatch: &GithubActionWorkflowDispatch{},
      Push: &GithubActionPush{
        Branches: []string{
          "master",
        },
        Paths: []string{
          filepath.Join(dir, b.GetSpec().GetContext(), "**"),
          filepath.Join(dir, "app.sha"),
        },
      },
    },
    Jobs: map[string]GithubActionJob{},
  }

  buildSpec := BuildImageGithubActionSpec{
    Name:             appName,
    Tags:             b.Spec.Tags,
    ImageDescription: b.Description,
    BuildArgs:        b.Spec.BuildArgs,
  }
  if b.Spec.Context != nil {
    buildSpec.Context = filepath.Join(dir, b.GetSpec().GetContext())
  } else {
    buildSpec.Context = dir
  }
  buildImageStep := BuildImageGithubAction(buildSpec)

  if b.Trigger != nil {
    act.On.Schedule = []GithubActionScheduleCron{
      {
        Cron: "0 */12 * * *",
      },
    }
    if b.Trigger.GithubRelease != nil {
      st, _ := GetLatestGithubReleaseAction(b.Trigger.GithubRelease.Repo)
      job.Steps = append(job.Steps, st)
    }
  }

  job.Steps = append(job.Steps, buildImageStep)

  act.Jobs[jobName] = job

  return act, nil
}
