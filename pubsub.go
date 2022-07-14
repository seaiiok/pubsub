package pubsub

import (
	"sync"

	"github.com/google/uuid"
)

type PubSub interface {
	Publish(topic string, payload interface{})
	Subscribe(topic string, queue chan interface{}) (id string)
	Unsubscribe(topic string, id string)
}

func New() PubSub {
	return &pubSubImpl{
		subscriptions: map[string]map[string]chan interface{}{},
	}
}

type pubSubImpl struct {
	subscriptions map[string]map[string]chan interface{}
	mutex         sync.RWMutex
}

func (pubSub *pubSubImpl) Publish(topic string, payload interface{}) {
	pubSub.mutex.RLock()
	defer pubSub.mutex.RUnlock()

	for _, ch := range pubSub.subscriptions[topic] {
		ch <- payload
	}
}

func (pubSub *pubSubImpl) Subscribe(topic string, queue chan interface{}) string {
	pubSub.mutex.Lock()
	defer pubSub.mutex.Unlock()

	if _, ok := pubSub.subscriptions[topic]; !ok {
		pubSub.subscriptions[topic] = make(map[string]chan interface{})
	}

	id := uuid.New().String()
	pubSub.subscriptions[topic][id] = queue

	return id
}

func (pubSub *pubSubImpl) Unsubscribe(topic string, id string) {
	pubSub.mutex.Lock()
	defer pubSub.mutex.Unlock()

	if subs, ok := pubSub.subscriptions[topic]; ok {
		if queue, ok := subs[id]; ok {
			close(queue)
			delete(subs, id)
		}
	}
}
