package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	"github.com/go-graphql-api-mysql-gorm/graph"
	"github.com/go-graphql-api-mysql-gorm/graph/generated"
	"github.com/go-graphql-api-mysql-gorm/graph/model"
)

var db *gorm.DB;
const defaultPort = "5050"

func initDB()  {
	var err error
	dataSourceName := "root:admin@tcp(localhost:3306)/test_db?parseTime=True"

	db, err := gorm.Open("mysql", dataSourceName)

	if err != nil {
		fmt.Println(err)
		panic("Failed to connect to database")
	}

	db.LogMode(true)

	// Create the database. This is a one-time step.
    // Comment out if running multiple times - You may see an error otherwise
    // db.Exec("CREATE DATABASE test_db")
    // db.Exec("USE test_db")

	// Migration to create tables for Order and Item schema
	db.AutoMigrate(&model.Order{}, &model.Item{})

}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	initDB()
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
