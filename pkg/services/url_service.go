package service

import (
	"log"
	"github.com/sammatcha/urlshortener/pkg/models"
	"gorm.io/gorm"
)
	//interface for interacting with URLs in the db
type URLService struct{
	DB *gorm.DB
}

	//URLService creates new instance with database connection
func NewURL(db *gorm.DB) *URLService {
	return &URLService{DB:db}
}
	//Save url
func (service *URLService) SaveURL(longURL, shortURL string) error {
	url := models.Url{
		LongURL: longURL,
		ShortURL: shortURL,
	}
	result := service.DB.Create(&url)
	if result.Error != nil {
		log.Println("Failed to save url", result.Error)
		return result.Error
	}
	log.Println("Saved to db:", url.ShortURL, "->", url.LongURL)
	return nil
}

	//get url by short url
func (service *URLService) GetURL(shortURL string) (*models.Url, error){
	var url models.Url
	result := service.DB.Where("ShortURL=?", shortURL).First(&url)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
		log.Println("Short URL not found", shortURL)
		return nil,nil
	}
	log.Println("Error retrieving URL:", result.Error)
	return nil, result.Error	
	}	
	return &url, nil
} 	