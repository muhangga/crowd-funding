package campaign

import "github.com/muhangga/entity"

type CampaignRepository interface {
	FindAll() ([]entity.Campaign, error)
	FindByUserID(userID int) (entity.Campaign, error)
}
