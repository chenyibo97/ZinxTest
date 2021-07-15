package main

import "studygo2/zinxtest/znet"

func main() {
	server := znet.NewServer("[zinxV0.1]")
	server.Server()
}
