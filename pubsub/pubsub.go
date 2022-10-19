package pubsub

import "context"

type Topic string

type Pubsub interface {
	/*
	channel: When publish message, need to know which topic it belongs to
	*/
	Publish(ctx context.Context, channel Topic, data *Message) error
	/*
	Subscribe return:
		channel: only read cannot write
		close: close channel when neccessary (unsubscribe)
	*/
	Subscribe(ctx context.Context, channel Topic) (ch <-chan *Message, close func())
}
