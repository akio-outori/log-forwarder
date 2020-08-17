package cmd

import (
  "github.com/spf13/cobra"
  "github.com/spf13/cobra/doc"
)

func init() {
  rootCmd.AddCommand(docs)
}

var docs = &cobra.Command {
  Use:   "docs",
  Short: "Generate Markdown documentation",
  Long:  `Command should generate documentation for all commands in the cmd package`,
  Run:   func(cmd *cobra.Command, args []string) {

    commands := [...]*cobra.Command{cpuinfo, meminfo, cmd}

    for _, command := range commands { 
      err := doc.GenMarkdownTree(command, "docs")
      checkError(err)
    }
  },
}
