package singleton

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

// Settings holds the directory and file pattern values
type Settings struct {
	Directory   string `json:"directory"`
	FilePattern string `json:"filePattern"`
}

// singleton
var instance *Settings
var once sync.Once

func GetInstance() *Settings {
	once.Do(func() {
		instance = &Settings{}
		instance.Load()
	})
	return instance
}

// saveValues saves the values of the entries to a JSON file
func (settings *Settings) Save() {
	file, err := os.Create("settings.json")
	if err != nil {
		fmt.Println("Error creating settings file:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // Pretty print
	if err := encoder.Encode(settings); err != nil {
		fmt.Println("Error encoding settings to JSON:", err)
	}
}

// loadValues loads previously saved values from a JSON file
func (settings *Settings) Load() {
	file, err := os.Open("settings.json")
	if err != nil {
		if os.IsNotExist(err) {
			// File does not exist, ignore error
			return
		}
		fmt.Println("Error opening settings file:", err)
		return
	}
	defer file.Close()

	var tempSettings Settings
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&tempSettings); err != nil {
		fmt.Println("Error decoding settings from JSON:", err)
		return
	}

	settings.Directory = tempSettings.Directory
	settings.FilePattern = tempSettings.FilePattern
}
