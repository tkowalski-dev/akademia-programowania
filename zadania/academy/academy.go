package academy

import "math"

type Student struct {
	Name      string
	Grades    []int
	Project   int
	Attendace []bool
}

// AverageGrade returns an average grade given a
// slice containing all grades received during a
// semester, rounded to the nearest integer.
func AverageGrade(grades []int) int {
	// panic("not implemented")
	if len(grades) == 0 {
		return 0
	}
	var sum int
	for _, v := range grades {
		sum += v
	}
	return int(math.Round(float64(sum) / float64(len(grades))))
}

// AttendancePercentage returns a percentage of class
// attendance, given a slice containing information
// whether a student was present (true) of absent (false).
//
// The percentage of attendance is represented as a
// floating-point number ranging from  0 to 1,
// with 2 digits of precision.
func AttendancePercentage(attendance []bool) float64 {
	// panic("not implemented")
	y, n, s := 0, 0, 0
	for _, v := range attendance {
		if v {
			y++
		} else {
			n++
		}
		s++
	}
	if s == 0 {
		return 0
	}
	return float64(y) / float64(s)
}

// FinalGrade returns a final grade achieved by a student,
// ranging from 1 to 5.
//
// The final grade is calculated as the average of a project grade
// and an average grade from the semester, with adjustments based
// on the student's attendance. The final grade is rounded
// to the nearest integer.

// If the student's attendance is below 80%, the final grade is
// decreased by 1. If the student's attendance is below 60%, average
// grade is 1 or project grade is 1, the final grade is 1.
func FinalGrade(s Student) int {
	// panic("not implemented")
	avGrades := float64(AverageGrade(s.Grades))
	avgOut := (avGrades + float64(s.Project)) / 2
	atte := AttendancePercentage(s.Attendace)
	switch {
	case atte < 0.6 || avGrades == 1 || s.Project == 1:
		avgOut = 1
	case atte < 0.8:
		avgOut--
	}
	return int(math.Round(avgOut))
}

// GradeStudents returns a map of final grades for a given slice of
// Student structs. The key is a student's name and the value is a
// final grade.
func GradeStudents(students []Student) map[string]uint8 {
	// panic("not implemented")
	ret := make(map[string]uint8, 1)
	for _, s := range students {
		ret[s.Name] = uint8(FinalGrade(s))
	}
	return ret
}
