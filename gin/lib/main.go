package main

func main() {
	r := makeRoutes()

	r.Run(":8888")
}
