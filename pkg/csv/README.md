# CSV Package

This package provides utilities for reading and writing CSV files in Go. It uses [gocsv](https://github.com/gocarina/gocsv) under the hood for efficient CSV handling.

## Installation

```bash
go get github.com/blueberry-guy/go-kit/csv
```

## Usage

### Reading CSV Files

```go
package main

import (
    "fmt"
    "github.com/blueberry-guy/go-kit/csv"
)

// Define your struct with csv tags
type Record struct {
    Name   string `csv:"name"`
    Age    int    `csv:"age"`
    City   string `csv:"city"`
}

func main() {
    // Create a reader
    reader := csv.NewReader(false) // Use default implementation

    // Prepare a slice to hold the records
    var records []Record

    // Read from CSV file
    err := reader.ReadFromFile("data.csv", &records)
    if err != nil {
        panic(err)
    }

    // Process the records
    for _, record := range records {
        fmt.Printf("Name: %s, Age: %d, City: %s\n", record.Name, record.Age, record.City)
    }
}
```

### Writing CSV Files

```go
package main

import (
    "github.com/blueberry-guy/go-kit/csv"
)

type Record struct {
    Name   string `csv:"name"`
    Age    int    `csv:"age"`
    City   string `csv:"city"`
}

func main() {
    // Create a writer
    writer := csv.NewWriter(false) // Use default implementation

    // Prepare some records
    records := []Record{
        {Name: "John Doe", Age: 30, City: "New York"},
        {Name: "Jane Smith", Age: 25, City: "Los Angeles"},
    }

    // Write to CSV file
    err := writer.WriteToFile(records, "output.csv")
    if err != nil {
        panic(err)
    }
}
```

## Features

- Simple and intuitive API
- Support for struct tags
- Automatic type conversion
- Error handling
- File-based operations

## Notes

- The package uses struct tags to map CSV columns to struct fields
- The `csv` tag is used to specify the column name
- Supports basic Go types (string, int, float64, etc.)
- Files are automatically opened and closed
- Thread-safe operations

## License

This project is licensed under the MIT License - see the LICENSE file for details. 