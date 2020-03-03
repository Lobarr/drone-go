package core

import "fmt"

//fileContainer holds file fragment keys
type fileContainer struct {
	fileName               string
	transactionID          string
	receivedFragmentsCount int
	totalFragments         int
}

type fileContainerInterface interface {
	addFragment()
	generateKey(int32) string
	getFileName() string
	getReceivedFragmentsCount() int
	getTotalFragments() int
	isComplete() bool
}

func newFileContainer(fileName string, transactionID string, totalFragments int) fileContainerInterface {
	return &fileContainer{
		fileName:       fileName,
		transactionID:  transactionID,
		totalFragments: totalFragments,
	}
}

func (fc fileContainer) generateKey(index int32) string {
	return fmt.Sprintf("%s:%s:%d:", fc.transactionID, fc.fileName, index)
}

func (fc *fileContainer) addFragment() {
	fc.receivedFragmentsCount++
}

func (fc fileContainer) isComplete() bool {
	return fc.totalFragments == fc.receivedFragmentsCount
}

func (fc fileContainer) getFileName() string {
	return fc.fileName
}

func (fc fileContainer) getTotalFragments() int {
	return fc.totalFragments
}

func (fc fileContainer) getReceivedFragmentsCount() int {
	return fc.receivedFragmentsCount
}
