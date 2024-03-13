package config


import (
	"fmt"
	"dbo-backend-oviek/routes"

	"github.com/gin-gonic/gin"
	grm "gorm.io/gorm"

)
type SetDatabaseConfig struct {
	MainDatabaseGRM *grm.DB
}

func setConfiguration(configPath string) {
	Setup(configPath)
	/*db.SetupDB()*/
	gin.SetMode(GetConfig().Server.Mode)
}

func SetDatabase() *SetDatabaseConfig {
	return &SetDatabaseConfig{
		MainDatabaseGRM: SetupMainDBGORM(),
	}
}

func Run(configPath string) {

	setConfiguration(configPath)
	conf := GetConfig()
	setDatabase := SetDatabase()
	fmt.Println("SERVICE IS RUNNING ON " + conf.Server.Port)

	routes.Routers.Use(func(c *gin.Context) {
		c.Set("db", setDatabase.MainDatabaseGRM)
		c.Next()
	})
	routes.Setup()
	gin.SetMode(conf.Server.Mode)
	routes.Routers.Run(":" + conf.Server.Port)
}

var (
	Conf = GetConfig()
)
