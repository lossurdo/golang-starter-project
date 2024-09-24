package cmd

import (
	"database/sql"
	"os"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/lossurdo/golang-starter-project/utils"
	"github.com/spf13/cobra"
)

var mysqlCmd = &cobra.Command{
	Use:   "mysql",
	Short: "A MySQL example",
	Run: func(cmd *cobra.Command, args []string) {
		key := os.Getenv("KEY")
		pass := os.Getenv("PASSWORD")
		conn := strings.Replace(os.Getenv("CONNECTION"), "xxx", utils.Decrypt(pass, key), 1)

		db, err := sql.Open("mysql", conn)
		if err != nil {
			panic(err)
		}
		db.SetConnMaxLifetime(time.Minute * 3)
		db.SetMaxOpenConns(10)
		db.SetMaxIdleConns(10)
		defer db.Close()

		myField := "myValue"
		rows, err := db.Query("SELECT id, data FROM table WHERE field=?", myField)
		if err != nil {
			panic(err)
		}
		defer rows.Close()

		var id int
		var data string
		for rows.Next() {
			err := rows.Scan(&id, &data)
			if err != nil {
				panic(err)
			}
		}

		if id == 0 || data == "" {
			panic("No data found")
		}

		_, err = db.Exec("UPDATE table SET field1 = ? WHERE field2 = ?", myField, myField)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(mysqlCmd)
}
