package client

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/NdoleStudio/lemonsqueezy-go/internal/helpers"
	"github.com/NdoleStudio/lemonsqueezy-go/internal/stubs"
	"github.com/stretchr/testify/assert"
)

func TestSubscriptionsService_Get(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusOK, stubs.SubscriptionGetResponse())
	client := New(WithBaseURL(server.URL))

	// Act
	subscription, response, err := client.Subscriptions.Get(context.Background(), "1")

	// Assert
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)
	assert.Equal(t, stubs.SubscriptionGetResponse(), *response.Body)

	assert.Equal(t, &ApiResponse[Subscription]{
		Jsonapi: ApiResponseJSONAPI{
			Version: "1.0",
		},
		Links: ApiResponseLink{
			Self: "https://api.lemonsqueezy.com/v1/subscriptions/1",
		},
		Data: ApiResponseData[Subscription]{
			Type: "subscriptions",
			ID:   "1",
			Attributes: Subscription{
				StoreID:         1,
				OrderID:         1,
				OrderItemID:     1,
				ProductID:       1,
				VariantID:       1,
				ProductName:     "Example Product",
				VariantName:     "Example Variant",
				UserName:        "Darlene Daugherty",
				UserEmail:       "gernser@yahoo.com",
				Status:          "active",
				StatusFormatted: "Active",
				Pause:           nil,
				Cancelled:       false,
				TrialEndsAt:     nil,
				BillingAnchor:   12,
				Urls: SubscriptionURLs{
					UpdatePaymentMethod: "https://app.lemonsqueezy.com/my-orders/2ba92a4e-a00a-45d2-a128-16856ffa8cdf/subscription/8/update-payment-method?expires=1666869343&signature=9985e3bf9007840aeb3951412be475abc17439c449c1af3e56e08e45e1345413",
				},
				RenewsAt:  time.Date(2022, time.November, 12, 0, 0, 0, 0, time.UTC),
				EndsAt:    nil,
				CreatedAt: time.Date(2021, time.August, 11, 13, 47, 27, 0, time.UTC),
				UpdatedAt: time.Date(2021, time.August, 11, 13, 54, 19, 0, time.UTC),
				TestMode:  false,
			},
			Relationships: ApiResponseRelationships{
				Store: ApiResponseLinks{
					Links: ApiResponseLink{
						Related: "https://api.lemonsqueezy.com/v1/subscriptions/1/store",
						Self:    "https://api.lemonsqueezy.com/v1/subscriptions/1/relationships/store",
					},
				},
				Order: ApiResponseLinks{
					Links: ApiResponseLink{
						Related: "https://api.lemonsqueezy.com/v1/subscriptions/1/order",
						Self:    "https://api.lemonsqueezy.com/v1/subscriptions/1/relationships/order",
					},
				},
				OrderItem: ApiResponseLinks{
					Links: ApiResponseLink{
						Related: "https://api.lemonsqueezy.com/v1/subscriptions/1/order-item",
						Self:    "https://api.lemonsqueezy.com/v1/subscriptions/1/relationships/order-item",
					},
				},
				Product: ApiResponseLinks{
					Links: ApiResponseLink{
						Related: "https://api.lemonsqueezy.com/v1/subscriptions/1/product",
						Self:    "https://api.lemonsqueezy.com/v1/subscriptions/1/relationships/product",
					},
				},
				Variant: ApiResponseLinks{
					Links: ApiResponseLink{
						Related: "https://api.lemonsqueezy.com/v1/subscriptions/1/variant",
						Self:    "https://api.lemonsqueezy.com/v1/subscriptions/1/relationships/variant",
					},
				},
			},
			Links: ApiResponseLink{
				Related: "",
				Self:    "https://api.lemonsqueezy.com/v1/subscriptions/1",
			},
		},
	}, subscription)

	// Teardown
	server.Close()
}

func TestSubscriptionsService_GetWithError(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusInternalServerError, nil)
	client := New(WithBaseURL(server.URL))

	// Act
	_, response, err := client.Subscriptions.Get(context.Background(), "1")

	// Assert
	assert.NotNil(t, err)

	assert.Equal(t, http.StatusInternalServerError, response.HTTPResponse.StatusCode)

	// Teardown
	server.Close()
}

func TestSubscriptionsService_Cancel(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusOK, stubs.SubscriptionCancelResponse())
	client := New(WithBaseURL(server.URL))

	// Act
	subscription, response, err := client.Subscriptions.Cancel(context.Background(), "1")

	// Assert
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)
	assert.Equal(t, stubs.SubscriptionCancelResponse(), *response.Body)

	assert.Equal(t, &ApiResponse[Subscription]{
		Jsonapi: ApiResponseJSONAPI{
			Version: "1.0",
		},
		Links: ApiResponseLink{
			Self: "https://api.lemonsqueezy.com/v1/subscriptions/1",
		},
		Data: ApiResponseData[Subscription]{
			Type: "subscriptions",
			ID:   "1",
			Attributes: Subscription{
				StoreID:         1,
				OrderID:         1,
				OrderItemID:     1,
				ProductID:       1,
				VariantID:       1,
				ProductName:     "Example Product",
				VariantName:     "Example Variant",
				UserName:        "Darlene Daugherty",
				UserEmail:       "gernser@yahoo.com",
				Status:          "cancelled",
				StatusFormatted: "Cancelled",
				Pause:           nil,
				Cancelled:       true,
				TrialEndsAt:     nil,
				BillingAnchor:   12,
				Urls: SubscriptionURLs{
					UpdatePaymentMethod: "https://app.lemonsqueezy.com/my-orders/2ba92a4e-a00a-45d2-a128-16856ffa8cdf/subscription/8/update-payment-method?expires=1666869343&signature=9985e3bf9007840aeb3951412be475abc17439c449c1af3e56e08e45e1345413",
				},
				RenewsAt: time.Date(2022, time.November, 12, 0, 0, 0, 0, time.UTC),
				EndsAt: (func() *time.Time {
					val := time.Date(2022, time.November, 12, 0, 0, 0, 0, time.UTC)
					return &val
				})(),
				CreatedAt: time.Date(2021, time.August, 11, 13, 47, 27, 0, time.UTC),
				UpdatedAt: time.Date(2021, time.August, 11, 13, 54, 19, 0, time.UTC),
				TestMode:  false,
			},
			Relationships: ApiResponseRelationships{
				Store: ApiResponseLinks{
					Links: ApiResponseLink{
						Related: "https://api.lemonsqueezy.com/v1/subscriptions/1/store",
						Self:    "https://api.lemonsqueezy.com/v1/subscriptions/1/relationships/store",
					},
				},
				Order: ApiResponseLinks{
					Links: ApiResponseLink{
						Related: "https://api.lemonsqueezy.com/v1/subscriptions/1/order",
						Self:    "https://api.lemonsqueezy.com/v1/subscriptions/1/relationships/order",
					},
				},
				OrderItem: ApiResponseLinks{
					Links: ApiResponseLink{
						Related: "https://api.lemonsqueezy.com/v1/subscriptions/1/order-item",
						Self:    "https://api.lemonsqueezy.com/v1/subscriptions/1/relationships/order-item",
					},
				},
				Product: ApiResponseLinks{
					Links: ApiResponseLink{
						Related: "https://api.lemonsqueezy.com/v1/subscriptions/1/product",
						Self:    "https://api.lemonsqueezy.com/v1/subscriptions/1/relationships/product",
					},
				},
				Variant: ApiResponseLinks{
					Links: ApiResponseLink{
						Related: "https://api.lemonsqueezy.com/v1/subscriptions/1/variant",
						Self:    "https://api.lemonsqueezy.com/v1/subscriptions/1/relationships/variant",
					},
				},
			},
			Links: ApiResponseLink{
				Related: "",
				Self:    "https://api.lemonsqueezy.com/v1/subscriptions/1",
			},
		},
	}, subscription)

	// Teardown
	server.Close()
}

func TestSubscriptionsService_CancelWithError(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusInternalServerError, nil)
	client := New(WithBaseURL(server.URL))

	// Act
	_, response, err := client.Subscriptions.Cancel(context.Background(), "1")

	// Assert
	assert.NotNil(t, err)

	assert.Equal(t, http.StatusInternalServerError, response.HTTPResponse.StatusCode)

	// Teardown
	server.Close()
}