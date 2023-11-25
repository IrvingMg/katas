// Reverse Array
//
// Description:
// Write a function that reverses each string in an array. The input is given
// as an array of strings.
//
// Examples:
//
//	Input: s = ["alpha", "beta", "charlie", "delta", "foxtrot"]
//	Output: ["ahpla", "ateb", "eilrahc", "atled", "tortxof"]
package reversearray

import (
	"context"
	"log"
	"runtime"
	"slices"
)

type ReversedString struct {
	Str   string
	Index int
}

func ReverseMultiString(data []string) []string {
	for i := range data {
		data[i] = reverseString(data[i])
	}

	return data
}

func reverseString(str string) string {
	runes := []rune(str)
	slices.Reverse(runes)

	return string(runes)
}

func ReverseMultiStringWithGenerator(data []string) chan ReversedString {
	stream := make(chan ReversedString)

	go func() {
		defer close(stream)
		for i := range data {
			stream <- ReversedString{
				Str:   reverseString(data[i]),
				Index: i,
			}
		}
	}()

	return stream
}

func ReverseMultiStringWithCancellation(ctx context.Context, data []string) chan ReversedString {
	stream := make(chan ReversedString)

	//time.Sleep(1 * time.Second)
	go func() {
		defer close(stream)
		for i := range data {
			val := ReversedString{
				Str:   reverseString(data[i]),
				Index: i,
			}
			select {
			case stream <- val:
				log.Printf("val: %+v sent", val)
			case <-ctx.Done():
				log.Println("work cancelled")
				return
			}

		}
	}()

	return stream
}

func ReverseMultiStringWithWorkLimit(ctx context.Context, data []string) chan chan ReversedString {
	workers := len(data)
	stream := make(chan chan ReversedString, workers)
	defer close(stream)

	//time.Sleep(1 * time.Second)
	g := runtime.GOMAXPROCS(0)
	sem := make(chan bool, g)
	for i := 0; i < workers; i++ {
		sem <- true
		{
			ch := make(chan ReversedString)
			go func(i int) {
				defer close(ch)
				val := ReversedString{
					Str:   reverseString(data[i]),
					Index: i,
				}
				select {
				case ch <- val:
					log.Printf("val: %+v sent", val)
				case <-ctx.Done():
					log.Println("work cancelled")
				}
			}(i)
			stream <- ch
		}
		<-sem
	}

	return stream
}
