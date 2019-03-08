package cmd

import (
	"github.com/kudobuilder/kudo/version"
	"github.com/spf13/cobra"
)

func NewKudoctlCmd() *cobra.Command {
	cmd := &cobra.Command{
		// Workaround or Compromise as "kubectl kudo" would result in Usage print out "kubectl install <name> [flags]"
		Use:   "kubectl-kudo",
		Short: "CLI to manipulate, inspect and troubleshoot KUDO-specific CRDs.",
		Long: `
KUDO CLI and future sub-commands can be used to manipulate, inspect and troubleshoot KUDO-specific CRDs
and serves as an API aggregation layer.
`,
		Example: `
	# Install a KUDO package from the official GitHub repo.
	kubectl kudo install <name> [flags]

	# View plan history of a specific package
	kubectl kudo plan history <name> [flags]

	# View all plan history of a specific package
	kubectl kudo plan history [flags]

	# List instances
	kubectl kudo list instances [flags]

	# View plan status
	kubectl kudo plan status [flags]

`,
		Version: version.Version,
	}

	cmd.AddCommand(NewCmdInstall())
	cmd.AddCommand(NewListCmd())
	cmd.AddCommand(NewPlanCmd())

	return cmd
}
