package subscriber

import (
	"context"
	"food-delivery/common"
	component "food-delivery/components"
	"food-delivery/components/asyncjob"
	"food-delivery/pubsub"
	"food-delivery/skio"
	"log"
)

type consumerJob struct {
	Title string
	Hld   func(ctx context.Context, message *pubsub.Message) error // handler
}

type consumerEngine struct {
	appCtx         component.AppContext
	realtimeEngine skio.RealtimeEngine
}

func NewEngine(appCtx component.AppContext, realtimeEngine skio.RealtimeEngine) *consumerEngine {
	return &consumerEngine{appCtx: appCtx,	realtimeEngine: realtimeEngine}
}

// Put all subscriber here
func (engine *consumerEngine) Start() error {
	engine.startSubTopic(
		common.TopicUserLikeRestaurant,
		true,
		RunIncreaseLikeCountAfterUserLikeRestaurant(engine.appCtx),
		EmitRealtimeAfterUserLikeRestaurant(engine.appCtx, engine.realtimeEngine),
	)

	engine.startSubTopic(
		common.TopicUserUnlikeRestaurant,
		true,
		RunDecreaseLikeCountAfterUserUnlikeRestaurant(engine.appCtx),
		EmitRealtimeAfterUserUnlikeRestaurant(engine.appCtx, engine.realtimeEngine),
	)

	return nil
}

type GroupJob interface {
	Run(ctx context.Context) error
}

func (engine *consumerEngine) startSubTopic(topic pubsub.Topic, isConcurrent bool, consumerJobs ...consumerJob) error {
	c, _ := engine.appCtx.GetPubsub().Subscribe(context.Background(), topic)

	for _, item := range consumerJobs {
		log.Println("Setup consumer for:", item.Title)
	}

	getJobHandler := func(job *consumerJob, message *pubsub.Message) asyncjob.JobHandler {
		return func(ctx context.Context) error {
			log.Println("running job for", job.Title, ". Value: ", message.Data())
			return job.Hld(ctx, message)
		}
	}

	go func() {
		for {
			msg := <-c

			jobHdlArr := make([]asyncjob.Job, len(consumerJobs))

			for i := range consumerJobs {
				jobHdl := getJobHandler(&consumerJobs[i], msg)
				jobHdlArr[i] = asyncjob.NewJob(jobHdl)
			}

			group := asyncjob.NewGroup(isConcurrent, jobHdlArr...)

			if err := group.Run(context.Background()); err != nil {
				log.Println(err)
			}
		}
	}()

	return nil
}
