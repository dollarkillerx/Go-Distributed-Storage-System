/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-13
* Time: 上午11:01
* */
package test

import (
	"Go-Distributed-Storage-System/utils"
	"log"
	"strconv"
	"testing"
)

func TestOne(t *testing.T)  {
	i, _ := strconv.Atoi(utils.TimeGetNowTimeStr())
	log.Println(i)
}


