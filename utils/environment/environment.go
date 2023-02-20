package environment

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

type Environment interface {
	Get(key string) string
	GetUint(key string, defaultValue uint) uint
}

func Default() *OsEnvironment {
	return new(OsEnvironment)
}

type OsEnvironment bool

func (OsEnvironment) Get(key string) string {
	return os.Getenv(key)
}

func (e OsEnvironment) GetUint(key string, defaultValue uint) uint {
	str := os.Getenv(key)
	value, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		log.Println(fmt.Errorf("strconv.ParseUint: %w", err))
		value = uint64(defaultValue)
	}

	return uint(value)
}
