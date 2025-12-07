// package routers

// import (
// 	"github.com/akhil/go-bookstore/pkg/controllers"
// 	"github.com/gorilla/mux"
// )

// var RegisterBookstoreRoutes = func(router *mux.Router) {
// 	router.HandleFunc("/book", controllers.CreateBook).Methods("POST")
// 	router.HandleFunc("/book", controllers.GetBook).Methods("GET")
// 	router.HandleFunc("/book/{bookId}", controllers.GetBookByID).Methods("GET")
// 	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
// 	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
// }
package routes

import (
    "bookmanagementmysql/controllers"
    "github.com/gofiber/fiber/v2"
)

func BookRoutes(app *fiber.App) {

    route := app.Group("/books")

    route.Get("/", controllers.GetBooks)
    route.Post("/", controllers.CreateBook)
    route.Get("/:id", controllers.GetBookByID)
    route.Put("/:id", controllers.UpdateBook)
    route.Delete("/:id", controllers.DeleteBook)
}
