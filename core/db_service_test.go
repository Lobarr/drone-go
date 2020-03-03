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

func TestDBService(t *testing.T) {
	mockFragmentID := "some-id"
	mockDbConn := new(dbConnectionMock)
	testDbService := dbService{
		conn: mockDbConn,
	}

	t.Run("should put file fragment content in db", func(t *testing.T) {
		mockFileFragment := new(FileFragment)

		mockDbConn.On("Put", mock.Anything, mock.Anything, mock.Anything).Return(nil)
		assert.NoError(t, testDbService.putFileFragmentContent(mockFragmentID, mockFileFragment))
		mockDbConn.AssertCalled(t, "Put", []byte(mockFragmentID), mock.Anything, mock.Anything)
	})

	t.Run("should get file fragment content from db", func(t *testing.T) {
		expectedFragmentContent := []byte{}

		mockDbConn.On("Get", mock.Anything, mock.Anything).Return(expectedFragmentContent, nil)
		fragmentContent, err := testDbService.getFileFragmentContent(mockFragmentID)

		assert.NoError(t, err)
		assert.Equal(t, expectedFragmentContent, fragmentContent)
		mockDbConn.AssertCalled(t, "Get", []byte(mockFragmentID), mock.Anything)
	})

	t.Run("should remove file fragments from db", func(t *testing.T) {
		mockFileContainer := &fileContainer{
			fileName:               "some-name",
			transactionID:          "some-transaction-id",
			receivedFragmentsCount: 1,
			totalFragments:         1,
		}
		mockDbConn.On("Delete", mock.Anything, mock.Anything).Return(nil)

		errs := testDbService.removeFileFragments(mockFileContainer)
		assert.Empty(t, errs)
		mockDbConn.AssertCalled(t, "Delete", mock.Anything, mock.Anything)
	})
}
