package main

import "github.com/jb55/go-twitter"
import "fmt"
import "time"
import "math/rand"

const kMaxDepth = 25
const kStart = "jb55"

var api *twitter.Api
var r *rand.Rand
var done chan bool

func main() {
	api = twitter.NewApi()
	done = make(chan bool)
	r = rand.New(rand.NewSource(time.Now().Unix()))
	crawl(kStart, 0)
	<-done
}

func crawl(userName string, level int) {
	// Get the user's status
	text := (<-api.GetUser(userName)).GetStatus().GetText()

	for i := 0; i < level; i++ {
		fmt.Printf("  ")
	}

	fmt.Printf("%s: %s\n", userName, text)

	level++
	if level > kMaxDepth {
		done <- true
		return
	}

	// Get the user's friends
	friends := <-api.GetFriends(userName, 1)
	length := len(friends)

	if length == 0 {
		done <- true
		return
	}

	rVal := r.Intn(length - 1)
	// Choose a random friend for the next user
	nextUser := friends[rVal].GetScreenName()

	go crawl(nextUser, level)
}
