package metrics

import (
	"bytes"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"locgame-mini-server/internal/config"
	"math"
	"net/http"
	"strconv"
	"time"
)

type ZapLogger struct {
	matchLogger  *zap.Logger
	matchHistory *zap.Logger
	lcLogger     *zap.Logger
	playerLogger *zap.Logger
	logger       *zap.Logger
}

func (z *ZapLogger) Log(message string) {
	defer z.logger.Sync()
	go z.logger.Info(message)
}

func (z *ZapLogger) Error(message error) {
	defer z.logger.Sync()
	go z.logger.Error(fmt.Sprint(message))
}

func (z *ZapLogger) LogGameTimeAndMode(user interface{}, opponent interface{}, gameTime string,
	gameMode string, gameResult string, gameReward interface{}, progress ...int32) {
	defer z.matchLogger.Sync()

	time64, err := strconv.ParseInt(gameTime, 10, 64)
	if err != nil {
		fmt.Println(err)
		time64 = 0
	}

	var missionProgress int32 = 0
	var currentLevel int32 = 0
	if len(progress) == 2 {
		missionProgress = progress[0]
		currentLevel = progress[1]
	}

	go z.matchLogger.Info("Game Result",
		zap.Any("user", user),
		zap.Any("opponent", opponent),
		zap.Int64("gameTime", time64),
		zap.String("gameMode", gameMode),
		zap.String("gameResult", gameResult),
		zap.Any("gameReward", gameReward),
		zap.Any("progress", missionProgress),
		zap.Any("currentLevel", currentLevel),
	)
}

func (z *ZapLogger) LCEarned(user interface{}, lc int32, reason ...string) {
	defer z.lcLogger.Sync()

	go z.lcLogger.Info("inventory-repository",
		zap.Any("user", user),
		zap.Int32("earned", lc),
		zap.Any("reason", reason),
	)
}

func (z *ZapLogger) LCSpent(user interface{}, lc int32, reason ...string) {
	defer z.lcLogger.Sync()

	go z.lcLogger.Info("inventory-repository",
		zap.Any("user", user),
		zap.Int32("spent", int32(math.Abs(float64(lc)))),
		zap.Any("reason", reason),
	)
}

func (z *ZapLogger) LastLogin(user interface{}, lastLogin time.Time) {
	defer z.playerLogger.Sync()

	go z.playerLogger.Info("last-login",
		zap.Any("user", user),
	)
}

func (z *ZapLogger) LogGameMove(valid string, gameId string, playerId string, movement string, moveResult string, err ...error) {
	defer z.matchHistory.Sync()

	go z.matchHistory.Info(valid,
		zap.String("gameId", gameId),
		zap.String("playerId", playerId),
		zap.String("movement", movement),
		zap.String("moveResult", moveResult),
		zap.Any("error", err),
	)
}

func NewZapLogger(config *config.Metric) Metric {
	esCfg := elasticsearch.Config{
		Addresses: []string{config.APIAddress},
		Username:  config.Username,
		Password:  config.Password,
		Transport: &transport{http.DefaultTransport},
	}
	esClient, err := elasticsearch.NewClient(esCfg)
	if err != nil {
		fmt.Println(err)
	}

	encoderConfig := zap.NewProductionConfig()
	encoderConfig.EncoderConfig.TimeKey = "@timestamp"
	encoderConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.OutputPaths = []string{"stdout"}
	encoderConfig.ErrorOutputPaths = []string{"stderr"}

	matchLogger := zap.New(newCore("match-info", encoderConfig.EncoderConfig, esClient))
	matchHistory := zap.New(newCore("match-history", encoderConfig.EncoderConfig, esClient))
	lcLogger := zap.New(newCore("lc-info", encoderConfig.EncoderConfig, esClient))
	playerLogger := zap.New(newCore("player-info", encoderConfig.EncoderConfig, esClient))
	logger := zap.New(newCore("locg-logs", encoderConfig.EncoderConfig, esClient))

	return &ZapLogger{
		matchLogger:  matchLogger,
		matchHistory: matchHistory,
		lcLogger:     lcLogger,
		playerLogger: playerLogger,
		logger:       logger,
	}
}

func newCore(index string, encoderConfig zapcore.EncoderConfig, esClient *elasticsearch.Client) zapcore.Core {
	matchCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(&ElasticsearchWriter{
			client: esClient,
			index:  index,
		})),
		zap.InfoLevel,
	)
	return matchCore
}

type ElasticsearchWriter struct {
	client *elasticsearch.Client
	index  string
}

func (w *ElasticsearchWriter) Write(p []byte) (n int, err error) {
	res, err := w.client.Index(
		w.index,
		bytes.NewReader(p),
		//w.client.Index.WithDocumentID("locg-logger"),
		//w.client.Index.WithRefresh("true"),
	)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	defer res.Body.Close()

	// Check response status here and handle errors appropriately

	return len(p), nil
}

type transport struct {
	base http.RoundTripper
}

func (t *transport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("User-Agent", "elastic-logger")
	req.Header.Set("Content-Type", "application/json")

	res, err := t.base.RoundTrip(req)
	if err != nil {
		return nil, err
	}

	// Check for X-Elastic-Product header indicating non-Elasticsearch server
	if res.Header.Get("X-Elastic-Product") != "Elasticsearch" {
		res.Header.Set("X-Elastic-Product", "Elasticsearch")
	}

	return res, err
}
