package db

import (
	"fmt"
	"os"

	"github.com/deta/deta-go/deta"
	"github.com/deta/deta-go/service/base"
)

var Deta *deta.Deta
var BaseUser *base.Base
var BaseRating *base.Base
var BasePost *base.Base

func init() {
	d, err := deta.New(deta.WithProjectKey(os.Getenv("DETA_COLLECTION_KEY")))
	if err != nil {
		fmt.Println("failed to init new Deta instance:", err)
		return
	}

	Deta = d

	user, err := base.New(d, "User")
	if err != nil {
		fmt.Println("failed to init new Base instance:", err)
		return
	}

	BaseUser = user

	rating, err := base.New(d, "Rating")
	if err != nil {
		fmt.Println("failed to init new Base instance:", err)
		return
	}

	BaseRating = rating

	post, err := base.New(d, "Post")
	if err != nil {
		fmt.Println("failed to init new Base instance:", err)
		return
	}

	BasePost = post
}
