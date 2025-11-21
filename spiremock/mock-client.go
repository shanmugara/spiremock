package spiremock

import (
	"context"
	"net/http"
	"time"

	spireauth "github.com/shanmugara/spireauthlib"
	"github.com/sirupsen/logrus"
	"github.com/spiffe/go-spiffe/v2/bundle/jwtbundle"
	"github.com/spiffe/go-spiffe/v2/svid/jwtsvid"
)

var Logger = logrus.New()

func NewTlsMockClient() (*http.Client, error) {
	ctx := context.Background()
	cauth := &spireauth.ClientAuth{Logger: logrus.New()}
	tlsclient, err := cauth.GetTlsClient(ctx)
	if err != nil {
		Logger.Error("error getting tls client", err)
		return nil, err
	}
	Logger.Infof("tls client created: %v", tlsclient)

	return tlsclient, nil

}

func NewJWTMockClient() (*jwtbundle.Set, *jwtsvid.SVID, error) {
	ctx := context.Background()
	cauth := &spireauth.ClientAuth{Logger: logrus.New()}

	jbundle, jwt, err := cauth.GetJWT(ctx)
	if err != nil {
		Logger.Error("error getting jwt svid", err)
		return nil, nil, err
	}
	Logger.Infof("jwt svid created: %+v", jwt)
	Logger.Infof("jwt svid verified: %+v", jbundle.Bundles())

	time.Sleep(5 * time.Second)
	return jbundle, jwt, nil
}
