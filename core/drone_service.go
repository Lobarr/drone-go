package core

import (
	"context"
	"io"
	"sync"

	"github.com/uw-labs/sync/rungroup"
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
			filesMap: make(map[string]fileContainerInterface),
			mutex:    &sync.Mutex{},
			db:       dbService,
		},
	}, nil
}

//CloseDB closes level db conn
func (droneService *DroneService) CloseDB() {
	droneService.fc.db.close()
}

//ReceiveFile receives incoming files
func (droneService *DroneService) ReceiveFile(stream Drone_ReceiveFileServer) error {
	runGroup, ctx := rungroup.New(context.Background())

	for {
		fileFragment, err := stream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		runGroup.Go(func() error {
			return droneService.fc.addFileFragment(ctx, fileFragment)
		})
	}

	if err := runGroup.Wait(); err != nil {
		return err
	}

	return stream.SendAndClose(&Status{
		StatusCode: 200,
		Message:    "OK",
	})
}
