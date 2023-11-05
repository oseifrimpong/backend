// include all the code that is neccessary to run the server
/*
 - godotenv
 - llms running check
 - router
 - logger
 - DB connection
 - vector store connection
*/
package server

import "os"

func InitializeServer() {
	//Router
	router := InitializeRouter()
	router.Run(":" + os.Getenv("APP_PORT"))
}
