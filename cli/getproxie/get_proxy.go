package getproxie

import (
	"apiproxy/cli/filecreate"
	"log"
)

//func GetProxy() error {
//	country := "US"
//	url := fmt.Sprintf("https://api.asocks.com/v2/proxy/search?apiKey=%s&country=%s", api.ApiKey, country)
//	resp, err := http.Get(url)
//	if err != nil {
//		log.Printf("[ошибка в получении адреса]")
//		return err
//	}
//	GetBody(resp)
//
//	//if err := json.Unmarshal(data, &arr); err != nil {
//	//	log.Printf("[ошибка unmarshal]")
//	//}
//	return nil
//}
//
//func GetBody(resp *http.Response) {
//	data, err := io.ReadAll(resp.Body)
//	if err != nil {
//		log.Printf("[ошибка чтения тела - body]")
//	}
//	defer resp.Body.Close()
//	BodyUnmarshal(&data)
//}

func BodyUnmarshal() {
	arr := []string{"121221", "saslasl", "sasas", "434234", "sak210323", "569ygjfi", "4023403"}
	//if err := json.Unmarshal(*data, &arr); err != nil {
	//	log.Printf("[ошибка unmarshal]")
	//}
	//
	if ok := filecreate.CreateDir(); ok {
		log.Printf("[директория создана]")
	}
	filecreate.PayloadInFile(&arr)
}
