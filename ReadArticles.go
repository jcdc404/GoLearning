package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/go-yaml/yaml"
)

// " = 0x22, ' = 0x27, : = 0x3A, CR=0x0D,LF=0x0A
type defaults struct {
	filesFolder string
}
type article struct {
	Title           string            `yaml:"Title"`
	Author          string            `yaml:"Author"`
	Date            string            `yaml:"Date"`
	BackgroundImage string            `yaml:"BackgroundImage"`
	Links           map[string]string `yaml:"Links,flow"`
	Section         map[string]string `yaml:"Sections,flow"`
}

func init() {

}
func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func main() {
	var art article
	file, err := ioutil.ReadFile("First-Article Title.yml")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = yaml.Unmarshal(file, &art)
	check(err)
	//Prints out the Section title, then the Section conetnt
	for k, v := range art.Section {
		fmt.Println(k)
		fmt.Println(v)
	}
}
