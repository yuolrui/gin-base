package main

import (
	"github.com/yuolrui/gin-base/internal/router"
)

func main() {
	r := router.InitRouter()
	r.Run(":8080")
}
