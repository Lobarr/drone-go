package core

import (
	"context"
	"os"
	"path"
	"sync"

	"github.com/google/uuid"
	"github.com/kpango/glg"
)

// coordinates receiving and sending files
type fileController struct {
	filesMap map[string]*fileContainer
	mutex    *sync.Mutex
	db       *dbService
}

func (fc fileController) createDownloadsFolder() error {
	return os.Mkdir(getDroneDownloadsPath(), os.FileMode(0755))
}

func (fc *fileController) putFileFragment(ctx context.Context, fileFragment *FileFragment) error {
	var err error
	done := make(chan struct{}, 1)

	go func() {
		glg.Debugf("[FileController] Putting fileFragment %d of %s", fileFragment.GetFragmentId(), fileFragment.GetFileName())
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
					fc.assembleFile(fileContainer.fileName)
				}()
			}
		}

		done <- struct{}{}
	}()

	select {
	case <-done:
		if err != nil {
			glg.Error(err.Error())
			return err
		}
	case <-ctx.Done():
		glg.Error(ctx.Err())
	}
	return nil
}

func (fc *fileController) createFileContainer(fileFragment *FileFragment) {
	glg.Debugf("[FileController] Initializing receipt of %s", fileFragment.GetFileName())
	fc.filesMap[fileFragment.GetFileName()] = newFileContainer(fileFragment.GetFileName(), int(fileFragment.GetTotalFragments()))
}

func (fc *fileController) assembleFile(fileName string) {
	glg.Debugf("[FileController] Assembing file %s", fileName)

	if !fc.inMap(fileName) {
		glg.Error("[FileController] Unable to assemble invalid file")
		return
	}

	fileContainer := fc.filesMap[fileName]
	if !fileExists(getDroneDownloadsPath()) {
		err := fc.createDownloadsFolder()
		if err != nil {
			glg.Fatalf("[Filecontroller] %s", err.Error())
			return
		}
	}

	file, err := os.Create(path.Join(getDroneDownloadsPath(), fileName))
	if err != nil {
		glg.Error("[FileController] Unable to create file", err)
		return
	}

	for _, fragmenID := range fileContainer.fragmentIDs {
		fileFragmentContent, err := fc.db.getFileFragmentContent(fragmenID)
		if err != nil {
			glg.Error("[FileController] Unable to get file fragment")
		} else {
			_, err := file.Write(fileFragmentContent)
			if err != nil {
				glg.Error("[FileController] Unable to write file fragment to file")
			}
		}
	}

	file.Close()
	fc.db.removeFileFragments(fileContainer.fragmentIDs...)
	delete(fc.filesMap, fileName)
}

func (fc fileController) inMap(fileName string) bool {
	if _, ok := fc.filesMap[fileName]; !ok {
		return false
	}
	return true
}
