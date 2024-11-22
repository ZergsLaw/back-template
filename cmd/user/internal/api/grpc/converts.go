package grpc

import (
	user_pb "github.com/Bar-Nik/back-template/api/user/v1"
	"github.com/Bar-Nik/back-template/cmd/user/internal/app"
	"github.com/Bar-Nik/back-template/internal/dom"
)

func toUser(u app.User) *user_pb.User {
	return &user_pb.User{
		Id:       u.ID.String(),
		Username: u.Name,
		Email:    u.Email,
		FullName: u.FullName,
		AvatarId: u.AvatarID.String(),
		Kind:     dom.UserStatusToAPI(u.Status),
	}
}

func toUserFile(f app.AvatarInfo) *user_pb.UserAvatar {
	return &user_pb.UserAvatar{
		UserId: f.OwnerID.String(),
		FileId: f.FileID.String(),
	}
}
