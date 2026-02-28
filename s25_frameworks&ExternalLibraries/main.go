package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/hello", func(c *gin.Context) {
		// Return JSON response
		c.JSON(http.StatusOK, gin.H{
			"msg": "Hello from the gin server.",
		})
	})

	r.Run() // Because that line, line 18 run is staying there with the thread forever.When you start the server, the thread on that server stops there.
}

// Intro:

// now as to make the api, it is lot of work. so, we will use the tool called as GIN WEB FRAMEWORK.

// https://gin-gonic.com/

//So it will make your life easier for working with API for example what else we are we are not validating a lot of JSON.We cannot make like complex routing manually.
// also, we are not adding the cors here. we should have to add the header for cors manually. So the cross origin resource sharing that's a JavaScript security restriction that we need to add manually those headers in our API.
// But in this case, using these frameworks is easy.

// so, we can install it and read all the documentation to do the task and all.

// Define a simple GET endpoint
//   r.GET("/ping", func(c *gin.Context) {
//     // Return JSON response
//     c.JSON(http.StatusOK, gin.H{
//       "message": "pong",
//     })
//   })

// It seems clear because you don't receive the writer and the request.You have received something known as a context that inside you havethe request and the response.I have some quick functions.

// In this case, c.JSON is responding with JSON.You have c.TXT, you have a lot of utility functions, and then you run a server.
// So at the end, it's the same, butyou write less code if you use the Gin Web Framework.And you are filtering the method directly when you hook your handler.

// Usage :

// Steps :
// 1. install it using the command as in the current folder where go mod is done as:
// $ go get -u github.com/gin-gonic/gin

// now, what we do have is a change in our Mod file. We have a require with the list of all the dependencies that we have right now,because we installed it. So, it's changing go.mod actually.

// and in the go.mod, And here you have all the versions and also you have some actions check for upgrades.That is just vs code is helping me executing some CLI commands to do that andthe next step is to import it and that's all.

// 2. then import it as where to use say the main.go :
// import "github.com/gin-gonic/gin"

// 3. write the basic code for the GET method and any basic route and then in terminal run 'go run .'
// go to localhost:8080/hello and you will get the output.
// Because that line, line 18 run is staying there with the thread forever.When you start the server, the thread on that server stops there.
