// Package convert is the package needed for data transformation.
package convert

import (
	"time"

	apiv1 "buf.build/gen/go/wanho/security-proof-api/protocolbuffers/go/api/v1"
	"google.golang.org/protobuf/types/known/timestamppb"

	"security-proof/internal/db/security_proof/user/model"
)

// ControllerConverter interface is defining data transformation between the controller layer and the service layer.
// goverter:converter
// goverter:output:file ./controllerConvert.go
// goverter:output:package security-proof/internal/user/convert
// goverter:extend TimeToTimestamppb TimeToPTimestamppb TimestampppbToTime TimestampppbToPTime
type ControllerConverter interface {
	// goverter:ignore state sizeCache unknownFields Idx CreatedAt UpdatedAt
	CreateRequestToUser(*apiv1.CreateUserRequest) *apiv1.User

	// goverter:ignore state sizeCache unknownFields Id CreatedAt UpdatedAt Passwd
	UpdateRequestToUser(*apiv1.UpdateUserRequest) *apiv1.User

	// goverter:ignore state sizeCache unknownFields Id Passwd CreatedAt UpdatedAt Name Email Role
	DeleteRequestToUser(*apiv1.DeleteUserRequest) *apiv1.User

	// goverter:ignore state sizeCache unknownFields Idx CreatedAt UpdatedAt Name Email Role
	SignInRequestToUser(*apiv1.SignInRequest) *apiv1.User
}

// ServiceConverter interface is defining data transformation between the service layer and the repository layer.
// goverter:converter
// goverter:output:file ./serviceConvert.go
// goverter:output:package security-proof/internal/user/convert
// goverter:extend TimeToTimestamppb TimeToPTimestamppb TimestampppbToTime TimestampppbToPTime
type ServiceConverter interface {
	// goverter:map Id ID
	ProtoToModel(*apiv1.User) *model.User
	// goverter:ignore state sizeCache unknownFields
	// goverter:map ID Id
	ModelToProto(*model.User) *apiv1.User
}

// TimeToTimestamppb function is returning a Timestamp accepting Time.
func TimeToTimestamppb(t time.Time) timestamppb.Timestamp {
	return *timestamppb.New(t)
}

// TimeToPTimestamppb function is returning a Timestamp pointer accepting Time.
func TimeToPTimestamppb(t time.Time) *timestamppb.Timestamp {
	return timestamppb.New(t)
}

// TimestampppbToTime function is returning a Time accepting Timestamp pointer.
func TimestampppbToTime(t *timestamppb.Timestamp) time.Time {
	return t.AsTime()
}

// TimestampppbToPTime functions is returning a Time pointer accepting Timestamp pointer.
func TimestampppbToPTime(t *timestamppb.Timestamp) *time.Time {
	result := t.AsTime()
	return &result
}
