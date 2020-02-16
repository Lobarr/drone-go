package core

import (
	"github.com/kpango/glg"
	"github.com/syndtr/goleveldb/leveldb"
)

//DBService interacts with db
type dbService struct {
	conn *leveldb.DB
}

//b=newDBService create a new db service
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

	glg.Debugf("[DBService] Putting file fragment %s of %s", fragmentID, fileFragment.GetFileName())
	return db.conn.Put(
		[]byte(fragmentID),
		fileFragment.GetFragmentContent(),
		nil,
	)

}

func (db dbService) getFileFragmentContent(fragmentID string) ([]byte, error) {
	glg.Debugf("[DBService] Getting file fragment %s", fragmentID)
	return db.conn.Get([]byte(fragmentID), nil)
}

func (db dbService) removeFileFragments(fragmentIDs ...string) []error {
	glg.Debugf("[DBService] Removing %d file fragments", len(fragmentIDs))
	errs := []error{}
	for _, fragmentID := range fragmentIDs {
		if err := db.conn.Delete([]byte(fragmentID), nil); err != nil {
			errs = append(errs, err)
		}
	}
	return errs
}
