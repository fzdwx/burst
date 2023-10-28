package handler

import (
	"github.com/fzdwx/burst/internal/protocal"
	"github.com/fzdwx/burst/internal/server/api/ws/handler/internetResponse"
)

func Dispatch(burst protocal.Burst) {
	switch burst.Type {
	case protocal.IntranetResponseType:
		internetResponse.Handle(burst.IntranetResponse)
	}
}
