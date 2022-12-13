package test_version1

import (
	"context"
	"os"
	"testing"

	"github.com/pip-services-infrastructure2/client-email-go/version1"
	"github.com/pip-services3-gox/pip-services3-commons-gox/config"
)

type emailGrpcClientV1Test struct {
	client  *version1.EmailGrpcClientV1
	fixture *EmailClientFixtureV1
}

func newEmailGrpcClientV1Test() *emailGrpcClientV1Test {
	return &emailGrpcClientV1Test{}
}

func (c *emailGrpcClientV1Test) setup(t *testing.T) {
	var GRPC_HOST = os.Getenv("GRPC_HOST")
	if GRPC_HOST == "" {
		GRPC_HOST = "localhost"
	}
	var GRPC_PORT = os.Getenv("GRPC_PORT")
	if GRPC_PORT == "" {
		GRPC_PORT = "8090"
	}

	var httpConfig = config.NewConfigParamsFromTuples(
		"connection.protocol", "http",
		"connection.host", GRPC_HOST,
		"connection.port", GRPC_PORT,
	)

	c.client = version1.NewEmailGrpcClientV1()
	c.client.Configure(context.Background(), httpConfig)
	c.client.Open(context.Background(), "")

	c.fixture = NewEmailClientFixtureV1(c.client)
}

func (c *emailGrpcClientV1Test) teardown(t *testing.T) {
	c.client.Close(context.Background(), "")
}

func TestGrpcClientSendEmail(t *testing.T) {
	c := newEmailGrpcClientV1Test()
	c.setup(t)
	defer c.teardown(t)

	c.fixture.TestSendEmailToAddress(t)
	c.fixture.TestSendEmailToRecipient(t)
}
