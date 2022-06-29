package main

func main() {
	m := make(chan struct{})

	close(m)
	close(m)
}
