package grpcserver

import (
	"context"
	"github.com/lekan-pvp/short/internal/cookies"
	"github.com/lekan-pvp/short/internal/makeshort"
	"github.com/lekan-pvp/short/internal/models"
	pb "github.com/lekan-pvp/short/internal/shortengrpc"
	"github.com/lekan-pvp/short/internal/storage/dbrepo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type UsersServer struct {
	pb.UnimplementedShortenGrpcServer
	users dbrepo.DBRepo
}

func (s *UsersServer) AddShort(ctx context.Context, in *pb.PostRequest) (*pb.ShortResponse, error) {
	var response pb.ShortResponse
	var err error
	var id string

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.DataLoss, "UnaryEcho: failed to get metadata")
	}
	if t, ok := md["token"]; !ok {
		cookie := cookies.New()
		token := cookie.Value
		header := metadata.Pairs("token", token)
		grpc.SendHeader(ctx, header)
		trailer := metadata.Pairs("token", token)
		grpc.SetTrailer(ctx, trailer)
	} else {
		for _, e := range t {
			id = e
		}
	}

	short := makeshort.GenerateShortLink(in.Id.Id, in.Url.Url)

	rec := models.Storage{
		UUID:        id,
		ShortURL:    short,
		OriginalURL: in.Url.Url,
		DeleteFlag:  false,
	}

	response.Short.Short, err = s.users.PostURL(ctx, rec)

	if err != nil {
		response.Error = err.Error()
		return &response, err
	}

	return &response, nil
}

func (s *UsersServer) GetShortURL(ctx context.Context, in *pb.GetRequest) (*pb.OriginResponse, error) {
	var response pb.OriginResponse
	var err error

	res, err := s.users.GetOriginal(ctx, in.Short.Short)

	if err != nil {
		response.Error = err.Error()
		return &response, err
	}

	response.Url.Url = res.URL
	response.Deleted = res.Deleted

	return &response, nil
}

func (s *UsersServer) GetURLsList(ctx context.Context) (*pb.ListResponse, error) {
	var response pb.ListResponse
	var err error
	var id string

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.DataLoss, "UnaryEcho: failed to get metadata")
	}
	if t, ok := md["token"]; !ok {
		cookie := cookies.New()
		token := cookie.Value
		header := metadata.Pairs("token", token)
		grpc.SendHeader(ctx, header)
		trailer := metadata.Pairs("token", token)
		grpc.SetTrailer(ctx, trailer)
	} else {
		for _, e := range t {
			id = e
		}
	}

	res, err := s.users.GetURLsList(ctx, id)
	if err != nil {
		response.Error = err.Error()
		return &response, err
	}
	for i, v := range res {
		response.ShortUrl[i] = v.ShortURL
		response.OriginalUrl[i] = v.ShortURL
	}
	return &response, nil
}

func (s *UsersServer) PostBatchURLs(ctx context.Context, in *pb.BatchRequest) (*pb.BatchResponse, error) {
	var response pb.BatchResponse
	var err error
	var id string

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.DataLoss, "UnaryEcho: failed to get metadata")
	}
	if t, ok := md["token"]; !ok {
		cookie := cookies.New()
		token := cookie.Value
		header := metadata.Pairs("token", token)
		grpc.SendHeader(ctx, header)
		trailer := metadata.Pairs("token", token)
		grpc.SetTrailer(ctx, trailer)
	} else {
		for _, e := range t {
			id = e
		}
	}

	var req []models.BatchRequest

	for i, v := range in.CorrelationId {
		req[i].CorrelationID = v
	}

	for i, v := range in.OriginalUrl {
		req[i].OriginalURL = v
	}

	res, err := s.users.BatchShorten(ctx, id, req)
	if err != nil {
		response.Error = err.Error()
		return &response, err
	}

	for _, v := range res {
		response.CorrelationId = append(response.CorrelationId, v.CorrelationID)
		response.ShortUrl = append(response.ShortUrl, v.ShortURL)
	}
	return &response, nil
}

func (s *UsersServer) SoftDelURLs(ctx context.Context, in *pb.DelRequest) (*pb.DelResponse, error) {
	var response pb.DelResponse
	var err error

	var req []string

	var id string

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.DataLoss, "UnaryEcho: failed to get metadata")
	}
	if t, ok := md["token"]; !ok {
		cookie := cookies.New()
		token := cookie.Value
		header := metadata.Pairs("token", token)
		grpc.SendHeader(ctx, header)
		trailer := metadata.Pairs("token", token)
		grpc.SetTrailer(ctx, trailer)
	} else {
		for _, e := range t {
			id = e
		}
	}

	for _, v := range in.Short {
		req = append(req, v)
	}

	err = s.users.SoftDelete(ctx, req, id)
	if err != nil {
		response.Error = err.Error()
		return &response, err
	}

	response.Error = ""

	return &response, nil
}
