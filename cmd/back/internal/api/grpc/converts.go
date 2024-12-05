package grpc

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	user_pb "github.com/ZergsLaw/back-template/api/user/v1"
	"github.com/ZergsLaw/back-template/cmd/back/internal/app"
	dom "github.com/ZergsLaw/back-template/internal/dom"
)

func toUser(u app.User) *user_pb.User {
	return &user_pb.User{
		Id:        u.ID.String(),
		Username:  u.Name,
		Email:     u.Email,
		FullName:  u.FullName,
		AvatarId:  u.AvatarID.String(),
		Kind:      dom.UserStatusToAPI(u.Status),
		CreatedAt: timestamppb.New(u.CreatedAt),
		UpdatedAt: timestamppb.New(u.UpdatedAt),
	}
}

func toUserFile(f app.AvatarInfo) *user_pb.UserAvatar {
	return &user_pb.UserAvatar{
		UserId: f.OwnerID.String(),
		FileId: f.FileID.String(),
	}
}
