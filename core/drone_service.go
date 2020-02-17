package core

import (
	"io"
	"sync"
)

//DroneService implements the drone grpc service
type DroneService struct {
	fc *fileController
}

//NewDroneService creates a new drone service
func NewDroneService(dbFilePath string) (*DroneService, error) {
	dbService, err := newDBService(dbFilePath)
	if err != nil {
		return nil, err
	}
	return &DroneService{
		fc: &fileController{
			filesMap: make(map[string]*fileContainer),
			mutex:    &sync.Mutex{},
			db:       dbService,
		},
	}, nil
}

//CloseDB closes level db conn
func (droneService *DroneService) CloseDB() {
	droneService.fc.db.conn.Close()
}

//ReceiveFile receives incoming files
func (droneService *DroneService) ReceiveFile(stream Drone_ReceiveFileServer) error {
	for {
		fileFragment, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&Status{
				StatusCode: 200,
				Message:    "OK",
			})
		} else if err != nil {
			return err
		}

		droneService.fc.putFileFragment(stream.Context(), fileFragment)
	}
}
