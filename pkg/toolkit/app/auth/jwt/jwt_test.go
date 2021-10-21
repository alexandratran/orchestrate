package jwt_test

import (
	"context"
	"crypto/x509"
	"fmt"
	"testing"
	"time"

	"github.com/consensys/orchestrate/pkg/toolkit/app/auth/jwt"
	jwtgenerator "github.com/consensys/orchestrate/pkg/toolkit/app/auth/jwt/generator"
	authutils "github.com/consensys/orchestrate/pkg/toolkit/app/auth/utils"
	"github.com/consensys/orchestrate/pkg/toolkit/app/multitenancy"
	"github.com/consensys/orchestrate/pkg/toolkit/tls/certificate"
	tlstestutils "github.com/consensys/orchestrate/pkg/toolkit/tls/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestJWT(t *testing.T) {
	tests := []struct { //nolint:maligned // reason
		name                       string
		cfg                        *jwt.Config
		genCfg                     *jwtgenerator.Config
		jwtTenantID                string
		jwtTTL                     time.Duration
		tenantID                   string
		expectedValidAuth          bool
		expectedImpersonatedTenant string
		expectedAllowedTenants     []string
	}{
		{
			"invalid signature",
			&jwt.Config{
				OrchestrateClaimPath: "orchestrate.test",
				Certificates:         certificates([]byte(tlstestutils.OneLineRSACertPEMB)),
			},
			&jwtgenerator.Config{
				KeyPair: &certificate.KeyPair{
					Cert: []byte(tlstestutils.OneLineRSACertPEMA),
					Key:  []byte(tlstestutils.OneLineRSAKeyPEMA),
				},
				OrchestrateClaimPath: "orchestrate.test",
			},
			"foo",
			10 * time.Hour,
			"foo",
			false,
			"",
			nil,
		},
		{
			"expired token",
			&jwt.Config{
				OrchestrateClaimPath: "orchestrate.test",
				Certificates:         certificates([]byte(tlstestutils.OneLineRSACertPEMA)),
			},
			&jwtgenerator.Config{
				KeyPair: &certificate.KeyPair{
					Cert: []byte(tlstestutils.OneLineRSACertPEMA),
					Key:  []byte(tlstestutils.OneLineRSAKeyPEMA),
				},
				OrchestrateClaimPath: "orchestrate.test",
			},
			"foo",
			-10 * time.Second,
			"foo",
			false,
			"",
			nil,
		},
		{
			"distinct jwt custom claims namespace",
			&jwt.Config{
				OrchestrateClaimPath: "orchestrate.prod",
				Certificates:         certificates([]byte(tlstestutils.OneLineRSACertPEMA)),
			},
			&jwtgenerator.Config{
				KeyPair: &certificate.KeyPair{
					Cert: []byte(tlstestutils.OneLineRSACertPEMA),
					Key:  []byte(tlstestutils.OneLineRSAKeyPEMA),
				},
				OrchestrateClaimPath: "orchestrate.dev",
			},
			"foo",
			10 * time.Hour,
			"foo",
			false,
			"",
			nil,
		},
		{
			"empty tenant id",
			&jwt.Config{
				OrchestrateClaimPath: "orchestrate.test",
				Certificates:         certificates([]byte(tlstestutils.OneLineRSACertPEMA)),
			},
			&jwtgenerator.Config{
				KeyPair: &certificate.KeyPair{
					Cert: []byte(tlstestutils.OneLineRSACertPEMA),
					Key:  []byte(tlstestutils.OneLineRSAKeyPEMA),
				},
				OrchestrateClaimPath: "orchestrate.test",
			},
			"",
			10 * time.Hour,
			"",
			false,
			"",
			nil,
		},
		{
			"JWT foo accessing empty tenant",
			&jwt.Config{
				OrchestrateClaimPath: "orchestrate.test",
				Certificates:         certificates([]byte(tlstestutils.OneLineRSACertPEMA)),
			},
			&jwtgenerator.Config{
				KeyPair: &certificate.KeyPair{
					Cert: []byte(tlstestutils.OneLineRSACertPEMA),
					Key:  []byte(tlstestutils.OneLineRSAKeyPEMA),
				},
				OrchestrateClaimPath: "orchestrate.test",
			},
			"foo",
			10 * time.Hour,
			"",
			true,
			"foo",
			[]string{"foo", multitenancy.DefaultTenant},
		},
		{
			"JWT foo accessing foo tenant",
			&jwt.Config{
				OrchestrateClaimPath: "orchestrate.test",
				Certificates:         certificates([]byte(tlstestutils.OneLineRSACertPEMA)),
			},
			&jwtgenerator.Config{
				KeyPair: &certificate.KeyPair{
					Cert: []byte(tlstestutils.OneLineRSACertPEMA),
					Key:  []byte(tlstestutils.OneLineRSAKeyPEMA),
				},
				OrchestrateClaimPath: "orchestrate.test",
			},
			"foo",
			10 * time.Hour,
			"foo",
			true,
			"foo",
			[]string{"foo"},
		},
		{
			"JWT foo accessing bar tenant",
			&jwt.Config{
				OrchestrateClaimPath: "orchestrate.test",
				Certificates:         certificates([]byte(tlstestutils.OneLineRSACertPEMA)),
			},
			&jwtgenerator.Config{
				KeyPair: &certificate.KeyPair{
					Cert: []byte(tlstestutils.OneLineRSACertPEMA),
					Key:  []byte(tlstestutils.OneLineRSAKeyPEMA),
				},
				OrchestrateClaimPath: "orchestrate.test",
			},
			"foo",
			10 * time.Hour,
			"bar",
			false,
			"",
			nil,
		},
		{
			"JWT * accessing empty tenant",
			&jwt.Config{
				OrchestrateClaimPath: "orchestrate.test",
				Certificates:         certificates([]byte(tlstestutils.OneLineRSACertPEMA)),
			},
			&jwtgenerator.Config{
				KeyPair: &certificate.KeyPair{
					Cert: []byte(tlstestutils.OneLineRSACertPEMA),
					Key:  []byte(tlstestutils.OneLineRSAKeyPEMA),
				},
				OrchestrateClaimPath: "orchestrate.test",
			},
			"*",
			10 * time.Hour,
			"",
			true,
			"_",
			[]string{multitenancy.Wildcard},
		},
		{
			"JWT * accessing foo tenant",
			&jwt.Config{
				OrchestrateClaimPath: "orchestrate.test",
				Certificates:         certificates([]byte(tlstestutils.OneLineRSACertPEMA)),
			},
			&jwtgenerator.Config{
				KeyPair: &certificate.KeyPair{
					Cert: []byte(tlstestutils.OneLineRSACertPEMA),
					Key:  []byte(tlstestutils.OneLineRSAKeyPEMA),
				},
				OrchestrateClaimPath: "orchestrate.test",
			},
			"*",
			10 * time.Hour,
			"foo",
			true,
			"foo",
			[]string{"foo", multitenancy.DefaultTenant},
		},
		{
			"JWT * accessing default tenant",
			&jwt.Config{
				OrchestrateClaimPath: "orchestrate.test",
				Certificates:         certificates([]byte(tlstestutils.OneLineRSACertPEMA)),
			},
			&jwtgenerator.Config{
				KeyPair: &certificate.KeyPair{
					Cert: []byte(tlstestutils.OneLineRSACertPEMA),
					Key:  []byte(tlstestutils.OneLineRSAKeyPEMA),
				},
				OrchestrateClaimPath: "orchestrate.test",
			},
			"*",
			10 * time.Hour,
			multitenancy.DefaultTenant,
			true,
			multitenancy.DefaultTenant,
			[]string{multitenancy.DefaultTenant},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			checker, err := jwt.New(tt.cfg)
			require.NoError(t, err)

			gen, err := jwtgenerator.New(tt.genCfg)
			require.NoError(t, err)

			token, err := gen.GenerateAccessTokenWithTenantID(tt.jwtTenantID, []string{"*:*"}, tt.jwtTTL)
			require.NoError(t, err)

			ctx := authutils.WithAuthorization(context.Background(), fmt.Sprintf("Bearer %v", token))
			ctx = multitenancy.WithTenantID(ctx, tt.tenantID)

			checkedCtx, err := checker.Check(ctx)
			if tt.expectedValidAuth {
				require.NoError(t, err, "Authentication check should succeeds")
				assert.Equal(t, tt.expectedImpersonatedTenant, multitenancy.TenantIDFromContext(checkedCtx))
				assert.Equal(t, tt.expectedAllowedTenants, multitenancy.AllowedTenantsFromContext(checkedCtx))
			} else {
				require.Error(t, err, "Authentication should fail")
			}
		})
	}
}

func certificates(content []byte) []*x509.Certificate {
	bCert, _ := certificate.Decode(content, "CERTIFICATE")
	cert, _ := x509.ParseCertificate(bCert[0])
	return []*x509.Certificate{cert}
}