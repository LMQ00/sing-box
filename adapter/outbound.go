package adapter

import (
	"context"
	"net/netip"
	"time"

	"github.com/sagernet/sing-box/log"
	"github.com/sagernet/sing-box/option"
	"github.com/sagernet/sing-tun"
	N "github.com/sagernet/sing/common/network"
)

// Note: for proxy protocols, outbound creates early connections by default.

type Outbound interface {
	Type() string
	Tag() string
	Network() []string
	Dependencies() []string
	N.Dialer
}

type OutboundWithPreferredRoutes interface {
	Outbound
	PreferredDomain(metadata *InboundContext, domain string) bool
	PreferredAddress(metadata *InboundContext, address netip.Addr) bool
}

type FlowOutbound interface {
	Outbound
	tun.Port
	PreMatchFlow(network string, destination netip.Addr) PreMatchAction
}

type OutboundRegistry interface {
	option.OutboundOptionsRegistry
	CreateOutbound(ctx context.Context, router Router, logger log.ContextLogger, tag string, outboundType string, options any) (Outbound, error)
}

type OutboundManager interface {
	Lifecycle
	Outbounds() []Outbound
	Outbound(tag string) (Outbound, bool)
	Default() Outbound
	Remove(tag string) error
	Create(ctx context.Context, router Router, logger log.ContextLogger, tag string, outboundType string, options any) error
}

// DirectRouteContext is the context for direct route connections.
// Defined here as the reF1nd sing-tun fork provided this type;
// when using upstream sing-tun, this is an opaque placeholder.
type DirectRouteContext = any

// DirectRouteDestination is the result of creating a direct route connection.
type DirectRouteDestination = any

// DirectRouteOutbound is an outbound that supports direct route connections.
type DirectRouteOutbound interface {
	Outbound
	NewDirectRouteConnection(metadata InboundContext, routeContext any, timeout time.Duration) (any, error)
}
