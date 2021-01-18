package main

import (
	"encoding/json"
	"fmt"
	"time"
)
// json структура
type jsonType struct {
	Status string `json:"Status"`
	TimeStart string `json:"TimeStart"`
	TimeEnd string `json:"TimeEnd"`
	TimeNow string `json:"TimeNow"`
}

// создание json
func blacklight() string  {
	var dateStart = time.Date(
						time.Now().Year(),
						time.Now().Month(),
						time.Now().Day(),
						9,
						0,
						0,
						0,
						time.Local)
	var dateEnd = time.Date(
						time.Now().Year(),
						time.Now().Month(),
						time.Now().Day(),
						17,
						0,
						0,
						0,
						time.Local)

	item := jsonType{
			Status: getStatus(dateStart,dateEnd),
			TimeStart: dateStart.String(),
			TimeEnd: dateEnd.String(),
			TimeNow: time.Now().String()}

	jitem, err := json.Marshal(item)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	fmt.Println(string(jitem))
	return string(jitem)
}

// получить режим работы лампы
func getStatus(start,end time.Time) string {
	if start.Before(time.Now().Local()) && end.After(time.Now().Local()) {
		return "ON"
	} else {
		return "OFF"
	}
}