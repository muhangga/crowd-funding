package campaign

import "gorm.io/gorm"

import "github.com/muhangga/entity"

type repository struct {
	db *gorm.DB
}

func NewCampaignRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]entity.Campaign, error) {
	var campaigns []entity.Campaign
	if err := r.db.Find(&campaigns).Error; err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func (r *repository) FindByUserID(id int) (entity.Campaign, error) {
	var campaign entity.Campaign
	if err := r.db.Where("user_id = ?", id).Find(&campaign).Error; err != nil {
		return campaign, err
	}

	return campaign, nil
}