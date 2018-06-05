package util

import (
	"net/url"
	"github.com/catchdoll/conf"
	"crypto/md5"
	"encoding/hex"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

func GetLbsSN( uri string, params url.Values )(string, error){
	var urlPack *url.URL

	urlPack, err := url.Parse(uri)
	//urlPack.Scheme = "http"
	if err != nil {
		panic("invalid url")
	}
	urlPack.RawQuery = params.Encode()
	encoded := urlPack.String()+conf.GlobalConf.LbsSk
	//fmt.Println("to be encoded:",encoded)
	//fmt.Println("encoded:",url.QueryEscape(encoded))
	hasher := md5.New()
	hasher.Write([]byte(url.QueryEscape(encoded)))
	return hex.EncodeToString(hasher.Sum(nil)), nil
}

func CreateLbsAddress(address string, longitude string, latitude string)(string, error){
	uri := "/geodata/v4/poi/create"
	form := url.Values{
		"ak":{"6wVrtUncrfIq5SE4AlYtrCRZtDtsT1kP"},
		//"id":{"190008"},
		//"name":{"doll_machine_address"},
		"coord_type":{"3"},
		"geotable_id":{"1000003990"},
		"latitude":{"22.56124"},
		"longitude":{"114.106875"},
		"title":{"复制"},
	}
	SNCode, err := GetLbsSN(uri,form)
	if err != nil{
		return "", err
	}
	form.Add("sn",SNCode)
	resp, err := http.PostForm("http://api.map.baidu.com/geodata/v4/poi/create",form)
	if err != nil{
		return "", err
	}
	result, err := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(result))
	var jsonReturn LbsCreateAddrReturn
	err = json.Unmarshal(result, &jsonReturn)
	if err != nil{
		return "", err
	}
	return jsonReturn.id, nil

}

type(
	LbsCreateAddrReturn struct {
		status uint32
		message string
		id string
	}
)

