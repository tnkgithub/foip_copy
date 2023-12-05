package livekit

import (
	"fmt"
	"foip/core/config"
	"time"

	"gopkg.in/square/go-jose.v2"
	"gopkg.in/square/go-jose.v2/jwt"
)

type accessToken struct {
	apiKey   string
	secret   string
	grant    claimGrants
	validFor time.Duration
}

type claimGrants struct {
	Identity string      `json:"-"`
	Name     string      `json:"name,omitempty"`
	Video    *videoGrant `json:"video,omitempty"`

	// for verifying integrity of the message body
	Sha256   string `json:"sha256,omitempty"`
	Metadata string `json:"metadata,omitempty"`
}

type videoGrant struct {
	RoomJoin bool   `json:"roomJoin,omitempty"`
	Room     string `json:"room,omitempty"`
}

func GenerateAccessToken(roomID, userID string, cfg config.LivekitConfig) (string, error) {
	at := accessToken{
		apiKey:   cfg.AccessKey,
		secret:   cfg.SecretKey,
		validFor: cfg.TokenValidDuration,
	}

	//set grant field
	grant := videoGrant{RoomJoin: true, Room: roomID}
	at.grant.Identity = userID
	//NOTE: JWT Claim Name: user name
	//at.grant.Name
	at.grant.Video = &grant

	return at.toJWT()
}

func (t *accessToken) toJWT() (string, error) {
	if t.apiKey == "" || t.secret == "" {
		return "", fmt.Errorf("key value is empty")
	}

	sig, err := jose.NewSigner(jose.SigningKey{Algorithm: jose.HS256, Key: []byte(t.secret)},
		(&jose.SignerOptions{}).WithType("JWT"))
	if err != nil {
		return "", err
	}

	cl := jwt.Claims{
		Expiry:    jwt.NewNumericDate(time.Now().Add(t.validFor)),
		Issuer:    t.apiKey,
		ID:        t.grant.Identity,
		NotBefore: jwt.NewNumericDate(time.Now()),
		Subject:   t.grant.Identity,
	}
	return jwt.Signed(sig).Claims(cl).Claims(&t.grant).CompactSerialize()
}
