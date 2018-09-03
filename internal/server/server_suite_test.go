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

	// How will we do multi-API setup? ??

	//trackClient := track_pb.NewTrackServiceProtobufClient("http://localhost:8080", &http.Client{})
	//userClient := user_pb.NewUserServiceProtobufClient("http://localhost:8080", &http.Client{})

	// newTrack := &track_pb.Track{
	// 	Title:       "track title",
	// 	Status:      "free",
	// 	CreatorId:   "b86517a0-afba-41ca-820b-a4f6599d0edb",
	// 	UserGroupId: uuid.NewV4().String(),
	// 	TrackNumber: 1,
	// }
	// newTrack, err := trackClient.CreateTrack(context.Background(), newTrack)
	// Expect(err).NotTo(HaveOccurred())

	// newTrackId, err := internal.GetUuidFromString(newTrack.Id)

	newTrackData = &models.TrackData{
		TrackId:   uuid.NewV4(),
		UserId:    uuid.NewV4(),
		StorageId: "4_z134ab1f7e45796cc6950011e_f117076c66da42a22_d20180903_m010708_c002_v0001108_t0017",
	}
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
