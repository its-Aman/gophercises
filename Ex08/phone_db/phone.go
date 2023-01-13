package phone_db

import (
	"database/sql"
	"fmt"
)

const (
	host     = "localhost"
	port     = 5432
	username = "postgres"
	password = "mysecretpassword"
	dbname   = "gophercises_phone"
)

var DB *sql.DB
var err error

type Phone struct {
	Id     int
	Number string
}

func NewConn() error {
	connString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, username, password, dbname)
	DB, err = sql.Open("postgres", connString)
	return err
}

func ResetAll() error {
	// err := resetDB(db)

	// if err != nil {
	// 	return err
	// }

	return resetPhoneNumberTable()
}

func Seed() error {
	numbers := []any{
		"1234567890",
		"123 456 7891",
		"(123) 456 7892",
		"(123) 456-7893",
		"123-456-7894",
		"123-456-7890",
		"1234567892",
		"(123)456-7892",
	}
	return insertPhoneNumbers(numbers)
}

func GetAllPhoneNumbers() ([]Phone, error) {
	rows, err := DB.Query(`SELECT id, value FROM phone_numbers`)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var ret []Phone

	for rows.Next() {
		var p Phone

		if err = rows.Scan(&p.Id, &p.Number); err != nil {
			return nil, err
		}
		ret = append(ret, p)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return ret, nil
}

func FindPhone(number string) (*Phone, error) {
	var p Phone
	row := DB.QueryRow(`SELECT * FROM phone_numbers WHERE value=$1`, number)
	err := row.Scan(&p.Id, &p.Number)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} else {
			return nil, err
		}
	}

	return &p, nil
}

func UpdatePhoneNumber(existing Phone) error {
	q := `UPDATE phone_numbers SET value=$2 WHERE id=$1`
	_, err := DB.Exec(q, existing.Id, existing.Number)
	return err
}

func DeletePhoneNumber(id int) error {
	q := `DELETE FROM phone_numbers WHERE id=$1`
	_, err := DB.Exec(q, id)
	return err
}

func GetPhone(id int) (string, error) {
	var number string
	q := `SELECT value FROM phone_numbers WHERE id=$1`
	row := DB.QueryRow(q, id)
	err := row.Scan(&number)

	if err != nil {
		return "", err
	}
	return number, nil
}

func insertPhoneNumbers(numbers []any) error {
	q := `INSERT INTO phone_numbers (value) VALUES `

	for i := range numbers {
		q += fmt.Sprintf(`($%d)`, i+1)
		if i < len(numbers)-1 {
			q += ","
		}
	}

	_, err := DB.Exec(q, numbers...)

	if err != nil {
		return err
	}

	return nil
}

func resetDB() error {
	_, err := DB.Exec("DROP DATABASE " + dbname)
	if err != nil {
		return err
	}
	return createDB()
}

func createDB() error {
	_, err := DB.Exec("CREATE DATABASE " + dbname)
	return err
}

func resetPhoneNumberTable() error {
	q := `DROP TABLE phone_numbers`

	_, err := DB.Exec(q)

	if err != nil {
		return err
	}

	return createPhoneNumberTable()
}

func createPhoneNumberTable() error {
	q := `
		CREATE TABLE phone_numbers 
		(
			id SERIAL,
			value VARCHAR(225)
		)
	`

	_, err := DB.Exec(q)
	return err
}
