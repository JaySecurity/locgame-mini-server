package jobs

import (
	"context"

	"locgame-mini-server/internal/jobs"
)

func init() {
	jobs.Register("0 0 */1 * *", &DeleteUnpaidOrdersJob{})
}

type DeleteUnpaidOrdersJob struct {
	jobs.BaseJob
}

func (j *DeleteUnpaidOrdersJob) Run() error {
	deletedCount, err := j.GetStore().Orders.DeleteUnpaidOrders(context.Background())
	if err != nil {
		return err
	}
	j.GetLogger().Info("Deleted orders:", deletedCount)
	return nil
}
