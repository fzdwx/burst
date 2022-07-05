package handler

import (
	"github.com/fzdwx/burst/pkg/protocal"
	"github.com/fzdwx/burst/server/api/ws/handler/internetResponse"
)

func Dispatch(burst protocal.Burst) {
	switch burst.Type {
	case protocal.IntranetResponseType:
		internetResponse.Handle(burst.IntranetResponse)
	}
}