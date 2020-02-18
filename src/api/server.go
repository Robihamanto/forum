package api

import (
	"api/routes"
	"fmt"
	"log"
	"net/http"
)

func listen(p int) {
	fmt.Println("Server starting at port:", p)
	port := fmt.Sprintf(":%d", p)
	r := routes.NewRouter()
	log.Fatal(http.ListenAndServe(port, r))
}
