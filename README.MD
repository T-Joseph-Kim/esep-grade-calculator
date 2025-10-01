# ESEP Grade Calculator

This library is a simple utility to calculate your grades.

## Tech Stack
- **Language:** Go (Golang) 1.24+  
- **Dependencies:** Go Modules (`go.mod`)  
- **Testing:** Built-in Go `testing` package

## File Overview
- **grade_calculator.go** – Core implementation of the grade calculator. Defines the `GradeCalculator` and `Grade` structs, manages assignments/exams/essays, and computes weighted averages and letter grades.  
- **grade_calculator_test.go** – Unit tests for `GetFinalGrade`. Includes cases for A, B, and F grades.  
- **go.mod** – Go module definition (`esep/grade-calculator`) specifying Go 1.24.4, with no external dependencies.  
- **README.md** – Project documentation, requirements, and usage instructions.  

## Tests
Tests are included in `grade_calculator_test.go`. They verify that `GradeCalculator` computes the correct final letter grade given different set of grades.

Runtime requirements:
Golang 1.24+

Run unit tests:
```
cd grade_calculator/
go test ./...
```

Run unit tests with code coverage:
```
cd grade_calculator/
go test -coverprofile=coverage.txt ./...
```