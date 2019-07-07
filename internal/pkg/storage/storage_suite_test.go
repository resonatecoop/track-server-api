package storage_test

import (
	"testing"
	"path/filepath"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
  "github.com/blushi/user-api/pkg/config"

  "track-server-api/internal/pkg/storage"
)

var sc *storage.StorageConnection

func TestStorage(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Storage Suite")
}

var _ = BeforeSuite(func() {
	var err error

	cfgPath, err := filepath.Abs("./../../../conf.local.yaml")
	Expect(err).NotTo(HaveOccurred())

	cfg, err := config.Load(cfgPath)
	Expect(err).NotTo(HaveOccurred())

	sc, err = storage.New(
		cfg.Storage.AccountId,
		cfg.Storage.Key,
		cfg.Storage.AuthEndpoint,
		cfg.Storage.FileEndpoint,
		cfg.Storage.UploadEndpoint,
		cfg.Storage.BucketId,
		cfg.Storage.Timeout)
  Expect(err).NotTo(HaveOccurred())
})

var _ = AfterSuite(func() {
})
