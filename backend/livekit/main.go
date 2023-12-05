package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/square/go-jose.v2"
	"gopkg.in/square/go-jose.v2/jwt"
	"gopkg.in/yaml.v2"
)

const (
	ACCESS_TOKEN_VALID_DURATION = 30 * 24 * time.Hour

	LIVEKIT_CONFIG_FILE = "livekit.yaml"
)

var (
	APP_KEY    string
	APP_SECRET string
)

type Config struct {
	Keys map[string]string `yaml:"keys"`
}

func init() {
	cfg := Config{}
	file, err := os.ReadFile(LIVEKIT_CONFIG_FILE)
	if err != nil {
		panic(err)
	}
	if err := yaml.Unmarshal(file, &cfg); err != nil {
		panic(err)
	}

	for key, secret := range cfg.Keys {
		APP_KEY = key
		APP_SECRET = secret
		break
	}
}

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

func GenerateAccessToken(roomID, userID, key, secret string) (string, error) {
	if roomID == "" || userID == "" {
		return "", fmt.Errorf("request field value is empty")
	}

	at := accessToken{
		apiKey:   key,
		secret:   secret,
		validFor: ACCESS_TOKEN_VALID_DURATION,
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

func main() {
	app := gin.Default()

	app.GET("/token", func(ctx *gin.Context) {
		room, user := ctx.Query("room"), ctx.Query("user")
		if room == "" || user == "" {
			return
		}

		token, err := GenerateAccessToken(room, user, APP_KEY, APP_SECRET)
		if err != nil {
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"token": token})
	})

	if err := app.Run(); err != nil {
		panic(err)
	}
}
