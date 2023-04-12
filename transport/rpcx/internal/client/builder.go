package client

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/bitini111/mvta/errors"
	"github.com/bitini111/mvta/registry"
	"github.com/bitini111/mvta/transport/rpcx/internal/resolver"
	"github.com/bitini111/mvta/transport/rpcx/internal/resolver/direct"
	"github.com/bitini111/mvta/transport/rpcx/internal/resolver/discovery"
	cli "github.com/smallnest/rpcx/client"
	proto "github.com/smallnest/rpcx/protocol"
	"net/url"
	"os"
	"sync"
)

const defaultBuilder = "direct"

type Builder struct {
	err      error
	opts     cli.Option
	pools    sync.Map
	builders map[string]resolver.Builder
}

type Options struct {
	CertFile   string
	ServerName string
	Discovery  registry.Discovery
	FailMode   cli.FailMode
}

func NewBuilder(opts *Options) *Builder {
	b := &Builder{}
	b.builders = make(map[string]resolver.Builder)
	b.opts = cli.DefaultOption
	b.opts.CompressType = proto.Gzip
	b.RegisterBuilder(direct.NewBuilder())
	if opts.Discovery != nil {
		b.RegisterBuilder(discovery.NewBuilder(opts.Discovery))
	}

	if opts.CertFile != "" && opts.ServerName != "" {
		b.opts.TLSConfig, b.err = newClientTLSFromFile(opts.CertFile, opts.ServerName)
	}

	return b
}

func newClientTLSFromFile(certFile string, serverName string) (*tls.Config, error) {
	b, err := os.ReadFile(certFile)
	if err != nil {
		return nil, err
	}

	cp := x509.NewCertPool()
	if !cp.AppendCertsFromPEM(b) {
		return nil, fmt.Errorf("credentials: failed to append certificates")
	}

	return &tls.Config{ServerName: serverName, RootCAs: cp}, nil
}

// RegisterBuilder 注册构建器
func (b *Builder) RegisterBuilder(builder resolver.Builder) {
	b.builders[builder.Scheme()] = builder
}

// Build 建立Discovery
func (b *Builder) Build(target string) (*cli.OneClient, error) {
	if b.err != nil {
		return nil, b.err
	}

	u, err := url.Parse(target)
	if err != nil {
		return nil, err
	}

	if u.Scheme == "" {
		u.Scheme = defaultBuilder
		target = u.String()
	}

	val, ok := b.pools.Load(target)
	if ok {
		return val.(*cli.OneClientPool).Get(), nil
	}

	var builder resolver.Builder
	if u.Scheme == "" {
		builder, ok = b.builders[defaultBuilder]
	} else {
		builder, ok = b.builders[u.Scheme]
	}
	if !ok {
		return nil, errors.New("missing resolver")
	}

	dis, err := builder.Build(u)
	if err != nil {
		return nil, err
	}

	pool := cli.NewOneClientPool(10, cli.Failtry, cli.RandomSelect, dis, b.opts)
	b.pools.Store(target, pool)

	return pool.Get(), nil
}
