package errorPith

import (
	"log"
)


func Errs(err error){
	if err != nil {
		log.Fatalln(err)
	}
}