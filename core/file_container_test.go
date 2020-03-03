package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFileContainer(t *testing.T) {
	mockFileName := "some-filename"
	mockTransactionID := "some-transaction-id"

	t.Run("should add fragment", func(t *testing.T) {
		mockTotalFragments := 1
		fileController := newFileContainer(mockFileName, mockTransactionID, mockTotalFragments)

		fileController.addFragment()
		assert.Greater(t, fileController.getReceivedFragmentsCount(), 0)
	})

	t.Run("should be complete", func(t *testing.T) {
		mockTotalFragments := 0
		fileContainer := newFileContainer(mockFileName, mockTransactionID, mockTotalFragments)

		assert.True(t, fileContainer.isComplete())
	})

	t.Run("should not be complete", func(t *testing.T) {
		mockTotalFragments := 1
		fileContainer := newFileContainer(mockFileName, mockTransactionID, mockTotalFragments)

		assert.False(t, fileContainer.isComplete())
	})
}
