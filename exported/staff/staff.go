package staff

import "log"

var OverPaidLimit = 75000  // Exported
var underPaidLimit = 20000 // Not exported

type Employee struct {
	FirstName string
	LastName  string
	Salary    int
	FullTime  bool
}

type Office struct {
	AllStaff []Employee
}

func (e *Office) All() []Employee {
	return e.AllStaff
}

func (e *Office) OverPaid() []Employee {
	var overpaid []Employee
	myCustomFunction()
	for _, emp := range e.AllStaff {
		if emp.Salary > OverPaidLimit {
			overpaid = append(overpaid, emp)
		}
	}
	return overpaid
}

func (e *Office) UnderPaid() []Employee {
	var underpaid []Employee
	myCustomFunction()
	for _, emp := range e.AllStaff {
		if emp.Salary < underPaidLimit {
			underpaid = append(underpaid, emp)
		}
	}
	return underpaid
}

func (e *Office) notVisible() {
	log.Println("Not visible function called")
}

func myCustomFunction() {
	log.Println("I'm inside custom function")
}
