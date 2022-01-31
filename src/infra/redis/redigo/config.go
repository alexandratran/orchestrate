package redigo

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/consensys/orchestrate/pkg/toolkit/tls"
	"github.com/consensys/orchestrate/pkg/toolkit/tls/certificate"
	"github.com/gomodule/redigo/redis"
)

type Config struct {
	Host          string
	Port          string
	User          string
	Password      string
	Database      int
	MaxIdle       int
	IdleTimeout   time.Duration
	TLSCert       string
	TLSKey        string
	TLSCA         string
	TLSSkipVerify bool
}

func (cfg *Config) ToRedisOptions() ([]redis.DialOption, error) {
	var options []redis.DialOption
	if cfg.Database != -1 {
		options = append(options, redis.DialDatabase(cfg.Database))
	}

	if cfg.User != "" {
		options = append(options, redis.DialUsername(cfg.User))
	}

	if cfg.Password != "" {
		options = append(options, redis.DialPassword(cfg.Password))
	}

	if cfg.TLSCert != "" && cfg.TLSKey != "" {
		tlsOption, err := cfg.getTLSOption()
		if err != nil {
			return nil, err
		}

		c, err := tls.NewConfig(tlsOption)
		if err != nil {
			return nil, err
		}

		options = append(options, redis.DialTLSConfig(c), redis.DialUseTLS(true))
	}

	return options, nil
}

func (cfg *Config) getTLSOption() (*tls.Option, error) {
	tlsOption := &tls.Option{}

	cert, err := ioutil.ReadFile(cfg.TLSCert)
	if err != nil {
		return nil, err
	}

	key, err := ioutil.ReadFile(cfg.TLSKey)
	if err != nil {
		return nil, err
	}

	tlsOption.Certificates = []*certificate.KeyPair{{Cert: cert, Key: key}}

	if cfg.TLSCA != "" {
		ca, err := ioutil.ReadFile(cfg.TLSCA)
		if err != nil {
			return nil, err
		}

		tlsOption.CAs = [][]byte{ca}
	}

	if cfg.TLSSkipVerify {
		tlsOption.InsecureSkipVerify = true
	}

	return tlsOption, nil
}

func (cfg *Config) URL() string {
	return fmt.Sprintf("%v:%v", cfg.Host, cfg.Port)
}
