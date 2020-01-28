package gen

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type Inputs struct {
	TemplatePath string
	ValuesPath string
	OutputPath string
}

func Start(inputs Inputs) {
	config := readFile(inputs.ValuesPath)
	t, err := template.New(filepath.Base(inputs.TemplatePath)).ParseFiles(inputs.TemplatePath)
	if err != nil {
		log.Fatal(err)
	}

	if inputs.OutputPath != "" {
		f, err := os.Create(inputs.OutputPath)
		if err != nil {
			log.Fatal(err)
		}
		err = t.Execute(f, config)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		err = t.Execute(os.Stdout, config)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func readFile(filePath string) map[string]string {

	config := make(map[string]string)

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if equal := strings.Index(line, "="); equal >= 0 {
			if key := strings.TrimSpace(line[:equal]); len(key) > 0 {
				value := ""
				if len(line) > equal {
					value = strings.TrimSpace(line[equal+1:])
				}
				config[key] = value
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return config
}
