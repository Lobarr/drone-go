package core

//fileContainer holds file fragment keys
type fileContainer struct {
	fileName       string
	fragmentIDs    []string
	totalFragments int
}

func newFileContainer(fileName string, totalFragments int) *fileContainer {
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
