package main

import (
	"bufio"
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/cuigh/auxo/util/lazy"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/masx200/to-do-list-go-sql-vue/backend/configs"
	"github.com/masx200/to-do-list-go-sql-vue/backend/database"
	"github.com/masx200/to-do-list-go-sql-vue/backend/models"
	"github.com/masx200/to-do-list-go-sql-vue/backend/routers"
	"gorm.io/gorm"
)

func main() {
	var create_sql = ""
	var LoadConfig = configs.LoadConfig
	var model = &models.ToDoItem{}
	var configfile = ""
	// 绑定命令行参数到变量上
	flag.StringVar(&configfile, "config_file", "", "Path to the configuration file.")
	flag.StringVar(&create_sql, "create_sql", "", "SQL statement for creating database objects.")

	// 解析命令行参数
	flag.Parse()

	// 使用传入的参数
	if configfile == "" || create_sql == "" {
		fmt.Println("Error: Both -config_file and -create_sql flags are required.")
		fmt.Println("Usage: yourprogram -config_file path/to/config -create_sql 'your sql statement'")
		return
	}

	fmt.Println("Configuration File Path:", configfile)
	fmt.Println("Create SQL Statement:", create_sql)
	config := LoadConfig(configfile)

	var lazyDB = lazy.Value[*gorm.DB]{New: func() (db *gorm.DB, err error) {

		db, err = database.ConnectDatabase(config.Dsn, &models.ToDoItem{}, "to_do_items", config.Debug)
		return
	}}

	if !config.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	var GetDB = func() (*gorm.DB, error) {
		var db, err = lazyDB.Get()
		if err != nil {
			return nil, err
		}
		if db == nil {
			return nil, errors.New("db is  nil pointer")
		}
		return db, err
	}

	db, err := sql.Open("mysql", config.Dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	file, err := os.Open(create_sql)
	if err != nil {
		fmt.Println("Error opening file:", err)
		log.Fatal(err)
		return
	}
	defer file.Close()
	var sqlinit string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		sqlinit += scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		log.Fatal(err)
	}
	for _, sqline := range strings.Split(sqlinit, ";") {
		sqline = strings.TrimSpace(sqline)
		if len(sqline) != 0 {
			fmt.Println(sqline)
			result, err := db.Exec(sqline)
			if err != nil {
				log.Fatal(err)
			}
			log.Println("LastInsertId")
			log.Println(result.LastInsertId())
			log.Println("RowsAffected")
			log.Println(result.RowsAffected())
		}

	}

	r := gin.Default()
	r.Use(gzip.Gzip(gzip.DefaultCompression))
	routers.TodoRoute(r, GetDB, "/todoitem", model)
	r.Run(":" + strconv.Itoa(config.Port))

}
