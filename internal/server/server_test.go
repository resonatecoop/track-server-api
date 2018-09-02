package trackdataserver_test

import (
	"context"
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	uuid "github.com/satori/go.uuid"
	"github.com/twitchtv/twirp"

	pb "track-server-api/rpc"
)

var _ = Describe("Track data server", func() {
	const invalid_argument_code twirp.ErrorCode = "invalid_argument"
	const not_found_code twirp.ErrorCode = "not_found"
	const internal_code twirp.ErrorCode = "internal"

	Describe("Stream ", func() {
		Context("with valid track and user uuid", func() {
			It("should respond with track stream if track exists", func() {
				userTrackPB := &pb.UserTrack{TrackId: newTrackData.TrackId.String(), UserId: newTrackData.UserId.String()}
				_, err := service.StreamTrackData(context.Background(), userTrackPB)
				Expect(err).NotTo(HaveOccurred())
			})
			It("should respond with not_found error if track does not exist", func() {
				userTrackPB := &pb.UserTrack{TrackId: uuid.NewV4().String(), UserId: uuid.NewV4().String()}
				_, err := service.StreamTrackData(context.Background(), userTrackPB)
				Expect(err).To(HaveOccurred())
				fmt.Printf("err%v xx %v \n", err, userTrackPB)
				twerr := err.(twirp.Error)
				Expect(twerr.Code()).To(Equal(not_found_code))
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
