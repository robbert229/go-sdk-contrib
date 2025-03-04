package service

import (
	"errors"

	"github.com/bufbuild/connect-go"
	"github.com/open-feature/go-sdk/pkg/openfeature"
	"golang.org/x/net/context"

	flagdModels "github.com/open-feature/flagd/pkg/model"
	log "github.com/sirupsen/logrus"
	schemaV1 "go.buf.build/open-feature/flagd-connect/open-feature/flagd/schema/v1"
	"google.golang.org/protobuf/types/known/structpb"
)

type ServiceConfiguration struct {
	Port            uint16
	Host            string
	CertificatePath string
	SocketPath      string
}

// Service handles the client side  interface for the flagd server
type Service struct {
	Client iClient
}

const ConnectionError = "connection not made"

// ResolveBoolean handles the flag evaluation response from the flagd ResolveBoolean rpc
func (s *Service) ResolveBoolean(ctx context.Context, flagKey string, evalCtx map[string]interface{}) (*schemaV1.ResolveBooleanResponse, error) {
	client := s.Client.Instance()
	if client == nil {
		return &schemaV1.ResolveBooleanResponse{
			Reason: string(openfeature.ErrorReason),
		}, openfeature.NewProviderNotReadyResolutionError(ConnectionError)
	}
	evalCtxF, err := structpb.NewStruct(evalCtx)
	if err != nil {
		log.Error(err)
		return &schemaV1.ResolveBooleanResponse{
			Reason: string(openfeature.ErrorReason),
		}, openfeature.NewParseErrorResolutionError(err.Error())
	}
	res, err := client.ResolveBoolean(ctx, connect.NewRequest(&schemaV1.ResolveBooleanRequest{
		FlagKey: flagKey,
		Context: evalCtxF,
	}))
	if err != nil {
		return &schemaV1.ResolveBooleanResponse{
			Reason: string(openfeature.ErrorReason),
		}, handleError(err)
	}
	return res.Msg, nil
}

// ResolveString handles the flag evaluation response from the  flagd interface ResolveString rpc
func (s *Service) ResolveString(ctx context.Context, flagKey string, evalCtx map[string]interface{}) (*schemaV1.ResolveStringResponse, error) {
	client := s.Client.Instance()
	if client == nil {
		return &schemaV1.ResolveStringResponse{
			Reason: flagdModels.ErrorReason,
		}, openfeature.NewProviderNotReadyResolutionError(ConnectionError)
	}
	evalCtxF, err := structpb.NewStruct(evalCtx)
	if err != nil {
		log.Error(err)
		return &schemaV1.ResolveStringResponse{
			Reason: flagdModels.ErrorReason,
		}, openfeature.NewParseErrorResolutionError(err.Error())
	}
	res, err := client.ResolveString(ctx, connect.NewRequest(&schemaV1.ResolveStringRequest{
		FlagKey: flagKey,
		Context: evalCtxF,
	}))
	if err != nil {
		return &schemaV1.ResolveStringResponse{
			Reason: string(openfeature.ErrorReason),
		}, handleError(err)
	}
	return res.Msg, nil
}

// ResolveFloat handles the flag evaluation response from the  flagd interface ResolveFloat rpc
func (s *Service) ResolveFloat(ctx context.Context, flagKey string, evalCtx map[string]interface{}) (*schemaV1.ResolveFloatResponse, error) {
	client := s.Client.Instance()
	if client == nil {
		return &schemaV1.ResolveFloatResponse{
			Reason: flagdModels.ErrorReason,
		}, openfeature.NewProviderNotReadyResolutionError(ConnectionError)
	}
	evalCtxF, err := structpb.NewStruct(evalCtx)
	if err != nil {
		log.Error(err)
		return &schemaV1.ResolveFloatResponse{
			Reason: flagdModels.ErrorReason,
		}, openfeature.NewParseErrorResolutionError(err.Error())
	}
	res, err := client.ResolveFloat(ctx, connect.NewRequest(&schemaV1.ResolveFloatRequest{
		FlagKey: flagKey,
		Context: evalCtxF,
	}))
	if err != nil {
		return &schemaV1.ResolveFloatResponse{
			Reason: string(openfeature.ErrorReason),
		}, handleError(err)
	}
	return res.Msg, nil
}

// ResolveInt handles the flag evaluation response from the  flagd interface ResolveNumber rpc
func (s *Service) ResolveInt(ctx context.Context, flagKey string, evalCtx map[string]interface{}) (*schemaV1.ResolveIntResponse, error) {
	client := s.Client.Instance()
	if client == nil {
		return &schemaV1.ResolveIntResponse{
			Reason: flagdModels.ErrorReason,
		}, openfeature.NewProviderNotReadyResolutionError(ConnectionError)
	}
	evalCtxF, err := structpb.NewStruct(evalCtx)
	if err != nil {
		log.Error(err)
		return &schemaV1.ResolveIntResponse{
			Reason: flagdModels.ErrorReason,
		}, openfeature.NewParseErrorResolutionError(err.Error())
	}
	res, err := client.ResolveInt(ctx, connect.NewRequest(&schemaV1.ResolveIntRequest{
		FlagKey: flagKey,
		Context: evalCtxF,
	}))
	if err != nil {
		return &schemaV1.ResolveIntResponse{
			Reason: string(openfeature.ErrorReason),
		}, handleError(err)
	}
	return res.Msg, nil
}

// ResolveObject handles the flag evaluation response from the  flagd interface ResolveObject rpc
func (s *Service) ResolveObject(ctx context.Context, flagKey string, evalCtx map[string]interface{}) (*schemaV1.ResolveObjectResponse, error) {
	client := s.Client.Instance()
	if client == nil {
		return &schemaV1.ResolveObjectResponse{
			Reason: flagdModels.ErrorReason,
		}, openfeature.NewProviderNotReadyResolutionError(ConnectionError)
	}
	evalCtxF, err := structpb.NewStruct(evalCtx)
	if err != nil {
		log.Error(err)
		return &schemaV1.ResolveObjectResponse{
			Reason: flagdModels.ErrorReason,
		}, openfeature.NewParseErrorResolutionError(err.Error())
	}
	res, err := client.ResolveObject(ctx, connect.NewRequest(&schemaV1.ResolveObjectRequest{
		FlagKey: flagKey,
		Context: evalCtxF,
	}))
	if err != nil {
		return &schemaV1.ResolveObjectResponse{
			Reason: string(openfeature.ErrorReason),
		}, handleError(err)
	}
	return res.Msg, nil
}

func handleError(err error) openfeature.ResolutionError {
	connectErr := &connect.Error{}
	errors.As(err, &connectErr)
	switch connectErr.Code() {
	case connect.CodeUnavailable:
		return openfeature.NewProviderNotReadyResolutionError(ConnectionError)
	case connect.CodeNotFound:
		return openfeature.NewFlagNotFoundResolutionError(err.Error())
	case connect.CodeInvalidArgument:
		return openfeature.NewTypeMismatchResolutionError(err.Error())
	case connect.CodeDataLoss:
		return openfeature.NewParseErrorResolutionError(err.Error())
	}
	return openfeature.NewGeneralResolutionError(err.Error())
}
