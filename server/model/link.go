package model

func GetAllLinks() ([]Link, error) {
	var links []Link

	tx := db.Find(&links)
	if tx.Error != nil {
		return []Link{}, tx.Error
	}

	return links, nil
}

func GetLink(id uint64) (Link, error) {
	var link Link

	tx := db.Where("id = ?", id).First(&link)
	if tx.Error != nil {
		return Link{}, tx.Error
	}

	return link, nil
}

func CreateLink(link Link) error {
	tx := db.Create(&link)

	return tx.Error
}

func UpdateLink(link Link) error {
	tx := db.Save(&link)

	return tx.Error
}

func DeleteLink(id uint64) error {
	tx := db.Unscoped().Delete(&Link{}, id)

	return tx.Error
}

func FindByHash(hash string) (Link, error) {
	var link Link
	tx := db.Where("hash = ?", hash).First(&link)
	if tx.Error != nil {
		return Link{}, tx.Error
	}

	return link, nil
}
