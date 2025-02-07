package getProxy

import (
	"apiproxy/cli/api"
	"fmt"
	"io"
	"log"
	"net/http"
)

func GetProxy() error {
	country := "US"
	url := fmt.Sprintf("https://api.asocks.com/v2/proxy/search?apiKey=%s&country=%s", api.ApiKey, country)
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("[ошибка в получении адреса]")
		return err
	}
	GetBody(resp)

	//var arr []string
	//if err := json.Unmarshal(data, &arr); err != nil {
	//	log.Printf("[ошибка unmarshal]")
	//}
	return nil
}

func GetBody(resp *http.Response) {
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("[ошибка чтения тела - body]")
	}
	defer resp.Body.Close()
	fmt.Println(string(data))
}

//func BodyUnmarshal(data *[]byte) {
//	var arr []string
//	if err := json.Unmarshal(*data, &arr); err != nil {
//		log.Printf("[ошибка unmarshal]")
//	}
//
//	fileCreate.CreateFile(arr)
//}
