package main

import (
	"github.com/chrismeyers/petstore/http"
	"github.com/chrismeyers/petstore/postgres"
)

func main() {
	db, err := postgres.Open()

	if err != nil {
		panic(err)
	}

	ps := postgres.PetService{DB: db}
	pHandler := http.NewPetHandler(&ps)
	handler := http.Handler{PetHandler: pHandler}

	s := http.NewServer()
	s.Handler = &handler
	s.ListenAndServe()
}
