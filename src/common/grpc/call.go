package grpc

import (
	"context"
	"errors"
	"github.com/gogo/protobuf/proto"
	oGrpc "google.golang.org/grpc"
	"juggernaut/lib/grpc"
	"strconv"
	"strings"
	"time"
)

const DefaultPoolSize = 256

type Config struct {
	Idle     time.Duration      `toml:"idle"`
	Ttl      time.Duration      `toml:"ttl"`
	PoolSize int                `toml:"pool_size"`
	Servers  map[string]*Server `toml:"servers"`
}

type Server struct {
	Host string `toml:"host"`
	Port int    `toml:"port"`
}

func (s *Server) GetAddr() string {
	return s.Host + ":" + strconv.Itoa(s.Port)
}

var srvConfig *Config
var grpcPool *grpc.Pool

func Dialer(addr string) (*oGrpc.ClientConn, error) {
	return oGrpc.Dial(addr, oGrpc.WithInsecure())
}

func Init(c *Config) {
	srvConfig = c

	if c.PoolSize <= 0 {
		c.PoolSize = DefaultPoolSize
	}

	grpcPool = grpc.NewPool(Dialer, c.PoolSize, c.Idle*time.Second, c.Ttl*time.Second)
}

func CallAddr(addr string, ctx context.Context, method string, req, rsp proto.Message, opts ...oGrpc.CallOption) (err error) {
	conn, err := grpcPool.Get(addr)

	if err != nil {
		return
	}

	defer conn.Close()

	return conn.GetConn().Invoke(ctx, method, req, rsp, opts...)
}

func Call(ctx context.Context, method string, req, rsp proto.Message, opts ...oGrpc.CallOption) (err error) {
	items := strings.Split(method, "/")

	if len(items) < 3 {
		return errors.New("undefined grpc service")
	}

	srvItems := strings.Split(items[1], ".")

	if len(srvItems) < 2 {
		return errors.New("undefined grpc service")
	}

	srvName := strings.Join(srvItems[:len(srvItems)-1], "_")
	srvConf, ok := srvConfig.Servers[srvName]

	if !ok {
		return errors.New("undefined grpc service")
	}

	return CallAddr(srvConf.GetAddr(), ctx, method, req, rsp, opts...)
}

func GetCfg() *Config {
	return srvConfig
}
