package main

import (
	"fmt"
	"go-design-pattern/structural-pattern/proxy"
)

func CallProxyMethod() {

	var nginxServer proxy.Server

	nginxServer = proxy.NewNginxServer()
	appStatusURL := "/app/status"
	createUserURL := "/create/user"

	httpCode, body := nginxServer.HandleRequest(appStatusURL, "GET")
	fmt.Printf("\nnUrl: %s\nnHttpCode: %d\nnBody: %s\n", appStatusURL, httpCode, body)

	httpCode, body = nginxServer.HandleRequest(appStatusURL, "GET")
	fmt.Printf("\nnUrl: %s\nnHttpCode: %d\nnBody: %s\n", appStatusURL, httpCode, body)

	httpCode, body = nginxServer.HandleRequest(appStatusURL, "GET")
	fmt.Printf("\nnUrl: %s\nnHttpCode: %d\nnBody: %s\n", appStatusURL, httpCode, body)

	httpCode, body = nginxServer.HandleRequest(createUserURL, "POST")
	fmt.Printf("\nnUrl: %s\nnHttpCode: %d\nnBody: %s\n", appStatusURL, httpCode, body)

	httpCode, body = nginxServer.HandleRequest(appStatusURL, "GET")
	fmt.Printf("\nnUrl: %s\nnHttpCode: %d\nnBody: %s\n", appStatusURL, httpCode, body)
}
