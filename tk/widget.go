// Copyright 2017 visualfc. All rights reserved.

package tk

import (
	"fmt"
	"os"
	"strings"
)

type WidgetId string

type Widget interface {
	Id() string
	Type() string
}

var (
	globalWidgetMap = make(map[string]Widget)
)

func RegisterWidget(w Widget) {
	if w == nil {
		return
	}
	globalWidgetMap[w.Id()] = w
}

func LookupWidget(id string) (w Widget, ok bool) {
	w, ok = globalWidgetMap[id]
	return
}

func ParentOfWidget(w Widget) Widget {
	if w == nil {
		return nil
	}
	id := w.Id()
	if id == "." {
		return nil
	}
	pos := strings.LastIndex(id, ".")
	if pos == -1 {
		return nil
	} else if pos == 0 {
		return mainWindow
	}
	return globalWidgetMap[id[:pos]]
}

func ChildrenOfWidget(w Widget) (list []Widget) {
	if w == nil {
		return nil
	}
	id := w.Id()
	if id == "." {
		for k, v := range globalWidgetMap {
			if strings.HasPrefix(k, id) {
				if k == "." {
					continue
				} else if strings.Index(k[1:], ".") >= 0 {
					continue
				}
				list = append(list, v)
			}
		}
	} else {
		id = id + "."
		offset := len(id)
		for k, v := range globalWidgetMap {
			if strings.HasPrefix(k, id) {
				if strings.Index(k[offset:], ".") >= 0 {
					continue
				}
				list = append(list, v)
			}
		}
	}
	return
}

func removeWidget(id string) {
	if id == "." {
		globalWidgetMap = make(map[string]Widget)
	} else {
		delete(globalWidgetMap, id)
		id = id + "."
		var list []string
		for k, _ := range globalWidgetMap {
			if strings.HasPrefix(k, id) {
				list = append(list, k)
			}
		}
		for _, k := range list {
			delete(globalWidgetMap, k)
		}
	}
}

func IsValidWidget(w Widget) bool {
	if w == nil {
		return false
	}
	_, ok := globalWidgetMap[w.Id()]
	return ok
}

func DestroyWidget(w Widget) error {
	if !IsValidWidget(w) {
		return os.ErrInvalid
	}
	id := w.Id()
	eval(fmt.Sprintln("destroy %v", id))
	removeWidget(id)
	return nil
}

var (
	fnGenWidgetId = NewGenInt64Func(1024)
)

func MakeWidgetId(id string, parent Widget) string {
	if len(id) == 0 {
		id = fmt.Sprintf("gotk_id%v", <-fnGenWidgetId())
	} else if id[0] == '.' {
		return id
	}
	if parent != nil {
		return string(parent.Id()) + "." + id
	}
	return "." + id
}