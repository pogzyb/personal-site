package website

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path"
	"sync"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	once 		sync.Once
	database 	*gorm.DB
)

type (
	Article struct {
		ID 				uint		`gorm:"primaryKey"`
		Title			string		`json:"title"`
		Path			string		`json:"path"`
		Date			*time.Time	`json:"date"`
		Template		string		`json:"template"`
		HTMLTemplate	string		`json:"html_template"`
		FAIcon			string		`json:"fa_icon"`
	}

	Project struct {
		ID				uint		`gorm:"primaryKey"`
		Title			string		`json:"title"`
		FAIcon			string		`json:"fa_icon"`
		Description		string		`json:"description"`
		Body			string		`json:"body"`
		Link			string		`json:"link"`
	}
)

func initDatabase() {
	once.Do(func() {
		var err error
		database, err = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
		if err != nil {
			log.Fatalf("could not connect to database: %v", err)
		}
	})
}

func migrateDatabase() {
	err := database.AutoMigrate(&Article{}, &Project{})
	if err != nil {
		log.Fatalf("could not run gorm schema migrations: %v", err)
	}
	populateArticles()
	populateProjects()
}


// todo: DRY out "populate" code
func populateArticles() {
	// check if database is empty
	var article Article
	result := database.First(&article)

	if result.RowsAffected == 0 {
		// struct for marshalling articles from json file
		var data = struct {
			Articles []Article `json:"articles"`
		}{}
		articlesDir := os.Getenv("DATA_DIR")
		articlesFile := os.Getenv("ARTICLES_FILENAME")
		fileLocation := path.Join(articlesDir, articlesFile)
		f, err := os.Open(fileLocation)
		if err != nil {
			log.Fatalf("could not open json file [%s]: %v", fileLocation, err)
		}
		defer f.Close()
		// marshal the results into "data", which is a slice of "Article" types
		j, _ := ioutil.ReadAll(f)
		err = json.Unmarshal(j, &data);
		if err != nil {
			log.Fatalf("could not unmarhsall json: %v", err)
		}
		// insert each article into database
		for _, article := range data.Articles {
			database.Create(&article)
			log.Printf("successfully inserted article: %s", article.Title)
		}
	}
}

func populateProjects() {
	// check if database is empty
	var project Project
	result := database.First(&project)

	if result.RowsAffected == 0 {
		// struct for marshalling the json file
		var data = struct {
			Projects []Project `json:"projects"`
		}{}
		// read in projects data from file
		projectsDir := os.Getenv("DATA_DIR")
		projectsFile := os.Getenv("PROJECTS_FILENAME")
		fileLocation := path.Join(projectsDir, projectsFile)
		f, err := os.Open(fileLocation)
		if err != nil {
			log.Fatalf("could not open json file [%s]: %v", fileLocation, err)
		}
		defer f.Close()
		// marshal the results into "data", which is a slice of "Project" types
		j, _ := ioutil.ReadAll(f)
		err = json.Unmarshal(j, &data);
		if err != nil {
			log.Fatalf("could not unmarhsall json: %v", err)
		}
		// insert each article into database
		for _, project := range data.Projects {
			database.Create(&project)
			log.Printf("successfully inserted project: %s", project.Title)
		}
	}
}