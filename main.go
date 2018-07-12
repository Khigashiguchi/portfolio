package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"

	"github.com/Khigashiguchi/profile-site/presenter/handler"
	"github.com/Khigashiguchi/profile-site/storage"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	// Open Database
	db, err := sql.Open("sqlite3", "tmp/storage.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	storage.Migrate(db)

	// Routing
	router := mux.NewRouter()

	bh := handler.BlogHandler{DB: db}

	// router.Methods("GET").Path("/").Handler(http.FileServer(http.Dir("public/front")))
	// router.Methods("GET").Path("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	http.Redirect(w, r, "/web", http.StatusMovedPermanently)
	// })
	router.Methods("GET").Path("/").Handler(http.FileServer(http.Dir("public/admin")))
	router.Methods("GET").Path("/api/blogs").HandlerFunc(bh.GetBlogsHandler)
	router.Methods("POST").Path("/api/blog").HandlerFunc(bh.PostBlogHandler)

	// adminRouter := mux.NewRouter().PathPrefix("/admin").Subrouter().StrictSlash(true)
	// adminRouter.HandlerFunc("/", handler.AdminIndexHandler)
	// router.PathPrefix("/admin").Handler(negroni.New(
	// 	// TODO: Auth middleware
	// 	negroni.Wrap(adminRouter),
	// ))

	http.Handle("/", handlers.CombinedLoggingHandler(os.Stderr, router))

	// Start server
	fmt.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
