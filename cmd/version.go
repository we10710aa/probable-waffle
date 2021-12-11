package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "print version of waffle",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("waffle v0.0.1")
	},
}
