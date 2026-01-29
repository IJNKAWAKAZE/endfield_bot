package utils

import "encoding/json"

type Operator struct {
	Name         string `json:"name"`         // 名字
	ItemId       string `json:"itemId"`       // 编号
	Cover        string `json:"cover"`        // 半身像
	Illustration string `json:"illustration"` // 立绘
}

func GetOperators() []Operator {
	var operators []Operator
	operatorsJson := RedisGet("operatorList")
	json.Unmarshal([]byte(operatorsJson), &operators)
	return operators
}
