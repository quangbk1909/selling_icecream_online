package model

const SecretKey string = "saomabietduoc"

//Chiếu đến bảng user trong database
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

// chiếu đến bảng store trong database
type Store struct {
	ID            int            `json:"id"`
	Name          string         `json:"name"`
	Address       string         `json:"address"`
	ImagePath     string         `json:"image_path"`
	Latitude      float64        `json:"latitude"`
	Longitude     float64        `json:"longitude"`
	CreatedAt     string         `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	IceCreamItems []IceCreamItem `json:"ice_cream_items" gorm:"many2many:item_store"`
}

// chiếu đến bảng rating trong database
type Rating struct {
	ID         int     `json:"id"`
	RatingStar float64 `json:"rating_star"`
	Comment    string  `json:"comment"`
	ItemID     int     `json:"item_id"`
	UserID     int     `json:"user_id"`
	CreatedAt  string  `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
}

// chiếu đến bảng order trong database
type Order struct {
	ID            int            `json:"id"`
	UserID        int            `json:"user_id"`
	Status        int            `json:"status"`
	TotalFee      int            `json:"total_fee"`
	CreatedAt     string         `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	IceCreamItems []IceCreamItem `json:"ice_cream_items" gorm:"many2many:order_item"`
}

// chiếu đến bảng ice_cream_item trong database
type IceCreamItem struct {
	ID         int                      `json:"id"`
	Name       string                   `json:"name"`
	Type       string                   `json:"type"`
	ImagePaths []string                 `json:"image_paths"`
	Price      int                      `json:"price"`
	CreatedAt  string                   `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	Stores     []Store                  `json:"stores" gorm:"many2many:item_store"`
	Ratings    []map[string]interface{} `json:"ratings"`
}

// chiếu đến bảng order_item
type OrderItem struct {
	ID             int `json:"id"`
	OrderID        int `json:"order_id"`
	IceCreamItemId int `json:"ice_cream_item_id"`
	Quantity       int `json:"quantity"`
}

func (OrderItem) TableName() string {
	return "order_item"
}

// chiếu đến bảng item_store
type ItemStore struct {
	ID             int `json:"id"`
	IceCreamItemID int `json:"ice_cream_item_id"`
	StoreID        int `json:"store_id"`
	Status         int `json:"status"`
}

func (ItemStore) TableName() string {
	return "item_store"
}

// chiếu đến bảng item_image: danh sách các ảnh ứng với các item
type ItemImage struct {
	ID             int    `json:"id"`
	IceCreamItemID int    `json:"ice_cream_item_id"`
	ImagePath      string `json:"image_path"`
}

// kiểu dữ liệu trả về của matadata trong response
type MetaDataResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// form data trả về của mỗi request
type ResponseForm struct {
	Data interface{}      `json:"data"`
	Meta MetaDataResponse `json:"meta"`
}

// dữ liệu của item, số lượng item trong order
type ItemInOrder struct {
	ItemInfo IceCreamItem `json:"item_info"`
	Quantity int          `json:"quantity"`
}

// chi tiết dữ liệu order trả về
type OrderDetail struct {
	OrderInfo Order         `json:"order_info"`
	Items     []ItemInOrder `json:"items"`
}

// dạng json dữ liệu item trong order nhận được từ client
type ItemOrderJson struct {
	ItemID   int `json:"item_id"`
	Quantity int `json:"quantity"`
}

// dạng json dữ liệu của 1 order nhận được từ client
type OrderJson struct {
	UserID   int             `json:"user_id"`
	TotalFee int             `json:"total_fee"`
	Items    []ItemOrderJson `json:"items"`
}

// json thông tin xác thực của người dùng, dùng cho register + login
type AuthenticationJson struct {
	PhoneNumber string `json:"phone_number" binding:"required"`
	Password    string `json:"password" binding:"required"`
}

// json thông tin người dùng muốn update
type UserInfoJson struct {
	FullName string `json:"full_name"`
	Address  string `json:"address"`
}
