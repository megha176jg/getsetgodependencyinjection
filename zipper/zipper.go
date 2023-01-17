package zipper

import (
	"archive/zip"
	"bytes"
	"io"
	"os"
	"path/filepath"
	"strings"

	"bitbucket.org/junglee_games/getsetgo/files"
)

type Zipper struct {
	buffer *bytes.Buffer
	writer *zip.Writer
}

type FilesData map[string]string

func NewZipper() *Zipper {
	buffer := new(bytes.Buffer)
	zipWriter := zip.NewWriter(buffer)
	return &Zipper{buffer: buffer, writer: zipWriter}
}

func (z *Zipper) AddFile(name string, data string) error {
	w1, err := z.writer.Create(name)
	if err != nil {
		return err
	}
	iow := strings.NewReader(data)
	if _, err := io.Copy(w1, iow); err != nil {
		return err
	}
	return nil
}

func (z *Zipper) AddFiles(data FilesData) error {
	for name, data := range data {
		w1, err := z.writer.Create(name)
		if err != nil {
			return err
		}
		iow := strings.NewReader(data)
		if _, err := io.Copy(w1, iow); err != nil {
			return err
		}
	}
	return nil
}

// please close the zipper before calling this method
func (z *Zipper) GetString() string {
	return z.buffer.String()
}

// please close the zipper before calling this method
func (z *Zipper) SaveFile(filename string) error {
	return files.SaveFile(".", filename, z.buffer.String())
}

func (z *Zipper) Close() {
	z.writer.Close()
}

func unzip(src, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer rc.Close()

		path := filepath.Join(dest, f.Name)
		if f.FileInfo().IsDir() {
			os.MkdirAll(path, f.Mode())
		} else {
			f, err := os.OpenFile(
				path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer f.Close()

			_, err = io.Copy(f, rc)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func Unzipper(data []byte) ([]*os.File, error) {
	var files []*os.File

	// Create a new zip reader
	reader, err := zip.NewReader(bytes.NewReader(data), int64(len(data)))
	if err != nil {
		return nil, err
	}

	// Iterate through the files in the zip archive
	for _, file := range reader.File {
		rc, err := file.Open()
		if err != nil {
			return nil, err
		}
		defer rc.Close()

		path := file.Name
		// Create the file on disk
		f, err := os.Create(path)
		if err != nil {
			return nil, err
		}
		defer f.Close()

		// Copy the file data
		_, err = io.Copy(f, rc)
		if err != nil {
			return nil, err
		}

		files = append(files, f)
	}

	return files, nil

}

// for reference

// func main() {
// 	zipper, err := NewZipper("myzip.zip")
// 	if err != nil {
// 		panic(err)
// 	}
// 	zipper.AddFile("test/pavan.txt", "Hello there !!!")
// 	zipper.AddFile("test/pavan1.txt", "Hello there !!!")
// 	zipper.Close()
// }
