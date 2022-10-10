package restaurantmodel

import (
	"food-delivery/common"
	"strings"
)

const EntityName = "Restaurant"

var (
	ErrNameCannotBeEmpty = common.NewCustomError(nil, "restaurant name can't be blank", "ErrNameCannotBeEmpty")
)

type Restaurant struct {
	common.SQLModel `json:",inline"` // embed struct
	Name            string           `json:"name" gorm:"column:name;"`
	Addr            string           `json:"address" gorm:"column:addr;"`
	Logo            *common.Image    `json:"logo" gorm:"column:logo;"`
	Cover           *common.Images   `json:"cover" gorm:"column:cover;"`
}

func (Restaurant) TableName() string {
	return "restaurants"
}

type RestaurantCreate struct {
	common.SQLModel `json:",inline"` // embed struct
	Name            string           `json:"name" gorm:"column:name;"`
	Addr            string           `json:"address" gorm:"column:addr;"`
	Logo            *common.Image    `json:"logo" gorm:"column:logo;"`
	Cover           *common.Images   `json:"cover" gorm:"column:cover;"`
}

func (RestaurantCreate) TableName() string {
	return Restaurant{}.TableName()
}

type RestaurantUpdate struct {
	Name  *string        `json:"name" gorm:"column:name;"`
	Addr  *string        `json:"address" gorm:"column:addr;"`
	Logo  *common.Image  `json:"logo" gorm:"column:logo;"`
	Cover *common.Images `json:"cover" gorm:"column:cover;"`
}

func (RestaurantUpdate) TableName() string {
	return Restaurant{}.TableName()
}

func (res *RestaurantCreate) Validate() error {
	res.Name = strings.TrimSpace(res.Name)

	if len(res.Name) == 0 {
		return ErrNameCannotBeEmpty
	}

	return nil
}

// If admin or owner, show all information
func (data *Restaurant) Mask(isAdminOrOwner bool) {
	data.GenUID(common.DbTypeRestaurant)
}
