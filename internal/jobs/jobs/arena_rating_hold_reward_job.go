package jobs

import (
	"locgame-mini-server/internal/jobs"
)

func init() {
	jobs.Register("0 21 */1 * *", &ArenaRatingHoldRewardJob{})
}

type ArenaRatingHoldRewardJob struct {
	jobs.BaseJob
}

func (j *ArenaRatingHoldRewardJob) Run() error {
	// TODO
	return nil
}
