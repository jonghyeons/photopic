package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "photopic",
	Short: "Photopic is JPG, RAW file sorting program in terminal",
	Long:  `Photopic is JPG, RAW file sorting program in terminal`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
