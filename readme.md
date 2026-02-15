# creating the module : 
go mod init <name-here>

- this will create the module, which will contain all the configuration and name of the packages used.
- this will also contain all the file names / dependencies used in the project.

# run command :

go run .

- this will find the main function go file with main package in it (with any name typically main.go) and execute it.

# package :
- a group of files in the same folder and same name of the package at the top of every go file is called a package.
- the compiler sees the package not the each file.