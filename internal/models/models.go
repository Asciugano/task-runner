package models

import "time"

type Task struct {
	Name      string   `yaml:"name"`
	Command   string   `yaml:"command"`
	DependsOn []string `yaml:"depents_on"`
}

type Config struct {
	Tasks []Task `yaml:"tasks"`
}

type CLIOptions struct {
	TaskName string

	ConfigPath      string
	Concurrency     int
	Parallel        bool
	ContinueOnError bool

	Verbose    bool
	DryRun     bool
	OutputFile string

	GlobalTimeout time.Duration

	Vars map[string]string

	List    bool
	Graph   bool
	Version bool
}
