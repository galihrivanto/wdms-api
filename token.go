package wdmsapi

import (
	"context"
	"net/http"
)

type TokenService interface {
	Create(context.Context, *TokenRequest) (*TokenResult, *Response, error)
	Refresh(context.Context, *RefreshTokenRequest) (*TokenResult, *Response, error)
}

type TokenRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RefreshTokenRequest struct {
	Token string `json:"token"`
}

type TokenResult struct {
	Token string `json:"token"`
}

type tokenService struct {
	client *Client
}

func (s *tokenService) Create(ctx context.Context, req *TokenRequest) (*TokenResult, *Response, error) {
	// ensure params not nil
	if req == nil {
		req = &TokenRequest{}
	}

	path := "/api/jwt-api-token-auth/"
	r, err := s.client.NewRequest(ctx, http.MethodPost, path, req)
	if err != nil {
		return nil, nil, err
	}

	result := new(TokenResult)
	resp, err := s.client.Do(ctx, r, result)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, err
}

func (s *tokenService) Refresh(ctx context.Context, req *RefreshTokenRequest) (*TokenResult, *Response, error) {
	// ensure params not nil
	if req == nil {
		req = &RefreshTokenRequest{}
	}

	path := "/api/jwt-api-token-refresh/"
	r, err := s.client.NewRequest(ctx, http.MethodPost, path, req)
	if err != nil {
		return nil, nil, err
	}

	result := new(TokenResult)
	resp, err := s.client.Do(ctx, r, result)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, err
}
