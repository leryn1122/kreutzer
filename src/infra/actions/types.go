package actions

type Workflow struct {
	Name        string            `yaml:"name,omitempty"`
	On          *On               `yaml:"on"`
	Env         map[string]string `yaml:"env,omitempty"`
	Environment *Environment      `yaml:"environment,omitempty"`
	Defaults    *Defaults         `yaml:"defaults,omitempty"`
	Concurrency *Concurrency      `yaml:"concurrency,omitempty"`
	Jobs        map[string]Job    `yaml:"jobs"`
	RunName     *RunName          `yaml:"run-name,omitempty"`
	Permissions *Permissions      `yaml:"permissions,omitempty"`
}

type Concurrency struct {
	Group            string `yaml:"group"`
	CancelInProgress bool   `yaml:"cancelInProgress"`
}

type Defaults struct {
	Run Run `yaml:"run"`
}

type Environment struct {
	Name string `yaml:"name"`
	URL  string `yaml:"url,omitempty"`
}

type Job struct {
	Name            string            `yaml:"name,omitempty"`
	Needs           []string          `yaml:"needs,omitempty"`
	Permissions     *Permissions      `yaml:"permissions,omitempty"`
	RunsOn          *RunsOn           `yaml:"runs-on"`
	Environment     *Environment      `yaml:"environment,omitempty"`
	Outputs         string            `yaml:"outputs,omitempty"`
	Env             map[string]string `yaml:"env,omitempty"`
	Defaults        *Defaults         `yaml:"defaults,omitempty"`
	If              string            `yaml:"if,omitempty"`
	Steps           []Step            `yaml:"steps,omitempty"`
	TimeoutMinutes  uint              `yaml:"timeout-minutes,omitempty"`
	Strategy        *Strategy         `yaml:"strategy,omitempty"`
	ContinueOnError bool              `yaml:"continue-on-error,omitempty"`
	Container       string            `yaml:"container,omitempty"`
	Concurrency     *Concurrency      `yaml:"concurrency,omitempty"`
}

type Matrix struct {
}

type On struct {
	Push Push `yaml:"push,omitempty"`
}

type Permissions struct {
}

type Push struct {
	Branches []string `yaml:"branches"`
}

type Run struct {
	Shell            []string `yaml:"shell"`
	WorkingDirectory string   `yaml:"working-directory"`
}

type RunName struct {
}

type RunsOn struct {
}

type Step struct {
}

type Strategy struct {
	Matrix      *Matrix `yaml:"matrix"`
	FailFast    bool    `yaml:"fail-fast,omitempty"`
	MaxParallel uint    `yaml:"max-parallel"`
}
