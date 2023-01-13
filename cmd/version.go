package cmd

import (
	"fmt"
	"github.com/jonghyeons/photopic/models/constant"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Output the version number.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(fmt.Sprintf("Photopic version %s", constant.Version))
	},
}
