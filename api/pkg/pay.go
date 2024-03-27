package pkg

import (
	"fmt"
	"github.com/smartwalle/alipay/v3"
	"github.com/spf13/viper"
	"net/url"
)

//func getClient() (*alipay.Client, error) {
//	viper.SetConfigFile("./etc/shop.yaml")
//	viper.ReadInConfig()
//	appID := viper.GetString("AppID")
//	privateKey := viper.GetString("PrivateKey")
//	publicKey := viper.GetString("PublicKey")
//	client, err := alipay.New(appID, privateKey, false)
//	if err != nil {
//		return nil, err
//	}
//	err = client.LoadAliPayPublicKey(publicKey)
//	if err != nil {
//		return nil, err
//	}
//	return client, nil
//}
//
//func GetWebPayUrl(title, orderNo, amount string) (string, error) {
//	viper.SetConfigFile("./etc/shop.yaml")
//	viper.ReadInConfig()
//	notifyUrl := viper.GetString("NotifyURL")
//	returnUrl := viper.GetString("ReturnURL")
//	var p = alipay.TradePagePay{
//		Trade: alipay.Trade{
//			NotifyURL:   notifyUrl,
//			ReturnURL:   returnUrl,
//			Subject:     title,
//			OutTradeNo:  orderNo,
//			TotalAmount: amount,
//			ProductCode: "FAST_INSTANT_TRADE_PAY",
//		},
//	}
//	cli, err := getClient()
//	if err != nil {
//		return "", err
//	}
//	pay, err := cli.TradePagePay(p)
//	if err != nil {
//		return "", err
//	}
//	return pay.String(), nil
//}
//
//func VerifySign(values url.Values) (*alipay.Notification, error) {
//	cli, err := getClient()
//	if err != nil {
//		return nil, err
//	}
//	return cli.DecodeNotification(values)
//}

func getClient() (*alipay.Client, error) {
	viper.SetConfigFile("./etc/shop.yaml")
	viper.ReadInConfig()
	appID := viper.GetString("AppID")
	privateKey := viper.GetString("PrivateKey")
	publicKey := viper.GetString("PublicKey")
	client, err := alipay.New(appID, privateKey, false)
	if err != nil {
		return nil, err
	}
	err = client.LoadAliPayPublicKey(publicKey)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func GetWebPayUrl(title, orderNo, amount string) (string, error) {
	viper.SetConfigFile("./etc/shop.yaml")
	viper.ReadInConfig()
	notifyUrl := viper.GetString("NotifyURL")
	returnUrl := viper.GetString("ReturnURL")
	var p = alipay.TradeWapPay{
		Trade: alipay.Trade{
			NotifyURL:   notifyUrl,
			ReturnURL:   returnUrl,
			Subject:     title,
			OutTradeNo:  orderNo,
			TotalAmount: amount,
			ProductCode: "FAST_INSTANT_TRADE_PAY",
		},
	}

	client, err := getClient()
	if err != nil {
		return "", err
	}

	url, err := client.TradeWapPay(p)
	if err != nil {
		fmt.Println(err)
	}
	return url.String(), nil
}

func VerifySign(value url.Values) (*alipay.Notification, error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}
	return client.DecodeNotification(value)
}
