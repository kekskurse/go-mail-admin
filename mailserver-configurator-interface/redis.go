package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
	"github.com/rs/zerolog/log"
)

type redisConnection struct {
	pool   *redis.Pool
	config Config
}

func newRedisConnection(config Config) redisConnection {
	r := redisConnection{config: config}
	r.initPool()
	r.ping()
	return r
}

func (r *redisConnection) initPool() {

	r.pool = &redis.Pool{
		MaxIdle:   80,
		MaxActive: 12000,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial(r.config.RedisNetwork, r.config.RedisAddress)
			if err != nil {
				log.Fatal().Err(err).Msg("Fail init redis pool")
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
		log.Fatal().Err(err).Msg("Fail init redis pool")
	}
}

func (r *redisConnection) set(key string, val string, ttl int) {
	conn := r.pool.Get()
	defer conn.Close()
	_, err := conn.Do("SET", key, val, "EX", ttl)
	if err != nil {
		log.Fatal().Err(err).Msg(fmt.Sprintf("Failed to set key >%s< with value >%s<", key, val))
	}
}

func (r *redisConnection) get(key string) (string, error) {
	// get conn and put back when exit from method
	conn := r.pool.Get()
	defer conn.Close()

	s, err := redis.String(conn.Do("GET", key))
	if err != nil {
		log.Warn().Err(err).Msg(fmt.Sprintf("Failed to get redis key >%s<", key))
		return "", err
	}

	return s, nil
}

func (r *redisConnection) delete(key string) error {
	// get conn and put back when exit from method
	conn := r.pool.Get()
	defer conn.Close()

	_, err := redis.String(conn.Do("DEL", key))
	if err != nil {
		log.Warn().Err(err).Msg(fmt.Sprintf("Failed to delete redis key >%s<", key))
		return err
	}

	return nil
}
