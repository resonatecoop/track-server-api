package playserver_test

import (
	"context"

	"github.com/go-pg/pg"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	uuid "github.com/satori/go.uuid"
	"github.com/twitchtv/twirp"

	pb "track-server-api/rpc"
)

var _ = Describe("Play server", func() {
	const invalid_argument_code twirp.ErrorCode = "invalid_argument"
	const not_found_code twirp.ErrorCode = "not_found"
	const internal_code twirp.ErrorCode = "internal"

	Describe("Play", func() {
		Context("with valid track and user uuid", func() {
			It("should respond with track stream if track exists", func() {
				userTrackPB := &pb.UserTrack{TrackId: "0", UserId: "0"}
				resp, err := service.Play(context.Background(), userTrackPB)
				Expect(err).NotTo(HaveOccurred())
			})
			It("should respond with not_found error if track does not exist", func() {
				userTrackPB := &pb.UserTrack{TrackId: "0", UserId: "0"}
				resp, err := service.Play(context.Background(), userTrackPB)
				Expect(err).To(HaveOccurred())
				twerr := err.(twirp.Error)
				Expect(twerr.Code()).To(Equal(not_found_code))
			})
		})
		Context("with invalid track uuid", func() {
			It("should respond with invalid_argument error", func() {
			})
		})
		Context("with invalid user uuid", func() {
			It("should respond with invalid_argument error", func() {
			})
		})
	})
})
