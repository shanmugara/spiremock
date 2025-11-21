package spiremock

import (
	"context"
	"time"

	spireauth "github.com/shanmugara/spireauthlib"
	"github.com/sirupsen/logrus"
)

var Logger = logrus.New()

func NewMockClient() error {
	ctx := context.Background()
	cauth := &spireauth.ClientAuth{Logger: logrus.New()}
	tlsclient, err := cauth.GetTlsClient(ctx)
	if err != nil {
		Logger.Error("error getting tls client", err)
		return err
	}
	Logger.Infof("tls client created: %v", tlsclient)

	jwt, err := cauth.GetJWT(ctx)
	if err != nil {
		Logger.Error("error getting jwt svid", err)
		return err
	}
	Logger.Infof("jwt svid created: %s", jwt)

	time.Sleep(2 * time.Second)
	return nil

}
