package util

import "io"

type ProgressReader struct {
	io.Reader
	Progress int64
}

func (pr *ProgressReader) Read(p []byte) (n int, err error) {
	n, err = pr.Reader.Read(p)
	pr.Progress += int64(n)
	return
}

func NewProgressReader(reader io.Reader) *ProgressReader {
	return &ProgressReader{
		Reader: reader,
	}
}
