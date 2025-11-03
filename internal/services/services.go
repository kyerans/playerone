package services

import (
	"context"
	"encoding/base64"
	"errors"

	servicepb "github.com/kyerans/playerone/api/services/v1"
	"github.com/kyerans/playerone/internal/repos"
)

var _ servicepb.PlayerOneServiceServer = (*Service)(nil)

func New() *Service {
	return &Service{
		repo: repos.New(),
	}
}

type Service struct {
	repo *repos.Repository
}

func (s *Service) Register(ctx context.Context,
	req *servicepb.RegisterRequest) (*servicepb.RegisterResponse, error) {

	s.repo.Set(req.GetKid(), req.GetKid())

	return &servicepb.RegisterResponse{}, nil
}

func (s *Service) License(ctx context.Context,
	req *servicepb.LicenseRequest) (*servicepb.LicenseResponse, error) {

	resp := &servicepb.LicenseResponse{}

	for _, kid := range req.Kids {
		key := s.repo.Get(kid)
		if key == nil {
			return nil, errors.New("key not found for kid: " + kid)
		}

		resp.Keys = append(resp.Keys, &servicepb.LicenseResponse_Key{
			Kty: "oct",
			K:   base64.StdEncoding.EncodeToString(key),
			Kid: kid,
		})
	}

	return resp, nil
}

func (s *Service) LicenseRelease(ctx context.Context,
	req *servicepb.LicenseReleaseRequest) (*servicepb.LicenseReleaseResponse, error) {
	return nil, nil
}
