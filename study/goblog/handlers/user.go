package handlers

import (
	"github.com/gorilla/mux"
	"io"
	"net/http"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	//Get user from DB by id
	params := mux.Vars(r)
	id := params["id"]
	io.WriteString(w, "return user info with id="+id)
}
