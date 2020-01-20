package events

import (
	"fmt"
	"testing"
	"time"
)

// 事件常量
const TEST_EVENT = "test"

// 事件功能
func TestEvent(t *testing.T) {
	dispatcher := NewEventDispatcher()
	listener := NewEventListener(myEventListener)
	dispatcher.AddEventListener(TEST_EVENT, listener)

	time.Sleep(time.Second * 2)
	//dispatcher.RemoveEventListener(TEST_EVENT, listener)

	dispatcher.DispatchEvent(NewEvent(TEST_EVENT, nil))
}

// 触发动作
func myEventListener(event Event) {
	fmt.Println(event.Type, event.Object, event.Target)
}
