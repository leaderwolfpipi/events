package events

import "fmt"

// 事件类型基类
type Event struct {
	//事件触发实例
	Target IEventDispatcher
	//事件类型
	Type string
	//事件携带数据源
	Object interface{}
}

// 事件调度器中存放的单元
type EventSaver struct {
	Type      string
	Listeners []*EventListener
}

// 创建事件
func NewEvent(eventType string, object interface{}) Event {
	e := Event{Type: eventType, Object: object}
	return e
}

// 克隆事件
func (this *Event) Clone() *Event {
	e := new(Event)
	e.Type = this.Type
	e.Target = e.Target
	return e
}

func (this *Event) ToString() string {
	return fmt.Sprintf("Event Type %v", this.Type)
}
