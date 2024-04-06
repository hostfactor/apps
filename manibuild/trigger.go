package manibuild

import (
	"errors"
	"fmt"
	"manibuild/gen/go/manifest"
	"manibuild/steam"
	"os"
	"time"
)

func SyncTrigger(maniPath string) error {
	mani, err := loadMani(maniPath)
	if err != nil {
		return err
	}

	changed, err := RunTriggers(mani)
	if err != nil {
		return err
	}

	if !changed {
		fmt.Println("No changes for", mani.GetName()+".")
		return nil
	}

	_, _ = fmt.Fprintf(os.Stdout, "Detected change for %s. Saving output.", mani.GetName())
	fmt.Println("Detected change for", mani.GetName()+".", " Saving output.")
	b, err := MarshalYaml(mani, Marshaller)
	if err != nil {
		return err
	}

	err = os.WriteFile(maniPath, b, 0666)
	if err != nil {
		return errors.Join(errors.New("failed to save outputs to file"), err)
	}

	return nil
}

func RunTriggers(mani *manifest.AppManifest) (changed bool, err error) {
	for _, build := range mani.GetBuilds() {
		if build.Trigger == nil {
			continue
		}

		if build.Trigger.Outputs == nil {
			build.Trigger.Outputs = &manifest.BuildTrigger_Outputs{}
		}

		if build.GetTrigger().GetGithubRelease() != nil {
			gh := build.GetTrigger().GetGithubRelease()
			resp, err := GetLatestGitHubRelease(gh.GetRepo())
			if err != nil {
				return false, fmt.Errorf("failed to get github release for build %s: %w", build.Name, err)
			}

			build.Trigger.Outputs.LatestTag = &resp.TagName
			build.Trigger.Outputs.ReleaseLink = &resp.HtmlUrl
			build.Trigger.Outputs.CreatedAt = &resp.CreatedAt
			if build.Trigger.Outputs.GetLatestTag() != resp.GetTagName() {
				changed = true
			}
		}

		if build.GetTrigger().GetSteamFeed() != nil {
			ver, err := steam.FetchLatestVersion(build.GetTrigger().GetSteamFeed())
			if err != nil {
				return false, fmt.Errorf("failed to execute steam feed trigger for build %s: %w", build.Name, err)
			}

			if ver.Name != build.GetTrigger().GetOutputs().GetSteamNewsVersion() {
				build.Trigger.Outputs.SteamNewsVersion = &ver.Name
				build.Trigger.Outputs.ReleaseLink = &ver.Link
				changed = true
			}

			if !ver.Date.IsZero() {
				date := ver.Date.Format(time.RFC3339)
				build.Trigger.Outputs.CreatedAt = &date

				if date != build.GetTrigger().GetOutputs().GetCreatedAt() {
					changed = true
				}
			}
		}
	}

	return
}
