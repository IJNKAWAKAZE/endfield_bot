package datasource

import (
	"endfield_bot/utils"
	"github.com/spf13/viper"
	"github.com/starudream/go-lib/core/v2/codec/json"
	"github.com/tidwall/gjson"
	"io"
	"log"
	"net/http"
)

// UpdateDataSource 更新数据源
func UpdateDataSource() {
	go UpdateDataSourceRunner()
}

// UpdateDataSourceRunner 更新数据源
func UpdateDataSourceRunner() {
	log.Println("开始更新数据源...")
	var operators []utils.Operator
	api := viper.GetString("api.skport")
	result, _ := getData(api + "/web/v1/wiki/item/catalog?typeMainId=1&typeSubId=1")
	gjson.ParseBytes(result).Get("data.catalog.0.typeSub.0.items").ForEach(func(key, value gjson.Result) bool {
		var operator utils.Operator
		operator.Name = utils.T2S(value.Get("name").String())
		operator.ItemId = value.Get("itemId").String()
		operator.Cover = value.Get("brief.cover").String()
		// 获取立绘
		result, _ := getData(api + "/web/v1/wiki/item/info?id=" + operator.ItemId)
		illustration := gjson.ParseBytes(result).Get("data.item.document.extraInfo.illustration").String()
		operator.Illustration = illustration
		operators = append(operators, operator)
		return true
	})
	utils.RedisSet("operatorList", json.MustMarshalString(operators), 0)
	log.Println("数据源更新完毕")
}

func getData(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Sk-Language", "zh_Hant")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	resBody, _ := io.ReadAll(res.Body)
	defer res.Body.Close()
	return resBody, nil
}
