package repository

import (
	"fmt"
	"github.com/GoGinApi/v2/entity"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var (
	username = "goland"   //os.Getenv("db_user")
	password = "goland"   //os.Getenv("db_pass")
	dbName   = "goland"   //os.Getenv("db_name")
	dbHost   = "localhost" //os.Getenv("db_host")
)

type VideoRepository interface {
	Save(video entity.Video)
	Update(video entity.Video)
	Delete(video entity.Video)
	FindAll() []entity.Video
	CloseDB()
}

type Database struct {
	connection *gorm.DB
}

func (db Database) CloseDB() {
	err := db.connection.Close()
	if err != nil {
		panic("failed to close connection")
	}
}

func (db Database) Save(video entity.Video) {
	db.connection.Create(&video)
}

func (db Database) Update(video entity.Video) {
	db.connection.Save(&video)
}

func (db Database) Delete(video entity.Video) {
	db.connection.Delete(&video)
}

func (db Database) FindAll() []entity.Video {
	var videos []entity.Video
	db.connection.Set("gorm:auto_preload", true).Find(&videos)
	return videos
}

func NewVideoRepository() VideoRepository {
	e := godotenv.Load()
	if e != nil {
		fmt.Printf("Error while loading the env file%v", e)
	}

	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)
	fmt.Println(dbUri)

	db, err := gorm.Open("postgres", dbUri)
	if err != nil {
		fmt.Printf("Error while opening db connection%v", err)
	}

	db.AutoMigrate(&entity.Video{}, &entity.Person{})

	return &Database{
		connection: db,
	}
}
