// Package controller is returning the data needed for the proof based on user input.
package controller

import (
	"context"
	"errors"
	"net/http"
	"strconv"
	"strings"

	apiv1 "buf.build/gen/go/wanho/security-proof-api/protocolbuffers/go/api/v1"
	"connectrpc.com/connect"

	goverter "security-proof/internal/proof/convert"
	"security-proof/internal/proof/service"
	"security-proof/pkg/constants"
)

var conv = goverter.ControllerConverterImpl{}

// ProofController struct is composed of command and query from the service layer.
type ProofController struct {
	proofCommand *service.ProofCommand
	proofQuery   *service.ProofQuery
}

// NewProofController function is returning a ProofController struct that accept command and query from the service layer.
func NewProofController(proofCommand *service.ProofCommand, proofQuery *service.ProofQuery) *ProofController {
	return &ProofController{proofCommand: proofCommand, proofQuery: proofQuery}
}

// CreateProof method is returning a CreateProofResponse and an error, accepting a CreateProofRequest and a context.
func (c *ProofController) CreateProof(ctx context.Context, req *connect.Request[apiv1.CreateProofRequest]) (*connect.Response[apiv1.CreateProofResponse], error) {
	accessToken := req.Header().Get("accessToken")

	idx, err := c.proofCommand.CreateProof(ctx, conv.CreateRequestToProof(req.Msg), accessToken)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	res := connect.NewResponse(&apiv1.CreateProofResponse{
		Idx: idx,
	})
	return res, nil
}

// ReadProof method is returning a ReadProofResponse and an error, accepting a ReadProofRequest and a context.
func (c *ProofController) ReadProof(ctx context.Context, req *connect.Request[apiv1.ReadProofRequest]) (*connect.Response[apiv1.ReadProofResponse], error) {
	accessToken := req.Header().Get("accessToken")

	proof, err := c.proofQuery.ReadProof(ctx, req.Msg.Idx, accessToken)
	if err != nil {
		return nil, connect.NewError(connect.CodeNotFound, err)
	}
	res := connect.NewResponse(&apiv1.ReadProofResponse{
		Proof: proof,
	})
	return res, nil
}

// UpdateProof method is returning a UpdateProofResponse and an error, accepting a UpdateProofRequest and a context.
func (c *ProofController) UpdateProof(ctx context.Context, req *connect.Request[apiv1.UpdateProofRequest]) (*connect.Response[apiv1.UpdateProofResponse], error) {
	accessToken := req.Header().Get("accessToken")

	idx, err := c.proofCommand.UpdateProof(ctx, conv.UpdateRequestToProof(req.Msg), accessToken)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	res := connect.NewResponse(&apiv1.UpdateProofResponse{
		Idx: idx,
	})

	return res, nil
}

// DeleteProof method is returning a DeleteProofResponse and an error, accepting a DeleteProofRequest and a context.
func (c *ProofController) DeleteProof(ctx context.Context, req *connect.Request[apiv1.DeleteProofRequest]) (*connect.Response[apiv1.DeleteProofResponse], error) {
	accessToken := req.Header().Get("accessToken")

	err := c.proofCommand.DeleteProof(ctx, req.Msg.Idx, accessToken)
	if err != nil {
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	res := connect.NewResponse(&apiv1.DeleteProofResponse{})
	return res, nil
}

// ListProof method is returning a ListProofResponse and an error, accepting a ListProofRequest and a context.
func (c *ProofController) ListProof(ctx context.Context, req *connect.Request[apiv1.ListProofRequest]) (*connect.Response[apiv1.ListProofResponse], error) {
	// TODO : 서버 페이징 작업 필요
	accessToken := req.Header().Get("accessToken")

	proofs, err := c.proofQuery.ListProofs(ctx, req.Msg.Category, accessToken)
	if errors.Is(err, constants.ErrItemNotFound) {
		return nil, connect.NewError(connect.CodeNotFound, err)
	} else if err != nil {
		return nil, connect.NewError(connect.CodeUnknown, err)
	}

	res := connect.NewResponse(&apiv1.ListProofResponse{
		Proofs: proofs,
	})

	return res, nil
}

// UploadProof method is returning a UploadProofResponse and an error, accepting a UploadProofRequest and a context.
func (c *ProofController) UploadProof(ctx context.Context, req *connect.Request[apiv1.UploadProofRequest]) (*connect.Response[apiv1.UploadProofResponse], error) {
	accessToken := req.Header().Get("accessToken")

	idx, err := c.proofCommand.UploadProof(ctx, req.Msg.Idx, req.Msg.FirstImage, req.Msg.SecondImage, accessToken)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	res := connect.NewResponse(&apiv1.UploadProofResponse{
		Idx: idx,
	})

	return res, nil
}

// ConfirmProof method is returning a ConfirmProofResponse and an error, accepting a ConfirmProofRequest and a context.
func (c *ProofController) ConfirmProof(ctx context.Context, req *connect.Request[apiv1.ConfirmProofRequest]) (*connect.Response[apiv1.ConfirmProofResponse], error) {
	accessToken := req.Header().Get("accessToken")

	err := c.proofCommand.ConfirmProof(ctx, req.Msg.Idx, accessToken)
	if err != nil {
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	res := connect.NewResponse(&apiv1.ConfirmProofResponse{})
	return res, nil
}

// ConfirmUpdateProof method is returning a ConfirmUpdateProofResponse and an error, accepting a ConfirmUpdateProofRequest and a context.
func (c *ProofController) ConfirmUpdateProof(ctx context.Context, req *connect.Request[apiv1.ConfirmUpdateProofRequest]) (*connect.Response[apiv1.ConfirmUpdateProofResponse], error) {
	accessToken := req.Header().Get("accessToken")

	err := c.proofCommand.ConfirmUpdateProof(ctx, req.Msg.Idx, accessToken)
	if err != nil {
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	res := connect.NewResponse(&apiv1.ConfirmUpdateProofResponse{})
	return res, nil
}

// ReadLog method is returning a ReadLogResponse and an error, accepting a ReadLogRequest and a context.
func (c *ProofController) ReadLog(ctx context.Context, req *connect.Request[apiv1.ReadLogRequest]) (*connect.Response[apiv1.ReadLogResponse], error) {
	accessToken := req.Header().Get("accessToken")

	proof, err := c.proofQuery.ReadProofLog(ctx, req.Msg.Idx, accessToken)
	if err != nil {
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	res := connect.NewResponse(&apiv1.ReadLogResponse{
		Log: proof.LogPath,
	})
	return res, nil
}

// ReadFirstImage method is returning an image file, accepting a proof index.
func (c *ProofController) ReadFirstImage(w http.ResponseWriter, r *http.Request) {
	pathParts := strings.Split(r.URL.Path, "/")

	if len(pathParts) != 4 || pathParts[2] != "readFirstImage" {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	accessToken := r.Header.Get("accessToken")

	idxInt64, err := strconv.ParseInt(pathParts[3], 10, 32)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}

	path, err := c.proofQuery.ReadFirstProofImage(r.Context(), int32(idxInt64), accessToken)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}

	filePath := path
	http.ServeFile(w, r, filePath)
}

// ReadSecondImage method is returning an image file, accepting a proof index.
func (c *ProofController) ReadSecondImage(w http.ResponseWriter, r *http.Request) {

	pathParts := strings.Split(r.URL.Path, "/")

	if len(pathParts) != 4 || pathParts[2] != "readSecondImage" {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	accessToken := r.Header.Get("accessToken")

	idxInt64, err := strconv.ParseInt(pathParts[3], 10, 32)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}

	path, err := c.proofQuery.ReadSecondProofImage(r.Context(), int32(idxInt64), accessToken)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}

	filePath := path
	http.ServeFile(w, r, filePath)
}
