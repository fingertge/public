// ***********************************************************************************************
// ***                                     G O L A N D                                         ***
// ***********************************************************************************************
// * Auth: ColeCai
// * Date: 2023/10/25 00:39:45
// * Proj: work
// * Pack: redis
// * File: redis.go
// *----------------------------------------------------------------------------------------------
// * Overviews:
// *----------------------------------------------------------------------------------------------
// * Functions:
// * - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
package redis

import (
	"github.com/fingertge/public/tools"
	redis2 "github.com/gomodule/redigo/redis"
)

func ScanStruct(mp map[string]string, dest interface{}) error {
	var values []interface{}
	for k, v := range mp {
		values = append(values, tools.StringToBytes(k), tools.StringToBytes(v))
	}
	return redis2.ScanStruct(values, dest)
}
