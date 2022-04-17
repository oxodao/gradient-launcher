package cmd

import (
	cmd_modpack "github.com/oxodao/gradient-builder/cmd/modpack"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gradient-builder",
	Short: "Minecraft modpack builder",
	Long:  `Gradient Builder is a tool for building Minecraft modpacks.`,
	//Run: func(cmd *cobra.Command, args []string) {
	// @TODO: Later this should open the GUI
	//},
}

func init() {
	cmd_modpack.RegisterModpackCmd(rootCmd)
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
