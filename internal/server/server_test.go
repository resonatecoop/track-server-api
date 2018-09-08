package trackdataserver_test

import (
	"context"
	"fmt"
	"io/ioutil"

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

	Describe("StreamTrackData", func() {
		Context("with valid track and user uuid", func() {
			It("should respond with track stream if track exists", func() {
				userTrackPB := &pb.UserTrack{TrackId: newTrackData.TrackId.String(), UserId: newTrackData.UserId.String()}
				trackDataStream, err := service.StreamTrackData(context.Background(), userTrackPB)
				Expect(err).NotTo(HaveOccurred())
				count := 0
				for tdOrErr := range trackDataStream {
					Expect(tdOrErr.Err).NotTo(HaveOccurred())
					td := tdOrErr.Msg
					Expect(td.StartPosition).To(Equal(trackdataserver.BytesPerRead * int32(count)))
					Expect(td.NumBytes).NotTo(Equal(0))
					fmt.Printf("%d Streaming data at pos:%v len:%v\n", count, td.StartPosition, td.NumBytes)
					count += 1
				}
				Expect(count).To(Equal(2)) // number of 60000 byte chunks in test file
			})
			It("should respond with not_found error if track does not exist", func() {
				userTrackPB := &pb.UserTrack{TrackId: uuid.NewV4().String(), UserId: uuid.NewV4().String()}
				_, err := service.StreamTrackData(context.Background(), userTrackPB)
				Expect(err).To(HaveOccurred())
			})
		})
		Context("with invalid track uuid", func() {
			It("should respond with invalid_argument error", func() {
				userTrackPB := &pb.UserTrack{TrackId: "0", UserId: newTrackData.UserId.String()}
				_, err := service.StreamTrackData(context.Background(), userTrackPB)
				Expect(err).To(HaveOccurred())
			})
		})
		Context("with invalid user uuid", func() {
			It("should respond with invalid_argument error", func() {
				userTrackPB := &pb.UserTrack{TrackId: newTrackData.TrackId.String(), UserId: "0"}
				_, err := service.StreamTrackData(context.Background(), userTrackPB)
				Expect(err).To(HaveOccurred())
			})
		})
	})
	Describe("UploadTrackData", func() {
		Context("with valid track and user uuid", func() {
			It("should respond with track stream if track exists", func() {
				dat, err := ioutil.ReadFile("/Users/jhno/Desktop/bendy_13s.mp4")
				Expect(err).NotTo(HaveOccurred())

				trackChunk := &pb.TrackChunk{
					StartPosition: 0,
					NumBytes:      int32(len(dat)),
					Data:          dat,
				}

				resp, err := service.UploadTrackData(context.Background(), trackChunk)
				Expect(err).NotTo(HaveOccurred())

				fmt.Printf("resp %v\n", resp)

				Expect(resp.TrackServerId).NotTo(BeNil())
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
