package util

import (
	"github.com/Unknwon/goconfig"
)

var Cfg *goconfig.ConfigFile

func init() {
	var err error
	Cfg, err = goconfig.LoadConfigFile("./conf/conf.ini")
	if err != nil {
		panic(err)
	}
}
