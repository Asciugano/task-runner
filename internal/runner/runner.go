package runner

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"gopkg.in/yaml.v3"

	"github.com/Asciugano/taskRunner/internal/models"
)

func Init(options models.CLIOptions) {
	tasks, err := LoadTasks(options.ConfigPath)
	if err != nil {
		fmt.Println("Error parsing the config file: ", err)
		os.Exit(1)
	}

	fmt.Println(tasks)

	for _, t := range tasks.Tasks {
		if options.TaskName == t.Name {
			RunTask(t, options)
		}
	}
}

func LoadTasks(path string) (models.Config, error) {
	var cfg models.Config

	data, err := os.ReadFile(path)
	if err != nil {
		return cfg, err
	}

	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return cfg, err
	}

	return cfg, nil
}

func RunTask(t models.Task, opts models.CLIOptions) error {
	fmt.Println("Running task: ", t.Name)
	fmt.Println("[*] => ", t.Command)

	if opts.DryRun {
		fmt.Println("[*][*] DryRun: not executing command")
		return nil
	}

	parts := strings.Fields(t.Command)
	if opts.OutputFile != "" {
		parts = append(parts[:2], append([]string{"-o", opts.OutputFile}, parts[2:]...)...)
	}
	cmd := exec.Command(parts[0], parts[1:]...)

	if opts.OutputFile != "" {
		f, err := os.OpenFile(opts.OutputFile, os.O_CREATE|os.O_EXCL|os.O_APPEND, 0644)
		if err != nil {
			return fmt.Errorf("cannot")
		}
		defer f.Close()

		mw := io.MultiWriter(os.Stdout, f)
		cmd.Stdout = mw
		cmd.Stderr = mw
	} else if opts.Verbose {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("[*][*] task %s failed: %w", t.Name, err)
	}

	return nil
}
