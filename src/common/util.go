package common

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gogo/protobuf/proto"
	"juggernaut/common/coder"
	"juggernaut/common/grpc"
	"juggernaut/lib/kafka"
	"runtime/debug"
	"strings"
	"time"
)

func convertStruct(a interface{}, b interface{}) error {
	data, err := json.Marshal(a)

	if err != nil {
		return err
	}

	err = json.Unmarshal(data, b)

	if err != nil {
		return err
	}

	return nil
}

func GetKey(prefix string, items ...interface{}) string {
	format := prefix + strings.Repeat(":%v", len(items))
	return fmt.Sprintf(format, items...)
}

func ConvertStruct(a interface{}, b interface{}) error {
	err := convertStruct(a, b)

	if err != nil {
		Logger.Debugf("convert data failed | data: %s | error: %s", a, err)
	}

	return err
}

func ConvertStructs(items ...fmt.Stringer) (err error) {
	for i := 0; i < len(items)-1; i += 2 {
		if err := ConvertStruct(items[i], items[i+1]); err != nil {
			return err
		}
	}

	return
}

func CatchPanic() {
	if err := recover(); err != nil {
		Logger.Fatalf("catch panic | %s\n%s", err, debug.Stack())
	}
}

func CallGrpcWithTimeout(method string, req, rsp proto.Message, delay time.Duration, alias ...string) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), delay)
	defer cancel()

	items := strings.Split(method, "/")

	if len(items) < 3 {
		return errors.New("undefined grpc service")
	}

	srvItems := strings.Split(items[1], ".")

	if len(srvItems) < 2 {
		return errors.New("undefined grpc service")
	}

	srvName := strings.Join(srvItems[:len(srvItems)-1], "_")

	if len(alias) > 0 && alias[0] != "" {
		srvName = alias[0]
	}

	srvConf, ok := grpc.GetCfg().Servers[srvName]

	if !ok {
		return errors.New("undefined grpc service")
	}

	err = grpc.CallAddr(srvConf.GetAddr(), ctx, method, req, rsp)

	if err != nil {
		Logger.Warningf("call grpc service failed | method: %s | req: %s | error: %s", method, req, err)
		return
	}

	Logger.Debugf("call grpc service | method: %s | req: %s | rsp: %s", method, req, rsp)

	return
}

func CreateKaMsg(producer *kafka.Producer, topic string, data interface{}) error {
	message, e := coder.JsonCoder.Marshal(data)

	if e != nil {
		return e
	}

	return producer.Send(topic, message)
}
