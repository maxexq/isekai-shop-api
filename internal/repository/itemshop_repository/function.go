package itemshoprepository

import (
	"strings"

	"github.com/maxexq/isekei-shop-api/entities"
	_itemShopModel "github.com/maxexq/isekei-shop-api/pkg/itemShop/model"
	"gorm.io/gorm"
)

func (r *RepositoryService) buildItemFilterQuery(itemFilter *_itemShopModel.ItemFilter) *gorm.DB {
	query := r.db.Connect().Model(&entities.Item{}).Where("is_archive = ?", false)

	if name := strings.TrimSpace(itemFilter.Name); name != "" {
		query = query.Where("name ILIKE ?", "%"+name+"%")
	}

	if desc := strings.TrimSpace(itemFilter.Description); desc != "" {
		query = query.Where("description ILIKE ?", "%"+desc+"%")
	}

	return query
}
