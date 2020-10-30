//----------------------------------------
// The code is automatically generated by the GenlibLcl tool.
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package vcl

import (
	. "github.com/dollarkillerx/govcl/vcl/api"
	. "github.com/dollarkillerx/govcl/vcl/types"
	"unsafe"
)

type TGridColumn struct {
	IObject
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

// 动态转换一个已存在的对象实例。
//
// Dynamically convert an existing object instance.
func AsGridColumn(obj interface{}) *TGridColumn {
	instance, ptr := getInstance(obj)
	if instance == 0 {
		return nil
	}
	return &TGridColumn{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// 新建一个对象来自已经存在的对象实例指针。
//
// Create a new object from an existing object instance pointer.
// Deprecated: use AsGridColumn.
func GridColumnFromInst(inst uintptr) *TGridColumn {
	return AsGridColumn(inst)
}

// 新建一个对象来自已经存在的对象实例。
//
// Create a new object from an existing object instance.
// Deprecated: use AsGridColumn.
func GridColumnFromObj(obj IObject) *TGridColumn {
	return AsGridColumn(obj)
}

// 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
//
// Create a new object from an unsecured address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsGridColumn.
func GridColumnFromUnsafePointer(ptr unsafe.Pointer) *TGridColumn {
	return AsGridColumn(ptr)
}

// -------------------------- Deprecated end --------------------------
// 返回对象实例指针。
//
// Return object instance pointer.
func (g *TGridColumn) Instance() uintptr {
	return g.instance
}

// 获取一个不安全的地址。
//
// Get an unsafe address.
func (g *TGridColumn) UnsafeAddr() unsafe.Pointer {
	return g.ptr
}

// 检测地址是否为空。
//
// Check if the address is empty.
func (g *TGridColumn) IsValid() bool {
	return g.instance != 0
}

// 检测当前对象是否继承自目标对象。
//
// Checks whether the current object is inherited from the target object.
func (g *TGridColumn) Is() TIs {
	return TIs(g.instance)
}

// 动态转换当前对象为目标对象。
//
// Dynamically convert the current object to the target object.
//func (g *TGridColumn) As() TAs {
//    return TAs(g.instance)
//}

// 获取类信息指针。
//
// Get class information pointer.
func TGridColumnClass() TClass {
	return GridColumn_StaticClassType()
}

// 复制一个对象，如果对象实现了此方法的话。
//
// Copy an object, if the object implements this method.
func (g *TGridColumn) Assign(Source IObject) {
	GridColumn_Assign(g.instance, CheckPtr(Source))
}

func (g *TGridColumn) FixDesignFontsPPI(ADesignTimePPI int32) {
	GridColumn_FixDesignFontsPPI(g.instance, ADesignTimePPI)
}

func (g *TGridColumn) ScaleFontsPPI(AToPPI int32, AProportion float64) {
	GridColumn_ScaleFontsPPI(g.instance, AToPPI, AProportion)
}

func (g *TGridColumn) IsDefault() bool {
	return GridColumn_IsDefault(g.instance)
}

// 获取类名路径。
//
// Get the class name path.
func (g *TGridColumn) GetNamePath() string {
	return GridColumn_GetNamePath(g.instance)
}

// 获取类的类型信息。
//
// Get class type information.
func (g *TGridColumn) ClassType() TClass {
	return GridColumn_ClassType(g.instance)
}

// 获取当前对象类名称。
//
// Get the current object class name.
func (g *TGridColumn) ClassName() string {
	return GridColumn_ClassName(g.instance)
}

// 获取当前对象实例大小。
//
// Get the current object instance size.
func (g *TGridColumn) InstanceSize() int32 {
	return GridColumn_InstanceSize(g.instance)
}

// 判断当前类是否继承自指定类。
//
// Determine whether the current class inherits from the specified class.
func (g *TGridColumn) InheritsFrom(AClass TClass) bool {
	return GridColumn_InheritsFrom(g.instance, AClass)
}

// 与一个对象进行比较。
//
// Compare with an object.
func (g *TGridColumn) Equals(Obj IObject) bool {
	return GridColumn_Equals(g.instance, CheckPtr(Obj))
}

// 获取类的哈希值。
//
// Get the hash value of the class.
func (g *TGridColumn) GetHashCode() int32 {
	return GridColumn_GetHashCode(g.instance)
}

// 文本类信息。
//
// Text information.
func (g *TGridColumn) ToString() string {
	return GridColumn_ToString(g.instance)
}

func (g *TGridColumn) Grid() *TStringGrid {
	return AsStringGrid(GridColumn_GetGrid(g.instance))
}

func (g *TGridColumn) DefaultWidth() int32 {
	return GridColumn_GetDefaultWidth(g.instance)
}

func (g *TGridColumn) StoredWidth() int32 {
	return GridColumn_GetStoredWidth(g.instance)
}

func (g *TGridColumn) WidthChanged() bool {
	return GridColumn_GetWidthChanged(g.instance)
}

// 获取文字对齐。
//
// Get Text alignment.
func (g *TGridColumn) Alignment() TAlignment {
	return GridColumn_GetAlignment(g.instance)
}

// 设置文字对齐。
//
// Set Text alignment.
func (g *TGridColumn) SetAlignment(value TAlignment) {
	GridColumn_SetAlignment(g.instance, value)
}

func (g *TGridColumn) ButtonStyle() TColumnButtonStyle {
	return GridColumn_GetButtonStyle(g.instance)
}

func (g *TGridColumn) SetButtonStyle(value TColumnButtonStyle) {
	GridColumn_SetButtonStyle(g.instance, value)
}

// 获取颜色。
//
// Get color.
func (g *TGridColumn) Color() TColor {
	return GridColumn_GetColor(g.instance)
}

// 设置颜色。
//
// Set color.
func (g *TGridColumn) SetColor(value TColor) {
	GridColumn_SetColor(g.instance, value)
}

func (g *TGridColumn) DropDownRows() int32 {
	return GridColumn_GetDropDownRows(g.instance)
}

func (g *TGridColumn) SetDropDownRows(value int32) {
	GridColumn_SetDropDownRows(g.instance, value)
}

func (g *TGridColumn) Expanded() bool {
	return GridColumn_GetExpanded(g.instance)
}

func (g *TGridColumn) SetExpanded(value bool) {
	GridColumn_SetExpanded(g.instance, value)
}

// 获取字体。
//
// Get Font.
func (g *TGridColumn) Font() *TFont {
	return AsFont(GridColumn_GetFont(g.instance))
}

// 设置字体。
//
// Set Font.
func (g *TGridColumn) SetFont(value *TFont) {
	GridColumn_SetFont(g.instance, CheckPtr(value))
}

func (g *TGridColumn) Layout() TTextLayout {
	return GridColumn_GetLayout(g.instance)
}

func (g *TGridColumn) SetLayout(value TTextLayout) {
	GridColumn_SetLayout(g.instance, value)
}

func (g *TGridColumn) MinSize() int32 {
	return GridColumn_GetMinSize(g.instance)
}

func (g *TGridColumn) SetMinSize(value int32) {
	GridColumn_SetMinSize(g.instance, value)
}

func (g *TGridColumn) MaxSize() int32 {
	return GridColumn_GetMaxSize(g.instance)
}

func (g *TGridColumn) SetMaxSize(value int32) {
	GridColumn_SetMaxSize(g.instance, value)
}

func (g *TGridColumn) PickList() *TStrings {
	return AsStrings(GridColumn_GetPickList(g.instance))
}

func (g *TGridColumn) SetPickList(value IStrings) {
	GridColumn_SetPickList(g.instance, CheckPtr(value))
}

// 获取只读。
func (g *TGridColumn) ReadOnly() bool {
	return GridColumn_GetReadOnly(g.instance)
}

// 设置只读。
func (g *TGridColumn) SetReadOnly(value bool) {
	GridColumn_SetReadOnly(g.instance, value)
}

func (g *TGridColumn) SizePriority() int32 {
	return GridColumn_GetSizePriority(g.instance)
}

func (g *TGridColumn) SetSizePriority(value int32) {
	GridColumn_SetSizePriority(g.instance, value)
}

// 获取对象标记。
//
// Get the control tag.
func (g *TGridColumn) Tag() int {
	return GridColumn_GetTag(g.instance)
}

// 设置对象标记。
//
// Set the control tag.
func (g *TGridColumn) SetTag(value int) {
	GridColumn_SetTag(g.instance, value)
}

func (g *TGridColumn) Title() *TGridColumnTitle {
	return AsGridColumnTitle(GridColumn_GetTitle(g.instance))
}

func (g *TGridColumn) SetTitle(value *TGridColumnTitle) {
	GridColumn_SetTitle(g.instance, CheckPtr(value))
}

// 获取宽度。
//
// Get width.
func (g *TGridColumn) Width() int32 {
	return GridColumn_GetWidth(g.instance)
}

// 设置宽度。
//
// Set width.
func (g *TGridColumn) SetWidth(value int32) {
	GridColumn_SetWidth(g.instance, value)
}

// 获取控件可视。
//
// Get the control visible.
func (g *TGridColumn) Visible() bool {
	return GridColumn_GetVisible(g.instance)
}

// 设置控件可视。
//
// Set the control visible.
func (g *TGridColumn) SetVisible(value bool) {
	GridColumn_SetVisible(g.instance, value)
}

func (g *TGridColumn) ValueChecked() string {
	return GridColumn_GetValueChecked(g.instance)
}

func (g *TGridColumn) SetValueChecked(value string) {
	GridColumn_SetValueChecked(g.instance, value)
}

func (g *TGridColumn) ValueUnchecked() string {
	return GridColumn_GetValueUnchecked(g.instance)
}

func (g *TGridColumn) SetValueUnchecked(value string) {
	GridColumn_SetValueUnchecked(g.instance, value)
}

func (g *TGridColumn) Collection() *TCollection {
	return AsCollection(GridColumn_GetCollection(g.instance))
}

func (g *TGridColumn) SetCollection(value *TCollection) {
	GridColumn_SetCollection(g.instance, CheckPtr(value))
}

func (g *TGridColumn) Index() int32 {
	return GridColumn_GetIndex(g.instance)
}

func (g *TGridColumn) SetIndex(value int32) {
	GridColumn_SetIndex(g.instance, value)
}

func (g *TGridColumn) DisplayName() string {
	return GridColumn_GetDisplayName(g.instance)
}

func (g *TGridColumn) SetDisplayName(value string) {
	GridColumn_SetDisplayName(g.instance, value)
}
