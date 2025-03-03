package csv

type Parser interface {
	ParseFrom(interface{}) (CsvRecords, error)
}
