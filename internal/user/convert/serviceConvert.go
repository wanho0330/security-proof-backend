// Code generated by github.com/jmattheis/goverter, DO NOT EDIT.
//go:build !goverter

package convert

import (
	v1 "buf.build/gen/go/wanho/security-proof-api/protocolbuffers/go/api/v1"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	model "security-proof/internal/db/security_proof/user/model"
	"time"
)

type ServiceConverterImpl struct{}

func (c *ServiceConverterImpl) ModelToProto(source *model.User) *v1.User {
	var pApiv1User *v1.User
	if source != nil {
		var apiv1User v1.User
		apiv1User.Idx = (*source).Idx
		apiv1User.Id = (*source).ID
		apiv1User.Passwd = (*source).Passwd
		apiv1User.CreatedAt = TimeToPTimestamppb((*source).CreatedAt)
		apiv1User.UpdatedAt = c.pTimeTimeToPTimestamppbTimestamp((*source).UpdatedAt)
		apiv1User.Name = (*source).Name
		apiv1User.Email = (*source).Email
		apiv1User.Role = (*source).Role
		pApiv1User = &apiv1User
	}
	return pApiv1User
}
func (c *ServiceConverterImpl) ProtoToModel(source *v1.User) *model.User {
	var pModelUser *model.User
	if source != nil {
		var modelUser model.User
		modelUser.Idx = (*source).Idx
		modelUser.ID = (*source).Id
		modelUser.Passwd = (*source).Passwd
		modelUser.CreatedAt = TimestampppbToTime((*source).CreatedAt)
		modelUser.UpdatedAt = TimestampppbToPTime((*source).UpdatedAt)
		modelUser.Name = (*source).Name
		modelUser.Email = (*source).Email
		modelUser.Role = (*source).Role
		pModelUser = &modelUser
	}
	return pModelUser
}
func (c *ServiceConverterImpl) pTimeTimeToPTimestamppbTimestamp(source *time.Time) *timestamppb.Timestamp {
	var pTimestamppbTimestamp *timestamppb.Timestamp
	if source != nil {
		timestamppbTimestamp := TimeToTimestamppb((*source))
		pTimestamppbTimestamp = &timestamppbTimestamp
	}
	return pTimestamppbTimestamp
}
