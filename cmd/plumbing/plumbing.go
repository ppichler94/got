package plumbing

import "github.com/spf13/cobra"

const groupId = "plumbing"

func InitCmds(rootCmd *cobra.Command) {
	rootCmd.AddGroup(&cobra.Group{
		ID:    groupId,
		Title: "Plumbing Commands",
	})

	rootCmd.AddCommand(catfileCmd)
	rootCmd.AddCommand(hashobjectCmd)
	rootCmd.AddCommand(updateIndexCmd)
}
