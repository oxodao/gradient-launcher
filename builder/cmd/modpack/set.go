package cmd_modpack

import (
	"fmt"

	"github.com/oxodao/gradient-builder/services"
	"github.com/spf13/cobra"
)

var setCmdProperties = []string{"name", "author", "mc_version", "forge_version", "fabric_loader_version"}

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set modpack properties",
	Long: `Set the modpack properties
	Allowed properties:
		- name
		- author
		- mc_version
		- forge_version
		- fabric_loader_version
	`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("Missing property name")
		}

		for _, prop := range setCmdProperties {
			if args[0] == prop {
				return nil
			}
		}

		return fmt.Errorf("Unknown property name: %s", args[0])
	},
	Run: func(cmd *cobra.Command, args []string) {
		services.GET.
	},
}
