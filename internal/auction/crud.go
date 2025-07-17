package auction

import "gorm.io/gorm"

func CreateAuction(db *gorm.DB, auction *Auction) error {
	return db.Create(auction).Error
}

func GetAuction(db *gorm.DB, auction *Auction) (*Auction, error) {
	var result Auction
	err := db.Where(auction).First(&result).Error
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func UpdateAuction(db *gorm.DB, auction *Auction) error {
	return db.Model(&Auction{}).Updates(auction).Error
}

func DeleteAuction(db *gorm.DB, auction *Auction) error {
	return db.Delete(&Auction{}, auction).Error
}
