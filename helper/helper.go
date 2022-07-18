package helper

import "log"

func CheckError(err error) {
	if err != nil {
		log.Fatalln("Error: ", err.Error())
	}
}
