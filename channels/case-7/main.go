package main

type Figure struct {
	Length int
	Width  int
	Height int
	Square int
	Volume int
}

const (
	n           = 2
	statusOK    = 0
	statusError = 1
)

func main() {
	errc := make(chan error)
	status := statusOK

	errGroup := sync.WaitGroup{}
	errGroup.Add(1)

	go func() {
		for err := range errc {
			status = statusError
			fmt.Printf("error processing the code: %s\n", err)
		}
		errGroup.Done()
	}()

	ff := []Figure{
		Figure{1, 2, -5, 0, 0},
		Figure{3, 2, 4, 0, 0},
		Figure{1, 10, 3, 0, 0},
		Figure{1, 10, -3, 0, 0},
		Figure{-1, 10, 3, 0, 0},
		Figure{1, 10, 3, 0, 0}}

	squarec := make(chan Figure, n)

	volumec := make(chan Figure, n)

	
}
