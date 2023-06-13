package main

import (
	"flag"
	"log"
	"os"
	"runtime"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/valeelim/mahchat/internal/cache"
	"github.com/valeelim/mahchat/internal/database"
	"github.com/valeelim/mahchat/pkg/controller"
	"github.com/valeelim/mahchat/pkg/middleware"
	"github.com/valeelim/mahchat/pkg/service"
	"github.com/valeelim/mahchat/seeds"

	_ "github.com/lib/pq"
)

func main() {
	handleArgs()

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
		log.Fatal("strconv stuff", err)
	}

	rdb, err := cache.New(rdbAddr, rdbPass, rdbDb)
	if err != nil {
		log.Fatal(err)
	}

	ctl := controller.New()

	registerUserService := service.RegisterUserService(db)
	loginUserService := service.LoginUserService(db, rdb)
	getAllUsersService := service.GetAllUsersService(db)

	createChannelService := service.CreateChannelService(db)
	getChannelByIDService := service.GetChannelByIDService(db)
	serveChannelService := service.ServeChannelService(db, rdb.Client)
	createGroupMessageService := service.CreateGroupMessageService(db)
	getGroupMessageByChannelIDService := service.GetGroupMessageByChannelIDService(db, db) // kinda lul

	router := gin.Default()

	router.POST("/register", ctl.RegisterController(registerUserService))
	router.POST("/login", ctl.LoginController(loginUserService))
	router.GET("/users", middleware.Authorized(rdb), ctl.GetAllUsersController(getAllUsersService))

	router.POST("/channels", ctl.CreateChannelController(createChannelService))
	router.GET("/channels/:id", ctl.GetChannelByIDController(getChannelByIDService))

	router.GET("/ws/channels/:id", ctl.ServeChannelController(serveChannelService))

	router.POST("/group-messages", ctl.CreateGroupMessageController(createGroupMessageService))
	router.GET("/group-messages/:id", ctl.GetGroupMessageByChannelIDController(getGroupMessageByChannelIDService))
	router.GET("/memory-usage", func(c *gin.Context) {
		var m runtime.MemStats

		runtime.GC()

		runtime.ReadMemStats(&m)
		// General memory statistics
		memAllocated := m.Alloc
		memHeapAlloc := m.HeapAlloc
		memHeapSys := m.HeapSys
		memHeapObjects := m.HeapObjects

		// Garbage collector statistics
		numGC := m.NumGC
		gcTime := m.PauseTotalNs

		c.JSON(200, gin.H{
			"AllocatedMemory":              memAllocated,
			"TotalMemoryAllocatedNotFreed": memHeapAlloc,
			"HeapMemorySize":               memHeapSys,
			"NumberOfHeapObjects":          memHeapObjects,
			"NumberOfGCCycles":             numGC,
			"TotalTimeSpentInGC":           gcTime,
		})
	})

	router.Run("localhost:8080")
}

func handleArgs() {
	flag.Parse()
	args := flag.Args()

	if len(args) >= 1 {
		switch args[0] {
		case "seed":
			seeds.Seed()
			os.Exit(0)
		}
	}
}
