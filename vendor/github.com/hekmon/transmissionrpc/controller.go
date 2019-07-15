package transmissionrpc

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"sync"
	"time"

	cleanhttp "github.com/hashicorp/go-cleanhttp"
)

const (
	// RPCVersion indicates the exact transmission RPC version this library is build against
	RPCVersion       = 15
	defaultPort      = 9091
	defaultRPCPath   = "/transmission/rpc"
	defaultTimeout   = 30 * time.Second
	defaultUserAgent = "github.com/hekmon/transmissionrpc"
)

// Client is the base object to interract with a remote transmission rpc endpoint.
// It must be created with New().
type Client struct {
	url             string
	user            string
	password        string
	sessionID       string
	sessionIDAccess sync.RWMutex
	userAgent       string
	rnd             *rand.Rand
	httpC           *http.Client
}

// AdvancedConfig handles options that are not mandatory for New().
// Default value for HTTPS is false, default port is 9091, default RPC URI is
// '/transmission/rpc', default HTTPTimeout is 30s.
type AdvancedConfig struct {
	HTTPS       bool
	Port        uint16
	RPCURI      string
	HTTPTimeout time.Duration
	UserAgent   string
}

// New returns an initialized and ready to use Controller
func New(host, user, password string, conf *AdvancedConfig) (c *Client, err error) {
	// Config
	if conf != nil {
		// Check custom config
		if conf.Port == 0 {
			conf.Port = defaultPort
		}
		if conf.RPCURI == "" {
			conf.RPCURI = defaultRPCPath
		}
		if conf.HTTPTimeout == 0 {
			conf.HTTPTimeout = defaultTimeout
		}
		if conf.UserAgent == "" {
			conf.UserAgent = defaultUserAgent
		}
	} else {
		// Spawn default config
		conf = &AdvancedConfig{
			// HTTPS false by default
			Port:        defaultPort,
			RPCURI:      defaultRPCPath,
			HTTPTimeout: defaultTimeout,
			UserAgent:   defaultUserAgent,
		}
	}
	// Build & validate URL
	var scheme string
	if conf.HTTPS {
		scheme = "https"
	} else {
		scheme = "http"
	}
	remoteURL, err := url.Parse(fmt.Sprintf("%s://%s:%d%s", scheme, host, conf.Port, conf.RPCURI))
	if err != nil {
		err = fmt.Errorf("can't build a valid URL: %v", err)
		return
	}
	// Initialize & return ready to use client
	c = &Client{
		url:       remoteURL.String(),
		user:      user,
		password:  password,
		userAgent: conf.UserAgent,
		rnd:       rand.New(rand.NewSource(time.Now().Unix())),
		httpC:     cleanhttp.DefaultPooledClient(),
	}
	c.httpC.Timeout = conf.HTTPTimeout
	return
}
