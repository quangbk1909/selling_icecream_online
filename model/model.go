package model

type User struct {
	ID          int    `json:"id"`
	FullName    string `json:"full_name"`
	PhoneNumber string `json:"phone_number"`
	Token       string `json:"token"`
	Address     string `json:"address"`
	Password    string `json:"password"`
	VinidPoint  int    `json:"vinid_point"`
	CreatedAt   string `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
}

type Store struct {
	ID            int            `json:"id"`
	Name          string         `json:"name"`
	Address       string         `json:"address"`
	ImagePath     string         `json:"image_path"`
	Latitude      float64        `json:"latitude"`
	Longitude     float64        `json:"longitude"`
	CreatedAt     string         `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	IceCreamItems []IceCreamItem `gorm:"many2many:item_store"`
}

type Rating struct {
	ID           int          `json:"id"`
	RatingStar   int          `json:"rating_start"`
	Comment      string       `json:"comment"`
	IteamID      int          `json:"item_id"`
	UserID       int          `json:"user_id"`
	CreatedAt    string       `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	User         User         `gorm:"foreignkey:UserID"`
	IceCreamItem IceCreamItem `gorm:"foreignkey:ItemID"`
}

type Order struct {
	ID            int            `json:"id"`
	UserID        int            `json:"user_id"`
	Status        int            `json:"status"`
	TotalFee      int            `json:"total_fee"`
	CreatedAt     string         `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	IceCreamItems []IceCreamItem `gorm:"many2many:order_item"`
}

type IceCreamItem struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Type      string  `json:"type"`
	ImagePath string  `json:"image_path"`
	Price     int     `json:"price"`
	CreatedAt string  `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	Stores    []Store `gorm:"many2many:item_store"`
}

type OrderItem struct {
	ID             int `json:"id"`
	OrderID        int `json:"order_id"`
	IceCreamItemId int `json:"ice_cream_item_id"`
	Quantity       int `json:"quantity"`
}

func (OrderItem) TableName() string {
	return "order_item"
}

type ItemStore struct {
	ID             int `json:"id"`
	IceCreamItemID int `json:"ice_cream_item_id"`
	StoreID        int `json:"store_id"`
	Status         int `json:"status"`
}

func (ItemStore) TableName() string {
	return "item_store"
}

type MetaDataResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ResponseForm struct {
	Data interface{}      `json:"data"`
	Meta MetaDataResponse `json:"meta"`
}

type ItemInOrder struct {
	ItemInfo IceCreamItem `json:"item_info"`
	Quantity int          `json:"quantity"`
}

type OrderDetail struct {
	OrderInfo Order         `json:"order_info"`
	Items     []ItemInOrder `json:"items"`
}

type ItemOrderJson struct {
	ItemID   int `json:"item_id"`
	Quantity int `json:"quantity"`
}

type OrderJson struct {
	UserID   int             `json:"user_id"`
	TotalFee int             `json:"total_fee"`
	Items    []ItemOrderJson `json:"items"`
}
