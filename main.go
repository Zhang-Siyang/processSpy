package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var dataFilePath = flag.String("c", "./programList.json", "数据文件路径")

type program struct {
	Name    string `json:"name"`
	Keyword string `json:"keyword"`
	Amount  int    `json:"amount"`
	Reality int    `json:"-"`
}

func takeArryPutMarkdown(programes *[]program) string {
	result := ""
	for _, program := range *programes {
		result += fmt.Sprintf("Hello %s\n", program.Name)
	}
	return result
}

func main() {
	log.SetPrefix("\n")
	log.SetFlags(log.LUTC)
	flag.Parse()
	dataFile, err := os.Open(*dataFilePath)
	if err != nil {
		log.Fatalln(err)
	}
	defer dataFile.Close()
	datainStr, err := ioutil.ReadAll(dataFile)
	if err != nil {
		log.Fatalln(err)
	}
	var programes []program
	err = json.Unmarshal([]byte(datainStr), &programes)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("# Input #")
	for i, pro := range programes {
		fmt.Printf("%d: %#v\n", i, pro)
	}
	log.Println("# Output #")
	fmt.Print(takeArryPutMarkdown(&programes))
}
