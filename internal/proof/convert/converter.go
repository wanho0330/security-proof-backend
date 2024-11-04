// Package convert is the package needed for data transformation.
package convert

import (
	"time"

	apiv1 "buf.build/gen/go/wanho/security-proof-api/protocolbuffers/go/api/v1"
	"google.golang.org/protobuf/types/known/timestamppb"

	"security-proof/internal/db/security_proof/proof/model"
)

// ControllerConverter interface is defining data transformation between the controller layer and the service layer.
// goverter:converter
// goverter:output:file ./controllerConvert.go
// goverter:output:package security-proof/internal/proof/convert
// goverter:extend TimeToTimestamppb TimeToPTimestamppb TimestampppbToTime TimestampppbToPTime
type ControllerConverter interface {
	// goverter:ignore state sizeCache unknownFields Idx FirstImagePath SecondImagePath LogPath CreatedUserIdx CreatedUserId CreatedAt UpdatedUserIdx UpdatedUserId UpdatedAt UploadedUserId UploadedAt Confirm TokenId
	CreateRequestToProof(*apiv1.CreateProofRequest) *apiv1.Proof

	// goverter:ignore state sizeCache unknownFields FirstImagePath SecondImagePath LogPath CreatedUserIdx CreatedUserId CreatedAt UpdatedUserIdx UpdatedUserId UpdatedAt UploadedUserId UploadedAt TokenId
	UpdateRequestToProof(*apiv1.UpdateProofRequest) *apiv1.Proof

	// goverter:ignore state sizeCache unknownFields Category Description FirstImagePath SecondImagePath LogPath CreatedUserIdx CreatedUserId CreatedAt UpdatedUserIdx UpdatedUserId UpdatedAt UploadedUserIdx UploadedUserId UploadedAt Confirm Num TokenId
	DeleteRequestToProof(*apiv1.DeleteProofRequest) *apiv1.Proof
}

// ServiceConverter interface is defining data transformation between the service layer and the repository layer.
// goverter:converter
// goverter:output:file ./serviceConvert.go
// goverter:output:package security-proof/internal/proof/convert
// goverter:extend TimeToTimestamppb TimeToPTimestamppb TimestampppbToTime TimestampppbToPTime PStringToString Pint32ToInt32
type ServiceConverter interface {
	// goverter:map TokenId TokenID
	ProtoToModel(*apiv1.Proof) *model.Proof
	// goverter:ignore state sizeCache unknownFields CreatedUserId UpdatedUserId UploadedUserId
	// goverter:map TokenID TokenId
	ModelToProto(*model.Proof) *apiv1.Proof
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

// PStringToString functions is returning a string accepting string pointer.
func PStringToString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

// Pint32ToInt32 function is returning a int32 accepting int32 pointer.
func Pint32ToInt32(i *int32) int32 {
	if i == nil {
		return 0
	}
	return *i
}
