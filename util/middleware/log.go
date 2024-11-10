package middleware

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
)

func UnaryLoggingInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	start := time.Now()
	resp, err := handler(ctx, req)
	log.Printf("[Success] = %s - [%s]",
		info.FullMethod, start.Format(time.RFC1123))
	return resp, err
}
