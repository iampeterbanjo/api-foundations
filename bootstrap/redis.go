package bootstrap

import (
	"time"

	"github.com/garyburd/redigo/redis"
)

var (
	connectionTimeout = redis.DialConnectTimeout(time.Second)
	readTimeout       = redis.DialReadTimeout(time.Second)
	writeTimeout      = redis.DialWriteTimeout(time.Second)
)

// GetRedis connection
func GetRedis() (redis.Conn, error) {
	return redis.Dial("tcp", "redis:6379", connectionTimeout, readTimeout, writeTimeout)
}
