package events

// 事件调度接口
type IEventDispatcher interface {
	//事件监听
	AddEventListener(eventType string, listener *EventListener)
	//移除事件监听
	RemoveEventListener(eventType string, listener *EventListener) bool
	//是否包含事件
	HasEventListener(eventType string) bool
	//事件派发
	DispatchEvent(event Event) bool
}
