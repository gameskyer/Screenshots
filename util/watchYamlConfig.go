package util

import (
	"Screenshots/constant"
	"Screenshots/module"
	"fmt"
	log "github.com/cihub/seelog"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var Config module.YamlConfig
var sentryCfg *viper.Viper

func ReadConfig() error {
	content, err := ioutil.ReadFile(constant.ConfigPath)
	if err != nil {
		fmt.Print(err)
		return err
	}
	err = yaml.Unmarshal(content, &Config)
	return err
}

//读指定位置配置文件
func ConfigViper(path string) *viper.Viper {
	v := viper.New()
	v.SetConfigFile(path)
	_ = v.ReadInConfig()
	return v
}

func WatchConfig() {
	//监控配置文件变化
	sentryCfg = ConfigViper(constant.ConfigPath)
	sentryCfg.WatchConfig()
	sentryCfg.OnConfigChange(func(e fsnotify.Event) {
		ReadConfig()
		log.Warn("配置文件已修改,重新加载配置文件")
		return
	})
}
