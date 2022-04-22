package main

import (
	"api/teamdev/controllers"
	_ "api/teamdev/docs"
	"api/teamdev/objects"
	"api/teamdev/utils"

	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func initDBConnection(cnf utils.DBConfiguration) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cnf.Host, cnf.User, cnf.Password, cnf.Name, cnf.Port)
	db, e := gorm.Open(cnf.Type, dsn)

	if e != nil {
		utils.Logger.Print("DB Connection failed")
		utils.Logger.Fatal(e)
	} else {
		utils.Logger.Print("DB Connection Established")
	}

	db.SingularTable(true)

	db.AutoMigrate(
		&objects.Recipe{}, &objects.Account{}, &objects.Category{},
		&objects.Ingredient{}, &objects.Step{},
	)

	return db
}

// @Title TeamDev API
// @Version 0.1
// @Description API for culinary recipes (BMSTU Team development project)
// @securityDefinitions.basic BasicAuth
func main() {
	rand.Seed(time.Now().UnixNano())

	if len(os.Args) == 1 {
		utils.InitConfig()
	} else {
		utils.InitConfig(os.Args[1])
	}

	utils.InitLogger()
	defer utils.CloseLogger()

	db := initDBConnection(utils.Config.DB)
	defer db.Close()

	r := controllers.InitRouter(db)
	controllers.RunSwagger(r)
	utils.Logger.Print("Server started")
	fmt.Printf("Server is running on http://localhost:%d\n", utils.Config.Port)
	code := controllers.RunRouter(r, utils.Config.Port)

	utils.Logger.Printf("Server ended with code %s", code)
}
