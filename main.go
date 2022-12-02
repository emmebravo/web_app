package main

import (
	"net/http"

	"github.com/emmebravo/web_app/controllers"
)

func main() {
	controllers.RegisterControllers()

	http.ListenAndServe(":3000", nil)
}
