package csv

// NewReader creates a new CSV reader instance
func NewReader() Reader {
	return newV1Reader()
}

// NewWriter creates a new CSV writer instance
func NewWriter() Writer {
	return newV1Writer()
}
