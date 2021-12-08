package main

import (
	"github.com/spf13/cobra"
	"github.com/we10710aa/waffle/police"
)

func main() {
	rootCmd := &cobra.Command{}
	rootCmd.AddCommand(police.ServeCmd)
	rootCmd.Execute()
}
