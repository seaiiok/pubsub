## pub/sub messages using channel

```go
func TestMain(t *testing.T) {
	pubsub := New()

	queue1 := make(chan interface{}, 100)
	clientID1 := pubsub.Subscribe("debug", queue1)

	queue2 := make(chan interface{}, 100)
	clientID2 := pubsub.Subscribe("debug", queue2)

	go func() {
		pubsub.Publish("debug", "It's debug.")
	}()

	//for msg := range queue1 {
	//	t.Log(msg)
	//}

	//for msg := range queue2 {
	//	t.Log(msg)
	//}

	t.Log(<-queue1)

	t.Log(<-queue2)

	pubsub.Unsubscribe("debug", clientID1)
	pubsub.Unsubscribe("debug", clientID2)

	_, ok := <-queue1
	t.Logf("its %v\n", ok)

	_, ok = <-queue2
	t.Logf("its %v\n", ok)
}

/*
=== RUN   TestMain
    pubsub_test.go:28: It's debug.
    pubsub_test.go:30: It's debug.
    pubsub_test.go:36: its false
    pubsub_test.go:39: its false
--- PASS: TestMain (0.00s)
PASS
*/
```
