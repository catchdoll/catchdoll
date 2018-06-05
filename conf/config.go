package conf

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"fmt"
)

type Configuration struct{
	Dsn string `yaml:"dsn"`
	RedisAddr string `yaml:"redis_addr"`
	RedisPwd string `yaml:"redis_password"`
	RedisDb int `yaml:"redis_db"`
	ServerPort string `yaml:"server_port"`
	WxAppid string `yaml:"wx_appid"`
	WxAppsecret string `yaml:"wx_appsecret"`
	JwtSecret string `yaml:"jwt_secret"`
	LbsAk string `yaml:"lbs_ak"`
	LbsSk string `yaml:"lbs_sk"`
}

var GlobalConf *Configuration

func InitConfig() error {
	fmt.Println()
	data, err := ioutil.ReadFile("conf/config.yaml")
	if err != nil{
		fmt.Println(err)
		return err
	}
	var config Configuration
	err = yaml.Unmarshal(data, &config)
	if err != nil{
		return err
	}
	GlobalConf = &config
	return nil
}


