package main

import (
	"database/sql"
	_ "database/sql"
	"fmt"
	_ "github.com/lib/pq"
)
var db *sql.DB
// подключение к БД
func initDB()  {
	connStr := "user=postgres password=123 dbname=test sslmode=disable"
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	}
}

// пишем в базу данных подговленные данные
func addDB(newData dataType)  {
	result, err := db.Exec("INSERT INTO home (data, temp, hum, light) VALUES($1, $2, $3, $4)",
									newData.Data, newData.Temp,newData.Hum,newData.Light)
	if err != nil {
		panic(err)
	}
	 result.RowsAffected()
}

// читаем с базы последнии 100 сообщений
func readSql() []dataType  {
	datas := []dataType{}
	query := fmt.Sprintf("select * from home ORDER BY id DESC LIMIT %v",100)
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}

	for rows.Next(){
		p := dataType{}
		err := rows.Scan(&p.Id, &p.Data, &p.Temp, &p.Hum, &p.Light)
		if err != nil{
			fmt.Println(err)
			continue
		}
		datas = append(datas, p)
	}
	return datas
}


