package models

// InputData хранит параметры ввода
type InputData struct {
	Value1 string
	Value2 string
	Value3 string
}

var CurrentInput InputData

func init() {
	CurrentInput = InputData{
		Value1: "500",
		Value2: "right",
		Value3: "Terminal",
	}
}
