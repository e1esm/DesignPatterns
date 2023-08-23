package main

import "fmt"

func main() {
	nginxServer := NewNginx()
	appStatusURL := "/app/status"
	createUserURL := "/create/user"

	htppCode, body := nginxServer.handleRequest(appStatusURL, "GET")
	fmt.Printf("\nURL:%s httpCode: %d body: %s\n", appStatusURL, htppCode, body)

	htppCode, body = nginxServer.handleRequest(appStatusURL, "GET")
	fmt.Printf("\nURL:%s httpCode: %d body: %s\n", appStatusURL, htppCode, body)

	htppCode, body = nginxServer.handleRequest(appStatusURL, "GET")
	fmt.Printf("\nURL:%s httpCode: %d body: %s\n", appStatusURL, htppCode, body)

	htppCode, body = nginxServer.handleRequest(createUserURL, "POST")
	fmt.Printf("\nURL:%s httpCode: %d body: %s\n", appStatusURL, htppCode, body)

	htppCode, body = nginxServer.handleRequest(createUserURL, "POST")
	fmt.Printf("\nURL:%s httpCode: %d body: %s\n", appStatusURL, htppCode, body)

}
