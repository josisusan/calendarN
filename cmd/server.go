package cmd

import (
	"github.com/samit22/calendarN/apiserver"
	"github.com/spf13/cobra"
)

var port int
var secret string
var serverCmd = &cobra.Command{
	Use:     "server",
	Short:   "run server at port 5001 (default)",
	PostRun: PostRunMsg,
	Run: func(cmd *cobra.Command, args []string) {
		if port == 0 {
			port = 5001
		}
		if secret == "" {
			secret = "secret"
		}
		apiserver.Run(port, secret)
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
	serverCmd.Flags().IntVarP(&port, "port", "p", port, "port")
	serverCmd.Flags().StringVarP(&secret, "token", "t", secret, "static api key")
}
