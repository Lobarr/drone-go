package core

import (
	"fmt"
	"net"

	"github.com/kpango/glg"
	"google.golang.org/grpc"
)

//StartServer starts drone server
func StartServer(port int) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		glg.Fatalf("[DroneServer] Failed to listen: %v", err)
	}

	server := grpc.NewServer()
	droneService, err := NewDroneService(getDroneDBPath())
	if err != nil {
		glg.Fatalf("[DroneServer] %s", err.Error())
	}

	defer droneService.CloseDB()
	RegisterDroneServer(server, droneService)

	glg.Infof("[DroneServer] Starting on 0.0.0.0:%d", port)
	if err := server.Serve(listener); err != nil {
		glg.Fatalf("[DroneServer] %s", err.Error())
	}
}
