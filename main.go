package main

import (
	_ "MyGram/docs"
	"MyGram/router"
	"fmt"
)

const PORT = ":8087"

// @title API Documentation	MyGram
// @version 0.1
// @description Documentation MyGram using Gin Framework
// @description inorder to use this api there bearer token which some of function are needed
// @description First you need regis and login
// @description Then Click "Authorize" at right and there pop-up will be appear and input your token and it will reveal true token e.g "Bearer e4udqw923....."
// @description Finally you can use some of function that already state before
// @BasePath /
//@host localhost:8087

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  togi.mare@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

// @swagg.NoModels
func main() {
	fmt.Println("Ready Server")
	router.ServerReady().Run(PORT)
}
