package must

import "log"

func Do(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
