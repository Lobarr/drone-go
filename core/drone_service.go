package core

import (
	"errors"
	"io"
	"sync"

	"github.com/panjf2000/ants/v2"
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
	errChan := make(chan error, 1)
	doneChan := make(chan struct{}, 1)
	pool, err := ants.NewPool(50)

	if err != nil {
		return errors.New("Unable to create worker pool")
	}

	go func() {
		for {
			fileFragment, err := stream.Recv()
			if err == io.EOF {
				doneChan <- struct{}{}
				break
			} else if err != nil {
				errChan <- err
			}

			pool.Submit(func() {
				err := droneService.fc.addFileFragment(stream.Context(), fileFragment)
				if err != nil {
					errChan <- err
				}
			})
		}
	}()

	select {
	case <-doneChan:
		pool.Release()
		return stream.SendAndClose(&Status{
			StatusCode: 200,
			Message:    "OK",
		})
	case err = <-errChan:
		pool.Release()
		return err
	}
}
