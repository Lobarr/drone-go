package core

import (
	"context"
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
	fileFragmentchan := make(chan *FileFragment, 3)
	workersCount := 3
	workersContext, cancel := context.WithCancel(stream.Context())
	defer cancel()

	for i := 0; i < workersCount; i++ {
		go func(ctx context.Context) {
			for {
				select {
				case <-ctx.Done():
					return
				case fileFragment := <-fileFragmentchan:
					err := droneService.fc.addFileFragment(workersContext, fileFragment)
					if err != nil {
						errChan <- err
					}
				}
			}
		}(workersContext)
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

			fileFragmentchan <- fileFragment
		}
	}()

	select {
	case <-doneChan:
		return stream.SendAndClose(&Status{
			StatusCode: 200,
			Message:    "OK",
		})
	case err := <-errChan:
		return err
	}
}
