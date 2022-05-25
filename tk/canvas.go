// Copyright 2018 visualfc. All rights reserved.

package tk

import "fmt"

// Create and manipulate 'canvas' hypergraphics drawing surface widgets
type Canvas struct {
	BaseWidget
	xscrollcommand *CommandEx
	yscrollcommand *CommandEx
}

func NewCanvas(parent Widget, attributes ...*WidgetAttr) *Canvas {
	theme := checkInitUseTheme(attributes)
	iid := makeNamedWidgetId(parent, "atk_canvas")
	info := CreateWidgetInfo(iid, WidgetTypeCanvas, theme, attributes)
	if info == nil {
		return nil
	}
	w := &Canvas{}
	w.id = iid
	w.info = info
	RegisterWidget(w)
	return w
}

func (w *Canvas) Attach(id string) error {
	info, err := CheckWidgetInfo(id, WidgetTypeCanvas)
	if err != nil {
		return err
	}
	w.id = id
	w.info = info
	RegisterWidget(w)
	return nil
}

// tk.NewCanvas(parent Widget) ==> tk.Canvas

// w.ConfigSet(key string, value string) error
// w.ConfigGet(key string) string
// 以下设置和属性获取使用ConfigSet和ConfigGet
// tk  key :
	// background,borderwidth,closeenough,confine,cursor,height,
	// highlightbackground,highlightcolor,highlightthickness,
	// insertbackground,insertborderwidth,insertofftime,insertontime,
	// insertwidth,offset,relief,scrollregion,selectbackground,
	// selectborderwidth,selectforeground,state,takefocus,width,
	// xscrollcommand,xscrollincrement,yscrollcommand,yscrollincrement

func (w *Canvas) SetXScrollIncrement(value int) error {
	return eval(fmt.Sprintf("%v configure -xscrollincrement {%v}", w.id, value))
}

func (w *Canvas) XScrollIncrement() int {
	r, _ := evalAsInt(fmt.Sprintf("%v cget -xscrollincrement", w.id))
	return r
}

func (w *Canvas) SetYScrollIncrement(value int) error {
	return eval(fmt.Sprintf("%v configure -yscrollincrement {%v}", w.id, value))
}

func (w *Canvas) YScrollIncrement() int {
	r, _ := evalAsInt(fmt.Sprintf("%v cget -yscrollincrement", w.id))
	return r
}

// **********新增****************

// line（线）, polygon（多边形）, oval（圆或椭圆形）, 矩形, 扇形, 

// pathName create line x1 y1... xn yn ?option value ...?
func (w *Canvas) PlotLine(pots [][2]int, options ...[2]string) error {
	// options: 
	// -dash -activedash -disableddash -dashoffset -fill -activefill -disabledfil -stipple 
	// -activestipple -disabledstipple -state -tags -width -activewidth -disabledwidth
	
	// extra options:
	// -arrow, -arrowshape, -capstyle, -joinstyle, -smooth, -splinesteps

	potall := ""
	for _,v := range pots {
		potall = potall + strconv.Itoa(v[0]) + " " + strconv.Itoa(v[1]) + " "
	}
	
	opts := ""
	for _,v := range options {
		opts = opts + "-" + v[0] + " " + v[1] + " "
	}
	
	return eval(fmt.Sprintf("%v create line %v %v", w.id, potall, opts))
}

// pathName create polygon x1 y1 ... xn yn ?option value ...?
func (w *Canvas) PlotPoly(pots [][2]int, options ...[2]string) error {
	// options: 
	// -dash  -activedash  -disableddash  -dashoffset  -fill  -activefill  -disabledfill
	// -offset  -outline  -activeoutline  -disabledoutline  -outlineoffset  -outlinestipple
	// -activeoutlinestipple  -disabledoutlinestipple  -stipple  -activestipple
	// -disabledstipple  -state  -tags  -width  -activewidth  -disabledwidth
	
	// extra options:
	// -joinstyle, -smooth, -splinesteps

	potall := ""
	for _,v := range pots {
		potall = potall + strconv.Itoa(v[0]) + " " + strconv.Itoa(v[1]) + " "
	}
	
	opts := ""
	for _,v := range options {
		opts = opts + "-" + v[0] + " " + v[1] + " "
	}
	
	return eval(fmt.Sprintf("%v create polygon %v %v", w.id, potall, opts))
}

// pathName create rectangle x1 y1 x2 y2 ?option value ...?
func (w *Canvas) PlotRec(rec [4]int, options ...[2]string) error {
	// options:
	// -dash  -activedash  -disableddash  -dashoffset  -fill  -activefill  -disabledfill  
	// -offset  -outline  -activeoutline  -disabledoutline  -outlineoffset  -outlinestipple  
	// -activeoutlinestipple  -disabledoutlinestipple  -stipple  -activestipple  
	// -disabledstipple  -state  -tags  -width  -activewidth  -disabledwidth 
	
	x1,y1,x2,y2 := rec[0],rec[1],rec[2],rec[3]
	opts := ""
	for _,v := range options {
		opts = opts + "-" + v[0] + " " + v[1] + " "
	}
	
	return eval(fmt.Sprintf("%v create rectangle %v %v %v %v %v", w.id, x1, y1, x2, y2, opts))
}

// pathName create oval x1 y1 x2 y2 ?option value ...?
func (w *Canvas) PlotOval(rec [4]int, options ...[2]string) error {
	// options:
	// -dash  -activedash  -disableddash  -dashoffset  -fill  -activefill  -disabledfill  
	// -offset  -outline  -activeoutline  -disabledoutline  -outlineoffset  -outlinestipple  
	// -activeoutlinestipple  -disabledoutlinestipple  -stipple  -activestipple  
	// -disabledstipple  -state  -tags  -width  -activewidth  -disabledwidth  

	x1,y1,x2,y2 := rec[0],rec[1],rec[2],rec[3]
	opts := ""
	for _,v := range options {
		opts = opts + "-" + v[0] + " " + v[1] + " "
	}
	
	return eval(fmt.Sprintf("%v create oval %v %v %v %v %v", w.id, x1, y1, x2, y2, opts))
}

// pathName create arc x1 y1 x2 y2 ?option value ...?
func (w *Canvas) PlotArc(rec [4]int, options ...[2]string) error {
	// options:
	// -dash  -activedash  -disableddash  -dashoffset  -fill  -activefill  -disabledfill  
	// -offset  -outline  -activeoutline  -disabledoutline  -outlineoffset  -outlinestipple  
	// -activeoutlinestipple  -disabledoutlinestipple  -stipple  -activestipple  
	// -disabledstipple  -state  -tags  -width  -activewidth  -disabledwidth  
	
	// extra options: 
	// -extent, -start, -style
	
	x1,y1,x2,y2 := rec[0],rec[1],rec[2],rec[3]
	opts := ""
	for _,v := range options {
		opts = opts + "-" + v[0] + " " + v[1] + " "
	}
	
	return eval(fmt.Sprintf("%v create arc %v %v %v %v %v", w.id, x1, y1, x2, y2, opts))
}

// 文字，图片，控件
// pathName create text x y ?option value ...?
func (w *Canvas) PlotText(pos [2]int, text string, options ...[2]string) error {
	// options:
	// -anchor -fill -activefill -disabledfill -stipple 
	// -activestipple -disabledstipple -state -tags
	
	// extra options: 
	// -angle, -font, -justify, -text, -underline, -width
	
	// text 值含有空格时使用{}, text string : "{A wonderful story}"
	x1,y1 := pos[0],pos[1]
	opts := ""
	for _,v := range options {
		opts = opts + "-" + v[0] + " " + v[1] + " "
	}
	
	return eval(fmt.Sprintf("%v create text %v %v -text %v %v", w.id, x1, y1, text, opts))
}

// pathName create image x y ?option value ...?
func (w *Canvas) PlotImg(pos [2]int, imageid string, options ...[2]string) error {
	// options:
	// -anchor -state -tags
	
	// extra options: 
	// -image, -activeimage, -disabledimage
	
	x1,y1 := pos[0],pos[1]
	opts := ""
	for _,v := range options {
		opts = opts + "-" + v[0] + " " + v[1] + " "
	}
	
	return eval(fmt.Sprintf("%v create image %v %v -image %v %v", w.id, x1, y1, imageid, opts))
}

// pathName create window x y ?option value ...?
func (w *Canvas) PlotWin(pos [2]int, widgetid string, options ...[2]string) error {
	// options:
	// -anchor -state -tags
	
	// extra options: 
	// -height, -width, -window
	
	x1,y1 := pos[0],pos[1]
	opts := ""
	for _,v := range options {
		opts = opts + "-" + v[0] + " " + v[1] + " "
	}
	
	return eval(fmt.Sprintf("%v create window %v %v -window %v %v", w.id, x1, y1, widgetid, opts))
}


// pathName bind tagOrId ?sequence? ?command?


// pathName delete ?tagOrId tagOrId ...?
func (w *Canvas) DeleteTags(tagnames []string) error {
	tags := ""
	for _,v := range tagnames {
		tags = tags + v + " "
	}
	return eval(fmt.Sprintf("%v delete %v", w.id, tags))
}

// pathName itemcget tagOrId option
func (w *Canvas) TagOptionGet(tagname string, option string) string {
	res, _ := evalAsString(fmt.Sprintf("%v itemcget %v -%v", w.id, tagname, option))
	return res
}

// pathName itemconfigure tagOrId ?option? ?value? ?option value ...?
func (w *Canvas) TagOptionsSet(tagname string, optvalue ...[2]string) error {
	optvalueset := ""
	for _,v := range optvalue {
		optvalueset = optvalueset + "-" +v[0] + " " + v[1] + " "
	}
	return eval(fmt.Sprintf("%v itemconfigure %v %v", w.id, tagname, optvalueset))
}

// pathName scale tagOrId xOrigin yOrigin xScale yScale
func (w *Canvas) TagScale(tagname string, x0 int, y0 int, xsc float64, ysc float64) error {

	return eval(fmt.Sprintf("%v scale %v %v %v %v %v", w.id, tagname, x0, y0, xsc, ysc))
}

// pathName coords tagOrId ?coordList?
func (w *Canvas) TagCoordsGet(tagname string) string {
	res, _ := evalAsString(fmt.Sprintf("%v coords %v", w.id, tagname))
	return res
}

// pathName raise tagOrId ?aboveThis?
func (w *Canvas) TagRaise(raiseTag string, aboveTag string) error {

	return eval(fmt.Sprintf("%v raise %v %v", w.id, raiseTag, aboveTag))
}

// pathName lower tagOrId ?belowThis?
func (w *Canvas) TagLower(lowerTag string, belowTag string) error {

	return eval(fmt.Sprintf("%v lower %v %v", w.id, lowerTag, belowTag))
}

// pathName move tagOrId xAmount yAmount
func (w *Canvas) TagMove(moveTag string, dx int, dy int) error {

	return eval(fmt.Sprintf("%v move %v %v %v", w.id, moveTag, dx, dy))
}

// pathName moveto tagOrId xPos yPos
func (w *Canvas) TagMoveTo(moveTag string, xpos int, ypos int) error {

	return eval(fmt.Sprintf("%v moveto %v %v %v", w.id, moveTag, xpos, ypos))
}

// **********新增****************
