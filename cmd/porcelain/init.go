package porcelain

import (
	"path/filepath"

	"github.com/ppichler94/got/internal/objectdb"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:     "init [<directory>]",
	Short:   "Initialize a new, empty repository",
	GroupID: groupId,
	Args:    cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		dir := "./"
		if len(args) > 0 {
			dir = args[0]
		}

		if err := objectdb.Init(dir); err != nil {
			return err
		}

		absDir, err := filepath.Abs(dir)
		if err != nil {
			return err
		}
		cmd.Printf("Initialized empty Got repository in %s\n", absDir)

		return nil
	},
}
