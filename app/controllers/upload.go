package controllers
import (
	"github.com/revel/revel"
)

type Upload struct {
	BaseController
}


func (c Upload)DoUploadFile() revel.Result {
	return nil;
}