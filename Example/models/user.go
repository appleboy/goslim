package models

import (
// "github.com/gowebf/goslim/model"
)

type TestUser struct {
	ID   int    `type:"Interger",prop:"autoincrement"`
	Name string `type:"Varchar",length:"10"`
}
