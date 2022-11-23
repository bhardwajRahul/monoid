package loader

import (
	"context"

	"github.com/brist-ai/monoid/config"
	"github.com/brist-ai/monoid/model"
	"github.com/graph-gophers/dataloader"
	"github.com/rs/zerolog/log"
)

// GetPropertyCategories wraps the associated dataloader
func GetPropertyCategories(ctx context.Context, propertyID string) ([]*model.Category, error) {
	loaders := For(ctx)
	thunk := loaders.PropertyCategoriesLoader.Load(ctx, dataloader.StringKey(propertyID))
	result, err := thunk()
	if err != nil {
		return nil, err
	}
	return result.([]*model.Category), nil
}

// CategoryReader reads categories from a database
type PropertyCategoryReader struct {
	conf *config.BaseConfig
}

// GetCategories gets all the categories for a number of properties
func (c *PropertyCategoryReader) GetPropertyCategories(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	propertyIDs := make([]string, len(keys))
	for ix, key := range keys {
		propertyIDs[ix] = key.String()
	}

	type propertyCategory struct {
		PropertyID string
		CategoryID string
		Category   *model.Category
	}

	pcs := []propertyCategory{}

	// Read from the property_categories table and get all associated
	// categories.
	if err := c.conf.DB.Where(
		"property_id IN ?",
		propertyIDs,
	).Preload("Category").Find(&pcs).Error; err != nil {
		log.Err(err).Msg("Error finding categories")
	}

	categoryMap := map[string][]*model.Category{}
	for _, c := range pcs {
		if categoryMap[c.PropertyID] == nil {
			categoryMap[c.PropertyID] = []*model.Category{}
		}

		categoryMap[c.PropertyID] = append(categoryMap[c.PropertyID], c.Category)
	}

	// Reassign output to an array of array results.
	output := make([]*dataloader.Result, len(keys))
	for index, catKey := range keys {
		cats, ok := categoryMap[catKey.String()]
		if ok {
			output[index] = &dataloader.Result{Data: cats, Error: nil}
		} else {
			output[index] = &dataloader.Result{Data: []*model.Category{}, Error: nil}
		}
	}

	return output
}