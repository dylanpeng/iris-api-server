package rocket_consumer

import (
	"fmt"
	"github.com/apache/rocketmq-client-go/v2/primitive"
)

func RocketConsumerTest(msg *primitive.MessageExt) error {
	fmt.Printf("in\n")
	return nil
}