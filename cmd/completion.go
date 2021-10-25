package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

func newCompletionCommand() *cobra.Command {
	cmd := &cobra.Command{
		Args:  cobra.ExactValidArgs(1),
		Short: "Output shell completion code for the specified shell (bash, zsh, fish or powershell)",
		Use:   "completion [bash|zsh|fish|powershell]",
		Example: `  # Load bash completion into current shell
  source <(vmkit completion bash)

  # Load zsh completion into current shell
  source <(vmkit completion zsh)

  # Load fish completion into current shell
  vmkit completion fish | source

  # Load powershell completion into current shell
  vmkit completion powershell | Out-String | Invoke-Expression`,
		ValidArgs: []string{"bash", "zsh", "fish", "powershell"},
		Run: func(cmd *cobra.Command, args []string) {
			switch args[0] {
			case "bash":
				cmd.Root().GenBashCompletion(os.Stdout)
			case "zsh":
				cmd.Root().GenZshCompletion(os.Stdout)
			case "fish":
				cmd.Root().GenFishCompletion(os.Stdout, true)
			case "powershell":
				cmd.Root().GenPowerShellCompletionWithDesc(os.Stdout)
			}
		},
	}

	return cmd
}
