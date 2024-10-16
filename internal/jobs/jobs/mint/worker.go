package mint

import (
	"locgame-mini-server/internal/blockchain"
	"locgame-mini-server/internal/jobs"
	"locgame-mini-server/pkg/log"
	"locgame-mini-server/pkg/pubsub"
	"reflect"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	awsSession "github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"

	storeDto "locgame-mini-server/pkg/dto/store"
)

type Worker struct {
	jobs.BaseWorker

	blockchain *blockchain.Blockchain
	queue      chan *storeDto.MintJobRequest

	s3 *s3.S3
}

func init() {
	worker := new(Worker)
	pubsub.RegisterHandler(worker)
	jobs.RegisterWorker(worker)
}

func (w *Worker) Run() {
	var err error

	cfg := &aws.Config{
		Region: aws.String(w.GetConfig().Aws.Region),
		Credentials: credentials.NewStaticCredentials(
			w.GetConfig().Aws.AccessKeyID,
			w.GetConfig().Aws.SecretAccessKey,
			"",
		),
	}
	sess := awsSession.Must(awsSession.NewSession(cfg))
	w.s3 = s3.New(sess, cfg)
	w.blockchain, err = blockchain.Connect(w.GetConfig().Blockchain)
	if err != nil {
		log.Fatal("Failed to connect blockchain:", err)
	}
	w.queue = make(chan *storeDto.MintJobRequest, 100)

	go func() {
		for {
			select {
			case req := <-w.queue:
				w.WaitGroup.Add(1)
				if req.MintType == storeDto.MintType_MintOrder {
					w.onMintOrderRequestReceived(req)
				}
				if req.MintType == storeDto.MintType_MintGift {
					w.onMintGiftRequestReceived(req)
				}
				if req.MintType == storeDto.MintType_MintUpgrade {
					w.onMintUpgradeOrderRequestReceived(req)
				}
				w.WaitGroup.Done()
			case <-w.BaseWorker.Done:
				log.Debug("Break")
				return
			}
		}
	}()
	return
}

func (w *Worker) Handle(data pubsub.MessageData) {
	req := data.(*storeDto.MintJobRequest)
	w.queue <- req
}

func (w *Worker) GetDataType() reflect.Type {
	return reflect.TypeOf(storeDto.MintJobRequest{})
}
