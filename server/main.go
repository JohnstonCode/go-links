package main

import (
	"github.com/JohnstonCode/go-links/http"
	"github.com/JohnstonCode/go-links/model"
)

func main() {
	model.Setup()

	http.SetupAndListen()
}
