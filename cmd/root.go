package cmd

import (
	"github.com/jonghyeons/photopic/models/constant"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "photopic",
	Short: "Photopic is JPG, RAW file sorting program in terminal",
	Long: `    ____  __          __              _     
   / __ \/ /_  ____  / /_____  ____  (_)____
  / /_/ / __ \/ __ \/ __/ __ \/ __ \/ / ___/
 / ____/ / / / /_/ / /_/ /_/ / /_/ / / /__  
/_/   /_/ /_/\____/\__/\____/ .___/_/\___/  
                           /_/              v` + constant.Version +
		`
Photopic is JPG, RAW file sorting program in terminal`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(runCmd)
	rootCmd.AddCommand(versionCmd)
}
