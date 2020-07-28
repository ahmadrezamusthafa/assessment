package cache

import (
	"github.com/ahmadrezamusthafa/assessment/common/errors"
	"github.com/ahmadrezamusthafa/assessment/common/logger"
	"github.com/ahmadrezamusthafa/assessment/config"
	"github.com/gomodule/redigo/redis"
	"time"
)

type AssessmentCache struct {
	Config config.Config `inject:"config"`
	RedisCache
}

func (sc *AssessmentCache) StartUp() {
	logger.Info("Initiating assessment redis cache... ")
	sc.startUpReaders()
	logger.Info("Connected to cache reader host %v", sc.Config.CacheReaderHost)

	sc.startUpWriters()
	logger.Info("Connected to cache writer host %v", sc.Config.CacheWriterHost)
}

func (sc *AssessmentCache) Shutdown() {
	logger.Info("Closing cache writer connection")
	sc.writerPool.Close()

	logger.Info("Closing cache reader connection")
	sc.readerPool.Close()
}

func (sc *AssessmentCache) startUpWriters() {
	url := sc.Config.CacheWriterHost
	pool := &redis.Pool{
		MaxIdle:     sc.Config.CacheMaxIdle,
		IdleTimeout: sc.Config.CacheIdleTimeout,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", url)
			if err != nil {
				logger.Err("Failed to create redis connection. \nErrors: %v", err)
				return nil, errors.AddTrace(err)
			}
			return conn, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			if err != nil {
				logger.Err("Failed to ping redis. \nErrors: %v", err)
			}
			return errors.AddTrace(err)
		},
	}
	sc.writerPool = pool
}

func (sc *AssessmentCache) startUpReaders() {
	url := sc.Config.CacheReaderHost
	pool := &redis.Pool{
		MaxIdle:     sc.Config.CacheMaxIdle,
		IdleTimeout: sc.Config.CacheIdleTimeout,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", url)
			if err != nil {
				logger.Err("Failed to create redis connection. \nErrors: %v", err)
			}
			return conn, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			if err != nil {
				logger.Err("Failed to ping redis. \nErrors: %v", err)
			}
			return err
		},
	}
	sc.readerPool = pool
}
