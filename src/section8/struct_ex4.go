package main

import "fmt"

type Employee struct {
	name   string
	salary float64
	bonus  float64
}

func (e Employee) Calculate() float64 {
	return e.salary + e.bonus
}

func (e Executive) Calculate() float64 {
	return e.salary + e.bonus + e.specialBonus
}

type Executive struct {
	Employee     // is a 관계
	specialBonus float64
}

func main() {

	ep1 := Employee{"kim", 5000, 500}
	ep2 := Employee{"lee", 6000, 600}
	ex := Executive{
		Employee{"park", 8000, 800},
		1000,
	}

	fmt.Println("Employee total:", ep1.Calculate())
	fmt.Println("Employee total:", ep2.Calculate())
	// Executive도 Employee를 포함하므로 Employee의 Calculate() 사용 가능 (오버라이딩 이전)
	// fmt.Println("Executive Total Compensation:", ex.Calculate()+ex.specialBonus)
	// Executive의 Calculate() (오버라이딩 이후)
	fmt.Println("Executive total(Overridden):", ex.Calculate())
	// 오버라이딩 이후 Employee의 Calculate() 사용하려면 아래와 같이 명시
	fmt.Println("Executive total:", ex.Employee.Calculate())
}
