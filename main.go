package main

import (
	"demo/src/data"
	prodDependcies "demo/src/products/dependencies"
	employeeDependencies "demo/src/employees/dependencies"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	
	mysql := data.NewMySQL()
	defer mysql.Close()
	
	router := gin.Default()
	prodDependcies := prodDependcies.NewProductDependencies(mysql.DB)
	prodDependcies.Execute(router)

	employeeDependencies := employeeDependencies.NewEmployeeDependencies(mysql.DB)
	employeeDependencies.Execute(router)

	log.Println("[Main] Servidor corriendo en http://localhost:8080")
	router.Run(":8080")
}
