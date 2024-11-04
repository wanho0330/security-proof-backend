// Code generated by github.com/jmattheis/goverter, DO NOT EDIT.
//go:build !goverter

package convert

import v1 "buf.build/gen/go/wanho/security-proof-api/protocolbuffers/go/api/v1"

type ControllerConverterImpl struct{}

func (c *ControllerConverterImpl) CreateRequestToProof(source *v1.CreateProofRequest) *v1.Proof {
	var pApiv1Proof *v1.Proof
	if source != nil {
		var apiv1Proof v1.Proof
		apiv1Proof.Num = (*source).Num
		apiv1Proof.Category = (*source).Category
		apiv1Proof.Description = (*source).Description
		apiv1Proof.UploadedUserIdx = (*source).UploadedUserIdx
		pApiv1Proof = &apiv1Proof
	}
	return pApiv1Proof
}
func (c *ControllerConverterImpl) DeleteRequestToProof(source *v1.DeleteProofRequest) *v1.Proof {
	var pApiv1Proof *v1.Proof
	if source != nil {
		var apiv1Proof v1.Proof
		apiv1Proof.Idx = (*source).Idx
		pApiv1Proof = &apiv1Proof
	}
	return pApiv1Proof
}
func (c *ControllerConverterImpl) UpdateRequestToProof(source *v1.UpdateProofRequest) *v1.Proof {
	var pApiv1Proof *v1.Proof
	if source != nil {
		var apiv1Proof v1.Proof
		apiv1Proof.Idx = (*source).Idx
		apiv1Proof.Num = (*source).Num
		apiv1Proof.Category = (*source).Category
		apiv1Proof.Description = (*source).Description
		apiv1Proof.UploadedUserIdx = (*source).UploadedUserIdx
		apiv1Proof.Confirm = (*source).Confirm
		pApiv1Proof = &apiv1Proof
	}
	return pApiv1Proof
}