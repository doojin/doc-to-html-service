package dirscanner

import (
	"io/ioutil"
	"path"
	"os"
)

// DirScannerInterface is an interface for DirScanner
type DirScannerInterface interface {
	ScanDir(string) []string
}

// DirScanner is scanner which finds .doc / .docx files
// within the directory
type DirScanner struct {}

// ScanDir returns a slice of .doc / .docx file names from directory
func (scanner DirScanner) ScanDir(dir string) (files []string) {
	items, _ := ioutil.ReadDir(dir)
	for _, item := range items {
		if !item.IsDir() && scanner.IsDoc(item) {
			files = append(files, dir + "/" + item.Name())
		}
	}
	return
}

// IsDoc returns true if file has valid DOC extension, otherwise returns false
func (scanner DirScanner) IsDoc(file os.FileInfo) bool {
	docExts := []string { ".doc", ".docx" }
	fileExt := path.Ext(file.Name())
	for _, ext := range docExts {
		if fileExt == ext {
			return true
		}
	}
	return false
}