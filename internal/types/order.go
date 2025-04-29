package types

import "time"

// OrderProduct 代表一个订单中的商品
type OrderProduct struct {
	ID          string `json:"id"`       // 商品ID
	ProductName string `json:"name"`     // 商品名称
	Quantity    int    `json:"quantity"` // 商品数量
	Price       int64  `json:"price"`    // 商品单价，单位分
}

// Order 代表订单信息
type Order struct {
	ID          string         `json:"id"`           // 订单ID
	UserID      string         `json:"user_id"`      // 用户ID
	Status      string         `json:"status"`       // 订单状态
	Currency    string         `json:"currency"`     // 货币类型
	Products    []OrderProduct `json:"products"`     // 订单包含的商品
	TotalAmount int64          `json:"total_amount"` // 总金额，单位分
	PaidAmount  int64          `json:"paid_amount"`  // 已支付金额，单位分
	Description string         `json:"description"`  // 订单描述
	Created     int64          `json:"created"`      // 创建时间（时间戳）
	Updated     int64          `json:"updated"`      // 更新时间（时间戳）
}

// CreateOrderRequest 创建订单请求
type CreateOrderRequest struct {
	Currency    string         `json:"currency" binding:"oneof=CNY USD"` // 货币类型
	Products    []OrderProduct `json:"products" binding:"required"`      // 订单商品
	TotalAmount int64          `json:"total_amount"`                     // 总金额，单位分
	Description string         `json:"description"`                      // 订单描述
}

// UpdateOrderRequest 更新订单请求
type UpdateOrderRequest struct {
	Currency    string         `json:"currency" binding:"oneof=CNY USD"` // 货币类型
	Products    []OrderProduct `json:"products"`                         // 订单商品
	TotalAmount int64          `json:"total_amount"`                     // 总金额，单位分
	Description string         `json:"description"`                      // 订单描述
}

// Helper method to convert totalAmount (in cents) to yuan (decimal)
func (o *Order) TotalAmountInYuan() float64 {
	return float64(o.TotalAmount) / 100
}

// Helper method to convert paidAmount (in cents) to yuan (decimal)
func (o *Order) PaidAmountInYuan() float64 {
	return float64(o.PaidAmount) / 100
}

// Helper method to format time as a readable string
func (o *Order) FormattedCreated() string {
	return time.Unix(o.Created, 0).Format("2006-01-02 15:04:05")
}

// Helper method to format time as a readable string
func (o *Order) FormattedUpdated() string {
	return time.Unix(o.Updated, 0).Format("2006-01-02 15:04:05")
}
