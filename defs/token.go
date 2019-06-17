/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-14
* Time: 下午4:08
* */
package defs

import "Go-Distributed-Storage-System/utils"

func GetToken(username string) string {
	// 这个感觉不建议阿! 标准是 头 + 载荷 + 签名 前面头和载荷是base64编码 签名是头+载荷通过特定算法签名的 防止用户修改

	// md5(username + timestamp + token_salt) + timestamp[:8]
	tim := utils.TimeGetNowTimeStr()
	tokenProfilx := utils.Md5Encode(username + tim + "_tokensalt")
	token := tokenProfilx + tim[:8]
	return token
}
