package drone

import (
	"github.com/Lobarr/drone-go/core"
	"github.com/spf13/cobra"
)

//flags
var port int

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Starts a drone server",
	Run: func(cmd *cobra.Command, args []string) {
		core.StartServer(port)
	},
}

func init() {
	serverCmd.Flags().IntVarP(&port, "port", "p", 9999, "Port to run drone server")
}
