package bid

import "gorm.io/gorm"

func CreateBid(db *gorm.DB, bid *Bid) error {
	return db.Create(bid).Error
}

func GetBid(db *gorm.DB, bid *Bid) (*Bid, error) {
	var result Bid
	err := db.Where(bid).First(&result).Error
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func UpdateBid(db *gorm.DB, bid *Bid) error {
	return db.Model(&Bid{}).Updates(bid).Error
}

func DeleteBid(db *gorm.DB, bid *Bid) error {
	return db.Delete(&Bid{}, bid).Error
}
