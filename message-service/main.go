package main

import (
	"fmt"
	"log"
	"net/http"

	"stewped-applet/common"
	"stewped-applet/common/metrics"
	"stewped-applet/message-service/routers"
)

func main() {
	router := routers.InitRoutes()
	port := fmt.Sprintf(":%s", common.GetEnv("PORT", "8080"))
	metrics.RegisterHandlerMetrics(router)

	log.Printf("server listening at %s", port)
	if err := http.ListenAndServe(port, router); err != nil {
		log.Panicf("error while serving: %s", err)
	}
}
