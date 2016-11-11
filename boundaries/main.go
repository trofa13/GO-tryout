package main


import (
	"fmt"
	"boundaries/model"
)


func main() {
	speakersList := model.GetSpeakers()

	for _, speaker := range speakersList {
		fmt.Println(speaker.Speak())
	}
}