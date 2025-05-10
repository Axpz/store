package payment

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/plutov/paypal/v4"
	"go.uber.org/zap"

	"github.com/Axpz/store/internal/config"
	"github.com/Axpz/store/internal/utils"
)

type PaymentProvider interface {
	CreateOrder(c *gin.Context, amount int, currency string) (string, error)
	CaptureOrder(c *gin.Context, orderID string) error
	HandleWebhook(c *gin.Context) error
}

type PayPalProvider struct {
	client *paypal.Client
	cfg    *config.Config
}

func NewPayPalProvider(cfg *config.Config) (PaymentProvider, error) {
	apiBase := paypal.APIBaseSandBox
	if cfg.PayPal.Environment == "production" {
		apiBase = paypal.APIBaseLive
	}

	client, err := paypal.NewClient(cfg.PayPal.ClientID, cfg.PayPal.ClientSecret, apiBase)
	if err != nil {
		return nil, errors.Wrap(err, "create PayPal client failed")
	}

	_, err = client.GetAccessToken(context.Background())
	if err != nil {
		return nil, errors.Wrap(err, "get PayPal access token failed")
	}

	return &PayPalProvider{
		client: client,
		cfg:    cfg,
	}, nil
}

func (p *PayPalProvider) CreateOrder(c *gin.Context, amount int, currency string) (string, error) {
	ctx := c.Request.Context()
	logger := utils.LoggerFromContext(ctx)

	priceStr := strconv.FormatFloat(float64(amount)/100, 'f', 2, 64)
	order, err := p.client.CreateOrder(ctx, paypal.OrderIntentCapture, []paypal.PurchaseUnitRequest{
		{
			Amount: &paypal.PurchaseUnitAmount{
				Currency: currency,
				Value:    priceStr,
			},
		},
	}, nil, nil)

	if err != nil {
		logger.Error("CreateOrder failed", zap.Error(err))
		return "", err
	}

	logger.Info("PayPal order created", zap.String("orderID", order.ID))
	return order.ID, nil
}

func (p *PayPalProvider) CaptureOrder(c *gin.Context, orderID string) error {
	ctx := c.Request.Context()
	logger := utils.LoggerFromContext(ctx)

	_, err := p.client.CaptureOrder(ctx, orderID, paypal.CaptureOrderRequest{})
	if err != nil {
		logger.Error("CaptureOrder failed", zap.Error(err))
		return err
	}

	logger.Info("Order captured", zap.String("orderID", orderID))
	return nil
}

func (p *PayPalProvider) HandleWebhook(c *gin.Context) error {
	ctx := c.Request.Context()
	logger := utils.LoggerFromContext(ctx)

	// 验签
	resp, err := p.client.VerifyWebhookSignature(ctx, c.Request, p.cfg.PayPal.WebhookID)
	if err != nil || resp.VerificationStatus != "SUCCESS" {
		logger.Error("Webhook verification failed", zap.Error(err))
		return errors.Wrap(err, "webhook verify failed")
	}

	var event paypal.Event
	// if err := json.Unmarshal(rawBody, &event); err != nil {
	// 	logger.Error("Webhook JSON unmarshal failed", zap.Error(err))
	// 	return err
	// }

	// 处理 webhook 类型
	switch event.EventType {
	case "PAYMENT.CAPTURE.COMPLETED":
		// logger.Info("Payment completed", zap.Any("resource", event.Resource))
	default:
		logger.Info("Unhandled event", zap.String("event_type", event.EventType))
	}

	return nil
}
