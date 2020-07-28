package cache

import (
	"errors"
	"github.com/ahmadrezamusthafa/assessment/config"
	"github.com/gomodule/redigo/redis"
	"github.com/json-iterator/go"
	"reflect"
)

type RedisCache struct {
	Config     *config.Config `inject:"config"`
	writerPool *redis.Pool
	readerPool *redis.Pool
}

func (r *RedisCache) IsAvailable() bool {
	redisReader := r.readerPool.Get()
	defer redisReader.Close()
	_, err := redis.Bytes(redisReader.Do("PING"))
	if err != nil {
		return false
	}

	redisWriter := r.writerPool.Get()
	defer redisWriter.Close()
	_, err = redis.Bytes(redisWriter.Do("PING"))
	if err != nil {
		return false
	}
	return true
}

func (r *RedisCache) Get(key string) (string, error) {
	redisReader := r.readerPool.Get()
	defer redisReader.Close()
	resp, err := redis.String(redisReader.Do("GET", key))
	if err != nil {
		return "", err
	}
	return resp, err
}

func (r *RedisCache) GetObj(key string, dest interface{}) error {
	resp, err := r.Get(key)
	if err != nil {
		return err
	}
	return jsoniter.Unmarshal([]byte(resp), dest)
}

func (r *RedisCache) GetSlice(key string, dest interface{}) error {
	redisReader := r.readerPool.Get()
	defer redisReader.Close()
	resp, err := redisReader.Do("SMEMBERS", key)
	values, err := redis.Values(resp, err)
	if err != nil {
		return err
	}
	return redis.ScanSlice(values, dest)
}

func (r *RedisCache) GetSliceDiff(firstKey, secondKey string, dest interface{}) error {
	redisReader := r.readerPool.Get()
	defer redisReader.Close()
	resp, err := redisReader.Do("SDIFF", firstKey, secondKey)
	values, err := redis.Values(resp, err)
	if err != nil {
		return err
	}
	return redis.ScanSlice(values, dest)
}

func (r *RedisCache) IsSliceContain(key string, value interface{}) bool {
	redisReader := r.readerPool.Get()
	defer redisReader.Close()
	resp, err := redis.Bool(redisReader.Do("SISMEMBER", key, value))
	if err != nil {
		return false
	}
	return resp
}

func (r *RedisCache) Set(key string, value interface{}) error {
	return r.SetEx(key, value, r.Config.CacheTTL)
}

func (r *RedisCache) SetEx(key string, value interface{}, ttl int) error {
	redisWriter := r.writerPool.Get()
	defer redisWriter.Close()
	_, err := redisWriter.Do("SETEX", key, ttl, value)
	return err
}

func (r *RedisCache) SetExObj(key string, value interface{}, ttl int) error {
	json, err := jsoniter.MarshalToString(value)
	if err != nil {
		return err
	}
	return r.SetEx(key, json, ttl)
}

func (r *RedisCache) SetSlice(key string, value interface{}, ttl int) error {
	rType := reflect.TypeOf(value)
	if rType.Kind() != reflect.Slice {
		return errors.New("Invalid format, value must be slice")
	}
	rValue := reflect.ValueOf(value)
	redisWriter := r.writerPool.Get()
	defer redisWriter.Close()
	redisWriter.Send("MULTI")
	for i := 0; i < rValue.Len(); i++ {
		val := rValue.Index(i)
		redisWriter.Send("SADD", key, val.Interface())
	}

	if ttl != 0 {
		redisWriter.Send("EXPIRE", key, ttl)
	}
	_, err := redisWriter.Do("EXEC")
	return err
}

func (r *RedisCache) ForceExpire(keys ...string) error {
	redisWriter := r.writerPool.Get()
	defer redisWriter.Close()
	redisWriter.Send("MULTI")
	for _, key := range keys {
		redisWriter.Send("EXPIRE", key, 0)
	}
	_, err := redisWriter.Do("EXEC")
	return err
}

func (r *RedisCache) Del(keys ...string) error {
	redisWriter := r.writerPool.Get()
	defer redisWriter.Close()
	keysInterface := make([]interface{}, len(keys))
	for i, key := range keys {
		keysInterface[i] = key
	}
	_, err := redisWriter.Do("DEL", keysInterface...)
	return err
}
