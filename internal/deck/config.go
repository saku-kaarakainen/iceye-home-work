package deck

import (
	"encoding/json"
	"io/ioutil"
)

// Load reads the configuration file with the given filename and returns a RootConfig.
// If the file is not found or cannot be read, Load panics.
func Load(filename string) (RootConfig, error) {
	jsonBlob, err := ioutil.ReadFile(filename)
	if err != nil {
		return RootConfig{}, err
	}

	var config RootConfig
	json.Unmarshal(jsonBlob, &config)

	return config, nil
}
