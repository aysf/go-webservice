package main

import (
	"go-getting-started/controllers"
	"net/http"
)

func main() {
	controllers.RegisterControllers()

	http.ListenAndServe(":3000", nil)
}
