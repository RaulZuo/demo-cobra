// add sub command for version
package cmd

import (
	"errors"
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
	// 自定义验证函数
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires at least one arg")
		}
		if args[0] != "" {
			return nil
		}
		return fmt.Errorf("invalid color specified: %s", args[0])
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hugo Version Show")
	},
}
