package main

import "io/ioutil"
import "./domain"

func main() {
	config := domain.Config{}

	data, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic(err)
	}

	err = config.SetFromBytes(data)
	if err != nil {
		panic(err)
	}
	//config.Get("applications")
}
