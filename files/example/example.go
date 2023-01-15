package main

import "bitbucket.org/junglee_games/getsetgo/files"

func main() {
	err := files.SaveFile(".", "myfile.txt", "This is my file!!!")
	if err != nil {
		panic(err)
	}
}
