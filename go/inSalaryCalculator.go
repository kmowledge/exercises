package main

import "fmt"

type EmployeeContract struct {
	PersonalID    int
	Renumeration  float64
	TaxPercentage float64
}

type FreelancerContract struct {
	PersonalID    int
	HourlyRate    float64
	HoursWorked   int
	TaxPercentage float64
}

func (ec EmployeeContract) CalculateSalary() float64 {
	return ec.Renumeration - ec.Renumeration*ec.TaxPercentage
}

func (fc FreelancerContract) CalculateSalary() float64 {
	return float64(fc.HoursWorked)*fc.HourlyRate - fc.TaxPercentage*float64(fc.HoursWorked)*fc.HourlyRate
}

type SalaryCalculator interface {
	CalculateSalary() float64 //Pamiętaj: wpisz co zwraca metoda.
}

func main() {
	panZdzislaw := EmployeeContract{
		PersonalID: 770312456451, Renumeration: 4100, TaxPercentage: 0.225,
	}
	fmt.Printf("pan Zdzislaw's salary: %.2f PLN\n", panZdzislaw.CalculateSalary())

	bryan := FreelancerContract{
		PersonalID: 92113076853, HoursWorked: 100, HourlyRate: 150, TaxPercentage: 0.12,
	}
	fmt.Printf("Bryan's salary: %.2f PLN\n", bryan.CalculateSalary())

	// var i SalaryCalculator
	employee := []SalaryCalculator{panZdzislaw, bryan}

	fmt.Printf("Monthly salaries expense: %.2f PLN\n", salaryExpense(employee))
}

func salaryExpense(salaries []SalaryCalculator) float64 {
	totalExpense := 0.0
	for _, v := range salaries {
		totalExpense += v.CalculateSalary()
	}
	return totalExpense
}
