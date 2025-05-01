package service

import (
	"context"
	"os"
	"strconv"

	"github.com/pkg/errors"
	"github.com/plutov/paypal/v4"
)

// PaymentProvider 定义支付提供商的接口
type PaymentProvider interface {
	CreateOrder(ctx context.Context, amount int, currency string) (string, error)
	CaptureOrder(ctx context.Context, orderID string) error
}

// PayPalProvider 实现 PayPal 支付
type PayPalProvider struct {
	client *paypal.Client
}

func NewPayPalProvider() (*PayPalProvider, error) {
	apiBase := paypal.APIBaseSandBox
	if os.Getenv("APP_ENV") == "production" {
		apiBase = paypal.APIBaseLive
	}

	client, err := paypal.NewClient(
		os.Getenv("PAYPAL_CLIENT_ID"),
		os.Getenv("PAYPAL_CLIENT_SECRET"),
		apiBase,
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create PayPal client")
	}

	_, err = client.GetAccessToken(context.Background())
	if err != nil {
		return nil, errors.Wrap(err, "failed to get PayPal access token")
	}

	return &PayPalProvider{
		client: client,
	}, nil
}

func (p *PayPalProvider) CreateOrder(ctx context.Context, amount int, currency string) (string, error) {
	priceInDecimal := float64(amount) / 100.0
	priceStr := strconv.FormatFloat(priceInDecimal, 'f', 2, 64)

	payPalOrder, err := p.client.CreateOrder(
		ctx,
		paypal.OrderIntentCapture,
		[]paypal.PurchaseUnitRequest{
			{
				Amount: &paypal.PurchaseUnitAmount{
					Currency: currency,
					Value:    priceStr,
				},
			},
		},
		nil,
		nil,
	)
	if err != nil {
		return "", errors.Wrap(err, "failed to create PayPal order")
	}

	return payPalOrder.ID, nil
}

func (p *PayPalProvider) CaptureOrder(ctx context.Context, orderID string) error {
	_, err := p.client.CaptureOrder(ctx, orderID, paypal.CaptureOrderRequest{})
	if err != nil {
		return errors.Wrap(err, "failed to capture PayPal order")
	}
	return nil
}
