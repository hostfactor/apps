package manibuild

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"manibuild/gen/go/manifest"
	"os"
	"path/filepath"
	"slices"
)

func NewApp() *cli.App {
	a := &cli.App{
		Name:  "manibuild",
		Usage: "Bringing app manifests to life.",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:      "manifest",
				Usage:     "A path to the app manifest that is being referenced.",
				Required:  true,
				Aliases:   []string{"m"},
				EnvVars:   []string{"APP_MANIFEST"},
				TakesFile: true,
			},
		},
		Commands: []*cli.Command{
			{
				Name:    "generate",
				Aliases: []string{"gen"},
				Usage:   "Generate files for apps.",
				Args:    true,
				Subcommands: []*cli.Command{
					{
						Name:    "action",
						Aliases: []string{"act", "a"},
						Usage:   "Generate GitHub action workflows.",
						Args:    true,
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:   "output",
								Usage:  "Directory to output the workflow files.",
								Value:  "./.github/workflows",
								Action: validateOutputDir,
							},
						},
						Action: GenerateActionWorkflowsCmd,
					},
				},
			},
			{
				Name:    "get",
				Aliases: []string{"g"},
				Usage:   "Get a resource for an app.",
				Subcommands: []*cli.Command{
					{
						Name:      "version",
						Aliases:   []string{"v", "ver"},
						Usage:     `Get the latest version for a build of an app. If not found, "latest" is returned.`,
						ArgsUsage: "[build name]",
						Args:      true,
						Action:    GetLatestVersionCmd,
					},
				},
			},
			{
				Name:    "sync",
				Aliases: []string{"s"},
				Usage:   "Sync external state for an app.",
				Args:    true,
				Subcommands: []*cli.Command{
					{
						Name:    "trigger",
						Aliases: []string{"t"},
						Usage:   `Checks the triggers for a particular app and updates the outputs.`,
						Args:    true,
						Action:  SyncTriggerCmd,
					},
				},
			},
		},
	}

	return a
}

func GenerateActionWorkflowsCmd(c *cli.Context) error {
	manifestPath := c.String("manifest")
	fmt.Println("Building", manifestPath)
	err := GenerateGithubActions(manifestPath, c.String("output"))
	if err != nil {
		return err
	}
	return nil
}

func GetLatestVersionCmd(c *cli.Context) error {
	manifestPath := c.String("manifest")
	build := c.Args().Get(0)

	mani, err := loadMani(manifestPath)
	if err != nil {
		return err
	}

	idx := slices.IndexFunc(mani.GetBuilds(), func(b *manifest.Build) bool {
		return b.GetName() == build
	})
	if idx < 0 {
		return fmt.Errorf("no build called %s for app %s", build, mani.GetName())
	}

	out := mani.Builds[idx].GetTrigger().GetOutputs()
	if out.GetLatestTag() != "" {
		fmt.Println(out.GetLatestTag())
	} else if out.GetSteamNewsVersion() != "" {
		fmt.Println(out.GetSteamNewsVersion())
	}

	return nil
}

func SyncTriggerCmd(c *cli.Context) error {
	manifestPath := c.String("manifest")
	fmt.Println("Syncing triggers for ", manifestPath)
	return SyncTrigger(manifestPath)
}

func validateOutputDir(c *cli.Context, s string) error {
	st, err := os.Stat(s)
	if err != nil {
		err = os.MkdirAll(s, 0666)
	}
	if st != nil {
		if !st.IsDir() {
			return fmt.Errorf("%s must be a directory", s)
		}
	}
	if err != nil {
		return err
	}

	if !filepath.IsAbs(s) {
		wd, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("failed to get your current directory: %w", err)
		}
		_ = c.Set("output", filepath.Join(wd, s))
	}

	return nil
}
