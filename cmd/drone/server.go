package drone

import (
	"github.com/Lobarr/drone-go/core"
	"github.com/kpango/glg"
	"github.com/spf13/cobra"
)

//flags
var (
	port        int // server port
	profilePort int
	profile     bool
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Starts a drone server",
	Run: func(cmd *cobra.Command, args []string) {
		if port == profilePort {
			glg.Get().Fatal("Unable to run profiler on same port as server")
		}

		core.StartServer(core.ServerOptions{
			ProfilePort: profilePort,
			Port:        port,
			Profile:     profile,
		})
	},
}

func init() {
	serverCmd.Flags().IntVarP(&port, "port", "p", 9999, "Port to run drone server")
	serverCmd.Flags().BoolVar(&profile, "profile", false, "Profiles the server")
	serverCmd.Flags().IntVar(&profilePort, "profilePort", 9998, "Port to run profiler")
}
