package grade

type Student struct {
	Name   string
	Grades []float64
}

func IsPassing(grade float64) bool {
	return grade >= 60
}