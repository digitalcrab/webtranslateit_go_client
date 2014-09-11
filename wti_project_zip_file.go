package webtranslateit_go_client

import (
	"archive/zip"
	"io/ioutil"
	"bytes"
	"fmt"
	"os"
)

type ProjectZipFile []byte

func (self ProjectZipFile) Size() int {
	return len(self)
}

func (self ProjectZipFile) SaveToPath(path string) error {
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("wti: error while creating project zip file %q %v", path, err)
	}
	defer file.Close()
	// Save
	return self.SaveToFile(file)
}

func (self ProjectZipFile) SaveToFile(file *os.File) error {
	if _, err := file.Write([]byte(self)); err != nil {
		return fmt.Errorf("wti: error while writing file %q %v", file.Name(), err)
	}
	return nil
}

func (self ProjectZipFile) BytesReader() *bytes.Reader {
	return bytes.NewReader([]byte(self))
}

func (self ProjectZipFile) ZipReader() (*zip.Reader, error) {
	return zip.NewReader(self.BytesReader(), int64(len(self)))
}

func (self ProjectZipFile) Extract() (map[string][]byte, error) {
	reader, err := self.ZipReader()
	if err != nil {
		return nil, fmt.Errorf("wti: error while extracting zip file %v", err)
	}

	res := make(map[string][]byte)

	for _, f := range reader.File {
		rc, err := f.Open()
		if err != nil {
			return nil, fmt.Errorf("wti: error while extracting %q from zip file %v", f.Name, err)
		}

		buff, err := ioutil.ReadAll(rc)
		if err != nil {
			rc.Close()
			return nil, fmt.Errorf("wti: error while reading data of %q from zip file %v", f.Name, err)
		}

		res[f.Name] = buff
		rc.Close()
	}

	return res, nil
}
