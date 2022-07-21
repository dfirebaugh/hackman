package config

import (
	_ "embed"
	"encoding/json"
	"flag"
	"os"

	"github.com/sirupsen/logrus"
)

type config struct {
	PaypalClientID     string `json:"paypalClientID"`
	PaypalClientSecret string `json:"paypalClientSecret"`
	PaypalURL          string `json:"paypalURL"`
	ServiceURL         string `json:"serviceURL"`
}

var (
	Config        config
	configPtr     = flag.String("config", "config.json", "path to config")
	serviceURLPtr = flag.String("url", "", "https url of server")
)

func init() {
	Config.ServiceURL = *serviceURLPtr
	if dat, err := os.ReadFile(*configPtr); err != nil {
		logrus.Errorf("error openings config file: %s", err)
	} else {
		json.Unmarshal(dat, &Config)
	}
}
