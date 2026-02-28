package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Define a simple GET endpoint
	r.GET("/ping", func(c *gin.Context) {
		// Return JSON response
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Start server on port 8080 (default)
	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
	r.Run()
}

// compiling the project :
/*
-> Compiling the project
		go build .
-> Compiling in one specific output folder
		go build . -o build/
-> Compiling for other platforms and OSs
		env GOOS=target-OS
    		GOARCH=target-architecture go build .
-> Compile and install
		go install .
*/

// The Build creates the .exe files for the windows and only the name file(file without extension name) for the mac and other platform.
// similary we can build in a particular folder as well.
// what is install?
// Go has a folder in your computer that is in the path. So if you're doing a CLI, this is only for CLIs. So if you want now your app to be available in the terminal everywhere,you install it And then you can run our app from anywhere.

// Packaging :
// What about packaging?
// Because that's the file there. I mean, if you're on a Mac/ windows, you know that to install an app, you typically download something, and then you need to put that into application folder.
// So do I need to just take that file as it is and put it there?
// No, that's packaging, and Go ends with a binary.That's all, the binary.
// It's your work to add a package for your operating system in case you need to.Because maybe you're going to deploy this in the cloud, and each provider willtell you the instructions on where to put the binary and how to do it.

// So you don't need to create the package.
// So, for example, the package is useful for embedding assets,because here, we have only the executable.
// But if we change the executable,our public folder will not go with the executable.
// So when we are building, we are just compiling the Go code. Our 'public' folder is not going to be in that executable file.
// So we are not embedding any assets.Does it make sense?Images, logo, PDF, whatever you need oryou are pointing in your code is completely separate from the binary.

// You can also embed assets into the binary using the 'go:embed' directive. More info on that here :
// https://gobyexample.com/embed-directive

// So for that, you need to create the package and add your embeddings.
// You can create the installer for Windows.There are a lot of tools out there that will let you do that.
// On the Mac, you need to learn how to create a DMG package.They can be actually just a package, or you can also create an installer.
// And something similar or different kind of packages forLinux based on the distribution, but the responsibility ends there.
