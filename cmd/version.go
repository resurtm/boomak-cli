package cmd

import (
	"github.com/spf13/cobra"
	jww "github.com/spf13/jwalterweatherman"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints version",
	Long:  `Long version command description.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		jww.FEEDBACK.Printf("v0.0.1")
		return nil
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
