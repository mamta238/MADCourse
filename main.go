
package main

import (
		"os"
		"net/http"
		"github.com/gorilla/mux"
		dbrepo "dbrepository"
		"usercrudhandler"
		logger "log"
		"mongoutils"
		)
		

func main(){
	
	h := mux.NewRouter()
	dbname := os.Args[1]
	mongoSession, _ := mongoutils.RegisterMongoSession(os.Getenv("MONGO_HOST"))
	repo := dbrepo.NewMongoRepository(mongoSession, dbname)
	usersrvc := dbrepo.NewService(repo)
	hndlr := usercrudhandler.NewUserCrudHandler(usersrvc)
	
	h.Handle("/user/",hndlr)
	h.Handle("/user/{id}",hndlr)
	logger.Println("Resource Setup Done.")
	logger.Fatal(http.ListenAndServe(":8080", h)) 

}
