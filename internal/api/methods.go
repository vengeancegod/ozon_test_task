package api

import (
	"context"
	"url-shortener/internal/models"
	desc "url-shortener/pkg/urlshortener_v3"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) OriginalToShort(ctx context.Context, req *desc.OriginalURLRequest) (*desc.ShortURLResponse, error) {
	if req.OriginalUrl == "" {
		return nil, status.Errorf(codes.InvalidArgument, "originalURL not found")
	}

	resp, err := i.URLService.OriginalToShort(models.ShortRequest{
		OriginalURL: req.OriginalUrl,
	})
	if err != nil {
		return nil, err
	}

	return &desc.ShortURLResponse{
		ShortUrl: resp.ShortURL,
	}, nil
}

func (i *Implementation) ShortToOriginal(ctx context.Context, req *desc.ShortURLRequest) (*desc.OriginalURLResponse, error) {
	originalRequest, err := i.URLService.ShortToOriginal(req.ShortUrl)

	if err != nil {
		return nil, err
	}

	return &desc.OriginalURLResponse{
		OriginalUrl: originalRequest.OriginalURL,
	}, nil
}
