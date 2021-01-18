package main

import (
	"log"
	"strconv"
	"time"
)

type dataType struct {
	Id int
	Data string
	Hum float32
	Temp float32
	Light int8
}

type printData struct {
	Data []string
	Temp []float32
	Hum []float32
	Light []int8
}

const formatTime = "02.01.2006 15:04:05"

// подготавливаем данные для записи в базу
func add(newHum string,newTemp string,newLight string)  {
	data := time.Now().Format(formatTime)
	fHum, erH := strconv.ParseFloat(newHum,32)
	fTemp, erT := strconv.ParseFloat(newTemp,32)
	iLight, erL := strconv.Atoi(newLight)

	if (erH != nil || erT != nil || erL != nil) {
		log.Fatal("Error input data")
	} else {
		newData := dataType{0, data,float32(fHum),float32(fTemp),int8(iLight)}
		addDB(newData)
	}
}

// формируем структуру для вывода данных в графике
func writeData(dt []dataType) printData {
	pd := printData{}

	for _ ,v := range dt {
		pd.Data = append(pd.Data, v.Data)
		pd.Temp = append(pd.Temp, v.Temp)
		pd.Hum = append(pd.Hum, v.Hum)
		pd.Light = append(pd.Light, v.Light)

	}
	return pd
}


