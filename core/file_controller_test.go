package core

import (
	"context"
	"os"
	"path"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type dbServiceMock struct {
	mock.Mock
}

func (mock *dbServiceMock) putFileFragmentContent(fragmentID string, fileFragment *FileFragment) error {
	args := mock.Called(fragmentID, fileFragment)
	return args.Error(0)
}

func (mock *dbServiceMock) getFileFragmentContent(fragmentID string) ([]byte, error) {
	args := mock.Called(fragmentID)
	return args.Get(0).([]byte), args.Error(1)
}

func (mock *dbServiceMock) removeFileFragments(fileContainer fileContainerInterface) []error {
	args := mock.Called(fileContainer)
	return args.Get(0).([]error)
}

func (mock *dbServiceMock) close() {
	mock.Called()
}

func TestFileControler(t *testing.T) {
	t.Run("should add file fragment", func(t *testing.T) {
		mockDbService := new(dbServiceMock)
		mockFileFragment := &FileFragment{
			FileName:        "test_file",
			FragmentID:      0,
			FragmentContent: []byte("something"),
			TotalFragments:  2, // made it 2 so as not to assemble file
		}
		fileController := fileController{
			filesMap: make(map[string]fileContainerInterface),
			mutex:    &sync.Mutex{},
			db:       mockDbService,
		}

		mockDbService.On("putFileFragmentContent", mock.Anything, mock.Anything).Return(nil)

		assert.False(t, fileController.inMap(mockFileFragment.GetFileName()))
		fileController.addFileFragment(context.Background(), mockFileFragment)
		assert.True(t, fileController.inMap(mockFileFragment.GetFileName()))
	})

	t.Run("should create file container", func(t *testing.T) {
		mockDbService := new(dbServiceMock)
		mockFileFragment := &FileFragment{
			FileName:        "test_file",
			FragmentID:      0,
			FragmentContent: []byte("something"),
			TotalFragments:  2, // made it 2 so as not to assemble file
		}
		fileController := fileController{
			filesMap: make(map[string]fileContainerInterface),
			mutex:    &sync.Mutex{},
			db:       mockDbService,
		}

		assert.False(t, fileController.inMap(mockFileFragment.GetFileName()))
		fileController.createFileContainer(mockFileFragment)
		assert.True(t, fileController.inMap(mockFileFragment.GetFileName()))
	})

	t.Run("should assemble file", func(t *testing.T) {
		mockDbService := new(dbServiceMock)
		mockFileFragment := &FileFragment{
			FileName:        "test_file",
			FragmentID:      0,
			FragmentContent: []byte("something"),
			TotalFragments:  2, // made it 2 so as not to assemble file
		}
		fileController := fileController{
			filesMap: make(map[string]fileContainerInterface),
			mutex:    &sync.Mutex{},
			db:       mockDbService,
		}
		testFilePath := path.Join(getDroneDownloadsPath(), mockFileFragment.GetFileName())

		mockDbService.On("putFileFragmentContent", mock.Anything, mock.Anything).Return(nil)
		mockDbService.On("getFileFragmentContent", mock.Anything).Return([]byte{}, nil)
		mockDbService.On("removeFileFragments", mock.Anything).Return([]error{})

		fileController.addFileFragment(context.Background(), mockFileFragment)
		fileController.assembleFile(mockFileFragment.GetFileName())
		assert.True(t, fileExists(testFilePath))
		os.Remove(testFilePath)
	})

	t.Run("should be in map", func(t *testing.T) {
		mockDbService := new(dbServiceMock)
		mockFileFragment := &FileFragment{
			FileName:        "test_file",
			FragmentID:      0,
			FragmentContent: []byte("something"),
			TotalFragments:  2, // made it 2 so as not to assemble file
		}
		fileController := fileController{
			filesMap: make(map[string]fileContainerInterface),
			mutex:    &sync.Mutex{},
			db:       mockDbService,
		}

		fileController.createFileContainer(mockFileFragment)
		assert.True(t, fileController.inMap(mockFileFragment.GetFileName()))
	})

	t.Run("should not be in map", func(t *testing.T) {
		mockDbService := new(dbServiceMock)
		mockFileFragment := &FileFragment{
			FileName:        "test_file",
			FragmentID:      0,
			FragmentContent: []byte("something"),
			TotalFragments:  2, // made it 2 so as not to assemble file
		}
		fileController := fileController{
			filesMap: make(map[string]fileContainerInterface),
			mutex:    &sync.Mutex{},
			db:       mockDbService,
		}

		assert.False(t, fileController.inMap(mockFileFragment.GetFileName()))
	})
}
