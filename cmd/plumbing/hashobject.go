package plumbing

import (
	"bufio"
	"errors"
	"io"
	"os"

	"github.com/ppichler94/got/internal/objectdb"
	"github.com/spf13/cobra"
)

var hashObjectFlags struct {
	stdin bool
	write bool
}

func init() {
	hashobjectCmd.Flags().BoolVar(&hashObjectFlags.stdin, "stdin", false, "Read from stdin")
	hashobjectCmd.Flags().BoolVarP(&hashObjectFlags.write, "write", "w", false, "Write to the database")
}

var hashobjectCmd = &cobra.Command{
	Use:     "hash-object",
	Short:   "Compute the object ID and optionally store it in the database.",
	GroupID: groupId,
	RunE: func(cmd *cobra.Command, args []string) error {
		var input string
		var err error
		if hashObjectFlags.stdin {
			reader := bufio.NewReader(os.Stdin)
			input, err = reader.ReadString(0)
			if err != nil {
				if !errors.Is(err, io.EOF) {
					return err
				}
			}
		}

		if len(args) > 0 {
			content, err := os.ReadFile(args[0])
			if err != nil {
				return err
			}
			input = string(content)
		}

		blob := objectdb.NewBlob(input)

		if hashObjectFlags.write {
			if err := blob.Write(); err != nil {
				return err
			}
		}

		cmd.Println(blob.Hash)

		return nil
	},
}
