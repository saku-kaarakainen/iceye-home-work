package json

import (
	"encoding/json"
	"io"
	"os"
)

// Load reads the configuration file with the given filename and returns a RootConfig.
// If the file is not found or cannot be read, Load panics.
func Load[Type any](filename string) (Type, error) {
	file, err := os.Open(filename)
	var retType Type
	if err != nil {
		return retType, err
	}
	defer file.Close()

	jsonBlob, err := io.ReadAll(file)
	if err != nil {
		return retType, err
	}

	json.Unmarshal(jsonBlob, &retType)
	return retType, nil
}
