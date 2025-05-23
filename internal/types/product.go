package types

type Product struct {
	ID          string   `json:"id"`          // 商品ID
	Name        string   `json:"name"`        // 商品名称
	Type        string   `json:"type"`        // 商品类型
	Image       string   `json:"image"`       // 商品图片
	Description string   `json:"description"` // 商品描述
	Price       int64    `json:"price"`       // 商品价格（单位分，比如1999表示¥19.99）
	Currency    string   `json:"currency"`    // 货币类型，比如 CNY
	Status      string   `json:"status"`      // 商品状态
	Content     []string `json:"content"`     // 商品内容
	Created     int64    `json:"created"`     // 创建时间
	Updated     int64    `json:"updated"`     // 更新时间
}
