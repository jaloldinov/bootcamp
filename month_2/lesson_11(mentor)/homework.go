package main

import (
	"fmt"
	"strings"
)

// -- 2. so'z(string) va gaplar([]string) berilgan. qaysi gaplarda shu so'z qatnashganini concurrent qilib topish kerak:
// --   a) so'z ixtiyoriy 1tasida topilsa shu gapni qaytarilsin
// --   b) so'z qatnashgan hamma gaplar hammasi chiqarilsin
type message struct {
	Text     string
	Receiver string
}

func main() {
	word := "apple"
	sentences := []string{
		"I love an apple",
		"Do you like apples?",
		"My friend likes banana",
		"HI how are you apple",
	}

	// Case a:
	fmt.Println("Case a:")
	foundSentence := make(chan string)
	go a(foundSentence, sentences, word)
	fmt.Println(<-foundSentence)

	// Case b:
	fmt.Println("Case b:")
	foundSentences := make(chan string)
	go b(foundSentences, sentences, word)

	for sentence := range foundSentences {
		fmt.Println(sentence)
	}
	fmt.Println()
	// ===================================================
	black_list := map[string]bool{
		"John":    true,
		"Omadbek": true,
	}
	smsChan := make(chan string)

	msg1 := message{
		Text:     "Hello, my friend",
		Receiver: "John",
	}
	msg2 := message{
		Text:     "Hello, my friend",
		Receiver: "Sarvarbek",
	}

	go serviceMessage(smsChan, msg1, black_list)
	go serviceMessage(smsChan, msg2, black_list)
	fmt.Println(<-smsChan)
	fmt.Println(<-smsChan)

}

// Case a:
func a(ch chan string, sentences []string, word string) {
	for _, str := range sentences {
		if strings.Contains(str, word) {
			ch <- str
			break
		}
	}
}

// Case b:
func b(ch chan string, sentences []string, word string) {
	for _, str := range sentences {
		if strings.Contains(str, word) {
			ch <- str
		}
	}
	close(ch)
}

// -- 3. Service: service message{text, receiver} larni qabul qiladi. Black list berilgan,
// shu black listda bo'lmaganlarga message yuboriladi, aks holda yuborilmaydi:
// -- black_list=[John]
// -- 1. message{text:'hello',receiver: 'John'}
// --    result: Message ignored
// -- 2. message{text:'hi',receiver: 'Adam'}
// --    result: Message sent

// fmt.Printf("Message ignored: %s -> %s\n", msg.Sender, msg.Receiver)

func serviceMessage(ch chan string, msg message, black_list map[string]bool) {
	if black_list[msg.Receiver] {
		ch <- fmt.Sprintf("Message ignored sending to %s", msg.Receiver)
	} else {
		ch <- fmt.Sprintf("Message sent to %s", msg.Receiver)
	}
}
