package main

import (
	"github.com/lfyr/go-api/app/admin/router"
	"os"
)

func main() {
	r := router.Router()
	r.Run("0.0.0.0:" + os.Getenv("PORT"))
}
