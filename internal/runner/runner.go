package runner

import (
	"errors"
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

	sorted, err := SortTasks(tasks)
	if err != nil {
		fmt.Errorf("Error in the task dependency: %w", err)
	}
	for _, t := range sorted {
		RunTask(t, options)
	}
}

func SortTasks(tasks models.Config) ([]models.Task, error) {
	taskMap := make(map[string]models.Task)
	for _, t := range tasks.Tasks {
		taskMap[t.Name] = t
	}

	visited := make(map[string]bool)
	tempMarked := make(map[string]bool)
	var sorted []models.Task

	var visit func(string) error
	visit = func(name string) error {
		if tempMarked[name] {
			return errors.New("Circular dependency detectred at task: " + name)
		}
		if visited[name] {
			return nil
		}
		tempMarked[name] = true

		task, ok := taskMap[name]
		if !ok {
			return fmt.Errorf("task %s not found", name)
		}

		for _, dep := range task.DependsOn {
			if err := visit(dep); err != nil {
				return err
			}
		}

		visited[name] = true
		tempMarked[name] = false
		sorted = append(sorted, task)
		return nil
	}

	for _, t := range tasks.Tasks {
		if !visited[t.Name] {
			if err := visit(t.Name); err != nil {
				return nil, err
			}
		}
	}

	return sorted, nil
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
