package conf

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

var Config config

type config struct {
	DataSource struct {
		Url          string `yaml:"url"`
		IdleCon      int    `yaml:"idleCon"`
		OpenCon      int    `yaml:"openCon"`
		ShowSQL      bool   `yaml:"showSQL"`
		ShowExecTime bool   `yaml:"showExecTime"`
	}
	rocketmq struct {
		nameSrv string `yaml:"nameSrv"`
		topic   string `yaml:"topic"`
	}
}

func Init() {
	bytes, err := ioutil.ReadFile("./conf/application.yml")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	err = yaml.Unmarshal(bytes, &Config)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
