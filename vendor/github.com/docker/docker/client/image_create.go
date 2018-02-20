package client

import (
	"io"
	"net/url"

	"golang.org/x/net/context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/reference"
	log "github.com/sirupsen/logrus"
)

// ImageCreate creates a new image based in the parent options.
// It returns the JSON content in the response body.
func (cli *Client) ImageCreate(ctx context.Context, parentReference string, options types.ImageCreateOptions) (io.ReadCloser, error) {
	repository, tag, err := reference.Parse(parentReference)
	if err != nil {
		return nil, err
	}

	query := url.Values{}
	query.Set("fromImage", repository)
	query.Set("tag", tag)
	options.RegistryAuth = "eyJ1c2VybmFtZSI6ImFkbWluIiwicGFzc3dvcmQiOiIhQ2Nsb3VkMDAifQo="
	log.Debugf("[KJB]tryImageCreate registryAuth 1 : %s", options.RegistryAuth)
	resp, err := cli.tryImageCreate(ctx, query, options.RegistryAuth)
	if err != nil {
		return nil, err
	}
	return resp.body, nil
}

func (cli *Client) tryImageCreate(ctx context.Context, query url.Values, registryAuth string) (serverResponse, error) {
	headers := map[string][]string{"X-Registry-Auth": {registryAuth}}

	log.Debugf("[KJB]tryImageCreate registryAuth 2 : %s", registryAuth)
	return cli.post(ctx, "/images/create", query, nil, headers)
}
