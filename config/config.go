package config

import (
	"encoding/json"
	"io/ioutil"
)

type config struct {
	Username       string   `json:"username"`
	Password       string   `json:"password"`
	Duration       int      `json:"duration"`
	Admin          string   `json:"admin"`
	Player         string   `json:"player"`
	Domain         string   `json:"domain"`
	Port           int      `json:"port"`
	Tip            tip      `json:"tip"`
	Limit_time     int64    `json:"limit_time"`
	Limit_requence int8     `json:"limit_requence"`
	Allow_url      []string `json:"allow_url"`
	Paras          int      `json:"paras"`
}

type tip struct {
	Time  string `json:"time"`
	Color string `json:"color"`
	Text  string `json:"text"`
}

func Read() (config, error) {
	var d config
	f, err := ioutil.ReadFile("config.json")
	if err != nil {
		return d, err
	}
	err = json.Unmarshal(f, &d)
	if err != nil {
		return d, err
	}
	return d, nil
}
