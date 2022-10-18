# Architecture

- Transport Layer (Parse data from request/socket)
- Business Layer (Do some logic)
- Storage Layer (Integrate with DB)
  ![image](./public/images/go_service_architecture.png)
- Modules:
  - Model: Entity of the SQL
  - Storage: Working with DB (for itself)
  - Biz: Do logic (Only receive data, Do nothing with DB - No gin, header, v.v)
  - Transport: Parse data (Setup dependencies like redis, database)
  - Repository: Provide data business model for business layer (like mocking LikeRestaurant to Restaurant)

# Middleware

- Middleware functions are functions that have access to the request data, and the next middleware function in application's request-response cycle - Execute any code - Make changes to the request and the response objects - End the request-response cycle - Call the next middleware function in the stack
  ![image](./public/images/gin_middleware.png)

# Golang

- Tất cả các hàm có io nên đặt context
- type lowercase for first letter is not public and otherwise. Use function to return (`see in modules/**/storage`)
- Dùng injection:

```
// Not public
type createRestaurantBiz struct {
	store CreateRestaurantStore
}

func NewCreateRestaurantBiz(store CreateRestaurantStore) *createRestaurantBiz {
	return &createRestaurantBiz{store: store}
}
```

- Interface xài ở đâu nên khai báo ở đó
- Tạo app_context để sau này có thêm redis hay mongo sẽ không phải thêm parameters (trong transportation)
- tag `form` dùng khi muốn nhận được luôn giá trị trên query string
- `ShouldBind`: nạp data vào nếu có
- Embed struct:

```
type Restaurant struct {
	common.SQLModel `json:",inline"` // embed struct, json:",inline" flating the struct
	Name string `json:"name" gorm:"column:name;"`
	Addr string `json:"address" gorm:"column:addr;"`
}
```

- GORM Preload: Not like join

```
type User struct {
    gorm.Model
    Username string
    Orders   []Order
}

type Order struct {
    gorm.Model
    UserID   uint
    Price    float64
}

// Preload Orders when find users
db.Preload("Orders").Find(&users)
// SELECT * FROM users;
// SELECT * FROM orders WHERE user_id IN (1, 2, 3, 4)

db.Preload("Orders").Preload("Profile").Preload("Role").Find(&users)
// SELECT * FROM users;
// SELECT * FROM orders WHERE user_id IN (1, 2, 3, 4); // has many
// SELECT * FROM profiles WHERE user_id IN (1, 2, 3, 4); // has one
// SELECT * FROM roles WHERE id IN (4, 5, 6) // belongs to
```
- Using ENUM in Golang:
```
const (
	StateInit JobState = iota // 0
	StateRunning // 1
	StateFailed // 2
	StateTimeout // 3
	StateCompleted // 4
	StateRetryFailed // 5
)
```