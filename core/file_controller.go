package core

import (
	"context"
	"fmt"
	"os"
	"path"
	"sync"

	"github.com/google/uuid"
	"github.com/kpango/glg"
)

var fileControllerLogTemplate = "[FileController] %s"

// coordinates receiving and sending files
type fileController struct {
	filesMap map[string]fileContainerInterface
	mutex    *sync.Mutex
	db       dbServiceInteface
}

func (fc fileController) createDownloadsFolder() error {
	return os.Mkdir(getDroneDownloadsPath(), os.FileMode(0755))
}

func (fc *fileController) addFileFragment(ctx context.Context, fileFragment *FileFragment) error {
	var err error
	done := make(chan struct{}, 1)

	go func() {
		glg.Debugf(fileControllerLogTemplate, fmt.Sprintf("Putting fileFragment %d of %s", fileFragment.GetFragmentId(), fileFragment.GetFileName()))

		fc.mutex.Lock()
		defer fc.mutex.Unlock()

		if !fc.inMap(fileFragment.GetFileName()) {
			fc.createFileContainer(fileFragment)
		}

		fragmentID := uuid.New().String()
		fileContainer := fc.filesMap[fileFragment.GetFileName()]
		err = fc.db.putFileFragmentContent(fragmentID, fileFragment)

		if err == nil {
			fileContainer.addFragment(fragmentID)
			if fileContainer.isComplete() {
				go func() {
					fc.assembleFile(fileContainer.getFileName())
				}()
			}
		}

		done <- struct{}{}
	}()

	select {
	case <-done:
		if err != nil {
			return err
		}
	case <-ctx.Done():
	}
	return nil
}

func (fc *fileController) createFileContainer(fileFragment *FileFragment) {
	glg.Debugf(fileControllerLogTemplate, fmt.Sprintf("Initializing receipt of %s", fileFragment.GetFileName()))
	fc.filesMap[fileFragment.GetFileName()] = newFileContainer(fileFragment.GetFileName(), int(fileFragment.GetTotalFragments()))
}

func (fc *fileController) assembleFile(fileName string) {
	glg.Debugf(fileControllerLogTemplate, fmt.Sprintf("Assembing file %s", fileName))

	if !fc.inMap(fileName) {
		glg.Errorf(fileControllerLogTemplate, "Unable to assemble invalid file")
		return
	}

	fileContainer := fc.filesMap[fileName]
	if !fileExists(getDroneDownloadsPath()) {
		err := fc.createDownloadsFolder()
		if err != nil {
			glg.Fatalf(fileControllerLogTemplate, err.Error())
			return
		}
	}

	file, err := os.Create(path.Join(getDroneDownloadsPath(), fileName))
	if err != nil {
		glg.Errorf(fileControllerLogTemplate, fmt.Sprintf("Unable to create file due to %s", err.Error()))
		return
	}
	defer file.Close()

	for _, fragmenID := range fileContainer.getFragmentIDs() {
		fileFragmentContent, err := fc.db.getFileFragmentContent(fragmenID)
		if err != nil {
			glg.Errorf(fileControllerLogTemplate, "Unable to get file fragment")
		} else {
			_, err := file.Write(fileFragmentContent)
			if err != nil {
				glg.Errorf(fileControllerLogTemplate, "Unable to write file fragment to file")
			}
		}
	}

	fc.db.removeFileFragments(fileContainer.getFragmentIDs()...)
	delete(fc.filesMap, fileName)
}

func (fc fileController) inMap(fileName string) bool {
	if _, ok := fc.filesMap[fileName]; !ok {
		return false
	}
	return true
}
