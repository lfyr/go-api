package main

import (
	"github.com/lfyr/go-api/app/admin/router"
	"github.com/lfyr/go-api/utils"
)

func main() {
	r := router.Router()
	r.Run("0.0.0.0:" + utils.GVA_CONFIG.System.Addr)
}
