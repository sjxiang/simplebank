package gapi

import (
	"context"

	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
)

const (
	grpcGatewayUserAgentHeader = "grpcgateway-user-agent"
	userAgentHeader            = "user-agent"
	xForwardedForHeader        = "x-forwarded-for"
)

type Metadata struct {
	UserAgent string
	ClientIP  string
}

func (server *Server) extractMetadata(ctx context.Context) *Metadata {
	// 没有就是 nil 零值 ""
	mtdt := &Metadata{}

	// 元信息
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		// 适用于 grpc gateway
		if userAgents := md.Get(grpcGatewayUserAgentHeader); len(userAgents) > 0 {
			mtdt.UserAgent = userAgents[0]
		}

		// 适用于 grpc clinet（evans）
		if userAgents := md.Get(userAgentHeader); len(userAgents) > 0 {
			mtdt.UserAgent = userAgents[0]
		}

		// 适用于 grpc gateway
		if clientIPs := md.Get(xForwardedForHeader); len(clientIPs) > 0 {
			mtdt.ClientIP = clientIPs[0]
		}
	}

	// 适用于 grpc clinet（evans）
	if p, ok := peer.FromContext(ctx); ok {
		mtdt.ClientIP = p.Addr.String()
	}

	return mtdt
}
