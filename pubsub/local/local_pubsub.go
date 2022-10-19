package pubsublocal

import (
	"context"
	"food-delivery/common"
	"food-delivery/pubsub"
	"log"
	"sync"
)

/*
 A pubsub run locally (in-mem)
 It has a queue (buffer channel) at it's core and many group of subscribers
 Because we want to send a message with a specific topic for many subscribers in a group can handle.
*/

type localPubSub struct {
	messageQueue chan *pubsub.Message
	mapChannel   map[pubsub.Topic][]chan *pubsub.Message // streaming
	locker       *sync.RWMutex                           // When concurrent is running, make sure 1 goroutine can access data
}

func NewPubSub() *localPubSub {
	ps := &localPubSub{
		messageQueue: make(chan *pubsub.Message, 10000),
		mapChannel:   make(map[pubsub.Topic][]chan *pubsub.Message),
		locker:       new(sync.RWMutex),
	}

	ps.run()

	return ps
}

// Enqueue
func (ps *localPubSub) Publish(
	ctx context.Context,
	topic pubsub.Topic,
	data *pubsub.Message,
) error {
	data.SetChannel(topic)

	go func() {
		// use go func because dont want it stuck when this messageQueue bottleneck
		defer common.AppRecover()
		ps.messageQueue <- data // enqueue message
		log.Println("New event published", data.String(), "with data", data.Data())
	}()

	return nil
}

func (ps *localPubSub) Subscribe(ctx context.Context, topic pubsub.Topic) (ch <-chan *pubsub.Message, close func()) {
	c := make(chan *pubsub.Message) // make channel with value inside is Message pointer

	// make sure it will not crash if there are many concurent calling this
	ps.locker.Lock()

	// if the topic already subscribe and already has channel, then append to the channel
	if val, ok := ps.mapChannel[topic]; ok {
		val = append(ps.mapChannel[topic], c)
		ps.mapChannel[topic] = val
	} else {
		ps.mapChannel[topic] = []chan *pubsub.Message{c}
	}

	ps.locker.Unlock()

	// Should not close channel in golang, do this instead
	return c, func() {
		log.Println("Unsubscribe")

		// Remove channel from the array
		if chans, ok := ps.mapChannel[topic]; ok {
			for i := range chans {
				if chans[i] == c {
					// remove element at index in chans
					chans = append(chans[:i], chans[i+1:]...)

					ps.locker.Lock()
					ps.mapChannel[topic] = chans
					ps.locker.Unlock()
					break
				}
			}
		}
	}
}

// Engine to push message to channel
func (ps *localPubSub) run() error {
	log.Println("Pubsub started")

	go func() {
		for { // for forever
			mess := <-ps.messageQueue // get message from queue
			log.Println("Message dequeue:", mess)

			if subs, ok := ps.mapChannel[mess.Channel()]; ok {
				for i := range subs {
					go func(c chan *pubsub.Message) {
						c <- mess // push message to each channel
					}(subs[i])
				}
			}
		}
	}()

	return nil
}
