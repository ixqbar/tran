package common

import (
	"encoding/xml"
	"errors"
	"fmt"
	"log"
	"os"
)

type config struct {
	Address string `xml:"address"`
	DataPath string `xml:"data"`
}

var Config *config

func ParseXmlConfig(file string) (*config, error) {
	if len(file) == 0 {
		return nil, errors.New("not found configure xml file")
	}

	n, err := GetFileSize(file)
	if err != nil || n == 0 {
		return nil, fmt.Errorf("not found config file %v", file)
	}

	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	data := make([]byte, n)

	m, err := f.Read(data)
	if err != nil {
		return nil, err
	}

	if int64(m) != n {
		return nil, fmt.Errorf("expect read configure xml file size %d but result is %d", n, m)
	}

	err = xml.Unmarshal(data, &Config)
	if err != nil {
		return nil, err
	}

	log.Printf("load config data %+v", Config)

	return Config, nil
}
