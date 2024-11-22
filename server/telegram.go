package server

import (
	"fmt"
	"github.com/b2network/tools/httputil"
)

func SendTgMessage(bootId, chatId, msg string) {
	url := "https://api.telegram.org/" + bootId + "/sendMessage?chat_id=" + chatId + "&text=" + msg
	fmt.Println(url)
	b := httputil.HttpGet(url)
	fmt.Println(string(b))
}
