package config

import (
	"github.com/ahmadrezamusthafa/assessment/common/logger"
	jsoniter "github.com/json-iterator/go"
	"github.com/vrischmann/envconfig"
	"io/ioutil"
	"os"
	"time"
)

type Config struct {
	HttpPort         string        `envconfig:"default=9505"`
	CacheWriterHost  string        `envconfig:"default=127.0.0.1:6379"`
	CacheReaderHost  string        `envconfig:"default=127.0.0.1:6379"`
	CacheMaxIdle     int           `envconfig:"default=5"`
	CacheIdleTimeout time.Duration `envconfig:"default=180s"`
	CacheTTL         int           `envconfig:"default=10000"`
	DatabaseHost     string        `envconfig:"default=127.0.0.1"`
	DatabasePort     string        `envconfig:"default=5432"`
	DatabaseName     string        `envconfig:"default=assessment"`
	DatabaseUsername string        `envconfig:"default=postgres"`
	DatabasePassword string        `envconfig:"default=reza"`
	NSQDHost         string        `envconfig:"default=127.0.0.1"`
	NSQDPort         int           `envconfig:"default=4150"`
	NSQLookupHost    string        `envconfig:"default=127.0.0.1"`
	NSQLookupPort    int           `envconfig:"default=4161"`
}

func readFromFileAndEnv(conf interface{}) (err error) {
	file, err := os.Open("appsettings.json")
	if err == nil {
		defer file.Close()
		data, inErr := ioutil.ReadAll(file)
		if inErr != nil {
			err = inErr
			return
		}
		maps := make(map[string]string)
		inErr = jsoniter.Unmarshal(data, &maps)
		if inErr != nil {
			err = inErr
			return
		}
		for k, v := range maps {
			inErr = os.Setenv(k, v)
			if inErr != nil {
				err = inErr
				return
			}
		}
	} else {
		logger.Warn("%v", err)
	}

	err = envconfig.Init(conf)
	if err != nil {
		return
	}
	return
}

func New() (conf *Config, err error) {
	conf = new(Config)
	err = readFromFileAndEnv(&conf)
	return
}
