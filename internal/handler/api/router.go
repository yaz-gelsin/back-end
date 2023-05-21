package api

import (
	"fmt"
	"net/http"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello World!")
}
