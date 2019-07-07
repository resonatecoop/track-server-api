package trackdataserver_test

import (
	"testing"
	"path/filepath"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	uuid "github.com/satori/go.uuid"
	"github.com/go-pg/pg"

	"github.com/blushi/user-api/pkg/config"
	"github.com/blushi/user-api/pkg/postgres"

	"track-server-api/internal/model"
	"track-server-api/internal/pkg/storage"
	trackdataserver "track-server-api/internal/server"
)

var (
	db            *pg.DB
	service       *trackdataserver.Server
	newTrackData  *model.TrackData
	longTrackData *model.TrackData
)

func TestTrackData(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Track data server Suite")
}

var _ = BeforeSuite(func() {
	var err error

	cfgPath, err := filepath.Abs("./../../conf.local.yaml")
	Expect(err).NotTo(HaveOccurred())

	cfg, err := config.Load(cfgPath)
	Expect(err).NotTo(HaveOccurred())

	db, err = pgsql.New(cfg.DB.Test.PSN, cfg.DB.Test.LogQueries, cfg.DB.Test.TimeoutSeconds)
	Expect(err).NotTo(HaveOccurred())

	sc, err := storage.New(
		cfg.Storage.AccountId,
		cfg.Storage.Key,
		cfg.Storage.AuthEndpoint,
		cfg.Storage.FileEndpoint,
		cfg.Storage.UploadEndpoint,
		cfg.Storage.BucketId,
		cfg.Storage.Timeout)

	service = trackdataserver.NewServer(db, sc)

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

	// newTrackId, err := uuidpkg.GetUuidFromString(newTrack.Id)

	newTrackData = &model.TrackData{
		TrackId:   uuid.NewV4(),
		UserId:    uuid.NewV4(),
		StorageId: "4_z134ab1f7e45796cc6950011e_f11730b579d55bb63_d20180908_m164303_c002_v0001108_t0009",
	}
	err = db.Insert(newTrackData)
	Expect(err).NotTo(HaveOccurred())
	longTrackData = &model.TrackData{
		TrackId:   uuid.NewV4(),
		UserId:    uuid.NewV4(),
		StorageId: "4_z134ab1f7e45796cc6950011e_f103176d1223dadcf_d20180910_m014723_c002_v0001108_t0010",
	}
	err = db.Insert(longTrackData)
	Expect(err).NotTo(HaveOccurred())
})

var _ = AfterSuite(func() {
	var tracks []model.TrackData
	err := db.Model(&tracks).Select()
	Expect(err).NotTo(HaveOccurred())
	_, err = db.Model(&tracks).Delete()
	Expect(err).NotTo(HaveOccurred())
})
