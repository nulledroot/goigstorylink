package igstorydl

import (
	"fmt"
	"github.com/siongui/goigstorylink"
)

func Download() {
	users, err := igstory.GetUnreadStories()
	if err != nil {
		panic(err)
	}

	for _, user := range users {
		for _, story := range user.Stories {
			fmt.Println(BuildOutputPath(user.Username, story.Url, story.Timestamp))
		}
	}
}
