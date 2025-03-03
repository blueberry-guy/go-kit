package example

// Person represents a record in the CSV file
type Person struct {
	Name   string  `csv:"name"`
	Age    int     `csv:"age"`
	City   string  `csv:"city"`
	Salary float64 `csv:"salary"`
}

// Employee represents a record with different field names
type Employee struct {
	FullName   string  `csv:"full_name"`
	YearsOld   int     `csv:"years_old"`
	Location   string  `csv:"location"`
	AnnualWage float64 `csv:"annual_wage"`
}
