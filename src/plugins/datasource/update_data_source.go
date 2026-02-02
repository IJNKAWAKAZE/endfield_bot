package datasource

import (
	"endfield_bot/plugins/skland"
	"endfield_bot/utils"
	"github.com/starudream/go-lib/core/v2/codec/json"
	"github.com/tidwall/gjson"
	"log"
)

// UpdateDataSource 更新数据源
func UpdateDataSource() {
	go UpdateDataSourceRunner()
}

// UpdateDataSourceRunner 更新数据源
func UpdateDataSourceRunner() {
	log.Println("开始更新数据源...")
	var operators []utils.Operator
	result, _ := getData("/web/v1/wiki/item/catalog?typeMainId=1&typeSubId=1")
	gjson.Parse(result).Get("data.catalog.0.typeSub.0.items").ForEach(func(key, value gjson.Result) bool {
		var operator utils.Operator
		operator.Name = value.Get("name").String()
		operator.ItemId = value.Get("itemId").String()
		operator.Cover = value.Get("brief.cover").String()
		// 获取立绘
		result, _ := getData("/web/v1/wiki/item/info?id=" + operator.ItemId)
		illustration := gjson.Parse(result).Get("data.item.document.extraInfo.illustration").String()
		operator.Illustration = illustration
		operators = append(operators, operator)
		return true
	})
	utils.RedisSet("operatorList", json.MustMarshalString(operators), 0)
	log.Println("数据源更新完毕")
}

func getData(url string) (string, error) {
	account := skland.Account{}
	tokenData, err := skland.SklandRequestPlayerData(skland.SKR(), "GET", "/web/v1/auth/refresh", account.Skland)
	if err != nil {
		return "", err
	}
	account.Skland.Token = gjson.Parse(tokenData).Get("data.token").String()
	data, err := skland.SklandRequestPlayerData(skland.SKR(), "GET", url, account.Skland)
	if err != nil {
		return "", err
	}
	return data, nil
}
