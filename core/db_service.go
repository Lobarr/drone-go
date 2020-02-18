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
	removeFileFragments(...string) []error
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
	glg.Debugf(dbServiceLogTemplate, fmt.Sprintf("Putting file fragment %s of %s", fragmentID, fileFragment.GetFileName()))
	return db.conn.Put(
		[]byte(fragmentID),
		fileFragment.GetFragmentContent(),
		nil,
	)
}

func (db dbService) getFileFragmentContent(fragmentID string) ([]byte, error) {
	glg.Debugf(dbServiceLogTemplate, fmt.Sprintf("Getting file fragment %s", fragmentID))
	return db.conn.Get([]byte(fragmentID), nil)
}

func (db dbService) removeFileFragments(fragmentIDs ...string) []error {
	glg.Debugf(dbServiceLogTemplate, fmt.Sprintf("Removing %d file fragments", len(fragmentIDs)))
	errs := []error{}
	for _, fragmentID := range fragmentIDs {
		if err := db.conn.Delete([]byte(fragmentID), nil); err != nil {
			errs = append(errs, err)
		}
	}
	return errs
}

func (db dbService) close() {
	db.conn.Close()
}
