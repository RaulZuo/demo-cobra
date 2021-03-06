package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:	"version",
	Short: "Print the version number of Hugo",
	Long: `All software has versions. This is Hugo's`,
	Args: cobra.MinimumNArgs(1),	// 至少1个非选项参数
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args)
		fmt.Println("Hugo Static Site Generator v0.9 -- HEAD", cfgFile)
	},
}
