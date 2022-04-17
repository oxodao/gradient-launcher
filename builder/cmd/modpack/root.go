package cmd_modpack

import (
	"os"

	"github.com/spf13/cobra"
)

var modpack string

var rootCmd = &cobra.Command{
	Use:   "modpack",
	Short: "Modpack related commands",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if modpack == "" {
			cmd.Help()
			os.Exit(1)
		}
	},
}

func RegisterModpackCmd(root *cobra.Command) {
	rootCmd.Flags().StringVar(&modpack, "modpack", "", "Modpack name")
	root.AddCommand(rootCmd)
}
