package deta

import (
	"fmt"
	"os"

	"github.com/deta/deta-go/deta"
	"github.com/deta/deta-go/service/base"
	"github.com/deta/deta-go/service/drive"
)

var Deta *deta.Deta
var BaseUser *base.Base
var BaseRating *base.Base
var BasePost *base.Base
var BaseLikeSortPost *base.Base

var DrivePost *drive.Drive

func init() {
	d, err := deta.New(deta.WithProjectKey(os.Getenv("DETA_COLLECTION_KEY")))
	if err != nil {
		fmt.Println("failed to init new Deta instance:", err)
		return
	}

	Deta = d

	userBase, err := base.New(Deta, "User")
	if err != nil {
		fmt.Println("failed to init new Base instance:", err)
		return
	}

	BaseUser = userBase

	ratingBase, err := base.New(Deta, "Rating")
	if err != nil {
		fmt.Println("failed to init new Base instance:", err)
		return
	}

	BaseRating = ratingBase

	postBase, err := base.New(Deta, "Post")
	if err != nil {
		fmt.Println("failed to init new Base instance:", err)
		return
	}

	BasePost = postBase

	likeSortPostBase, err := base.New(Deta, "Like_Sort_Post")
	if err != nil {
		fmt.Println("failed to init new Base instance:", err)
		return
	}

	BaseLikeSortPost = likeSortPostBase

	postDrive, err := drive.New(Deta, "Post")
	if err != nil {
		fmt.Println("failed to init new Drive instance:", err)
		return
	}

	DrivePost = postDrive
}
