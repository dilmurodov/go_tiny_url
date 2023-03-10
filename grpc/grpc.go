package grpc

import (
	"go_auth_api_gateway/config"
	"go_auth_api_gateway/genproto/auth_service"
	"go_auth_api_gateway/grpc/client"
	"go_auth_api_gateway/grpc/service"
	"go_auth_api_gateway/storage"

	"github.com/saidamir98/udevs_pkg/logger"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func SetUpServer(cfg config.Config, log logger.LoggerI, strg storage.StorageI, svcs client.ServiceManagerI) (grpcServer *grpc.Server) {
	grpcServer = grpc.NewServer()

	// auth_service.RegisterClientServiceServer(grpcServer, service.NewClientService(cfg, log, strg, svcs))
	// auth_service.RegisterPermissionServiceServer(grpcServer, service.NewPermissionService(cfg, log, strg, svcs))
	auth_service.RegisterUserServiceServer(grpcServer, service.NewUserService(cfg, log, strg, svcs))
	// auth_service.RegisterSessionServiceServer(grpcServer, service.NewSessionService(cfg, log, strg, svcs))

	auth_service.RegisterShortenerServiceServer(grpcServer, service.NewShortenerService(cfg, log, strg, svcs))

	reflection.Register(grpcServer)
	return
}
