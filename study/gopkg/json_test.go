package gopkg

import (
	"testing"

	"encoding/json"
	"fmt"
	"sort"
)

func TestJsonSort(t *testing.T) {

	map2 := getMapFromJson()
	fmt.Println(" map的长度：", len(map2))
	//1.定义一个slice
	s1 := make([]string, 0, len(map2))
	map3 := make(map[string]interface{})
	//2.遍历map获取key-->s1中
	for key := range map2 {
		s1 = append(s1, key)
	}
	//3.给s1进行排序
	//sort.Ints(s1) //使用sort包下的方法直接排序，不用自己写冒泡了。
	sort.Strings(s1)
	//4. 遍历s1，map
	for _, k := range s1 { // 先下标，再数值
		map3[k] = map2[k]
	}
	b, berror := json.Marshal(map3)
	if berror != nil {
		fmt.Print("berror:", berror)
	}
	fmt.Println(string(b))

}

func getMapFromJson() map[string]interface{} {
	var jsonBody = []byte(`{"akehi":"绝地求生","snifeni":["ddd","aaa"],"bndinfi":{"eee":"d","ccc":"c"},"zhondfi":"连连看"}`)
	var map4 = make(map[string]interface{})
	err := json.Unmarshal(jsonBody, &map4)
	if err != nil {
		println("unJsonerr:", err)
	} else {
		for k, v := range map4 {
			print("map4 k and Value :", k, v)
		}
	}
	return map4
}
