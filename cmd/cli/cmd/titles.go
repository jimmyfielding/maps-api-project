package cmd

import (
	"io"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	client *TitlesClient
)

// NewTitlesCommand creates the root `titles` command and its nested children.
func NewTitlesCommand(in io.Reader, out, errOut io.Writer) *cobra.Command {
	var cmd = &cobra.Command{
		Use:              "titles",
		PersistentPreRun: preRun,
	}

	cmd.AddCommand(NewCmdGenerate(out, errOut))

	return cmd
}

func preRun(cmd *cobra.Command, args []string) {
	log := logrus.New()
	client, _ = NewTitlesClient("http://127.0.0.1:8080", log)
}
