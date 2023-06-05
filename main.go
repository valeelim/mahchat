package main

import (
	"log"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/valeelim/mahchat/internal/cache"
	"github.com/valeelim/mahchat/internal/database"
	"github.com/valeelim/mahchat/pkg/controller"
	"github.com/valeelim/mahchat/pkg/middleware"
	"github.com/valeelim/mahchat/pkg/service"

	_ "github.com/lib/pq"
)

func main() {

	var (
		dbUser = os.Getenv("DB_USER")
		dbPort = os.Getenv("DB_PORT")
		dbPass = os.Getenv("DB_PASS")
		dbName = os.Getenv("DB_NAME")
		dbHost = os.Getenv("DB_HOST")

		rdbAddr  = os.Getenv("RDB_ADDR")
		rdbPass  = os.Getenv("RDB_PASS")
		rdbDBstr = os.Getenv("RDB_DB")
	)

	db, err := database.New(dbUser, dbPort, dbPass, dbName, dbHost)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	rdbDb, err := strconv.Atoi(rdbDBstr)
	if err != nil {
		log.Fatal(err)
	}

	rdb, err := cache.New(rdbAddr, rdbPass, rdbDb)
	if err != nil {
		log.Fatal(err)
	}

	ctl := controller.New()

	registerUserService := service.RegisterUserService(db)
	loginUserService := service.LoginUserService(db, rdb)
	getAllUsersService := service.GetAllUsersService(db)

	router := gin.Default()

	router.POST("/register", ctl.RegisterController(registerUserService))
	router.POST("/login", ctl.LoginController(loginUserService))

	router.GET("/users", middleware.Authorized(rdb), ctl.GetAllUsersController(getAllUsersService))

	router.Run("localhost:8080")
}
