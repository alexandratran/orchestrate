package client

import (
	"context"
	"fmt"
	"strings"

	"github.com/containous/traefik/v2/pkg/log"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/pkg/errors"
	clientutils "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/pkg/http/client-utils"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/pkg/http/httputil"
	types "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/pkg/types/api"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/pkg/types/entities"
)

func (c *HTTPClient) GetFaucet(ctx context.Context, uuid string) (*types.FaucetResponse, error) {
	reqURL := fmt.Sprintf("%v/faucets/%s", c.config.URL, uuid)
	resp := &types.FaucetResponse{}

	err := callWithBackOff(ctx, c.config.backOff, func() error {
		response, err := clientutils.GetRequest(ctx, c.client, reqURL)
		if err != nil {
			errMessage := "error while getting faucet"
			log.FromContext(ctx).WithError(err).Error(errMessage)
			return errors.ServiceConnectionError(errMessage).ExtendComponent(component)
		}

		defer clientutils.CloseResponse(response)
		return httputil.ParseResponse(ctx, response, resp)
	})

	return resp, err
}

func (c *HTTPClient) SearchFaucets(ctx context.Context, filters *entities.FaucetFilters) ([]*types.FaucetResponse, error) {
	reqURL := fmt.Sprintf("%v/faucets", c.config.URL)
	var resp []*types.FaucetResponse

	var qParams []string
	if len(filters.Names) > 0 {
		qParams = append(qParams, "names="+strings.Join(filters.Names, ","))
	}

	if filters.ChainRule != "" {
		qParams = append(qParams, "chain_rule="+filters.ChainRule)
	}

	if len(qParams) > 0 {
		reqURL = reqURL + "?" + strings.Join(qParams, "&")
	}

	err := callWithBackOff(ctx, c.config.backOff, func() error {
		response, err := clientutils.GetRequest(ctx, c.client, reqURL)
		if err != nil {
			errMessage := "error while searching faucets"
			log.FromContext(ctx).WithError(err).Error(errMessage)
			return errors.ServiceConnectionError(errMessage).ExtendComponent(component)
		}
		defer clientutils.CloseResponse(response)
		return httputil.ParseResponse(ctx, response, &resp)
	})

	return resp, err
}

func (c *HTTPClient) RegisterFaucet(ctx context.Context, request *types.RegisterFaucetRequest) (*types.FaucetResponse, error) {
	reqURL := fmt.Sprintf("%v/faucets", c.config.URL)
	resp := &types.FaucetResponse{}

	err := callWithBackOff(ctx, c.config.backOff, func() error {
		response, err := clientutils.PostRequest(ctx, c.client, reqURL, request)
		if err != nil {
			errMessage := "error while registering faucet"
			log.FromContext(ctx).WithError(err).Error(errMessage)
			return errors.ServiceConnectionError(errMessage).ExtendComponent(component)
		}
		defer clientutils.CloseResponse(response)
		return httputil.ParseResponse(ctx, response, resp)
	})

	return resp, err
}

func (c *HTTPClient) UpdateFaucet(ctx context.Context, uuid string, request *types.UpdateFaucetRequest) (*types.FaucetResponse, error) {
	reqURL := fmt.Sprintf("%v/faucets/%v", c.config.URL, uuid)
	resp := &types.FaucetResponse{}

	err := callWithBackOff(ctx, c.config.backOff, func() error {
		response, err := clientutils.PatchRequest(ctx, c.client, reqURL, request)
		if err != nil {
			errMessage := "error while updating faucet"
			log.FromContext(ctx).WithError(err).Error(errMessage)
			return errors.ServiceConnectionError(errMessage).ExtendComponent(component)
		}

		defer clientutils.CloseResponse(response)
		return httputil.ParseResponse(ctx, response, resp)
	})

	return resp, err
}

func (c *HTTPClient) DeleteFaucet(ctx context.Context, uuid string) error {
	reqURL := fmt.Sprintf("%v/faucets/%v", c.config.URL, uuid)

	response, err := clientutils.DeleteRequest(ctx, c.client, reqURL)
	if err != nil {
		errMessage := "error while deleting faucet"
		log.FromContext(ctx).WithError(err).Error(errMessage)
		return errors.ServiceConnectionError(errMessage).ExtendComponent(component)
	}

	defer clientutils.CloseResponse(response)
	return httputil.ParseEmptyBodyResponse(ctx, response)
}