// Package config TODO
package config

import (
	"encoding/json"
	"log"
	"sync"

	"github.com/spf13/viper"
)

// Address TODO
var Address = "localhost:9090"

// Conf TODO
type Conf struct {
	Viper *viper.Viper
	Lock  sync.RWMutex
}

var (
	// C TODO
	C        *Conf
	initOnce sync.Once
)

func init() {
	InitConfig()
	InitRedisPool()
}

// NewConf TODO
func NewConf() *Conf {
	initOnce.Do(func() {
		C = &Conf{
			Viper: viper.New(),
			Lock:  sync.RWMutex{},
		}
	})
	return C
}

// InitConfig TODO
func InitConfig() {
	C = NewConf()
	C.Viper.AddConfigPath("./config")
	C.Viper.SetConfigName("local")
	err := C.Viper.ReadInConfig()
	if err != nil {
		panic("failed to read in config")
	}
	log.Println("server config initialize success")
	log.Printf("yaml config:\n%s\n", C.MarshalIndent())
}

// MarshalIndent TODO
func (c *Conf) MarshalIndent() []byte {
	r, err := json.MarshalIndent(c.Viper.AllSettings(), "", " ")
	if err != nil {
		return []byte("")
	}
	return r
}
