package main

import (
	"github.com/inamura-nakamura-lab/timecard-api/utils/handler"
	"os"
)

func main() {
	err := handler.Router.Run(":" + os.Getenv("PORT"))
	if err != nil {
		panic(err)
	}
}