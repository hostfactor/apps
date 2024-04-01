package manibuild

import (
	"errors"
	"fmt"
	"google.golang.org/protobuf/encoding/protojson"
	"io/fs"
	"manibuild/gen/go/github"
	"manibuild/gen/go/manifest"
	"os"
	"path/filepath"
)

var manifestDir fs.FS

var (
	Unmarshaller = &protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: true,
	}
	Marshaller = &protojson.MarshalOptions{
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

	mani := new(manifest.AppManifest)
	err = UnmarshalYaml(b, mani, Unmarshaller)
	if err != nil {
		return nil, errors.Join(fmt.Errorf("failed to read manifest %s", name), err)
	}

	return mani, nil
}

func GenerateGithubActions(manifestFile, targetFolder string) error {
	mani, err := loadMani(manifestFile)
	if err != nil {
		return errors.Join(fmt.Errorf("failed to read manifest %s", manifestFile), err)
	}

	for _, v := range mani.GetBuilds() {
		act, err := BuildToGithubAction(mani.Name, manifestFile, v)
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

func writeGithubActionFile(fp string, build *manifest.Build, act *github.GithubAction) error {
	f, err := os.OpenFile(fp, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return errors.Join(fmt.Errorf("failed to write file %s", fp), err)
	}
	defer f.Close()

	fmt.Println("Writing Github action for feature", build.Name, "to", fp)
	b, err := MarshalYaml(act, Marshaller)
	if err != nil {
		return errors.Join(fmt.Errorf("failed to write file %s", fp), err)
	}

	err = os.WriteFile(fp, b, 0666)
	if err != nil {
		return errors.Join(fmt.Errorf("failed to write file %s", fp), err)
	}

	return nil
}

func BuildToGithubAction(appName, manifestFilePath string, b *manifest.Build) (*github.GithubAction, error) {
	dir := filepath.Dir(manifestFilePath)
	jobName := fmt.Sprintf("build_and_deploy_%s_%s", appName, b.Name)

	job := &github.GithubActionJob{
		Name: fmt.Sprintf("Build and deploy %s: %s", appName, b.Name),
		Steps: []*github.GithubActionJobStep{
			CheckoutGithubAction,
			CompileManibuild,
		},
		RunsOn: "ubuntu-latest",
	}

	act := &github.GithubAction{
		Name: fmt.Sprintf("%s-%s", appName, b.Name),
		On: &github.GithubActionTrigger{
			WorkflowDispatch: &github.GithubActionWorkflowDispatch{},
			Push: &github.GithubActionPush{
				Branches: []string{
					"master",
				},
				Paths: []string{
					filepath.Join(dir, b.GetSpec().GetContext(), "**"),
					filepath.Join(dir, "app.sha"),
				},
			},
		},
		Jobs: map[string]*github.GithubActionJob{},
		Env:  map[string]string{},
	}

	if v := b.GetTrigger().GetOutputs().GetLatestTag(); v != "" {
		act.Env["LATEST_TAG_NAME"] = v
	}

	if v := b.GetTrigger().GetOutputs().GetSteamNewsVersion(); v != "" {
		act.Env["LATEST_STEAM_VERSION"] = v
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

	job.Steps = append(job.Steps, buildImageStep)

	act.Jobs[jobName] = job

	return act, nil
}

func loadMani(manifestFile string) (*manifest.AppManifest, error) {
	dir, manifestFilename := filepath.Split(manifestFile)
	if manifestDir == nil {
		manifestDir = os.DirFS(dir)
	}

	mani, err := LoadAppManifest(manifestDir, manifestFilename)
	if err != nil {
		return nil, errors.Join(fmt.Errorf("failed to read manifest %s", manifestFile), err)
	}
	return mani, nil
}
