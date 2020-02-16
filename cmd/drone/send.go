package drone

import (
	"github.com/Lobarr/drone-go/core"
	"github.com/spf13/cobra"
)

// flags
var host string
var fragmentSize int

var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "Used to communicate with drone servers",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return core.SendFiles(args, host, fragmentSize)
	},
}

func init() {
	sendCmd.Flags().StringVarP(&host, "host", "", "0.0.0.0:9999", "Receipient drone server address")
	sendCmd.Flags().IntVarP(&fragmentSize, "fragmentSize", "", 2000, "Size of each fragment of a file in bytes")
}
