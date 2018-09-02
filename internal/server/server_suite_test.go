package trackdataserver_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	uuid "github.com/satori/go.uuid"

	"github.com/go-pg/pg"

	"track-server-api/internal/database"
	"track-server-api/internal/database/models"

	trackdataserver "track-server-api/internal/server"
)

var (
	db           *pg.DB
	service      *trackdataserver.Server
	newTrackData *models.TrackData
)

func TestTrackData(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Track data server Suite")
}

var _ = BeforeSuite(func() {
	testing := true
	db = database.Connect(testing)
	service = trackdataserver.NewServer(db)

	newTrackData = &models.TrackData{TrackId: uuid.NewV4(), UserId: uuid.NewV4()}
	err := db.Insert(newTrackData)
	Expect(err).NotTo(HaveOccurred())
})

var _ = AfterSuite(func() {
	var tracks []models.TrackData
	err := db.Model(&tracks).Select()
	Expect(err).NotTo(HaveOccurred())
	_, err = db.Model(&tracks).Delete()
	Expect(err).NotTo(HaveOccurred())
})
