package server

import (
	"github.com/go-kratos/kratos/v2/encoding/json"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"google.golang.org/protobuf/encoding/protojson"
)

func init() {
	json.MarshalOptions = protojson.MarshalOptions{
		UseEnumNumbers:  true, // Emit enum values as numbers instead of their string representation (default is string).
		UseProtoNames:   true, // Use the field names defined in the proto file as the output field names.
		EmitUnpopulated: true, // Emit fields even if they are unset or empty.
	}
}

// Server 服务
type Server struct {
	RpcSrv  *grpc.Server
	HttpSrv *http.Server
}

// GetServers 注册服务
func (s *Server) GetServers() []transport.Server {
	return []transport.Server{
		s.RpcSrv,
		s.HttpSrv,
	}
}
