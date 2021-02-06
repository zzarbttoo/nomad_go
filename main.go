package main

import (
	"fmt"
	"time"
)

func main() {

	//어떤 type 의 information 인지 적어줘야한다
	channel := make(chan string)
	people := [5]string{"nico", "flynn", "dal", "japangy", "larry"}

	//5 routines -> 5 messages
	for _, person := range people {
		go isSexy(person, channel)
	}

	for i := 0; i < len(people); i++ {
		fmt.Println(<-channel)
	}
}

//return 값이 의마가 없다 (go routines 는 어떤 값을 저장할 수 없게 한다)
// chan channel type
func isSexy(person string, channel chan string) {

	time.Sleep(time.Second * 10)
	//channel 에 true 를 보내라
	channel <- person + "is sexy"

}
