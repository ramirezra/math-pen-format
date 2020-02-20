package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

const (
	exitFail = 1
)

// Pseudocode:
// 1. Read file.
// 2. Extract studentscolumn. Discard other columns.
// 3. Transform students column into individual students (split by `;`) - one student per line
// 4. Transform individual student into separate columns "First Name", "Last Name", "Yes/No"
// 5. Write to New CSV File
func main() {
	if err := run(os.Args, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(exitFail)
	}
}

// Entry struct defined
type Entry struct {
	FirstName string
	LastName  string
	Attending string
}

func run(args []string, stdout io.Writer) error {
	if len(args) < 2 {
		return errors.New("no file")
	}
	for _, file := range args[1:] {
		fmt.Fprintf(stdout, "File name: %s\n", file)
	}

	students := readFile(args)
	writeFile(students)
	// Return no error to main
	return nil
}

// Submission Date,Please fill out for each of the students you have registered in Math Pentathlon.,First Name,Last Name,Parent/Emergency Contact Phone Number,Will parent/contact volunteer during Game Day?,Does parent/contact have clearances?
// 2020-02-16 9:58:28,"First Name: Name1, Last Name: Name2, Participating in Game Day?: Yes; First Name: Brother1, Last Name: Name2, Participating in Game Day?: No;",Parent,Best,(111) 123456y,Yes,No
// 2020-02-16 9:30:00,"First Name: Test, Last Name: TestLastName, Participating in Game Day?: Yes;",Parent,TestLastName,(234) 12345678,Yes,Yes

func readFile(args []string) []string {
	// Read csv file
	file, err := os.Open(args[1])
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	data, err := reader.ReadAll()
	if err != nil {
		log.Fatalln(err)
	}

	var students []string
	for i, item := range data {
		if i == 0 {
			continue
		} else {
			items := strings.Split(item[1], ";")
			for _, student := range items {
				students = append(students, strings.Trim(student, " "))
			}
		}
	}
	fmt.Println(students[1])
	return students
}

func writeFile(students []string) {
	csvFile, err := os.Create("result.csv")
	if err != nil {
		log.Fatalln(err)
	}
	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)
	// for _, student := range students {
	// 	line := student
	// 	err := writer.Write(line)
	// 	if err != nil {
	// 		log.Fatalln(err)
	// 	}
	// }
	err = writer.Write(students)
	if err != nil {
		log.Fatalln(err)
	}
	writer.Flush()
}
