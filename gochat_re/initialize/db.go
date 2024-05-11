package initialize

import (
	"fmt"
	"github.com/joho/godotenv"
	"gochat_re/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"strconv"
	"time"
)

func InitDB() {
	err := godotenv.Load("config/.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	Host := os.Getenv("DATABASE_HOST")
	Port, err := strconv.Atoi(os.Getenv("DATABASE_PORT"))
	DBName := os.Getenv("DATABASE_NAME")
	User := os.Getenv("DATABASE_USER")
	PassWord := os.Getenv("DATABASE_PASSWORD")
	fmt.Println(Host, Port, DBName, User, PassWord)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", User, PassWord, Host, Port, DBName)
	Newlogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second, //记录慢查询
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)
	global.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: Newlogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		fmt.Sprintln("connect database error.")
		panic(err)
	}

}
