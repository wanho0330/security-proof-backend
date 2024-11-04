// Package convert is the package needed for data transformation.
package convert

import (
	"time"

	apiv1 "buf.build/gen/go/wanho/security-proof-api/protocolbuffers/go/api/v1"
	"google.golang.org/protobuf/types/known/timestamppb"

	"security-proof/internal/db/security_proof/proof/model"
)

// ServiceConverter interface is defining data transformation between the service layer and the repository layer.
// goverter:converter
// goverter:output:file ./serviceConvert.go
// goverter:output:package security-proof/internal/dashboard/convert
// goverter:extend TimeToTimestamppb TimeToPTimestamppb TimestampppbToTime TimestampppbToPTime PStringToString Pint32ToInt32
type ServiceConverter interface {
	// goverter:ignore state sizeCache unknownFields CreatedUserId UpdatedUserId UploadedUserId
	ModelToNotConfirmProto(*model.Proof) *apiv1.NotConfirmProof

	// goverter:ignore state sizeCache unknownFields CreatedUserId UpdatedUserId UploadedUserId
	ModelToNotUploadProto(*model.Proof) *apiv1.NotUploadProof
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
