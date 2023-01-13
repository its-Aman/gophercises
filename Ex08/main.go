package main

import (
	"fmt"

	"PHONE_NUMBER_NORMALIZER/must"
	"PHONE_NUMBER_NORMALIZER/normalize"
	"PHONE_NUMBER_NORMALIZER/phone_db"

	_ "github.com/lib/pq"
)

func main() {
	must.Do(phone_db.NewConn())
	defer phone_db.DB.Close()

	must.Do(phone_db.ResetAll())
	must.Do(phone_db.Seed())

	phoneNumbers, err := phone_db.GetAllPhoneNumbers()
	must.Do(err)

	for _, p := range phoneNumbers {
		fmt.Printf("Working on %+v\n", p)
		number := normalize.DoRegex(p.Number)

		if number != p.Number {
			existing, err := phone_db.FindPhone(number)
			must.Do(err)
			if existing != nil {
				fmt.Printf("Removing... %+v\n", existing)
				must.Do(phone_db.DeletePhoneNumber(p.Id))
			} else {
				p.Number = number
				fmt.Printf("Updating... %+v\n", p)
				must.Do(phone_db.UpdatePhoneNumber(p))
			}
		} else {
			fmt.Println("No changes required")
		}
	}

	fmt.Println("Done!")
}
