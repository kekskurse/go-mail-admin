package main

import (
	"github.com/gomodule/redigo/redis"
	"log"
	"os"
)
type redisConnection struct {
	pool *redis.Pool
}

func newRedisConnection() redisConnection {
	r := redisConnection{}
	r.initPool()
	r.ping()
	return r
}

func (r *redisConnection) initPool() {
	r.pool = &redis.Pool{
		MaxIdle:   80,
		MaxActive: 12000,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial(getConfigVariableWithDefault("REDIS_NETWORK", "tcp"), getConfigVariableWithDefault("REDIS_ADDRESS", "localhost:6379"))
			if err != nil {
				log.Printf("ERROR: fail init redis pool: %s", err.Error())
				os.Exit(1)
			}
			return conn, err
		},
	}
}

func (r *redisConnection) ping() {
	conn := r.pool.Get()
	defer conn.Close()
	_, err := redis.String(conn.Do("PING"))
	if err != nil {
		log.Printf("ERROR: fail ping redis conn: %s", err.Error())
		os.Exit(1)
	}
}

func (r *redisConnection) set(key string, val string, ttl int) {
	conn := r.pool.Get()
	defer conn.Close()
	_, err := conn.Do("SET", key, val, "EX", ttl)
	if err != nil {
		panic("ERROR: fail set key "+key+", val "+val+", error "+err.Error())
	}
}

func (r *redisConnection) get(key string) (string, error) {
	// get conn and put back when exit from method
	conn := r.pool.Get()
	defer conn.Close()

	s, err := redis.String(conn.Do("GET", key))
	if err != nil {
		log.Printf("ERROR: fail get key %s, error %s", key, err.Error())
		return "", err
	}

	return s, nil
}

func (r *redisConnection) delete(key string) (error) {
	// get conn and put back when exit from method
	conn := r.pool.Get()
	defer conn.Close()

	_, err := redis.String(conn.Do("DEL", key))
	if err != nil {
		log.Printf("ERROR: fail get key %s, error %s", key, err.Error())
		return  err
	}

	return nil
}