package server

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/b2network/tools/httputil"
	"github.com/b2network/tools/initialize"
	"github.com/b2network/tools/types"
	"k8s.io/klog/v2"
)

func GetNewBlockNumByRpc(domain string) (string, error) {
	var blockNumStr types.BlockNumResponse
	b := httputil.HttpPost("https://"+domain, `{"method":"eth_getBlockByNumber","params":["latest", false],"id":1,"jsonrpc":"2.0"}`)
	err := json.Unmarshal(b, &blockNumStr)
	if err != nil {
		klog.Errorf("从rpc接口获取blocknum失败 domain: %v err: %v \n", domain, err)
		return "", err
	}
	fmt.Printf("%v 新区块 %v \n", domain, blockNumStr.Result.Number)
	//blockNum, _ := strconv.ParseInt(strings.Split(blockNumStr.Result.Number, "0x")[1], 16, 0)
	//fmt.Printf("%v 新区块 %v \n", domain, blockNum)
	return blockNumStr.Result.Number, nil
}
func GetOldBlockNumByRedis(ctx context.Context, rc *initialize.RedisClient, domain string) (string, error) {
	blockNumStr, err := rc.GetKey(ctx, domain)
	if err != nil {
		klog.Errorf("从redis获取blocknum失败 domain: %v err: %v \n", domain, err)
		return "", err
	}
	fmt.Printf("%v 老区块 %v \n", domain, blockNumStr)
	//blockNum, _ := strconv.ParseInt(strings.Split(blockNumStr, "0x")[1], 16, 32)
	//fmt.Printf("%v 老区块 %v \n", domain, blockNum)
	return blockNumStr, nil
}

func ComPareBlockNum(ctx context.Context, rc *initialize.RedisClient, domain string) (string, string, bool) {
	newblockNum, _ := GetNewBlockNumByRpc(domain)

	oldBlockNum, _ := GetOldBlockNumByRedis(ctx, rc, domain)
	if oldBlockNum == "" || newblockNum == "" || oldBlockNum == newblockNum {
		return oldBlockNum, newblockNum, false
	}
	return oldBlockNum, newblockNum, true
}
