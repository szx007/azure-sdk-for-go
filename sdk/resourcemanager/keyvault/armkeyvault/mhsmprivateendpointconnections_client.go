//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armkeyvault

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// MHSMPrivateEndpointConnectionsClient contains the methods for the MHSMPrivateEndpointConnections group.
// Don't use this type directly, use NewMHSMPrivateEndpointConnectionsClient() instead.
type MHSMPrivateEndpointConnectionsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewMHSMPrivateEndpointConnectionsClient creates a new instance of MHSMPrivateEndpointConnectionsClient with the specified values.
//   - subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
//     part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewMHSMPrivateEndpointConnectionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*MHSMPrivateEndpointConnectionsClient, error) {
	cl, err := arm.NewClient(moduleName+".MHSMPrivateEndpointConnectionsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &MHSMPrivateEndpointConnectionsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginDelete - Deletes the specified private endpoint connection associated with the managed hsm pool.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-10-01
//   - resourceGroupName - Name of the resource group that contains the managed HSM pool.
//   - name - Name of the managed HSM Pool
//   - privateEndpointConnectionName - Name of the private endpoint connection associated with the managed hsm pool.
//   - options - MHSMPrivateEndpointConnectionsClientBeginDeleteOptions contains the optional parameters for the MHSMPrivateEndpointConnectionsClient.BeginDelete
//     method.
func (client *MHSMPrivateEndpointConnectionsClient) BeginDelete(ctx context.Context, resourceGroupName string, name string, privateEndpointConnectionName string, options *MHSMPrivateEndpointConnectionsClientBeginDeleteOptions) (*runtime.Poller[MHSMPrivateEndpointConnectionsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, name, privateEndpointConnectionName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[MHSMPrivateEndpointConnectionsClientDeleteResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[MHSMPrivateEndpointConnectionsClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Deletes the specified private endpoint connection associated with the managed hsm pool.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-10-01
func (client *MHSMPrivateEndpointConnectionsClient) deleteOperation(ctx context.Context, resourceGroupName string, name string, privateEndpointConnectionName string, options *MHSMPrivateEndpointConnectionsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, name, privateEndpointConnectionName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *MHSMPrivateEndpointConnectionsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, name string, privateEndpointConnectionName string, options *MHSMPrivateEndpointConnectionsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/managedHSMs/{name}/privateEndpointConnections/{privateEndpointConnectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if privateEndpointConnectionName == "" {
		return nil, errors.New("parameter privateEndpointConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateEndpointConnectionName}", url.PathEscape(privateEndpointConnectionName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the specified private endpoint connection associated with the managed HSM Pool.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-10-01
//   - resourceGroupName - Name of the resource group that contains the managed HSM pool.
//   - name - Name of the managed HSM Pool
//   - privateEndpointConnectionName - Name of the private endpoint connection associated with the managed hsm pool.
//   - options - MHSMPrivateEndpointConnectionsClientGetOptions contains the optional parameters for the MHSMPrivateEndpointConnectionsClient.Get
//     method.
func (client *MHSMPrivateEndpointConnectionsClient) Get(ctx context.Context, resourceGroupName string, name string, privateEndpointConnectionName string, options *MHSMPrivateEndpointConnectionsClientGetOptions) (MHSMPrivateEndpointConnectionsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, name, privateEndpointConnectionName, options)
	if err != nil {
		return MHSMPrivateEndpointConnectionsClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MHSMPrivateEndpointConnectionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return MHSMPrivateEndpointConnectionsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *MHSMPrivateEndpointConnectionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, name string, privateEndpointConnectionName string, options *MHSMPrivateEndpointConnectionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/managedHSMs/{name}/privateEndpointConnections/{privateEndpointConnectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if privateEndpointConnectionName == "" {
		return nil, errors.New("parameter privateEndpointConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateEndpointConnectionName}", url.PathEscape(privateEndpointConnectionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *MHSMPrivateEndpointConnectionsClient) getHandleResponse(resp *http.Response) (MHSMPrivateEndpointConnectionsClientGetResponse, error) {
	result := MHSMPrivateEndpointConnectionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MHSMPrivateEndpointConnection); err != nil {
		return MHSMPrivateEndpointConnectionsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourcePager - The List operation gets information about the private endpoint connections associated with the
// managed HSM Pool.
//
// Generated from API version 2021-10-01
//   - resourceGroupName - Name of the resource group that contains the managed HSM pool.
//   - name - Name of the managed HSM Pool
//   - options - MHSMPrivateEndpointConnectionsClientListByResourceOptions contains the optional parameters for the MHSMPrivateEndpointConnectionsClient.NewListByResourcePager
//     method.
func (client *MHSMPrivateEndpointConnectionsClient) NewListByResourcePager(resourceGroupName string, name string, options *MHSMPrivateEndpointConnectionsClientListByResourceOptions) *runtime.Pager[MHSMPrivateEndpointConnectionsClientListByResourceResponse] {
	return runtime.NewPager(runtime.PagingHandler[MHSMPrivateEndpointConnectionsClientListByResourceResponse]{
		More: func(page MHSMPrivateEndpointConnectionsClientListByResourceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *MHSMPrivateEndpointConnectionsClientListByResourceResponse) (MHSMPrivateEndpointConnectionsClientListByResourceResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceCreateRequest(ctx, resourceGroupName, name, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return MHSMPrivateEndpointConnectionsClientListByResourceResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return MHSMPrivateEndpointConnectionsClientListByResourceResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return MHSMPrivateEndpointConnectionsClientListByResourceResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceHandleResponse(resp)
		},
	})
}

// listByResourceCreateRequest creates the ListByResource request.
func (client *MHSMPrivateEndpointConnectionsClient) listByResourceCreateRequest(ctx context.Context, resourceGroupName string, name string, options *MHSMPrivateEndpointConnectionsClientListByResourceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/managedHSMs/{name}/privateEndpointConnections"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceHandleResponse handles the ListByResource response.
func (client *MHSMPrivateEndpointConnectionsClient) listByResourceHandleResponse(resp *http.Response) (MHSMPrivateEndpointConnectionsClientListByResourceResponse, error) {
	result := MHSMPrivateEndpointConnectionsClientListByResourceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MHSMPrivateEndpointConnectionsListResult); err != nil {
		return MHSMPrivateEndpointConnectionsClientListByResourceResponse{}, err
	}
	return result, nil
}

// Put - Updates the specified private endpoint connection associated with the managed hsm pool.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-10-01
//   - resourceGroupName - Name of the resource group that contains the managed HSM pool.
//   - name - Name of the managed HSM Pool
//   - privateEndpointConnectionName - Name of the private endpoint connection associated with the managed hsm pool.
//   - properties - The intended state of private endpoint connection.
//   - options - MHSMPrivateEndpointConnectionsClientPutOptions contains the optional parameters for the MHSMPrivateEndpointConnectionsClient.Put
//     method.
func (client *MHSMPrivateEndpointConnectionsClient) Put(ctx context.Context, resourceGroupName string, name string, privateEndpointConnectionName string, properties MHSMPrivateEndpointConnection, options *MHSMPrivateEndpointConnectionsClientPutOptions) (MHSMPrivateEndpointConnectionsClientPutResponse, error) {
	req, err := client.putCreateRequest(ctx, resourceGroupName, name, privateEndpointConnectionName, properties, options)
	if err != nil {
		return MHSMPrivateEndpointConnectionsClientPutResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MHSMPrivateEndpointConnectionsClientPutResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return MHSMPrivateEndpointConnectionsClientPutResponse{}, runtime.NewResponseError(resp)
	}
	return client.putHandleResponse(resp)
}

// putCreateRequest creates the Put request.
func (client *MHSMPrivateEndpointConnectionsClient) putCreateRequest(ctx context.Context, resourceGroupName string, name string, privateEndpointConnectionName string, properties MHSMPrivateEndpointConnection, options *MHSMPrivateEndpointConnectionsClientPutOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/managedHSMs/{name}/privateEndpointConnections/{privateEndpointConnectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if privateEndpointConnectionName == "" {
		return nil, errors.New("parameter privateEndpointConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateEndpointConnectionName}", url.PathEscape(privateEndpointConnectionName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, properties)
}

// putHandleResponse handles the Put response.
func (client *MHSMPrivateEndpointConnectionsClient) putHandleResponse(resp *http.Response) (MHSMPrivateEndpointConnectionsClientPutResponse, error) {
	result := MHSMPrivateEndpointConnectionsClientPutResponse{}
	if val := resp.Header.Get("Retry-After"); val != "" {
		retryAfter32, err := strconv.ParseInt(val, 10, 32)
		retryAfter := int32(retryAfter32)
		if err != nil {
			return MHSMPrivateEndpointConnectionsClientPutResponse{}, err
		}
		result.RetryAfter = &retryAfter
	}
	if val := resp.Header.Get("Azure-AsyncOperation"); val != "" {
		result.AzureAsyncOperation = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.MHSMPrivateEndpointConnection); err != nil {
		return MHSMPrivateEndpointConnectionsClientPutResponse{}, err
	}
	return result, nil
}
