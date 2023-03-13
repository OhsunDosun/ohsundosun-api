package util

import (
	"ohsundosun-api/db"
	"ohsundosun-api/model"

	"github.com/deta/deta-go/service/base"
)

func VerifyEmail(email *string) bool {
	query := base.Query{
		{"email": email},
	}

	var result []*model.User

	db.BaseUser.Fetch(&base.FetchInput{
		Q:    query,
		Dest: &result,
	})

	return len(result) == 0
}

func VerifyNickname(nickname *string) bool {
	query := base.Query{
		{"nickname": nickname},
	}

	var result []*model.User

	db.BaseUser.Fetch(&base.FetchInput{
		Q:    query,
		Dest: &result,
	})

	return len(result) == 0
}
