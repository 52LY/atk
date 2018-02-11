// Copyright 2018 visualfc. All rights reserved.

package tk

var (
	typeMetaMap = make(map[WidgetType]*MetaType)
)

type MetaClass struct {
	Command    string
	Class      string
	Attributes []string
}

func (m *MetaClass) HasAttribute(attr string) bool {
	if attr == "" {
		return false
	}
	for _, v := range m.Attributes {
		if v == attr {
			return true
		}
	}
	return false
}

type MetaType struct {
	Type string
	Tk   *MetaClass
	Ttk  *MetaClass
}

func IsTtkClass(class string) bool {
	for _, v := range ttkClassList {
		if v == class {
			return true
		}
	}
	return false
}

func IsTkClass(class string) bool {
	for _, v := range tkClassList {
		if v == class {
			return true
		}
	}
	return false
}

var (
	tkClassList = []string{
		"Button",
		"Canvas",
		"Checkbutton",
		"Entry",
		"Frame",
		"Label",
		"Labelframe",
		"Listbox",
		"Menu",
		"Menubutton",
		"Panedwindow",
		"Radiobutton",
		"Scale",
		"Scrollbar",
		"Spinbox",
		"Text",
		"Toplevel",
	}
	ttkClassList = []string{
		"TButton",
		"TCheckbutton",
		"TCombobox",
		"TEntry",
		"TFrame",
		"TLabel",
		"TLabelframe",
		"TMenubutton",
		"TNotebook",
		"TPanedwindow",
		"TProgressbar",
		"TRadiobutton",
		"TScale",
		"Scrollbar",
		"TSeparator",
		"TSizegrip",
		"Treeview",
	}
)

func init() {
	typeMetaMap[WidgetTypeButton] =
		&MetaType{
			Type: "Button",
			Tk: &MetaClass{"tk::button", "Button",
				[]string{"activebackground",
					"activeforeground",
					"anchor",
					"background",
					"bitmap",
					"borderwidth",
					"command",
					"compound",
					"cursor",
					"default",
					"disabledforeground",
					"font",
					"foreground",
					"height",
					"highlightbackground",
					"highlightcolor",
					"highlightthickness",
					"image",
					"justify",
					"overrelief",
					"padx",
					"pady",
					"relief",
					"repeatdelay",
					"repeatinterval",
					"state",
					"takefocus",
					"text",
					"textvariable",
					"underline",
					"width",
					"wraplength"}},
			Ttk: &MetaClass{"ttk::button", "TButton",
				[]string{"command",
					"default",
					"takefocus",
					"text",
					"textvariable",
					"underline",
					"width",
					"image",
					"compound",
					"padding",
					"state",
					"cursor",
					"style",
					"class"}},
		}
	typeMetaMap[WidgetTypeCanvas] =
		&MetaType{
			Type: "Canvas",
			Tk: &MetaClass{"tk::canvas", "Canvas",
				[]string{"background",
					"borderwidth",
					"closeenough",
					"confine",
					"cursor",
					"height",
					"highlightbackground",
					"highlightcolor",
					"highlightthickness",
					"insertbackground",
					"insertborderwidth",
					"insertofftime",
					"insertontime",
					"insertwidth",
					"offset",
					"relief",
					"scrollregion",
					"selectbackground",
					"selectborderwidth",
					"selectforeground",
					"state",
					"takefocus",
					"width",
					"xscrollcommand",
					"xscrollincrement",
					"yscrollcommand",
					"yscrollincrement"}},
			Ttk: nil,
		}
	typeMetaMap[WidgetTypeCheckButton] =
		&MetaType{
			Type: "CheckButton",
			Tk: &MetaClass{"tk::checkbutton", "Checkbutton",
				[]string{"activebackground",
					"activeforeground",
					"anchor",
					"background",
					"bitmap",
					"borderwidth",
					"command",
					"compound",
					"cursor",
					"disabledforeground",
					"font",
					"foreground",
					"height",
					"highlightbackground",
					"highlightcolor",
					"highlightthickness",
					"image",
					"indicatoron",
					"justify",
					"offrelief",
					"offvalue",
					"onvalue",
					"overrelief",
					"padx",
					"pady",
					"relief",
					"selectcolor",
					"selectimage",
					"state",
					"takefocus",
					"text",
					"textvariable",
					"tristateimage",
					"tristatevalue",
					"underline",
					"variable",
					"width",
					"wraplength"}},
			Ttk: &MetaClass{"ttk::checkbutton", "TCheckbutton",
				[]string{"variable",
					"onvalue",
					"offvalue",
					"command",
					"takefocus",
					"text",
					"textvariable",
					"underline",
					"width",
					"image",
					"compound",
					"padding",
					"state",
					"cursor",
					"style",
					"class"}},
		}
	typeMetaMap[WidgetTypeComboBox] =
		&MetaType{
			Type: "ComboBox",
			Tk:   nil,
			Ttk: &MetaClass{"ttk::combobox", "TCombobox",
				[]string{"height",
					"postcommand",
					"values",
					"exportselection",
					"font",
					"invalidcommand",
					"justify",
					"show",
					"state",
					"textvariable",
					"validate",
					"validatecommand",
					"width",
					"xscrollcommand",
					"foreground",
					"background",
					"takefocus",
					"cursor",
					"style",
					"class"}},
		}
	typeMetaMap[WidgetTypeEntry] =
		&MetaType{
			Type: "Entry",
			Tk: &MetaClass{"tk::entry", "Entry",
				[]string{"background",
					"borderwidth",
					"cursor",
					"disabledbackground",
					"disabledforeground",
					"exportselection",
					"font",
					"foreground",
					"highlightbackground",
					"highlightcolor",
					"highlightthickness",
					"insertbackground",
					"insertborderwidth",
					"insertofftime",
					"insertontime",
					"insertwidth",
					"invalidcommand",
					"justify",
					"readonlybackground",
					"relief",
					"selectbackground",
					"selectborderwidth",
					"selectforeground",
					"show",
					"state",
					"takefocus",
					"textvariable",
					"validate",
					"validatecommand",
					"width",
					"xscrollcommand"}},
			Ttk: &MetaClass{"ttk::entry", "TEntry",
				[]string{"exportselection",
					"font",
					"invalidcommand",
					"justify",
					"show",
					"state",
					"textvariable",
					"validate",
					"validatecommand",
					"width",
					"xscrollcommand",
					"foreground",
					"background",
					"takefocus",
					"cursor",
					"style",
					"class"}},
		}
	typeMetaMap[WidgetTypeFrame] =
		&MetaType{
			Type: "Frame",
			Tk: &MetaClass{"tk::frame", "Frame",
				[]string{"borderwidth",
					"class",
					"relief",
					"background",
					"colormap",
					"container",
					"cursor",
					"height",
					"highlightbackground",
					"highlightcolor",
					"highlightthickness",
					"padx",
					"pady",
					"takefocus",
					"visual",
					"width"}},
			Ttk: &MetaClass{"ttk::frame", "TFrame",
				[]string{"borderwidth",
					"padding",
					"relief",
					"width",
					"height",
					"takefocus",
					"cursor",
					"style",
					"class"}},
		}
	typeMetaMap[WidgetTypeLabel] =
		&MetaType{
			Type: "Label",
			Tk: &MetaClass{"tk::label", "Label",
				[]string{"activebackground",
					"activeforeground",
					"anchor",
					"background",
					"bitmap",
					"borderwidth",
					"compound",
					"cursor",
					"disabledforeground",
					"font",
					"foreground",
					"height",
					"highlightbackground",
					"highlightcolor",
					"highlightthickness",
					"image",
					"justify",
					"padx",
					"pady",
					"relief",
					"state",
					"takefocus",
					"text",
					"textvariable",
					"underline",
					"width",
					"wraplength"}},
			Ttk: &MetaClass{"ttk::label", "TLabel",
				[]string{"background",
					"foreground",
					"font",
					"borderwidth",
					"relief",
					"anchor",
					"justify",
					"wraplength",
					"takefocus",
					"text",
					"textvariable",
					"underline",
					"width",
					"image",
					"compound",
					"padding",
					"state",
					"cursor",
					"style",
					"class"}},
		}
	typeMetaMap[WidgetTypeLabelFrame] =
		&MetaType{
			Type: "LabelFrame",
			Tk: &MetaClass{"tk::labelframe", "Labelframe",
				[]string{"borderwidth",
					"class",
					"font",
					"foreground",
					"labelanchor",
					"labelwidget",
					"relief",
					"text",
					"background",
					"colormap",
					"container",
					"cursor",
					"height",
					"highlightbackground",
					"highlightcolor",
					"highlightthickness",
					"padx",
					"pady",
					"takefocus",
					"visual",
					"width"}},
			Ttk: &MetaClass{"ttk::labelframe", "TLabelframe",
				[]string{"labelanchor",
					"text",
					"underline",
					"labelwidget",
					"borderwidth",
					"padding",
					"relief",
					"width",
					"height",
					"takefocus",
					"cursor",
					"style",
					"class"}},
		}
	typeMetaMap[WidgetTypeListBox] =
		&MetaType{
			Type: "ListBox",
			Tk: &MetaClass{"tk::listbox", "Listbox",
				[]string{"activestyle",
					"background",
					"borderwidth",
					"cursor",
					"disabledforeground",
					"exportselection",
					"font",
					"foreground",
					"height",
					"highlightbackground",
					"highlightcolor",
					"highlightthickness",
					"justify",
					"relief",
					"selectbackground",
					"selectborderwidth",
					"selectforeground",
					"selectmode",
					"setgrid",
					"state",
					"takefocus",
					"width",
					"xscrollcommand",
					"yscrollcommand",
					"listvariable"}},
			Ttk: nil,
		}
	typeMetaMap[WidgetTypeMenu] =
		&MetaType{
			Type: "Menu",
			Tk: &MetaClass{"menu", "Menu",
				[]string{"activebackground",
					"activeborderwidth",
					"activeforeground",
					"background",
					"borderwidth",
					"cursor",
					"disabledforeground",
					"font",
					"foreground",
					"postcommand",
					"relief",
					"selectcolor",
					"takefocus",
					"tearoff",
					"tearoffcommand",
					"title",
					"type"}},
			Ttk: nil,
		}
	typeMetaMap[WidgetTypeMenuButton] =
		&MetaType{
			Type: "MenuButton",
			Tk: &MetaClass{"tk::menubutton", "Menubutton",
				[]string{"activebackground",
					"activeforeground",
					"anchor",
					"background",
					"bitmap",
					"borderwidth",
					"cursor",
					"direction",
					"disabledforeground",
					"font",
					"foreground",
					"height",
					"highlightbackground",
					"highlightcolor",
					"highlightthickness",
					"image",
					"indicatoron",
					"justify",
					"menu",
					"padx",
					"pady",
					"relief",
					"compound",
					"state",
					"takefocus",
					"text",
					"textvariable",
					"underline",
					"width",
					"wraplength"}},
			Ttk: &MetaClass{"ttk::menubutton", "TMenubutton",
				[]string{"menu",
					"direction",
					"takefocus",
					"text",
					"textvariable",
					"underline",
					"width",
					"image",
					"compound",
					"padding",
					"state",
					"cursor",
					"style",
					"class"}},
		}
	typeMetaMap[WidgetTypeNoteBook] =
		&MetaType{
			Type: "NoteBook",
			Tk:   nil,
			Ttk: &MetaClass{"ttk::notebook", "TNotebook",
				[]string{"width",
					"height",
					"padding",
					"takefocus",
					"cursor",
					"style",
					"class"}},
		}
	typeMetaMap[WidgetTypePanedWindow] =
		&MetaType{
			Type: "PanedWindow",
			Tk: &MetaClass{"tk::panedwindow", "Panedwindow",
				[]string{"background",
					"borderwidth",
					"cursor",
					"handlepad",
					"handlesize",
					"height",
					"opaqueresize",
					"orient",
					"proxybackground",
					"proxyborderwidth",
					"proxyrelief",
					"relief",
					"sashcursor",
					"sashpad",
					"sashrelief",
					"sashwidth",
					"showhandle",
					"width"}},
			Ttk: &MetaClass{"ttk::panedwindow", "TPanedwindow",
				[]string{"orient",
					"width",
					"height",
					"takefocus",
					"cursor",
					"style",
					"class"}},
		}
	typeMetaMap[WidgetTypeProgressBar] =
		&MetaType{
			Type: "ProgressBar",
			Tk:   nil,
			Ttk: &MetaClass{"ttk::progressbar", "TProgressbar",
				[]string{"orient",
					"length",
					"mode",
					"maximum",
					"variable",
					"value",
					"phase",
					"takefocus",
					"cursor",
					"style",
					"class"}},
		}
	typeMetaMap[WidgetTypeRadioButton] =
		&MetaType{
			Type: "RadioButton",
			Tk: &MetaClass{"tk::radiobutton", "Radiobutton",
				[]string{"activebackground",
					"activeforeground",
					"anchor",
					"background",
					"bitmap",
					"borderwidth",
					"command",
					"compound",
					"cursor",
					"disabledforeground",
					"font",
					"foreground",
					"height",
					"highlightbackground",
					"highlightcolor",
					"highlightthickness",
					"image",
					"indicatoron",
					"justify",
					"offrelief",
					"overrelief",
					"padx",
					"pady",
					"relief",
					"selectcolor",
					"selectimage",
					"state",
					"takefocus",
					"text",
					"textvariable",
					"tristateimage",
					"tristatevalue",
					"underline",
					"value",
					"variable",
					"width",
					"wraplength"}},
			Ttk: &MetaClass{"ttk::radiobutton", "TRadiobutton",
				[]string{"variable",
					"value",
					"command",
					"takefocus",
					"text",
					"textvariable",
					"underline",
					"width",
					"image",
					"compound",
					"padding",
					"state",
					"cursor",
					"style",
					"class"}},
		}
	typeMetaMap[WidgetTypeScale] =
		&MetaType{
			Type: "Scale",
			Tk: &MetaClass{"tk::scale", "Scale",
				[]string{"activebackground",
					"background",
					"bigincrement",
					"borderwidth",
					"command",
					"cursor",
					"digits",
					"font",
					"foreground",
					"from",
					"highlightbackground",
					"highlightcolor",
					"highlightthickness",
					"label",
					"length",
					"orient",
					"relief",
					"repeatdelay",
					"repeatinterval",
					"resolution",
					"showvalue",
					"sliderlength",
					"sliderrelief",
					"state",
					"takefocus",
					"tickinterval",
					"to",
					"troughcolor",
					"variable",
					"width"}},
			Ttk: &MetaClass{"ttk::scale", "TScale",
				[]string{"command",
					"variable",
					"orient",
					"from",
					"to",
					"value",
					"length",
					"takefocus",
					"cursor",
					"style",
					"class"}},
		}
	typeMetaMap[WidgetTypeScrollBar] =
		&MetaType{
			Type: "ScrollBar",
			Tk: &MetaClass{"tk::scrollbar", "Scrollbar",
				[]string{"activebackground",
					"activerelief",
					"background",
					"borderwidth",
					"command",
					"cursor",
					"elementborderwidth",
					"highlightbackground",
					"highlightcolor",
					"highlightthickness",
					"jump",
					"orient",
					"relief",
					"repeatdelay",
					"repeatinterval",
					"takefocus",
					"troughcolor",
					"width"}},
			Ttk: &MetaClass{"ttk::scrollbar", "Scrollbar",
				[]string{"command",
					"orient",
					"takefocus",
					"cursor",
					"style",
					"class"}},
		}
	typeMetaMap[WidgetTypeSeparator] =
		&MetaType{
			Type: "Separator",
			Tk:   nil,
			Ttk: &MetaClass{"ttk::separator", "TSeparator",
				[]string{"orient",
					"takefocus",
					"cursor",
					"style",
					"class"}},
		}
	typeMetaMap[WidgetTypeSizeGrip] =
		&MetaType{
			Type: "SizeGrip",
			Tk:   nil,
			Ttk: &MetaClass{"ttk::sizegrip", "TSizegrip",
				[]string{"takefocus",
					"cursor",
					"style",
					"class"}},
		}
	typeMetaMap[WidgetTypeSpinBox] =
		&MetaType{
			Type: "SpinBox",
			Tk: &MetaClass{"tk::spinbox", "Spinbox",
				[]string{"activebackground",
					"background",
					"borderwidth",
					"buttonbackground",
					"buttoncursor",
					"buttondownrelief",
					"buttonuprelief",
					"command",
					"cursor",
					"disabledbackground",
					"disabledforeground",
					"exportselection",
					"font",
					"foreground",
					"format",
					"from",
					"highlightbackground",
					"highlightcolor",
					"highlightthickness",
					"increment",
					"insertbackground",
					"insertborderwidth",
					"insertofftime",
					"insertontime",
					"insertwidth",
					"invalidcommand",
					"justify",
					"relief",
					"readonlybackground",
					"repeatdelay",
					"repeatinterval",
					"selectbackground",
					"selectborderwidth",
					"selectforeground",
					"state",
					"takefocus",
					"textvariable",
					"to",
					"validate",
					"validatecommand",
					"values",
					"width",
					"wrap",
					"xscrollcommand"}},
			Ttk: nil,
		}
	typeMetaMap[WidgetTypeTextEdit] =
		&MetaType{
			Type: "TextEdit",
			Tk: &MetaClass{"tk::text", "Text",
				[]string{"autoseparators",
					"background",
					"blockcursor",
					"borderwidth",
					"cursor",
					"endline",
					"exportselection",
					"font",
					"foreground",
					"height",
					"highlightbackground",
					"highlightcolor",
					"highlightthickness",
					"inactiveselectbackground",
					"insertbackground",
					"insertborderwidth",
					"insertofftime",
					"insertontime",
					"insertunfocussed",
					"insertwidth",
					"maxundo",
					"padx",
					"pady",
					"relief",
					"selectbackground",
					"selectborderwidth",
					"selectforeground",
					"setgrid",
					"spacing1",
					"spacing2",
					"spacing3",
					"startline",
					"state",
					"tabs",
					"tabstyle",
					"takefocus",
					"undo",
					"width",
					"wrap",
					"xscrollcommand",
					"yscrollcommand"}},
			Ttk: nil,
		}
	typeMetaMap[WidgetTypeWindow] =
		&MetaType{
			Type: "Window",
			Tk: &MetaClass{"toplevel", "Toplevel",
				[]string{"borderwidth",
					"class",
					"menu",
					"relief",
					"screen",
					"use",
					"background",
					"colormap",
					"container",
					"cursor",
					"height",
					"highlightbackground",
					"highlightcolor",
					"highlightthickness",
					"padx",
					"pady",
					"takefocus",
					"visual",
					"width"}},
			Ttk: nil,
		}
	typeMetaMap[WidgetTypeTreeView] =
		&MetaType{
			Type: "TreeView",
			Tk:   nil,
			Ttk: &MetaClass{"ttk::treeview", "Treeview",
				[]string{"columns",
					"displaycolumns",
					"show",
					"selectmode",
					"height",
					"padding",
					"xscrollcommand",
					"yscrollcommand",
					"takefocus",
					"cursor",
					"style",
					"class"}},
		}
}
