// add sub command for version
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	versionCmd.AddCommand(showCmd)
}

var showCmd = &cobra.Command{
	Use:	"show",
	Short:	"Show the version number of Hugo",
	Long: `All software has versions. This is Hugo's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hugo Version Show")
	},
}
