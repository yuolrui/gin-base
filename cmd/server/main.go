package main

import (
	"github.com/yuolrui/gin-base/internal/bootstrap"
	"github.com/yuolrui/gin-base/internal/router"
)

func main() {
	bootstrap.Init()
	r := router.InitRouter()
	r.Run(":8080")
}
