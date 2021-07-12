package rediskv

import (
	"github.com/ashutoshpith/api"
	errorPith "github.com/ashutoshpith/error"
	"github.com/go-redis/redis"
)

// var url = "localhost:6379"
type RedisDrive struct {
	Api api.Api
}
func (rd *RedisDrive) Connect() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: rd.Api.Addr,
		Password: rd.Api.Password,
		DB: rd.Api.Db,
	})
	return client
}

func (rd *RedisDrive) Ping() string {
	res, err := rd.Connect().Ping().Result()
	errorPith.Errs(err)
	return res
}

func (rd *RedisDrive) Set(key string , value interface{}) *redis.StatusCmd {
	res := rd.Connect().Set(key,value,0)
	return res

}

func (rd *RedisDrive) Get(key string) *redis.StringCmd {
	res := rd.Connect().Get(key)
	return res
}
