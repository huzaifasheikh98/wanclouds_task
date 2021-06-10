package main
import (
	"database/sql"
	"encoding/csv"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"strconv"
)
type Person struct {
	firstName  string
	lastName   string
	age        string
	blood_group string
}
func main() {
	records, err := readData("person.csv")
	if err != nil {
		log.Fatal(err)
	}
	insertRows(records)
}
func readData(person string) ([][]string, error)  {
	f,err := os.Open(person)
	if err != nil {
		return [][]string{}, err
	}
	defer f.Close()
	r := csv.NewReader(f)
	if _, err := r.Read(); err != nil {
		return [][]string{}, err
	}
	records, err := r.ReadAll()

	if err != nil {
		return [][]string{}, err
	}

	return records, nil
}
func insertRows(records [][]string) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/task")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	for _, record := range records {
		person := Person{
			firstName:   record[0],
			lastName:    record[1],
			age:         record[2],
			blood_group: record[3],
		}
		fmt.Printf("%s %s  %s %s\n", person.firstName, person.lastName,
			person.age, person.blood_group)
		age , err := strconv.Atoi(person.age)
		if err != nil {
			panic(err.Error())
		}
		insert, err := db.Query(fmt.Sprintf("INSERT INTO person(firstname, lastname, age, bloodgroup) VALUES ('%s','%s',%d,'%s');",person.firstName,person.lastName,age,person.blood_group))
		if err != nil {
			panic(err.Error())
		}
		defer insert.Close()
	}


	}
