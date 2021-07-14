package main

import (
	app "github.com/nitinjangam/simplego/services"
)

func main() {
	a := app.App{}
	a.Initialize()
	defer a.DbConn.UserFile.Close()
	defer a.DbConn.TodoFile.Close()
	a.Run()
}
