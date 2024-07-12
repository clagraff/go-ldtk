package main

import (
	"fmt"
	"goldtk/quicktype"
	"log"
	"os"
)

func main() {
	fileData, err := os.ReadFile("C:\\Users\\curti\\GolandProjects\\github\\clagraff\\go-ldtk\\test\\ldtk\\openrogue.ldtk")
	if err != nil {
		log.Fatal(err)
	}

	project, err := quicktype.UnmarshalLdtkJSON(fileData)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(project)

}
