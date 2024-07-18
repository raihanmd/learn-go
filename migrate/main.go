package main

import (
	"net/http"

	"github.com/raihanmd/dependency_injection/helper"
)

func main() {
	router := InitializedRouter()

	err := http.ListenAndServe(":3000", router)
	helper.PanicIfError(err)
}
