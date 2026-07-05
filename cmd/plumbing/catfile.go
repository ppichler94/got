package plumbing

import (
	"github.com/ppichler94/got/internal/objectdb"
	"github.com/spf13/cobra"
)

var catFileFlags struct {
	prettyPrint bool
}

func init() {
	catfileCmd.Flags().BoolVarP(&catFileFlags.prettyPrint, "pretty-print", "p", false, "Pretty-print the contents of <object> based on its type.")
}

var catfileCmd = &cobra.Command{
	Use:     "cat-file <object>",
	Short:   "Provide contents or details of repository objects",
	GroupID: groupId,
	Args:    cobra.ExactArgs(1),
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
