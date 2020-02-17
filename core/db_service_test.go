package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/syndtr/goleveldb/leveldb/opt"
)

type dbConnectionMock struct {
	mock.Mock
}

func (mock *dbConnectionMock) Put(key, value []byte, wo *opt.WriteOptions) error {
	args := mock.Called(key, value, wo)
	return args.Error(0)
}

func (mock *dbConnectionMock) Get(key []byte, ro *opt.ReadOptions) (value []byte, err error) {
	args := mock.Called(key, ro)
	return args.Get(0).([]byte), args.Error(1)
}

func (mock *dbConnectionMock) Delete(key []byte, wo *opt.WriteOptions) error {
	args := mock.Called(key, wo)
	return args.Error(0)
}

func (mock *dbConnectionMock) Close() error {
	args := mock.Called()
	return args.Error(0)
}

func TestPutFileFragmentContent(t *testing.T) {
	mockFragmentID := "some-id"
	mockFileFragment := new(FileFragment)
	mockDbConn := new(dbConnectionMock)
	testDbService := dbService{
		conn: mockDbConn,
	}

	mockDbConn.On("Put", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	assert.NoError(t, testDbService.putFileFragmentContent(mockFragmentID, mockFileFragment))
	mockDbConn.AssertCalled(t, "Put", []byte(mockFragmentID), mock.Anything, mock.Anything)
}

func TestGetFileFragmentContent(t *testing.T) {
	mockFragmentID := "some-id"
	mockDbConn := new(dbConnectionMock)
	expectedFragmentContent := []byte{}
	testDbService := dbService{
		conn: mockDbConn,
	}

	mockDbConn.On("Get", mock.Anything, mock.Anything).Return(expectedFragmentContent, nil)

	fragmentContent, err := testDbService.getFileFragmentContent(mockFragmentID)

	assert.NoError(t, err)
	assert.Equal(t, expectedFragmentContent, fragmentContent)
	mockDbConn.AssertCalled(t, "Get", []byte(mockFragmentID), mock.Anything)
}

func TestRemoveFileFragments(t *testing.T) {
	mockFragmentID := "some-id"
	mockDbConn := new(dbConnectionMock)
	testDbService := dbService{
		conn: mockDbConn,
	}

	mockDbConn.On("Delete", mock.Anything, mock.Anything).Return(nil)

	errs := testDbService.removeFileFragments(mockFragmentID)
	assert.Empty(t, errs)
	mockDbConn.AssertCalled(t, "Delete", []byte(mockFragmentID), mock.Anything)
}
