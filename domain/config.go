package domain

import (
	"encoding/json"
	"fmt"
)

type msi = map[string]interface{}

type Config struct {
	config map[string]interface{}
}

func (c *Config) SetFromBytes(data []byte) error {
	var rawConfig interface{}
	if err := json.Unmarshal(data, &rawConfig); err != nil {
		return err
	}
	dump := fmt.Sprintf("%#v", rawConfig)
	fmt.Printf("%s", dump[:500])
	for k := range rawConfig.(msi) {
		fmt.Println(k)
	}
	fmt.Printf("%#v", rawConfig.(msi)["roles"])
	return nil
	/*
		config, ok := rawConfig.(map[interface{}]interface{})
		if !ok {
			return fmt.Errorf("config is not a map")
		}

		return nil
	*/
}
