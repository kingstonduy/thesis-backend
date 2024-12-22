package domain

type AddCommentRequest struct {
	Content string `json:"content"`
	Rating  int    `json:"rating"`
}

type AddCommentResponse struct {
	Success bool `json:"success"`
}

type CartItem struct {
	ProductID   string  `json:"product_id"`
	ProductName string  `json:"product_name"`
	ImageURL    string  `json:"image_url"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
}

type UpdateCartRequest struct {
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}

type UpdateCartResponse struct {
	Success bool `json:"success"`
}

type CheckoutResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type PurchasedProduct struct {
	OrderID        string `json:"order_id"`
	ProductName    string `json:"product_name"`
	ImageURL       string `json:"image_url"`
	DeliveryStatus string `json:"delivery_status"`
	PaymentStatus  string `json:"payment_status"`
}

type UserInformation struct {
	Username    string `json:"username"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Gender      string `json:"gender"`
	DateOfBirth string `json:"date_of_birth"`
	Street      string `json:"street"`
	City        string `json:"city"`
	District    string `json:"district"`
	Ward        string `json:"ward"`
}

type UpdateUserRequest struct {
	Username    string `json:"username"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Gender      string `json:"gender"`
	DateOfBirth string `json:"date_of_birth"`
	Street      string `json:"street"`
	City        string `json:"city"`
	District    string `json:"district"`
	Ward        string `json:"ward"`
}

type UpdateUserResponse struct {
	Success bool `json:"success"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token   string `json:"token"`
	Expires string `json:"expires"`
}

type ValidateTokenRequest struct {
	Token string `json:"token"`
}

type ValidateTokenResponse struct {
	IsValid bool `json:"is_valid"`
}
