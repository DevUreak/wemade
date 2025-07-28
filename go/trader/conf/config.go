package conf

import (
	"fmt"
	"os"
	"github.com/naoina/toml"
)

type Config struct {
	Common struct {
		ServiceId string
	}

	DataDirectory struct {
		Root     string
		Keystore string
		Journal  string
		Log      string
		ExAccKey string
	}

	Port struct {
		Server     string
		Http       int
		Prometheus int
	}

	Gclient struct {
		GrpcPort string
	}

	Gserver struct {
		ServerAddr string
	}

	Repositories map[string]map[string]interface{}

	Contracts map[string]map[string]interface{}

	Chains []string

	Log struct {
		Terminal struct {
			Use       bool
			Verbosity int
		}
		File struct {
			Use       bool
			Verbosity int
			FileName  string
		}
	}

	CoinMarketCapAPI struct {
		Url    string
		ApiKey string
	}
}

func NewConfig(file string) *Config {
	c := new(Config)

	if file, err := os.Open(file); err != nil {
		panic(err)
	} else {
		defer file.Close()
		if err := toml.NewDecoder(file).Decode(c); err != nil {
			panic(err)
		} else {
			fmt.Print(c.Repositories)
			return c
		}
	}
}
