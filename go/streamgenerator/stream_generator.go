// Stream Generator
//
// Description:
// Create a generator stream that will receive a slice with the letters of the
// alphabet and will return a channel signaling every letter.
package streamgenerator

func MakeGenerator(data []rune) <-chan rune {
	stream := make(chan rune)

	// Note: with buffered channel, go routine is unnecessary
	go func() {
		defer close(stream)
		for _, val := range data {
			stream <- val
		}
	}()

	return stream
}
