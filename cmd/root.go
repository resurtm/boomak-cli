package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "bmk",
	Short: "Boomak CLI Application",
	Long:  `Long root command description.`,
	//RunE: func(cmd *cobra.Command, args []string) error {
	//	return nil
	//},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
}
