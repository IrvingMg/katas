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
				continue
			case <-ctx.Done():
				log.Println("work cancelled")
				return
			}

		}
	}()

	return stream
}

func ReverseMultiStringConcurrentlyWithLimit(data []string, poolSize int) []string {
	out := make(chan string, poolSize)
	for i := range data {
		go func(i int) {
			out <- reverseString(data[i])
		}(i)
		data[i] = <-out
	}

	return data
}
