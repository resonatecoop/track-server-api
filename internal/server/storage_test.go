package trackdataserver_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
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

				storageId := "4_z134ab1f7e45796cc6950011e_f117076c66da42a22_d20180903_m010708_c002_v0001108_t0017"

				sc, err := trackdataserver.OpenStorageConnection()
				Expect(err).NotTo(HaveOccurred())

				const bytesPerRead int32 = 12000 // temporary fake value
				trackChunk := &pb.TrackChunk{
					StartPosition: 100,
					NumBytes:      bytesPerRead,
				}

				trackChunk, err = trackdataserver.GetTrackChunkFromStorage(storageId, trackChunk, sc)
				Expect(err).NotTo(HaveOccurred())
				Expect(trackChunk.NumBytes).To(Equal(bytesPerRead))
				Expect(trackChunk.Data).NotTo(BeNil())
				Expect(trackChunk.StartPosition).To(Equal(int32(100)))
			})
		})
		Context("with invalid arguments", func() {
			It("should respond with invalid_argument error", func() {
				storageId := "3_z134ab1f7e45796cc6950011e_f117076c66da42a22_d20180903_m010708_c002_v0001108_t0017"

				sc, err := trackdataserver.OpenStorageConnection()
				Expect(err).NotTo(HaveOccurred())

				const bytesPerRead int32 = 12000 // temporary fake value
				trackChunk := &pb.TrackChunk{
					StartPosition: 100,
					NumBytes:      bytesPerRead,
				}

				trackChunk, err = trackdataserver.GetTrackChunkFromStorage(storageId, trackChunk, sc)
				Expect(err).To(HaveOccurred())
			})
		})
		Context("with invalid user uuid", func() {
			It("should respond with invalid_argument error", func() {
			})
		})
	})
})
