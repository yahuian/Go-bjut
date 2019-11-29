package card

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/YahuiAn/Go-bjut/logger"

	"github.com/spf13/viper"
)

// 一个特别好用的Json-to-Go的在线工具
// https://mholt.github.io/json-to-go/
type telephone struct {
	Mobile     string `json:"mobile"`
	Nationcode string `json:"nationcode"`
}

type payload struct {
	Ext    string    `json:"ext"`
	Extend string    `json:"extend"`
	Params []string  `json:"params"`
	Sig    string    `json:"sig"`
	Sign   string    `json:"sign"`
	Tel    telephone `json:"tel"`
	Time   int64     `json:"time"`
	TplID  int       `json:"tpl_id"`
}

type outcome struct {
	Result int    `json:"result"`
	ErrMsg string `json:"errmsg"`
	Ext    string `json:"ext"`
	Fee    int    `json:"fee"`
	Sid    string `json:"sid"`
}

// 指定腾讯云模板单发短信
// 本例模板内容：{stuNumber}，您的一卡通在{location}，好心人为{registrant}
func sendMessage(telNumber, stuNumber, location, registrant string) bool {
	// 构造url
	sdkappid := viper.GetString("tencent.card_sms.sdkappid")
	random := string(rand.Int())
	url := "https://yun.tim.qq.com/v5/tlssmssvr/sendsms?sdkappid=" + sdkappid + "&random=" + random

	// 构造参数
	params := []string{stuNumber, location, registrant}
	appkey := viper.GetString("tencent.card_sms.appkey")
	nowTime := time.Now().Unix()
	sig := calculateSig(appkey, telNumber, random, nowTime)
	if sig == "" {
		logger.Error.Println("sig错误")
		return false
	}
	sign := viper.GetString("tencent.card_sms.sign")
	tel := telephone{
		Mobile:     telNumber,
		Nationcode: "86", // 中国大陆
	}
	tplId := viper.GetInt("tencent.card_sms.tpl_id")

	requestBody, _ := json.Marshal(payload{
		Params: params,
		Sig:    sig,
		Sign:   sign,
		Tel:    tel,
		Time:   nowTime,
		TplID:  tplId,
	})

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		logger.Error.Println(err.Error())
		return false
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error.Println(err.Error())
		return false
	}

	var res outcome
	_ = json.Unmarshal(body, &res)

	if res.Result != 0 {
		logger.Error.Println(res)
		return false
	}
	return true
}

func calculateSig(appKey, telNumber, random string, time int64) string {
	h := sha256.New()
	plaintext := "appkey=" + appKey + "&random=" + random + "&time=" + strconv.FormatInt(time, 10) + "&mobile=" + telNumber
	if _, err := h.Write([]byte(plaintext)); err != nil {
		return ""
	}
	return hex.EncodeToString(h.Sum(nil))
}
