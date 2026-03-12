package tyme

import (
	"fmt"
	"regexp"
)

const (
	EventManagerChars = "0123456789ABCDEFGHIJKLMNOPQRSTU_VWXYZabcdefghijklmnopqrstuvwxyz"
	EventManagerRegex = "(@[0-9A-Za-z_]{8})(%s)"
)

var EventManagerData = ""

// EventManager 事件管理器
type EventManager struct {
}

// Remove 通过事件名称移除事件
func (EventManager) Remove(name string) {
	pattern := regexp.MustCompile(fmt.Sprintf(EventManagerRegex, name))
	EventManagerData = pattern.ReplaceAllString(EventManagerData, "")
}

func saveOrUpdate(name string, data string) {
	o := fmt.Sprintf(EventManagerRegex, name)
	re := regexp.MustCompile(o)
	if re.MatchString(EventManagerData) {
		EventManagerData = re.ReplaceAllString(EventManagerData, data)
	} else {
		EventManagerData += data
	}
}

// UpdateEvent 新增或更新事件
func (EventManager) UpdateEvent(name string, event *Event) {
	eventName := event.GetName()
	if eventName == "" {
		eventName = name
	}
	saveOrUpdate(name, event.GetData()+eventName)
}

// UpdateEventData 新增或更新事件
func (EventManager) UpdateEventData(name string, data string) error {
	err := Event{}.Validate(data)
	if err != nil {
		return err
	}
	saveOrUpdate(name, data)
	return nil
}
