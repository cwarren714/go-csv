package main

import (
	"fmt"
	"os"

	"github.com/gocarina/gocsv"
)

type People struct {
	Lastname  string `csv:"Lastname"`
	Firstname string `csv:"Firstname"`
	SSN       string `csv:"SSN"`
}

func readFile(filepath string, people *[]People) error {
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer file.Close()
	return gocsv.Unmarshal(file, people)
}

// findDifference returns a list of people who are in the old list but not in the new list
// currently using firstname and lastname as the key but can use whatever is in the CSV
func findDifference(people_old []People, people_new []People) []People {
	// Create a map to store the people in people_new
	newPeopleMap := make(map[string]bool)
	for _, person := range people_new {
		key := person.Firstname + " " + person.Lastname
		newPeopleMap[key] = true
	}

	var difference []People
	// Iterate over people_old and find elements not in people_new
	for _, person := range people_old {
		key := person.Firstname + " " + person.Lastname
		if !newPeopleMap[key] {
			difference = append(difference, person)
		}
	}

	return difference
}

func main() {
	people_old := []People{}
	people_new := []People{}

	err := readFile("people.csv", &people_old)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = readFile("people_compare.csv", &people_new)
	if err != nil {
		fmt.Println(err)
		return
	}

	difference := findDifference(people_old, people_new)
	fmt.Println(difference)
}
