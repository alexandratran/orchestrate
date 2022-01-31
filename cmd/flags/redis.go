package flags

import (
	"fmt"
	"time"

	"github.com/consensys/orchestrate/src/infra/redis/redigo"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func init() {
	viper.SetDefault(HostViperKey, hostDefault)
	_ = viper.BindEnv(HostViperKey, hostEnv)
	viper.SetDefault(PortViperKey, portDefault)
	_ = viper.BindEnv(PortViperKey, portEnv)
	viper.SetDefault(UsernameViperKey, usernameDefault)
	_ = viper.BindEnv(UsernameViperKey, usernameEnv)
	viper.SetDefault(PasswordViperKey, passwordDefault)
	_ = viper.BindEnv(PasswordViperKey, passwordEnv)
	viper.SetDefault(DatabaseViperKey, databaseDefault)
	_ = viper.BindEnv(DatabaseViperKey, databaseEnv)
	viper.SetDefault(maxIdleViperKey, maxIdleDefault)
	_ = viper.BindEnv(maxIdleViperKey, maxidleEnv)
	viper.SetDefault(idleTimeoutViperKey, idleTimeoutDefault)
	_ = viper.BindEnv(idleTimeoutViperKey, idleTimeoutEnv)
	viper.SetDefault(TLSCertViperKey, tlsCertDefault)
	_ = viper.BindEnv(TLSCertViperKey, tlsCertEnv)
	viper.SetDefault(TLSKeyViperKey, tlsKeyDefault)
	_ = viper.BindEnv(TLSKeyViperKey, tlsKeyEnv)
	viper.SetDefault(TLSCAViperKey, tlsCADefault)
	_ = viper.BindEnv(TLSCAViperKey, tlsCAEnv)
	viper.SetDefault(TLSSkipVerifyViperKey, tlsSkipVerifyDefault)
	_ = viper.BindEnv(TLSSkipVerifyViperKey, tlsSkipVerifyEnv)
}

func RedisFlags(f *pflag.FlagSet) {
	host(f)
	port(f)
	username(f)
	database(f)
	password(f)
	maxIdle(f)
	idleTimeout(f)
	tlsCert(f)
	tlsKey(f)
	tlsCA(f)
	skipVerify(f)
}

const (
	hostFlag     = "redis-host"
	HostViperKey = "redis.host"
	hostDefault  = "localhost"
	hostEnv      = "REDIS_HOST"
)

func host(f *pflag.FlagSet) {
	desc := fmt.Sprintf(`Host (address) of Redis server to connect to.
Environment variable: %q`, hostEnv)
	f.String(hostFlag, hostDefault, desc)
	_ = viper.BindPFlag(HostViperKey, f.Lookup(hostFlag))
}

const (
	portFlag     = "redis-port"
	PortViperKey = "redis.port"
	portDefault  = "6379"
	portEnv      = "REDIS_PORT"
)

func port(f *pflag.FlagSet) {
	desc := fmt.Sprintf(`Port of Redis server to connect to.
Environment variable: %q`, portEnv)
	f.String(portFlag, portDefault, desc)
	_ = viper.BindPFlag(PortViperKey, f.Lookup(portFlag))
}

const (
	usernameFlag     = "redis-user"
	UsernameViperKey = "redis.user"
	usernameDefault  = ""
	usernameEnv      = "REDIS_USER"
)

func username(f *pflag.FlagSet) {
	desc := fmt.Sprintf(`Redis Username.
Environment variable: %q`, usernameEnv)
	f.String(usernameFlag, usernameDefault, desc)
	_ = viper.BindPFlag(UsernameViperKey, f.Lookup(usernameFlag))
}

const (
	passwordFlag     = "redis-password"
	PasswordViperKey = "redis.password"
	passwordDefault  = ""
	passwordEnv      = "REDIS_PASSWORD"
)

func password(f *pflag.FlagSet) {
	desc := fmt.Sprintf(`Redis Username password
Environment variable: %q`, passwordEnv)
	f.String(passwordFlag, passwordDefault, desc)
	_ = viper.BindPFlag(PasswordViperKey, f.Lookup(passwordFlag))
}

const (
	databaseFlag     = "redis-database"
	DatabaseViperKey = "redis.database"
	databaseDefault  = -1
	databaseEnv      = "REDIS_DATABASE"
)

func database(f *pflag.FlagSet) {
	desc := fmt.Sprintf(`Target Redis database name
Environment variable: %q`, databaseEnv)
	f.Int(databaseFlag, databaseDefault, desc)
	_ = viper.BindPFlag(DatabaseViperKey, f.Lookup(databaseFlag))
}

const (
	tlsCertFlag     = "redis-tls-cert"
	TLSCertViperKey = "redis.tls.cert"
	tlsCertDefault  = ""
	tlsCertEnv      = "REDIS_TLS_CERT"
)

func tlsCert(f *pflag.FlagSet) {
	desc := fmt.Sprintf(`TLS Certificate to connect to Redis
Environment variable: %q`, tlsCertEnv)
	f.String(tlsCertFlag, tlsCertDefault, desc)
	_ = viper.BindPFlag(TLSCertViperKey, f.Lookup(tlsCertFlag))
}

const (
	tlsKeyFlag     = "redis-tls-key"
	TLSKeyViperKey = "redis.tls.key"
	tlsKeyDefault  = ""
	tlsKeyEnv      = "REDIS_TLS_KEY"
)

func tlsKey(f *pflag.FlagSet) {
	desc := fmt.Sprintf(`TLS Private Key to connect to Redis
Environment variable: %q`, tlsKeyEnv)
	f.String(tlsKeyFlag, tlsKeyDefault, desc)
	_ = viper.BindPFlag(TLSKeyViperKey, f.Lookup(tlsKeyFlag))
}

const (
	tlsCAFlag     = "redis-tls-ca"
	TLSCAViperKey = "redis.tls.ca"
	tlsCADefault  = ""
	tlsCAEnv      = "REDIS_TLS_CA"
)

func tlsCA(f *pflag.FlagSet) {
	desc := fmt.Sprintf(`Trusted Certificate Authority
Environment variable: %q`, tlsCAEnv)
	f.String(tlsCAFlag, tlsCADefault, desc)
	_ = viper.BindPFlag(TLSCAViperKey, f.Lookup(tlsCAFlag))
}

const (
	tlsSkipVerifyFlag     = "redis-tls-skip-verify"
	TLSSkipVerifyViperKey = "redis.tls.skip-verify"
	tlsSkipVerifyDefault  = false
	tlsSkipVerifyEnv      = "REDIS_TLS_SKIP_VERIFY"
)

func skipVerify(f *pflag.FlagSet) {
	desc := fmt.Sprintf(`Skip service certificate verification
Environment variable: %q`, tlsSkipVerifyEnv)
	f.Bool(tlsSkipVerifyFlag, tlsSkipVerifyDefault, desc)
	_ = viper.BindPFlag(TLSSkipVerifyViperKey, f.Lookup(tlsSkipVerifyFlag))
}

const (
	maxIdleFlag     = "redis-max-idle"
	maxIdleViperKey = "redis.max.idle"
	maxIdleDefault  = 10000
	maxidleEnv      = "REDIS_MAX_IDLE"
)

func maxIdle(f *pflag.FlagSet) {
	desc := fmt.Sprintf(`Redis Max idle.
Environment variable: %q`, maxidleEnv)
	f.Int(maxIdleFlag, maxIdleDefault, desc)
	_ = viper.BindPFlag(maxIdleViperKey, f.Lookup(maxIdleFlag))
}

const (
	idleTimeoutFlag     = "redis-idle-timeout"
	idleTimeoutViperKey = "redis.idle-timeout"
	idleTimeoutDefault  = 240 * time.Second
	idleTimeoutEnv      = "REDIS_IDLE_TIMEOUT"
)

func idleTimeout(f *pflag.FlagSet) {
	desc := fmt.Sprintf(`Redis idle timeout duration.
Environment variable: %q`, idleTimeoutEnv)
	f.Duration(idleTimeoutFlag, idleTimeoutDefault, desc)
	_ = viper.BindPFlag(idleTimeoutViperKey, f.Lookup(idleTimeoutFlag))
}

func NewRedisConfig(vipr *viper.Viper) *redigo.Config {
	return &redigo.Config{
		Host:          vipr.GetString(HostViperKey),
		Port:          vipr.GetString(PortViperKey),
		User:          vipr.GetString(UsernameViperKey),
		Password:      vipr.GetString(PasswordViperKey),
		Database:      vipr.GetInt(DatabaseViperKey),
		MaxIdle:       vipr.GetInt(maxIdleViperKey),
		IdleTimeout:   vipr.GetDuration(idleTimeoutViperKey),
		TLSCert:       vipr.GetString(TLSCertViperKey),
		TLSKey:        vipr.GetString(TLSKeyViperKey),
		TLSCA:         vipr.GetString(TLSCAViperKey),
		TLSSkipVerify: vipr.GetBool(TLSSkipVerifyViperKey),
	}
}
