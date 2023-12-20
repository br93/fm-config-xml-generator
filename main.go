package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"strings"
)

func createRecord() Record {
	preload := BooleanElement{
		ID:    "preload",
		Value: "false",
	}

	amap := BooleanElement{
		ID:    "amap",
		Value: "false",
	}

	list := List{
		ID:   "maps",
		Maps: pictures(),
	}

	record := Record{
		Booleans: []BooleanElement{preload, amap},
		List:     list,
	}

	return record
}

func pictures() []Map {
	files, err := os.ReadDir(".")
	var maps []Map
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		id := strings.Split(file.Name(), ".")

		if strings.Contains(file.Name(), "png") {
			fmt.Println(file.Name())
			player := Map{
				From: id[0],
				To:   fmt.Sprintf("graphics/pictures/person/%s/portrait", id[0]),
			}

			maps = append(maps, player)
		}
	}

	return maps
}

func encode(record Record) {
	outputFile, err := os.Create("config.xml")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer outputFile.Close()
	encoder := xml.NewEncoder(outputFile)
	encoder.Indent("", "\t")
	if err := encoder.Encode(record); err != nil {
		fmt.Println("Error encoding XML to file:", err)
	}
}

func main() {
	record := createRecord()
	encode(record)
	fmt.Printf("Config.xml generated with %d files\n", len(record.List.Maps))
}
