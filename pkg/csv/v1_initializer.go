package csv

import "github.com/blueberry-guy/go-kit/pkg/csv/internal"

func newV1Reader() Reader {
	return &internal.GoCsvReaderImpl{}
}

func newV1Writer() Writer {
	return &internal.GoCsvWriterImpl{}
}
