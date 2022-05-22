// Copyright 2018 visualfc. All rights reserved.

package tk

import "fmt"

// separator
type Separator struct {
	BaseWidget
}

// orient Orient
// orient: "vertical", "horizontal"
func NewSeparator(parent Widget, orient string, attributes ...*WidgetAttr) *Separator {
	iid := makeNamedWidgetId(parent, "atk_separator")
	attributes = append(attributes, &WidgetAttr{"orient", orient})
	info := CreateWidgetInfo(iid, WidgetTypeSeparator, true, attributes)
	if info == nil {
		return nil
	}
	w := &Separator{}
	w.id = iid
	w.info = info
	RegisterWidget(w)
	return w
}

func (w *Separator) Attach(id string) error {
	info, err := CheckWidgetInfo(id, WidgetTypeSeparator)
	if err != nil {
		return err
	}
	w.id = id
	w.info = info
	RegisterWidget(w)
	return nil
}

// tk.NewSeparator(parent Widget, orient string) ==> ttk.Separator
// w.ConfigSet(key string, value string) error
// w.ConfigGet(key string) string
// ttk key : 
// tk key : 

// 以下设置和属性获取使用ConfigSet、ConfigGet
// func (w *Separator) SetOrient(orient Orient) error {
// 	return eval(fmt.Sprintf("%v configure -orient {%v}", w.id, orient))
// }

// func (w *Separator) Orient() Orient {
// 	r, err := evalAsString(fmt.Sprintf("%v cget -orient", w.id))
// 	return parserOrientResult(r, err)
// }

// func (w *Separator) SetTakeFocus(takefocus bool) error {
// 	return eval(fmt.Sprintf("%v configure -takefocus {%v}", w.id, boolToInt(takefocus)))
// }

// func (w *Separator) IsTakeFocus() bool {
// 	r, _ := evalAsBool(fmt.Sprintf("%v cget -takefocus", w.id))
// 	return r
// }

// func SeparatorAttrOrient(orient Orient) *WidgetAttr {
// 	return &WidgetAttr{"orient", orient}
// }

// func SeparatorAttrTakeFocus(takefocus bool) *WidgetAttr {
// 	return &WidgetAttr{"takefocus", boolToInt(takefocus)}
// }
