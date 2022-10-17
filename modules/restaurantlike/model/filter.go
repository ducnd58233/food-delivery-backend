package restaurantlikemodel

type Filter struct {
	RestaurantId int `json:"-" form:"restaurant_id"` // form: query string
	UserId       int `json:"-" form:"user_id"`
}
