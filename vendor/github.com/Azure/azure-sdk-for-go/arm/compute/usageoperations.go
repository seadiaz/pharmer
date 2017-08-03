package compute

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
// Code generated by Microsoft (R) AutoRest Code Generator 0.17.0.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"net/http"
)

// UsageOperationsClient is the the Compute Management Client.
type UsageOperationsClient struct {
	ManagementClient
}

// NewUsageOperationsClient creates an instance of the UsageOperationsClient
// client.
func NewUsageOperationsClient(subscriptionID string) UsageOperationsClient {
	return NewUsageOperationsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewUsageOperationsClientWithBaseURI creates an instance of the
// UsageOperationsClient client.
func NewUsageOperationsClientWithBaseURI(baseURI string, subscriptionID string) UsageOperationsClient {
	return UsageOperationsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// List gets, for the specified location, the current compute resource usage
// information as well as the limits for compute resources under the
// subscription.
//
// location is the location for which resource usage is queried.
func (client UsageOperationsClient) List(location string) (result ListUsagesResult, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: location,
			Constraints: []validation.Constraint{{Target: "location", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "compute.UsageOperationsClient", "List")
	}

	req, err := client.ListPreparer(location)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "compute.UsageOperationsClient", "List", nil, "Failure preparing request")
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "compute.UsageOperationsClient", "List", resp, "Failure sending request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.UsageOperationsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client UsageOperationsClient) ListPreparer(location string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"location":       autorest.Encode("path", location),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/usages", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client UsageOperationsClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client UsageOperationsClient) ListResponder(resp *http.Response) (result ListUsagesResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListNextResults retrieves the next set of results, if any.
func (client UsageOperationsClient) ListNextResults(lastResults ListUsagesResult) (result ListUsagesResult, err error) {
	req, err := lastResults.ListUsagesResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "compute.UsageOperationsClient", "List", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "compute.UsageOperationsClient", "List", resp, "Failure sending next results request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.UsageOperationsClient", "List", resp, "Failure responding to next results request")
	}

	return
}
