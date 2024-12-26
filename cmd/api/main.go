package main

import "net/http"

 type application struct{
	addr string
 }

func (app *application) getUserHandler() {
	w.Write([]byte("hi"))
}
func (app *application) createUserHandler() {
	w.Write([]byte("create"))
}

 func main(){
		s:= &application{addr:":8000"}

		mux := http.NewServeMux();

		srv := &http.Server{
			Addr: api.addr,
			Handler: mux,
		}
		mux.HandleFunc("GET /users", api.getUserHandler)
		mux.HandleFunc("POST /users",api.createUseraHandler)
		if err := http.ListenAndServe()
 } 