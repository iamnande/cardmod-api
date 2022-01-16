package magicapi

import (
	"context"

	"github.com/google/uuid"
	"github.com/iamnande/cardmod/internal/daos"
	"github.com/iamnande/cardmod/internal/domains/magic"
	"github.com/iamnande/cardmod/pkg/api/magicv1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// api is the service implementation of the generate MagicAPI gRPC service.
type api struct {
	magicv1.UnimplementedMagicAPIServer
	magicRepository daos.MagicDAO
}

// New initializes a new magic api instance.
func New(magicRepository daos.MagicDAO) api {
	return api{
		magicRepository: magicRepository,
	}
}

// ListMagics lists all available magic entities.
func (api *api) ListMagics(ctx context.Context, request *magicv1.ListMagicsRequest) (*magicv1.ListMagicsResponse, error) {

	// list: list the available magics
	magics, err := api.magicRepository.ListMagics(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list magics: %s", err)
	}

	// list: return response to caller
	return &magicv1.ListMagicsResponse{
		Magics: marshalMagics(magics),
	}, nil

}

// DescribeMagic describes an existing magic entity.
func (api *api) DescribeMagic(ctx context.Context, request *magicv1.DescribeMagicRequest) (*magicv1.DescribeMagicResponse, error) {

	// describe: parse input uuid
	id, err := uuid.Parse(request.GetMagicId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid magic id")
	}

	// describe: describe magic
	magic, err := api.magicRepository.DescribeMagic(ctx, id)
	if err != nil {
		return nil, err
	}

	// describe: return magic to caller
	return &magicv1.DescribeMagicResponse{
		Magic: marshalMagic(magic),
	}, nil

}

func marshalMagic(magic magic.Magic) *magicv1.Magic {
	return &magicv1.Magic{
		Name: magic.Name(),
		Id:   magic.ID().String(),
	}
}

func marshalMagics(records []magic.Magic) []*magicv1.Magic {
	magics := make([]*magicv1.Magic, len(records))
	for i := range magics {
		magics[i] = marshalMagic(records[i])
	}
	return magics
}
