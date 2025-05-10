package service

import (
	"github.com/Axpz/store/internal/config"
	"github.com/Axpz/store/internal/payment"
	"github.com/gin-gonic/gin"
)

type PaymentService struct {
	provider payment.PaymentProvider
}

func NewPaymentService(cfg *config.Config) (*PaymentService, error) {
	// switch cfg.PayProvider {
	// case "paypal":
	// 	provider, err = NewPayPalProvider(cfg)
	// default:
	// 	return nil, errors.New("unsupported payment provider")
	// }

	provider, err := payment.NewPayPalProvider(cfg)

	if err != nil {
		return nil, err
	}

	return &PaymentService{
		provider: provider,
	}, nil
}

func (s *PaymentService) CreateOrder(c *gin.Context, amount int, currency string) (string, error) {
	return s.provider.CreateOrder(c, amount, currency)
}

func (s *PaymentService) CaptureOrder(c *gin.Context, orderID string) error {
	return s.provider.CaptureOrder(c, orderID)
}

func (s *PaymentService) HandleWebhook(c *gin.Context, rawBody []byte, headers map[string]string) error {
	return s.provider.HandleWebhook(c)
}
