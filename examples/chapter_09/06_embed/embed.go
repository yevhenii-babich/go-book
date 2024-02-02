package main

import (
	"embed"
	"fmt"
	"gopkg.in/yaml.v3"
	"io/fs"
	"log/slog"
	"os"
	"strings"
)

//go:embed hello.txt
var helloFile embed.FS

//go:embed config
var configDir embed.FS

func main() {
	// Читання з файлу
	data, err := fs.ReadFile(helloFile, "hello.txt")
	if err != nil {
		slog.Error("Can't read embedded file:", "error", err)
	}
	fmt.Println(string(data))

	// Перегляд вмісту директорії
	dirEntries, err := fs.ReadDir(configDir, "config")
	if err != nil {
		slog.Error("Can't read embedded directory:", "error", err)
		os.Exit(1)
	}
	// Перегляд вмісту директорії
	fmt.Println("Embedded directory 'config':")
	for _, entry := range dirEntries {
		if !entry.IsDir() {
			// Читання з файлу
			fileData, _ := fs.ReadFile(configDir, "config/"+entry.Name())
			fmt.Println(entry.Name(), ":\n", string(fileData)) // Виведення вмісту файлу
			if strings.HasSuffix(entry.Name(), ".yaml") {
				decode(fileData)
			}
		} else {
			// Якщо це директорія
			fmt.Println(entry.Name(), ": is directory")
		}
	}
}

func decode(data []byte) {
	var config map[string]interface{}
	_ = yaml.Unmarshal(data, &config)
	fmt.Printf("Decoded YAML data: %+v\n", config)
}
