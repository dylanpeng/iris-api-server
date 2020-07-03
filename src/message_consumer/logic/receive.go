package logic

import (
	"fmt"
	"juggernaut/common"
)

func ReceiveMessage(msg []byte) error {
	defer common.CatchPanic()

	common.Logger.Debugf("receive message | msg: %s", string(msg))

	fmt.Printf("receive message1: %s \n", msg)

	return nil
}

func ReceiveMessage2(msg []byte) error {
	defer common.CatchPanic()

	common.Logger.Debugf("receive message | msg: %s", string(msg))

	fmt.Printf("receive message2: %s \n", msg)

	return nil
}