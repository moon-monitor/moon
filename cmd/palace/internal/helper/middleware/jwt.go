package middleware

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	jwtv5 "github.com/golang-jwt/jwt/v5"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do/system"
	"github.com/moon-monitor/moon/pkg/config"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
	"github.com/moon-monitor/moon/pkg/merr"
)

// JwtBaseInfo jwt base info
type JwtBaseInfo struct {
	UserID   uint32      `json:"user_id"`
	Username string      `json:"username"`
	Nickname string      `json:"nickname"`
	Avatar   string      `json:"avatar"`
	Gender   vobj.Gender `json:"gender"`
}

// JwtClaims jwt claims
type JwtClaims struct {
	signKey string
	*JwtBaseInfo
	*jwtv5.RegisteredClaims
}

// ParseJwtClaims parse jwt claims
func ParseJwtClaims(ctx context.Context) (*JwtClaims, bool) {
	claims, ok := jwt.FromContext(ctx)
	if !ok {
		return nil, false
	}
	jwtClaims, ok := claims.(*JwtClaims)
	if !ok {
		return nil, false
	}
	return jwtClaims, true
}

// JwtServer jwt server
func JwtServer(signKey string) middleware.Middleware {
	return jwt.Server(
		func(token *jwtv5.Token) (interface{}, error) {
			return []byte(signKey), nil
		},
		jwt.WithSigningMethod(jwtv5.SigningMethodHS256),
		jwt.WithClaims(func() jwtv5.Claims {
			return &JwtClaims{}
		}),
	)
}

// MustLogin must login
func MustLogin() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req any) (any, error) {
			claims, ok := ParseJwtClaims(ctx)
			if !ok {
				return nil, merr.ErrorUnauthorized("token error")
			}
			ctx = WithUserIDContext(ctx, claims.UserID)
			return handler(ctx, req)
		}
	}
}

// NewJwtClaims new jwt claims
func NewJwtClaims(c *config.JWT, userDo *system.User) *JwtClaims {
	base := &JwtBaseInfo{
		UserID:   userDo.ID,
		Username: userDo.Username,
		Nickname: userDo.Nickname,
		Avatar:   userDo.Avatar,
		Gender:   userDo.Gender,
	}
	return &JwtClaims{
		signKey:     c.GetSignKey(),
		JwtBaseInfo: base,
		RegisteredClaims: &jwtv5.RegisteredClaims{
			ExpiresAt: jwtv5.NewNumericDate(time.Now().Add(c.GetExpire().AsDuration())),
			Issuer:    c.GetIssuer(),
		},
	}
}

// GetToken get token
func (l *JwtClaims) GetToken() (string, error) {
	return jwtv5.NewWithClaims(jwtv5.SigningMethodHS256, l).SignedString([]byte(l.signKey))
}
