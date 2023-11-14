package main

type GithubAction struct {
	Name string                     `json:"name,omitempty" yaml:"name,omitempty"`
	On   *GithubActionTrigger       `json:"on,omitempty" yaml:"on,omitempty"`
	Jobs map[string]GithubActionJob `json:"jobs,omitempty" yaml:"jobs,omitempty"`
	Env  map[string]string          `json:"env,omitempty" yaml:"env,omitempty"`
}

type GithubActionJob struct {
	Uses   string                `json:"uses,omitempty" yaml:"uses,omitempty"`
	With   map[string]string     `json:"with,omitempty" yaml:"with,omitempty"`
	RunsOn string                `json:"runs-on,omitempty" yaml:"runs-on,omitempty"`
	Steps  []GithubActionJobStep `json:"steps,omitempty" yaml:"steps,omitempty"`
}

type GithubActionJobStep struct {
	Run string            `json:"run,omitempty" yaml:"run,omitempty"`
	Env map[string]string `json:"env,omitempty" yaml:"env,omitempty"`
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
