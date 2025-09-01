package cmd

import (
	"github.com/samit22/calendarN/apiserver"
	"github.com/spf13/cobra"
)

var port int
var serverCmd = &cobra.Command{
	Use:     "server",
	Short:   "run server at port 5001 (default)",
	PostRun: PostRunMsg,
	Run: func(cmd *cobra.Command, args []string) {
		if port == 0 {
			port = 5001
		}
		apiserver.Run(port)
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
	serverCmd.Flags().IntVarP(&port, "port", "p", port, "port")
}
