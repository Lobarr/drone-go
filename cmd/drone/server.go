package drone

import (
	"github.com/spf13/cobra"
	"github.com/Lobarr/drone-go/core"
)

//flags
var port int

var serverCmd = &cobra.Command{
	Use: "server",
	Short: "Used to start a drone server",
	Run: func(cmd *cobra.Command, args []string) {
		core.StartServer(port)
	},
}

func init() {
	serverCmd.Flags().IntVarP(&port, "port", "p", 9999, "Port to run drone server")
}


