/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-11
* Time: 下午10:11
* */
package err

func ErrPanic(err error) {
	if err != nil {
		panic(err.Error())
	}
}
