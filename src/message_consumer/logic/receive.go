package logic

import (
	"fmt"
	"juggernaut/common"
)

func ReceiveMessage(msg []byte) error {
	defer common.CatchPanic()

	common.Logger.Debugf("receive message | msg: %s", string(msg))

	fmt.Printf("receive message: %s \n", msg)

	return nil
}
