package playserver_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	playserver "track-server-api/internal/server"
)

var service *playserver.Server

func TestPlay(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Invite server Suite")
}

var _ = BeforeSuite(func() {
	testing := true
	service = playserver.NewServer(db)
})

var _ = AfterSuite(func() {
})
