package spiremock

import (
	"context"
	"net/http"

	spireauth "github.com/shanmugara/spireauthlib"
	"github.com/sirupsen/logrus"
	"github.com/spiffe/go-spiffe/v2/bundle/jwtbundle"
	"github.com/spiffe/go-spiffe/v2/svid/jwtsvid"
	"github.com/spiffe/spire-api-sdk/proto/spire/api/types"
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

func NewJWTMockClient(audience string) (*jwtbundle.Set, *jwtsvid.SVID, error) {
	ctx := context.Background()
	cauth := &spireauth.ClientAuth{Logger: logrus.New()}

	jbundle, jwt, err := cauth.GetJWT(ctx, audience)
	if err != nil {
		Logger.Error("error getting jwt svid", err)
		return nil, nil, err
	}
	Logger.Infof("jwt svid created: %+v", jwt)
	Logger.Infof("jwt svid verified: %+v", jbundle.Bundles())

	return jbundle, jwt, nil
}

func NewDLGJWTMockClient(selectors []*types.Selector, audience string) error {
	ctx := context.Background()
	cauth := &spireauth.DelegatedAuth{Logger: logrus.New()}
	Logger.Infof("dlg client created...")
	_, err := cauth.GetDelegatedJWT(ctx, selectors, audience)
	if err != nil {
		return err
	}
	return nil
}
