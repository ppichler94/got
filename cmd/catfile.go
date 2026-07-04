package cmd

import (
	"github.com/ppichler94/got/internal/objectdb"
	"github.com/spf13/cobra"
)

type catfileFlags struct {
	prettyPrint bool
}

var catFileFlags catfileFlags

func init() {
	rootCmd.AddCommand(catfileCmd)

	catfileCmd.Flags().BoolVarP(&catFileFlags.prettyPrint, "pretty-print", "p", false, "Pretty-print the contents of <object> based on its type.")
}

var catfileCmd = &cobra.Command{
	Use:   "cat-file <object>",
	Short: "Provide contents or details of repository objects",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if catFileFlags.prettyPrint {
			object, err := objectdb.Read(args[0])
			if err != nil {
				return err
			}

			cmd.Print(object.PrettyPrint())
		}

		return nil
	},
}
