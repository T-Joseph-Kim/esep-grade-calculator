package esepunittests

type GradeCalculator struct {
	grades []Grade
}

type GradeType int

const (
	Assignment GradeType = iota
	Exam
	Essay
)

var gradeTypeName = map[GradeType]string{
	Assignment: "assignment",
	Exam:       "exam",
	Essay:      "essay",
}

func (gt GradeType) String() string {
	return gradeTypeName[gt]
}

type Grade struct {
	Name  string
	Grade int
	Type  GradeType
}

func NewGradeCalculator() *GradeCalculator {
	return &GradeCalculator{
		grades: make([]Grade, 0),
	}
}

func (gc *GradeCalculator) GetFinalGrade() string {
	numericalGrade := gc.calculateNumericalGrade()

	switch {
	case numericalGrade >= 90:
		return "A"
	case numericalGrade >= 80:
		return "B"
	case numericalGrade >= 70:
		return "C"
	case numericalGrade >= 60:
		return "D"
	default:
		return "F"
	}
}

func (gc *GradeCalculator) AddGrade(name string, grade int, gradeType GradeType) {
	gc.grades = append(gc.grades, Grade{
		Name:  name,
		Grade: grade,
		Type:  gradeType,
	})
}

func (gc *GradeCalculator) calculateNumericalGrade() int {
	assignmentAvg := computeAverage(gc.getGradesByType(Assignment))
	examAvg := computeAverage(gc.getGradesByType(Exam))
	essayAvg := computeAverage(gc.getGradesByType(Essay))

	weightedGrade := float64(assignmentAvg)*0.5 +
		float64(examAvg)*0.35 +
		float64(essayAvg)*0.15

	return int(weightedGrade)
}

func (gc *GradeCalculator) getGradesByType(gt GradeType) []Grade {
	var filtered []Grade
	for _, g := range gc.grades {
		if g.Type == gt {
			filtered = append(filtered, g)
		}
	}
	return filtered
}

func computeAverage(grades []Grade) int {
	if len(grades) == 0 {
		return 0
	}

	sum := 0
	for _, g := range grades {
		sum += g.Grade
	}
	return sum / len(grades)
}
