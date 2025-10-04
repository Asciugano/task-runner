package cmd

import (
	"os"

	"github.com/Asciugano/taskRunner/internal/models"
	"github.com/Asciugano/taskRunner/internal/runner"
	"github.com/spf13/cobra"
)

var opts = models.CLIOptions{}

var rootCmd = &cobra.Command{
	Use:   "taskRunner",
	Short: "",
	Long:  ``,

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 && !opts.All {
			cmd.PrintErrln("Please specify a task")
			os.Exit(1)
		} else {
			if !opts.All {
				opts.TaskName = args[0]
			}
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}

	runner.Init(opts)
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&opts.Verbose, "verbose", "v", false, "Show verbose output")
	rootCmd.PersistentFlags().BoolVarP(&opts.All, "all", "A", false, "Run all tasks")
	rootCmd.PersistentFlags().IntVarP(&opts.Concurrency, "concurrency", "c", 1, "Number of tasks in parallel")
	rootCmd.PersistentFlags().BoolVarP(&opts.DryRun, "dry-run", "d", false, "Do not execute tasks")
	rootCmd.PersistentFlags().BoolVar(&opts.ContinueOnError, "continue-on-error", false, "Continue on error")
	rootCmd.PersistentFlags().StringVar(&opts.ConfigPath, "config-path", "./tasks.yaml", "Path to tasks config")
	rootCmd.PersistentFlags().StringVarP(&opts.OutputFile, "out-file", "o", "", "Output file")
	rootCmd.PersistentFlags().BoolVarP(&opts.Parallel, "parallel", "p", false, "Run multiple tasks in parallel")
	rootCmd.PersistentFlags().BoolVarP(&opts.List, "list", "l", false, "List all the tasks")
	rootCmd.PersistentFlags().BoolVarP(&opts.Graph, "graph", "g", false, "")
	rootCmd.PersistentFlags().BoolVar(&opts.Version, "version", false, "Display the version")
}
