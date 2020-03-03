package core

import (
	"fmt"

	"github.com/kpango/glg"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
)

var dbServiceLogTemplate = "[DBService] %s"

type dbConnection interface {
	Put([]byte, []byte, *opt.WriteOptions) error
	Get([]byte, *opt.ReadOptions) ([]byte, error)
	Delete([]byte, *opt.WriteOptions) error
	Close() error
}

//dBService interacts with db
type dbService struct {
	conn dbConnection
}

type dbServiceInteface interface {
	putFileFragmentContent(string, *FileFragment) error
	getFileFragmentContent(string) ([]byte, error)
	removeFileFragments(fileContainerInterface) []error
	close()
}

//newDBService create a new db service
func newDBService(dbFilePath string) (*dbService, error) {
	db, err := leveldb.OpenFile(dbFilePath, nil)

	if err != nil {
		return nil, err
	}

	return &dbService{
		conn: db,
	}, nil
}

func (db dbService) putFileFragmentContent(fragmentID string, fileFragment *FileFragment) error {
	glg.Get().Debugf(dbServiceLogTemplate, fmt.Sprintf("Putting file fragment %s of %s", fragmentID, fileFragment.GetFileName()))
	fragmentKey := []byte(fragmentID)
	return db.conn.Put(
		fragmentKey,
		fileFragment.GetFragmentContent(),
		nil,
	)
}

func (db dbService) getFileFragmentContent(fragmentID string) ([]byte, error) {
	glg.Get().Debugf(dbServiceLogTemplate, fmt.Sprintf("Getting file fragment %s", fragmentID))
	fragmentKey := []byte(fragmentID)
	return db.conn.Get(fragmentKey, nil)
}

func (db dbService) removeFileFragments(fileContainer fileContainerInterface) []error {
	glg.Get().Debugf(dbServiceLogTemplate, fmt.Sprintf("Removing %d file fragments", fileContainer.getTotalFragments()))
	errs := []error{}
	for fragmentID := 0; fragmentID < fileContainer.getTotalFragments(); fragmentID++ {
		fragmentKey := []byte(fileContainer.generateKey(int32(fragmentID)))
		if err := db.conn.Delete(fragmentKey, nil); err != nil {
			errs = append(errs, err)
		}
	}
	return errs
}

func (db dbService) close() {
	db.conn.Close()
}
