package shared

import (
	"context"

	"locgame-mini-server/pkg/dto/accounts"
)

type AccountInfoUpdateListener interface {
	OnAccountInfoChanged(ctx context.Context, info *accounts.AccountInfo)
}
