package serve

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"strconv"

	"google.golang.org/grpc"

	"github.com/ZergsLaw/back-template/internal/logger"
)

// GRPC starts gRPC server on addr, logged as service.
// It runs until failed or ctx.Done.
func GRPC(log *slog.Logger, host string, port uint16, srv *grpc.Server) func(context.Context) error {
	return func(ctx context.Context) error {
		ln, err := net.Listen("tcp", net.JoinHostPort(host, strconv.FormatUint(uint64(port), 10)))
		if err != nil {
			return fmt.Errorf("net.Listen: %w", err)
		}

		errc := make(chan error, 1)
		go func() { errc <- srv.Serve(ln) }()
		log.Info("started", slog.String(logger.Host.String(), host), slog.Uint64(logger.Port.String(), uint64(port)))

		defer log.Info("shutdown")

		select {
		case err = <-errc:
		case <-ctx.Done():
			srv.GracefulStop() // It will not interrupt streaming.
		}

		if err != nil {
			return fmt.Errorf("srv.Serve: %w", err)
		}

		return nil
	}
}
