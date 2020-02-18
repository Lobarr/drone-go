package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddFragment(t *testing.T) {
	mockFileName := "some-filename"
	mockFragmentID := "some-fragment-id"
	mockTotalFragments := 0
	fileController := newFileContainer(mockFileName, mockTotalFragments)

	fileController.addFragment(mockFragmentID)
	assert.Contains(t, fileController.getFragmentIDs(), mockFragmentID)
}

func TestFileContainer(t *testing.T) {
	mockFileName := "some-filename"

	t.Run("should add fragment", func(t *testing.T) {
		mockTotalFragments := 1
		fileController := newFileContainer(mockFileName, mockTotalFragments)
		mockFragmentID := "some-fragment-id"

		fileController.addFragment(mockFragmentID)

		assert.Contains(t, fileController.getFragmentIDs(), mockFragmentID)

	})

	t.Run("should be complete", func(t *testing.T) {
		mockTotalFragments := 0
		fileContainer := newFileContainer(mockFileName, mockTotalFragments)

		assert.True(t, fileContainer.isComplete())
	})

	t.Run("should not be complete", func(t *testing.T) {
		mockTotalFragments := 1
		fileContainer := newFileContainer(mockFileName, mockTotalFragments)

		assert.False(t, fileContainer.isComplete())
	})
}
