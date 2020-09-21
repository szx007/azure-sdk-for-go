package streamanalytics

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// OutputsClient is the stream Analytics Client
type OutputsClient struct {
	BaseClient
}

// NewOutputsClient creates an instance of the OutputsClient client.
func NewOutputsClient(subscriptionID string) OutputsClient {
	return NewOutputsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewOutputsClientWithBaseURI creates an instance of the OutputsClient client using a custom endpoint.  Use this when
// interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewOutputsClientWithBaseURI(baseURI string, subscriptionID string) OutputsClient {
	return OutputsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrReplace creates an output or replaces an already existing output under an existing streaming job.
// Parameters:
// output - the definition of the output that will be used to create a new output or replace the existing one
// under the streaming job.
// resourceGroupName - the name of the resource group. The name is case insensitive.
// jobName - the name of the streaming job.
// outputName - the name of the output.
// ifMatch - the ETag of the output. Omit this value to always overwrite the current output. Specify the
// last-seen ETag value to prevent accidentally overwriting concurrent changes.
// ifNoneMatch - set to '*' to allow a new output to be created, but to prevent updating an existing output.
// Other values will result in a 412 Pre-condition Failed response.
func (client OutputsClient) CreateOrReplace(ctx context.Context, output Output, resourceGroupName string, jobName string, outputName string, ifMatch string, ifNoneMatch string) (result Output, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OutputsClient.CreateOrReplace")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("streamanalytics.OutputsClient", "CreateOrReplace", err.Error())
	}

	req, err := client.CreateOrReplacePreparer(ctx, output, resourceGroupName, jobName, outputName, ifMatch, ifNoneMatch)
	if err != nil {
		err = autorest.NewErrorWithError(err, "streamanalytics.OutputsClient", "CreateOrReplace", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrReplaceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "streamanalytics.OutputsClient", "CreateOrReplace", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrReplaceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "streamanalytics.OutputsClient", "CreateOrReplace", resp, "Failure responding to request")
	}

	return
}

// CreateOrReplacePreparer prepares the CreateOrReplace request.
func (client OutputsClient) CreateOrReplacePreparer(ctx context.Context, output Output, resourceGroupName string, jobName string, outputName string, ifMatch string, ifNoneMatch string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"jobName":           autorest.Encode("path", jobName),
		"outputName":        autorest.Encode("path", outputName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-04-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}/outputs/{outputName}", pathParameters),
		autorest.WithJSON(output),
		autorest.WithQueryParameters(queryParameters))
	if len(ifMatch) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("If-Match", autorest.String(ifMatch)))
	}
	if len(ifNoneMatch) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("If-None-Match", autorest.String(ifNoneMatch)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrReplaceSender sends the CreateOrReplace request. The method will close the
// http.Response Body if it receives an error.
func (client OutputsClient) CreateOrReplaceSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateOrReplaceResponder handles the response to the CreateOrReplace request. The method always
// closes the http.Response Body.
func (client OutputsClient) CreateOrReplaceResponder(resp *http.Response) (result Output, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes an output from the streaming job.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// jobName - the name of the streaming job.
// outputName - the name of the output.
func (client OutputsClient) Delete(ctx context.Context, resourceGroupName string, jobName string, outputName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OutputsClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("streamanalytics.OutputsClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, jobName, outputName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "streamanalytics.OutputsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "streamanalytics.OutputsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "streamanalytics.OutputsClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client OutputsClient) DeletePreparer(ctx context.Context, resourceGroupName string, jobName string, outputName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"jobName":           autorest.Encode("path", jobName),
		"outputName":        autorest.Encode("path", outputName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-04-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}/outputs/{outputName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client OutputsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client OutputsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets details about the specified output.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// jobName - the name of the streaming job.
// outputName - the name of the output.
func (client OutputsClient) Get(ctx context.Context, resourceGroupName string, jobName string, outputName string) (result Output, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OutputsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("streamanalytics.OutputsClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, jobName, outputName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "streamanalytics.OutputsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "streamanalytics.OutputsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "streamanalytics.OutputsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client OutputsClient) GetPreparer(ctx context.Context, resourceGroupName string, jobName string, outputName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"jobName":           autorest.Encode("path", jobName),
		"outputName":        autorest.Encode("path", outputName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-04-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}/outputs/{outputName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client OutputsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client OutputsClient) GetResponder(resp *http.Response) (result Output, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByStreamingJob lists all of the outputs under the specified streaming job.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// jobName - the name of the streaming job.
// selectParameter - the $select OData query parameter. This is a comma-separated list of structural properties
// to include in the response, or "*" to include all properties. By default, all properties are returned except
// diagnostics. Currently only accepts '*' as a valid value.
func (client OutputsClient) ListByStreamingJob(ctx context.Context, resourceGroupName string, jobName string, selectParameter string) (result OutputListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OutputsClient.ListByStreamingJob")
		defer func() {
			sc := -1
			if result.olr.Response.Response != nil {
				sc = result.olr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("streamanalytics.OutputsClient", "ListByStreamingJob", err.Error())
	}

	result.fn = client.listByStreamingJobNextResults
	req, err := client.ListByStreamingJobPreparer(ctx, resourceGroupName, jobName, selectParameter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "streamanalytics.OutputsClient", "ListByStreamingJob", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByStreamingJobSender(req)
	if err != nil {
		result.olr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "streamanalytics.OutputsClient", "ListByStreamingJob", resp, "Failure sending request")
		return
	}

	result.olr, err = client.ListByStreamingJobResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "streamanalytics.OutputsClient", "ListByStreamingJob", resp, "Failure responding to request")
	}
	if result.olr.hasNextLink() && result.olr.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListByStreamingJobPreparer prepares the ListByStreamingJob request.
func (client OutputsClient) ListByStreamingJobPreparer(ctx context.Context, resourceGroupName string, jobName string, selectParameter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"jobName":           autorest.Encode("path", jobName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-04-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(selectParameter) > 0 {
		queryParameters["$select"] = autorest.Encode("query", selectParameter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}/outputs", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByStreamingJobSender sends the ListByStreamingJob request. The method will close the
// http.Response Body if it receives an error.
func (client OutputsClient) ListByStreamingJobSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByStreamingJobResponder handles the response to the ListByStreamingJob request. The method always
// closes the http.Response Body.
func (client OutputsClient) ListByStreamingJobResponder(resp *http.Response) (result OutputListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByStreamingJobNextResults retrieves the next set of results, if any.
func (client OutputsClient) listByStreamingJobNextResults(ctx context.Context, lastResults OutputListResult) (result OutputListResult, err error) {
	req, err := lastResults.outputListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "streamanalytics.OutputsClient", "listByStreamingJobNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByStreamingJobSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "streamanalytics.OutputsClient", "listByStreamingJobNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByStreamingJobResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "streamanalytics.OutputsClient", "listByStreamingJobNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByStreamingJobComplete enumerates all values, automatically crossing page boundaries as required.
func (client OutputsClient) ListByStreamingJobComplete(ctx context.Context, resourceGroupName string, jobName string, selectParameter string) (result OutputListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OutputsClient.ListByStreamingJob")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByStreamingJob(ctx, resourceGroupName, jobName, selectParameter)
	return
}

// Test tests whether an output’s datasource is reachable and usable by the Azure Stream Analytics service.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// jobName - the name of the streaming job.
// outputName - the name of the output.
// output - if the output specified does not already exist, this parameter must contain the full output
// definition intended to be tested. If the output specified already exists, this parameter can be left null to
// test the existing output as is or if specified, the properties specified will overwrite the corresponding
// properties in the existing output (exactly like a PATCH operation) and the resulting output will be tested.
func (client OutputsClient) Test(ctx context.Context, resourceGroupName string, jobName string, outputName string, output *Output) (result OutputsTestFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OutputsClient.Test")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("streamanalytics.OutputsClient", "Test", err.Error())
	}

	req, err := client.TestPreparer(ctx, resourceGroupName, jobName, outputName, output)
	if err != nil {
		err = autorest.NewErrorWithError(err, "streamanalytics.OutputsClient", "Test", nil, "Failure preparing request")
		return
	}

	result, err = client.TestSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "streamanalytics.OutputsClient", "Test", result.Response(), "Failure sending request")
		return
	}

	return
}

// TestPreparer prepares the Test request.
func (client OutputsClient) TestPreparer(ctx context.Context, resourceGroupName string, jobName string, outputName string, output *Output) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"jobName":           autorest.Encode("path", jobName),
		"outputName":        autorest.Encode("path", outputName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-04-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}/outputs/{outputName}/test", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if output != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithJSON(output))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// TestSender sends the Test request. The method will close the
// http.Response Body if it receives an error.
func (client OutputsClient) TestSender(req *http.Request) (future OutputsTestFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// TestResponder handles the response to the Test request. The method always
// closes the http.Response Body.
func (client OutputsClient) TestResponder(resp *http.Response) (result ResourceTestStatus, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Update updates an existing output under an existing streaming job. This can be used to partially update (ie. update
// one or two properties) an output without affecting the rest the job or output definition.
// Parameters:
// output - an Output object. The properties specified here will overwrite the corresponding properties in the
// existing output (ie. Those properties will be updated). Any properties that are set to null here will mean
// that the corresponding property in the existing output will remain the same and not change as a result of
// this PATCH operation.
// resourceGroupName - the name of the resource group. The name is case insensitive.
// jobName - the name of the streaming job.
// outputName - the name of the output.
// ifMatch - the ETag of the output. Omit this value to always overwrite the current output. Specify the
// last-seen ETag value to prevent accidentally overwriting concurrent changes.
func (client OutputsClient) Update(ctx context.Context, output Output, resourceGroupName string, jobName string, outputName string, ifMatch string) (result Output, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OutputsClient.Update")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("streamanalytics.OutputsClient", "Update", err.Error())
	}

	req, err := client.UpdatePreparer(ctx, output, resourceGroupName, jobName, outputName, ifMatch)
	if err != nil {
		err = autorest.NewErrorWithError(err, "streamanalytics.OutputsClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "streamanalytics.OutputsClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "streamanalytics.OutputsClient", "Update", resp, "Failure responding to request")
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client OutputsClient) UpdatePreparer(ctx context.Context, output Output, resourceGroupName string, jobName string, outputName string, ifMatch string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"jobName":           autorest.Encode("path", jobName),
		"outputName":        autorest.Encode("path", outputName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-04-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}/outputs/{outputName}", pathParameters),
		autorest.WithJSON(output),
		autorest.WithQueryParameters(queryParameters))
	if len(ifMatch) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("If-Match", autorest.String(ifMatch)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client OutputsClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client OutputsClient) UpdateResponder(resp *http.Response) (result Output, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
