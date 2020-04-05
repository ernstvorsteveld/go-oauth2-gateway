package oauth2_gateway

import(
    "github.com/ernstvorsteveld/go-oauth2")

type ClientMapRepository struct {
	idMap       map[models.IdType]models.User
	usernameMap map[string]models.User
	emailMap    map[string]models.User
}