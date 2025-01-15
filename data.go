package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Iot struct {
	Id     int64
	Time   time.Time
	Value  float64
	Offset float64
	CID    string
}

func (me *Iot) ValueOf(c int) interface{} {
	switch c {
	case 0:
		return me.Id
	case 1:
		return me.Time.Format("2006-01-02 15:04:05")
	case 2:
		return me.Value
	case 3:
		return me.Offset
	case 4:
		return me.CID
	}
	return ""
}

func fetchLatestIotData() ([]Iot, error) {

	db, err := OpenDb()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query := "SELECT id, time, value, offset, cid FROM test ORDER BY time DESC LIMIT 100"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var iots []Iot
	for rows.Next() {
		var iot Iot
		var timeStr string
		if err := rows.Scan(&iot.Id, &timeStr, &iot.Value, &iot.Offset, &iot.CID); err != nil {
			return nil, err
		}
		iot.Time, err = time.Parse("2006-01-02 15:04:05", timeStr)
		if err != nil {
			return nil, err
		}
		iots = append(iots, iot)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return iots, nil
}

func OpenDb() (*sql.DB, error) {

	dsn := "root:123456@tcp(192.168.1.128:3306)/test"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("Error opening database:", err)
		return nil, err
	}
	return db, nil
}
