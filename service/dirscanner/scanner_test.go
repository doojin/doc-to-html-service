package dirscanner

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/doojin/doc-to-html-service/test"
)

func Test_IsDoc_shouldReturnTrueIfFileHasOneOfValidDocExtensions(t *testing.T) {
	fileInfo := new(test.FileInfoMock)
	scanner := new(DirScanner)

	fileInfo.On("Name").Return("dummy-name.doc")
	assert.True(t, scanner.IsDoc(fileInfo))

	fileInfo.On("Name").Return("dummy-name.docx")
	assert.True(t, scanner.IsDoc(fileInfo))
}

func Test_IsDoc_shouldReturnTFalseIfFileHasNotValidDocExtension(t *testing.T) {
	fileInfo := new(test.FileInfoMock)
	scanner := new(DirScanner)

	fileInfo.On("Name").Return("dummy-name.dummy-extension")
	assert.False(t, scanner.IsDoc(fileInfo))
}