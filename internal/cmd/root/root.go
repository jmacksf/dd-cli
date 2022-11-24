package root

import (
	"github.com/spf13/cobra"

	"github.com/jmacksf/dd-cli/internal/cmd/search"
	"github.com/jmacksf/dd-cli/internal/cmd/version"
)

// NewCmdRoot is a root command.
func NewCmdRoot() *cobra.Command {
	cmd := cobra.Command{
		Use:   "dd-cli",
		Short: "Datadog helper cli",
		Long:  "Datadog helper cli.",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			subCmd := cmd.Name()
			if !cmdRequireToken(subCmd) {
				return
			}
		},
	}

	addChildCommands(&cmd)

	return &cmd
}

func addChildCommands(cmd *cobra.Command) {
	cmd.AddCommand(
		version.NewCmdVersion(),
		search.NewCmdSearch(),
	)
}

func cmdRequireToken(cmd string) bool {
	allowList := []string{
		"help",
		"version",
	}

	for _, item := range allowList {
		if item == cmd {
			return false
		}
	}

	return true
}
