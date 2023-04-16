package academy

import (
	"github.com/grupawp/akademia-programowania/Golang/zadania/academy2/mocks"
	"testing"
)

//func TestGradeYear(t *testing.T) {
//	t.Run("Empty slice", func(t *testing.T) {
//		grades := []int{}
//		student := Sophomore{grades: grades}
//		expected := 0
//		if result := student.averageGrade(); result != expected {
//			t.Errorf("Expected %d but got %d", expected, result)
//		}
//	})
//
//	//assert.Equal(t)
//
//}

func TestGradeStudent(t *testing.T) {
	t.Run("Invalid Grade", func(t *testing.T) {
		mock := mocks.NewRepository(t)
		mock.On("Get").Return(nil, nil)
		got1, _ := mock.Get("")
		if got1 != nil {
			t.Errorf("Blad")
		}
	})

}
