package porcelain

import "github.com/spf13/cobra"

const groupId = "porcelain"

func InitCmds(rootCmd *cobra.Command) {
	rootCmd.AddGroup(&cobra.Group{
		ID:    groupId,
		Title: "Porcelain Commands",
	})

	rootCmd.AddCommand(initCmd)
}
