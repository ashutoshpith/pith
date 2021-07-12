package main

import (
	"fmt"

	"github.com/ashutoshpith/api"
	rediskv "github.com/ashutoshpith/redisKv"
)



func main(){
	redisKp := rediskv.RedisDrive{
		Api: api.Api{
		  Addr: "localhost:6379",
		  Password: "",
		  Db: 0,
		},
	}
	p := redisKp.Ping()
	fmt.Println(p)
	redisKp.Set("kp", "oj")
	k := redisKp.Get("kp")
	fmt.Println(k)
	
	
}