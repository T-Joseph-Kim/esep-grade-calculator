package esepunittests

import "testing"

func TestGetGradeA(t *testing.T) {
	expected_value := "A"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 100, Assignment)
	gradeCalculator.AddGrade("exam 1", 100, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 100, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeB(t *testing.T) {
	expected_value := "B"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 80, Assignment)
	gradeCalculator.AddGrade("exam 1", 81, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 85, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeF(t *testing.T) {
	expected_value := "F"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 50, Assignment)
	gradeCalculator.AddGrade("exam 1", 45, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 35, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGradeTypeString(t *testing.T) {
	if Assignment.String() != "assignment" {
		t.Errorf("Expected 'assignment', got '%s'", Assignment.String())
	}
	if Exam.String() != "exam" {
		t.Errorf("Expected 'exam', got '%s'", Exam.String())
	}
	if Essay.String() != "essay" {
		t.Errorf("Expected 'essay', got '%s'", Essay.String())
	}
}

func TestAddEssayGrade(t *testing.T) {
	gradeCalculator := NewGradeCalculator()
	gradeCalculator.AddGrade("essay on testing", 70, Essay)

	if len(gradeCalculator.essays) != 1 {
		t.Errorf("Expected 1 essay grade, got %d", len(gradeCalculator.essays))
	}
	if gradeCalculator.essays[0].Grade != 70 {
		t.Errorf("Expected essay grade 70, got %d", gradeCalculator.essays[0].Grade)
	}
}

func TestEmptyCategory(t *testing.T) {
	gradeCalculator := NewGradeCalculator()
	gradeCalculator.AddGrade("exam 1", 80, Exam)

	finalGrade := gradeCalculator.GetFinalGrade()
	if finalGrade == "" {
		t.Errorf("Expected a valid grade, got empty string")
	}
}

func TestGetGradeC(t *testing.T) {
	expected_value := "C"

	gradeCalculator := NewGradeCalculator()
	gradeCalculator.AddGrade("assignment 1", 70, Assignment)
	gradeCalculator.AddGrade("exam 1", 72, Exam)
	gradeCalculator.AddGrade("essay 1", 75, Essay)

	actual_value := gradeCalculator.GetFinalGrade()
	if expected_value != actual_value {
		t.Errorf("Expected '%s'; got '%s'", expected_value, actual_value)
	}
}

func TestGetGradeD(t *testing.T) {
	expected_value := "D"

	gradeCalculator := NewGradeCalculator()
	gradeCalculator.AddGrade("assignment 1", 65, Assignment)
	gradeCalculator.AddGrade("exam 1", 62, Exam)
	gradeCalculator.AddGrade("essay 1", 60, Essay)

	actual_value := gradeCalculator.GetFinalGrade()
	if expected_value != actual_value {
		t.Errorf("Expected '%s'; got '%s'", expected_value, actual_value)
	}
}
