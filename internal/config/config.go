package config

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

var (
	ErrConfigNotSet    = fmt.Errorf("configuration path isn't set")
	ErrConfigNotExist  = fmt.Errorf("configuration path doesn't exist")
	ErrConfigBadFormat = fmt.Errorf("configuration file has bad format")
)

type Config struct {
	ServerHost string `json:"server_host"`
	ServerPort string `json:"server_port"`

	ChallengeTTL int64 `json:"challenge_ttl"`

	HashcashZerosAmount   uint   `json:"hashcash_zeros_amount"`
	HashcashMaxIterations uint64 `json:"hashcash_max_iterations"`
}

func ParseConfigFromCLI() (config *Config, err error) {
	path, err := parseParamsFromCLI()
	if err != nil {
		return nil, err
	}

	reader, err := readConfigFile(path)
	if err != nil {
		return nil, err
	}

	defer func() {
		e := reader.Close()
		if err != nil {
			err = fmt.Errorf("%w: %v", err, e)
		} else {
			err = e
		}
	}()

	config, err = parseConfigFile(reader)

	return config, err
}

func parseParamsFromCLI() (string, error) {
	var path string

	flag.StringVar(&path, "path", "", "-path=./config.json")
	flag.Parse()

	if path == "" {
		return "", ErrConfigNotSet
	}

	path, err := filepath.Abs(path)
	if err != nil {
		return "", ErrConfigNotExist
	}

	return path, nil
}

func readConfigFile(path string) (reader io.ReadCloser, err error) {
	reader, err = os.Open(path)
	if err != nil {
		return nil, ErrConfigNotExist
	}

	return reader, nil
}

func parseConfigFile(reader io.ReadCloser) (config *Config, err error) {
	err = json.NewDecoder(bufio.NewReader(reader)).Decode(&config)
	if err != nil {
		return nil, ErrConfigBadFormat
	}

	return config, nil
}
