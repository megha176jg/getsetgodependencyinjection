package main

import (
	"log"

	"bitbucket.org/junglee_games/getsetgo/zipper"
)

func main() {

	zipper := zipper.NewZipper()

	err := zipper.AddFile("test/pavan.txt", "Hello there !!!")
	if err != nil {
		panic(err)
	}

	err = zipper.AddFile("test/pavan1.txt", "Hello there !!!")
	if err != nil {
		panic(err)
	}

	zipper.Close()

	// get a zip file content as string
	content := zipper.GetString()
	log.Printf("Zip file content:%s", content)

	// you can save it to file as well
	zipper.SaveFile("myzip.zip")
}
