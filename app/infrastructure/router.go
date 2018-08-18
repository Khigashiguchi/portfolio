package infrastructure

import (
    "net/http"
    "github.com/gorilla/mux"
    "github.com/Khigashiguchi/portfolio/app/presenter/handler"
    "database/sql"
)

// Router return http.Handler interface.
func Router(db *sql.DB) http.Handler {
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
    return router
}
