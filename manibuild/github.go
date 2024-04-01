package manibuild

import (
	"fmt"
	"io"
	"manibuild/gen/go/github"
	"net/http"
	"strings"
)

var CheckoutGithubAction = &github.GithubActionJobStep{
	Uses: "actions/checkout@v4",
	Name: "Checkout",
}

var CompileManibuild = &github.GithubActionJobStep{
	Uses: "./.github/actions/compile_manibuild",
	Name: "Compile manibuild",
}

type BuildImageGithubActionSpec struct {
	Name             string
	File             string
	Context          string
	Tags             []string
	ImageDescription string
	BuildArgs        map[string]string
}

func BuildImageGithubAction(spec BuildImageGithubActionSpec) *github.GithubActionJobStep {
	out := &github.GithubActionJobStep{
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

func GetLatestGitHubRelease(repo string) (*github.GetLatestReleaseResponse, error) {
	resp, err := http.Get(fmt.Sprintf("https://api.github.com/repos/%s/releases/latest", repo))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	out := new(github.GetLatestReleaseResponse)
	err = Unmarshaller.Unmarshal(b, out)
	return out, err
}
