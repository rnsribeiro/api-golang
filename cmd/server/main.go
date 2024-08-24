package main

import "github.com/rnsribeiro/api-golang/configs"

func main() {
	config, _ := configs.LoadConfig(".")
	println(config.DBDriver)
}
