package utility

import "strings"


func GetName(value string) string {
	data := strings.Split(value, "/")
    name := data[len(data) -1]
	return name
}