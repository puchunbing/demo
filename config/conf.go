package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
)

var (
	newconfig *Config
)

type Config struct {
	MySql    MySqlConfig `json:"mysql"`
	Httpport string      `json:"httpport"`
}

type MySqlConfig struct {
	DbUrl  string `json:"db_url"`
	DbUser string `json:"db_user"`
	DbPass string `json:"db_pass"`
	DbName string `json:"db_name"`
}

type RedisConfig struct {
	RURI      string `json:"ruri"`
	RPassword string `json:"rpassword"`
	Db        int    `json:"db"`
}

type AnimalConfig struct {
	Count int `json:"count"`
}

type Loss struct {
	Jcoin int `json:"jcoin"`
	YCoin int `json:"ycoin"`
	TCoin int `json:"tcoin"`
}

func ConfigInit(con *Config) error {
	newconfig = con
	return nil
}

func atoiUtil(value string) int {
	if value == "" {
		return 0
	}
	i, _ := strconv.Atoi(value)
	return i
}

func StringToFloat(str string) float64 {
	in, _ := strconv.Atoi(str)
	return float64(in)
}

// 从配置文件中加载配置信息
func LoadConfig() (Config, error) {
	var config Config
	// 读取配置文件
	data, err := ioutil.ReadFile("config/conf.json")
	if err != nil {
		return config, fmt.Errorf("failed to read config file: %w", err)
	}

	// 解析配置文件
	err = json.Unmarshal(data, &config)
	if err != nil {
		return config, fmt.Errorf("failed to parse config file: %w", err)
	}
	newconfig = &config
	return *newconfig, nil
}

func GetConfig() *Config {
	return newconfig
}
