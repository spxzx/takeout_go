package main

import (
	"TakeOut/model"
	"TakeOut/router"
)

func main() {
	model.GobInit()
	model.InitDb()
	router.InitRouter()
}
