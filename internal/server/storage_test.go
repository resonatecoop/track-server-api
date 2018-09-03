package trackdataserver_test

import (
	"context"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	uuid "github.com/satori/go.uuid"
	"github.com/twitchtv/twirp"

	trackdataserver "track-server-api/internal/server"
	pb "track-server-api/rpc"
)

var _ = Describe("Track data server", func() {
	const invalid_argument_code twirp.ErrorCode = "invalid_argument"
	const not_found_code twirp.ErrorCode = "not_found"
	const internal_code twirp.ErrorCode = "internal"

	Describe("openStorageConnection", func() {
		Context("with no arguments", func() {
			It("should return a StorageConnection", func() {
				_, err := trackdataserver.OpenStorageConnection()
				Expect(err).NotTo(HaveOccurred())
			})
		})
	})
	Describe("getTrackChunkFromStorage", func() {
		Context("with valid arguments", func() {
			It("should return a filled TrackChunk", func() {

				sc, err := trackdataserver.OpenStorageConnection()
				Expect(err).NotTo(HaveOccurred())

				const bytesPerRead = 12000 // temporary fake value
				trackServerId := uuid.NewV4()
				trackChunk := &pb.TrackChunk{
					StartPosition: 0,
					NumBytes:      bytesPerRead,
				}

				trackChunk, err = trackdataserver.GetTrackChunkFromStorage(trackServerId.String(), trackChunk, sc)
				Expect(err).NotTo(HaveOccurred())
			})
		})
		Context("with invalid track uuid", func() {
			It("should respond with invalid_argument error", func() {
				userTrackPB := &pb.UserTrack{TrackId: "0", UserId: "0"}
				_, err := service.StreamTrackData(context.Background(), userTrackPB)
				Expect(err).To(HaveOccurred())
			})
		})
		Context("with invalid user uuid", func() {
			It("should respond with invalid_argument error", func() {
			})
		})
	})
})
