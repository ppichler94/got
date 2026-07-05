package cmd

import (
	"fmt"
	"os"

	"github.com/ppichler94/got/cmd/plumbing"
	"github.com/ppichler94/got/cmd/porcelain"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "got",
	Short: "A git implementation in Go",
}

func init() {
	porcelain.InitCmds(rootCmd)
	plumbing.InitCmds(rootCmd)
}

func Execute() {
	rootCmd.SetOut(os.Stdout)

	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
