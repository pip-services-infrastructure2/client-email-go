package test_version1

import (
	"context"
	"os"
	"testing"

	"github.com/pip-services-infrastructure2/client-email-go/version1"
	"github.com/pip-services3-gox/pip-services3-commons-gox/config"
)

type emailCommandableGrpcClientV1Test struct {
	client  *version1.EmailCommandableGrpcClientV1
	fixture *EmailClientFixtureV1
}

func newEmailCommandableGrpcClientV1Test() *emailCommandableGrpcClientV1Test {
	return &emailCommandableGrpcClientV1Test{}
}

func (c *emailCommandableGrpcClientV1Test) setup(t *testing.T) {
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

	c.client = version1.NewEmailCommandableGrpcClientV1()
	c.client.Configure(context.Background(), httpConfig)
	c.client.Open(context.Background(), "")

	c.fixture = NewEmailClientFixtureV1(c.client)
}

func (c *emailCommandableGrpcClientV1Test) teardown(t *testing.T) {
	c.client.Close(context.Background(), "")
}

func TestCommandableGrpcClientSendEmail(t *testing.T) {
	c := newEmailCommandableGrpcClientV1Test()
	c.setup(t)
	defer c.teardown(t)

	c.fixture.TestSendEmailToAddress(t)
	c.fixture.TestSendEmailToRecipient(t)
}
