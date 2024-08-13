package pluginArtifacts

import (
	"sort"
	"time"
)

type GenericImageDetail struct {
	Image         string    `json:"image"`
	ImageDigest   string    `json:"imageDigest"`
	LastUpdatedOn time.Time `json:"imagePushedAt"`
}

func (g *GenericImageDetail) SetImage(image *string) *GenericImageDetail {
	if image == nil {
		return g
	}
	g.Image = *image
	return g
}

func (g *GenericImageDetail) SetImageDigest(imageDigest *string) *GenericImageDetail {
	if imageDigest == nil {
		return g
	}
	g.ImageDigest = *imageDigest
	return g
}

func (g *GenericImageDetail) SetLastUpdatedOn(imagePushedAt *time.Time) *GenericImageDetail {
	if imagePushedAt == nil {
		return g
	}
	g.LastUpdatedOn = *imagePushedAt
	return g
}

func NewGenericImageDetailsFromPlugin() *GenericImageDetail {
	return &GenericImageDetail{}
}

type OrderBy string

const (
	Ascending  = "ASC"
	Descending = "DSC" // default
)

// SortGenericImageDetailsByCreatedOn is used to sort the list of GenericImageDetail by GenericImageDetail.LastUpdatedOn
//   - OrderBy - default value Descending
//   - Original Slice is not manipulated, returns a new slice
func SortGenericImageDetailsByCreatedOn(images []*GenericImageDetail, orderBy OrderBy) []*GenericImageDetail {
	if len(images) == 0 {
		return images
	}
	// don't modify the original slice
	sortedImages := make([]*GenericImageDetail, len(images))
	copy(sortedImages, images)
	// sort by createdOn in descending order
	sort.Slice(sortedImages, func(i, j int) bool {
		if orderBy == Ascending {
			return sortedImages[i].LastUpdatedOn.Before(sortedImages[j].LastUpdatedOn)
		}
		return sortedImages[i].LastUpdatedOn.After(sortedImages[j].LastUpdatedOn)
	})
	return sortedImages
}
