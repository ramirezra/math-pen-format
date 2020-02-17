package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

const (
	exitFail = 1
)

func main() {
	if err := run(os.Args, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(exitFail)
	}
}

// Entry struct defined
type Entry struct {
	SubmissionDate string
	Students       string
	FirstName      string
	LastName       string
	Phone          string
	Volunteering   string
	Clearances     string
}

func run(args []string, stdout io.Writer) error {
	if len(args) < 2 {
		// return errors.New("no file")
	}
	for _, file := range args[1:] {
		fmt.Fprintf(stdout, "File name: %s", file)
	}

	// Read csv file
	file, err := os.Open(args[1])
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 01
	record, err := reader.ReadAll()
	if err != nil {
		log.Fatalln(err)
	}

	var entries []Entry
	for _, item := range record {
		entry := Entry{
			SubmissionDate: item[0],
			Students:       item[1],
			FirstName:      item[2],
			LastName:       item[3],
			Phone:          item[4],
			Volunteering:   item[5],
			Clearances:     item[6],
		}
		entries = append(entries, entry)
	}
	fmt.Println(entries[0].Students)

	// Return no error to main
	return nil
}

// Submission Date,Please fill out for each of the students you have registered in Math Pentathlon.,First Name,Last Name,Parent/Emergency Contact Phone Number,Will parent/contact volunteer during Game Day?,Does parent/contact have clearances?
// 2020-02-16 9:58:28,"First Name: Name1, Last Name: Name2, Participating in Game Day?: Yes; First Name: Brother1, Last Name: Name2, Participating in Game Day?: No;",Parent,Best,(111) 123456y,Yes,No
// 2020-02-16 9:30:00,"First Name: Test, Last Name: TestLastName, Participating in Game Day?: Yes;",Parent,TestLastName,(234) 12345678,Yes,Yes
