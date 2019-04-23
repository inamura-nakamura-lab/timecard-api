package main

import (
	"github.com/inamura-nakamura-lab/timecard-api/interface/handler"
)

func main() {
	routerHandler := handler.NewRouterHandler()
	r := routerHandler.SetUpRouter()
	err := r.Run(":8888")
	if err != nil {
		panic(err)
	}
}