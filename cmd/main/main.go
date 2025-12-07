// package main

// import (
// 	"github.com/akhil/go-bookstore/pkg/routes"
// 	"github.com/akhil/go-bookstore/pkg/config"
// 	"log"
// 	"net/http"
// 	"github.com/gorilla/mux"
// 	_ "github.com/jinzhu/gorm/dialects/mysql"
// )

package main

import (
    "bookmanagementmysql/config"
    "bookmanagementmysql/routes"
    "github.com/gofiber/fiber/v2"
)

func main() {

    app := fiber.New()

    config.Connect()

    // Migrar tabela Book automaticamente
    config.DB.AutoMigrate(&models.Book{})

    routes.BookRoutes(app)

    app.Listen(":3000")
}
