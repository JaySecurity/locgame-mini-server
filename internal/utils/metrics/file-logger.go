package metrics

import (
	"fmt"
	"log"
	"os"
	"time"
)

type FileLogger struct{}

func (fl *FileLogger) Log(message string) {
	f, err := os.OpenFile("metrics.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	_, err = f.WriteString(message + "\n")
	if err != nil {
		log.Fatal(err)
	}
}

func (fl *FileLogger) Error(message error) {
	f, err := os.OpenFile("metrics-errors.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	_, err = f.WriteString(fmt.Sprint(message) + "\n")
	if err != nil {
		log.Fatal(err)
	}
}

func (fl *FileLogger) LCEarned(user interface{}, lc int32, reason ...string) {
	fl.Log(fmt.Sprint(user, " earned ", lc, reason))
}

func (fl *FileLogger) LCSpent(user interface{}, lc int32, reason ...string) {
	fl.Log(fmt.Sprint(user, " spent ", lc, reason))
}

func (fl *FileLogger) LogGameTimeAndMode(user interface{}, opponent interface{}, gameTime string,
	gameMode string, gameResult string, gameReward interface{}, progress ...int32) {
	fl.Log(fmt.Sprint(user, " ", opponent, " ", gameTime, " ", gameMode, " ", gameResult, " ", gameReward, " ", progress))
}

func (fl *FileLogger) LastLogin(user interface{}, lastLogin time.Time) {
	fl.Log(fmt.Sprint(user, " ", lastLogin))
}

func (fl *FileLogger) LogGameMove(valid string, gameId string, playerId string, movement string, moveResult string, err ...error) {
	fl.Log(fmt.Sprint(valid, " gameId: ", gameId, " playerId: ", playerId,
		" movement: ", movement, " moveResult ", moveResult, " errors: ", err))
}

func NewFileLogger() Metric {
	return &FileLogger{}
}
