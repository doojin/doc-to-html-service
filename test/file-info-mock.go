package test

import (
	"github.com/stretchr/testify/mock"
	"os"
	"time"
)

type FileInfoMock struct {
	mock.Mock
}

func (m *FileInfoMock) Name() string {
	args := m.Called()
	return args.String(0)
}

func (m *FileInfoMock) Size() int64 {
	args := m.Called()
	return args.Get(0).(int64)
}

func (m *FileInfoMock) Mode() os.FileMode {
	args := m.Called()
	return args.Get(0).(os.FileMode)
}

func (m *FileInfoMock) ModTime() time.Time {
	args := m.Called()
	return args.Get(0).(time.Time)
}

func (m *FileInfoMock) IsDir() bool {
	args := m.Called()
	return args.Bool(0)
}

func (m *FileInfoMock) Sys() interface{} {
	args := m.Called()
	return args.Get(0)
}
