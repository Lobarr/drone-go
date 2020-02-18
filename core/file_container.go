package core

//fileContainer holds file fragment keys
type fileContainer struct {
	fileName       string
	fragmentIDs    []string
	totalFragments int
}

type fileContainerInterface interface {
	addFragment(string)
	isComplete() bool
	getFileName() string
	getFragmentIDs() []string
	getTotalFragments() int
}

func newFileContainer(fileName string, totalFragments int) fileContainerInterface {
	return &fileContainer{
		fileName:       fileName,
		fragmentIDs:    []string{},
		totalFragments: totalFragments,
	}
}

func (fc *fileContainer) addFragment(fragmentID string) {
	fc.fragmentIDs = append(fc.fragmentIDs, fragmentID)
}

func (fc fileContainer) isComplete() bool {
	return fc.totalFragments == len(fc.fragmentIDs)
}

func (fc fileContainer) getFileName() string {
	return fc.fileName
}

func (fc fileContainer) getFragmentIDs() []string {
	return fc.fragmentIDs
}

func (fc fileContainer) getTotalFragments() int {
	return fc.totalFragments
}
