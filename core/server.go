package core

import (
	"fmt"
	"net"

	"net/http"
	_ "net/http/pprof"

	"github.com/kpango/glg"
	"google.golang.org/grpc"
)

var droneServerLogTemplate = "[DroneServer] %s"

// ServerOptions options passed into the application
type ServerOptions struct {
	Port        int
	Profile     bool
	ProfilePort int
}

//StartServer starts drone server
func StartServer(options ServerOptions) {
	if options.Profile {
		go func() {
			glg.Get().Log(http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", options.ProfilePort), nil))
		}()
	}

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", options.Port))
	if err != nil {
		glg.Get().Fatalf(droneServerLogTemplate, fmt.Sprintf("Failed to listen: %v", err))
	}

	server := grpc.NewServer()
	droneService, err := NewDroneService(getDroneDBPath())
	if err != nil {
		glg.Get().Fatalf(droneServerLogTemplate, err.Error())
	}

	defer droneService.CloseDB()
	RegisterDroneServer(server, droneService)

	glg.Get().Infof(droneServerLogTemplate, fmt.Sprintf("Starting on 0.0.0.0:%d", options.Port))
	if err := server.Serve(listener); err != nil {
		glg.Get().Fatalf(droneServerLogTemplate, err.Error())
	}
}
