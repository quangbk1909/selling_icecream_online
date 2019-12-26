package model

type User struct {
	ID          int    `json:"id"`
	FullName    string `json:"full_name"`
	PhoneNumber string `json:"phone_number"`
	Token       string `json:"token"`
	Address     string `json:"address"`
	Password    string `json:"password"`
	VinidPoint  int    `json:"vinid_point"`
	CreatedAt   string `json:"created_at"`
}

type Store struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Address   string  `json:"address"`
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
	CreatedAt string  `json:"created_at"`
}

type Rating struct {
	ID           int          `json:"id"`
	RatingStar   int          `json:"rating_start"`
	Comment      string       `json:"comment"`
	IteamID      int          `json:"item_id"`
	UserID       int          `json:"user_id"`
	CreatedAt    string       `json:"created_at"`
	User         User         `gorm:"foreignkey:UserID"`
	IceCreamItem IceCreamItem `gorm:"foreignkey:ItemID"`
}

type Order struct {
	ID            int            `json:"id"`
	UserID        int            `json:"user_id"`
	CreatedAt     string         `json:"created_at"`
	User          User           `gorm:"foreignkey:UserID"`
	IceCreamItems []IceCreamItem `gorm:"many2many:order_item"`
}

type IceCreamItem struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Type      string  `json:"type"`
	ImagePath string  `json:"image_path"`
	Price     int     `json:"price"`
	CreatedAt string  `json:"created_at"`
	Orders    []Order `gorm:"many2many:order_item"`
}

type OrderItem struct {
	ID       int `json:"id"`
	OrderID  int `json:"order_id"`
	ItemId   int `json:"item_id"`
	Quantity int `json:"quantity"`
}

func (OrderItem) TableName() string {
	return "order_item"
}
