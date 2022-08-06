package boxcli

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"go.jetpack.io/axiom/opensource/devbox"
)

func ShellCmd() *cobra.Command {
	command := &cobra.Command{
		Use:  "shell [<dir>]",
		Args: cobra.MaximumNArgs(1),
		RunE: runShellCmd,
	}
	return command
}

func runShellCmd(cmd *cobra.Command, args []string) error {
	path := pathArg(args)

	// Check the directory exists.
	box, err := devbox.Open(path)
	if err != nil {
		return errors.WithStack(err)
	}

	return box.Shell()
}