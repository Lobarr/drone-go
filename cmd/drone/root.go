package drone

import (
	"github.com/kpango/glg"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "drone",
	Short: "Drone is a simple file transfer program",
}

func init() {
	rootCmd.AddCommand(sendCmd)
	rootCmd.AddCommand(serverCmd)
}

//Execute entry point to drone cli app
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		glg.Fatalf("[DroneCmd] %s", err.Error())
	}
}
