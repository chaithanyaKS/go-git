/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/chaithanyaKS/go-git/cmd"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	cmd.Execute()
}
