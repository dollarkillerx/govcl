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

type TToggleBox struct {
	IWinControl
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

// 创建一个新的对象。
//
// Create a new object.
func NewToggleBox(owner IComponent) *TToggleBox {
	t := new(TToggleBox)
	t.instance = ToggleBox_Create(CheckPtr(owner))
	t.ptr = unsafe.Pointer(t.instance)
	return t
}

// 动态转换一个已存在的对象实例。
//
// Dynamically convert an existing object instance.
func AsToggleBox(obj interface{}) *TToggleBox {
	instance, ptr := getInstance(obj)
	if instance == 0 {
		return nil
	}
	return &TToggleBox{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// 新建一个对象来自已经存在的对象实例指针。
//
// Create a new object from an existing object instance pointer.
// Deprecated: use AsToggleBox.
func ToggleBoxFromInst(inst uintptr) *TToggleBox {
	return AsToggleBox(inst)
}

// 新建一个对象来自已经存在的对象实例。
//
// Create a new object from an existing object instance.
// Deprecated: use AsToggleBox.
func ToggleBoxFromObj(obj IObject) *TToggleBox {
	return AsToggleBox(obj)
}

// 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
//
// Create a new object from an unsecured address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsToggleBox.
func ToggleBoxFromUnsafePointer(ptr unsafe.Pointer) *TToggleBox {
	return AsToggleBox(ptr)
}

// -------------------------- Deprecated end --------------------------
// 释放对象。
//
// Free object.
func (t *TToggleBox) Free() {
	if t.instance != 0 {
		ToggleBox_Free(t.instance)
		t.instance, t.ptr = 0, nullptr
	}
}

// 返回对象实例指针。
//
// Return object instance pointer.
func (t *TToggleBox) Instance() uintptr {
	return t.instance
}

// 获取一个不安全的地址。
//
// Get an unsafe address.
func (t *TToggleBox) UnsafeAddr() unsafe.Pointer {
	return t.ptr
}

// 检测地址是否为空。
//
// Check if the address is empty.
func (t *TToggleBox) IsValid() bool {
	return t.instance != 0
}

// 检测当前对象是否继承自目标对象。
//
// Checks whether the current object is inherited from the target object.
func (t *TToggleBox) Is() TIs {
	return TIs(t.instance)
}

// 动态转换当前对象为目标对象。
//
// Dynamically convert the current object to the target object.
//func (t *TToggleBox) As() TAs {
//    return TAs(t.instance)
//}

// 获取类信息指针。
//
// Get class information pointer.
func TToggleBoxClass() TClass {
	return ToggleBox_StaticClassType()
}

// 是否可以获得焦点。
func (t *TToggleBox) CanFocus() bool {
	return ToggleBox_CanFocus(t.instance)
}

// 返回是否包含指定控件。
//
// it's contain a specified control.
func (t *TToggleBox) ContainsControl(Control IControl) bool {
	return ToggleBox_ContainsControl(t.instance, CheckPtr(Control))
}

// 返回指定坐标及相关属性位置控件。
//
// Returns the specified coordinate and the relevant attribute position control..
func (t *TToggleBox) ControlAtPos(Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) *TControl {
	return AsControl(ToggleBox_ControlAtPos(t.instance, Pos, AllowDisabled, AllowWinControls, AllLevels))
}

// 禁用控件的对齐。
//
// Disable control alignment.
func (t *TToggleBox) DisableAlign() {
	ToggleBox_DisableAlign(t.instance)
}

// 启用控件对齐。
//
// Enabled control alignment.
func (t *TToggleBox) EnableAlign() {
	ToggleBox_EnableAlign(t.instance)
}

// 查找子控件。
//
// Find sub controls.
func (t *TToggleBox) FindChildControl(ControlName string) *TControl {
	return AsControl(ToggleBox_FindChildControl(t.instance, ControlName))
}

func (t *TToggleBox) FlipChildren(AllLevels bool) {
	ToggleBox_FlipChildren(t.instance, AllLevels)
}

// 返回是否获取焦点。
//
// Return to get focus.
func (t *TToggleBox) Focused() bool {
	return ToggleBox_Focused(t.instance)
}

// 句柄是否已经分配。
//
// Is the handle already allocated.
func (t *TToggleBox) HandleAllocated() bool {
	return ToggleBox_HandleAllocated(t.instance)
}

// 插入一个控件。
//
// Insert a control.
func (t *TToggleBox) InsertControl(AControl IControl) {
	ToggleBox_InsertControl(t.instance, CheckPtr(AControl))
}

// 要求重绘。
//
// Redraw.
func (t *TToggleBox) Invalidate() {
	ToggleBox_Invalidate(t.instance)
}

// 绘画至指定DC。
//
// Painting to the specified DC.
func (t *TToggleBox) PaintTo(DC HDC, X int32, Y int32) {
	ToggleBox_PaintTo(t.instance, DC, X, Y)
}

// 移除一个控件。
//
// Remove a control.
func (t *TToggleBox) RemoveControl(AControl IControl) {
	ToggleBox_RemoveControl(t.instance, CheckPtr(AControl))
}

// 重新对齐。
//
// Realign.
func (t *TToggleBox) Realign() {
	ToggleBox_Realign(t.instance)
}

// 重绘。
//
// Repaint.
func (t *TToggleBox) Repaint() {
	ToggleBox_Repaint(t.instance)
}

// 按比例缩放。
//
// Scale by.
func (t *TToggleBox) ScaleBy(M int32, D int32) {
	ToggleBox_ScaleBy(t.instance, M, D)
}

// 滚动至指定位置。
//
// Scroll by.
func (t *TToggleBox) ScrollBy(DeltaX int32, DeltaY int32) {
	ToggleBox_ScrollBy(t.instance, DeltaX, DeltaY)
}

// 设置组件边界。
//
// Set component boundaries.
func (t *TToggleBox) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
	ToggleBox_SetBounds(t.instance, ALeft, ATop, AWidth, AHeight)
}

// 设置控件焦点。
//
// Set control focus.
func (t *TToggleBox) SetFocus() {
	ToggleBox_SetFocus(t.instance)
}

// 控件更新。
//
// Update.
func (t *TToggleBox) Update() {
	ToggleBox_Update(t.instance)
}

// 将控件置于最前。
//
// Bring the control to the front.
func (t *TToggleBox) BringToFront() {
	ToggleBox_BringToFront(t.instance)
}

// 将客户端坐标转为绝对的屏幕坐标。
//
// Convert client coordinates to absolute screen coordinates.
func (t *TToggleBox) ClientToScreen(Point TPoint) TPoint {
	return ToggleBox_ClientToScreen(t.instance, Point)
}

// 将客户端坐标转为父容器坐标。
//
// Convert client coordinates to parent container coordinates.
func (t *TToggleBox) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
	return ToggleBox_ClientToParent(t.instance, Point, CheckPtr(AParent))
}

// 是否在拖拽中。
//
// Is it in the middle of dragging.
func (t *TToggleBox) Dragging() bool {
	return ToggleBox_Dragging(t.instance)
}

// 是否有父容器。
//
// Is there a parent container.
func (t *TToggleBox) HasParent() bool {
	return ToggleBox_HasParent(t.instance)
}

// 隐藏控件。
//
// Hidden control.
func (t *TToggleBox) Hide() {
	ToggleBox_Hide(t.instance)
}

// 发送一个消息。
//
// Send a message.
func (t *TToggleBox) Perform(Msg uint32, WParam uintptr, LParam int) int {
	return ToggleBox_Perform(t.instance, Msg, WParam, LParam)
}

// 刷新控件。
//
// Refresh control.
func (t *TToggleBox) Refresh() {
	ToggleBox_Refresh(t.instance)
}

// 将屏幕坐标转为客户端坐标。
//
// Convert screen coordinates to client coordinates.
func (t *TToggleBox) ScreenToClient(Point TPoint) TPoint {
	return ToggleBox_ScreenToClient(t.instance, Point)
}

// 将父容器坐标转为客户端坐标。
//
// Convert parent container coordinates to client coordinates.
func (t *TToggleBox) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
	return ToggleBox_ParentToClient(t.instance, Point, CheckPtr(AParent))
}

// 控件至于最后面。
//
// The control is placed at the end.
func (t *TToggleBox) SendToBack() {
	ToggleBox_SendToBack(t.instance)
}

// 显示控件。
//
// Show control.
func (t *TToggleBox) Show() {
	ToggleBox_Show(t.instance)
}

// 获取控件的字符，如果有。
//
// Get the characters of the control, if any.
func (t *TToggleBox) GetTextBuf(Buffer *string, BufSize int32) int32 {
	return ToggleBox_GetTextBuf(t.instance, Buffer, BufSize)
}

// 获取控件的字符长，如果有。
//
// Get the character length of the control, if any.
func (t *TToggleBox) GetTextLen() int32 {
	return ToggleBox_GetTextLen(t.instance)
}

// 设置控件字符，如果有。
//
// Set control characters, if any.
func (t *TToggleBox) SetTextBuf(Buffer string) {
	ToggleBox_SetTextBuf(t.instance, Buffer)
}

// 查找指定名称的组件。
//
// Find the component with the specified name.
func (t *TToggleBox) FindComponent(AName string) *TComponent {
	return AsComponent(ToggleBox_FindComponent(t.instance, AName))
}

// 获取类名路径。
//
// Get the class name path.
func (t *TToggleBox) GetNamePath() string {
	return ToggleBox_GetNamePath(t.instance)
}

// 复制一个对象，如果对象实现了此方法的话。
//
// Copy an object, if the object implements this method.
func (t *TToggleBox) Assign(Source IObject) {
	ToggleBox_Assign(t.instance, CheckPtr(Source))
}

// 获取类的类型信息。
//
// Get class type information.
func (t *TToggleBox) ClassType() TClass {
	return ToggleBox_ClassType(t.instance)
}

// 获取当前对象类名称。
//
// Get the current object class name.
func (t *TToggleBox) ClassName() string {
	return ToggleBox_ClassName(t.instance)
}

// 获取当前对象实例大小。
//
// Get the current object instance size.
func (t *TToggleBox) InstanceSize() int32 {
	return ToggleBox_InstanceSize(t.instance)
}

// 判断当前类是否继承自指定类。
//
// Determine whether the current class inherits from the specified class.
func (t *TToggleBox) InheritsFrom(AClass TClass) bool {
	return ToggleBox_InheritsFrom(t.instance, AClass)
}

// 与一个对象进行比较。
//
// Compare with an object.
func (t *TToggleBox) Equals(Obj IObject) bool {
	return ToggleBox_Equals(t.instance, CheckPtr(Obj))
}

// 获取类的哈希值。
//
// Get the hash value of the class.
func (t *TToggleBox) GetHashCode() int32 {
	return ToggleBox_GetHashCode(t.instance)
}

// 文本类信息。
//
// Text information.
func (t *TToggleBox) ToString() string {
	return ToggleBox_ToString(t.instance)
}

func (t *TToggleBox) AnchorToNeighbour(ASide TAnchorKind, ASpace int32, ASibling IControl) {
	ToggleBox_AnchorToNeighbour(t.instance, ASide, ASpace, CheckPtr(ASibling))
}

func (t *TToggleBox) AnchorParallel(ASide TAnchorKind, ASpace int32, ASibling IControl) {
	ToggleBox_AnchorParallel(t.instance, ASide, ASpace, CheckPtr(ASibling))
}

// 置于指定控件的横向中心。
func (t *TToggleBox) AnchorHorizontalCenterTo(ASibling IControl) {
	ToggleBox_AnchorHorizontalCenterTo(t.instance, CheckPtr(ASibling))
}

// 置于指定控件的纵向中心。
func (t *TToggleBox) AnchorVerticalCenterTo(ASibling IControl) {
	ToggleBox_AnchorVerticalCenterTo(t.instance, CheckPtr(ASibling))
}

func (t *TToggleBox) AnchorSame(ASide TAnchorKind, ASibling IControl) {
	ToggleBox_AnchorSame(t.instance, ASide, CheckPtr(ASibling))
}

func (t *TToggleBox) AnchorAsAlign(ATheAlign TAlign, ASpace int32) {
	ToggleBox_AnchorAsAlign(t.instance, ATheAlign, ASpace)
}

func (t *TToggleBox) AnchorClient(ASpace int32) {
	ToggleBox_AnchorClient(t.instance, ASpace)
}

func (t *TToggleBox) ScaleDesignToForm(ASize int32) int32 {
	return ToggleBox_ScaleDesignToForm(t.instance, ASize)
}

func (t *TToggleBox) ScaleFormToDesign(ASize int32) int32 {
	return ToggleBox_ScaleFormToDesign(t.instance, ASize)
}

func (t *TToggleBox) Scale96ToForm(ASize int32) int32 {
	return ToggleBox_Scale96ToForm(t.instance, ASize)
}

func (t *TToggleBox) ScaleFormTo96(ASize int32) int32 {
	return ToggleBox_ScaleFormTo96(t.instance, ASize)
}

func (t *TToggleBox) Scale96ToFont(ASize int32) int32 {
	return ToggleBox_Scale96ToFont(t.instance, ASize)
}

func (t *TToggleBox) ScaleFontTo96(ASize int32) int32 {
	return ToggleBox_ScaleFontTo96(t.instance, ASize)
}

func (t *TToggleBox) ScaleScreenToFont(ASize int32) int32 {
	return ToggleBox_ScaleScreenToFont(t.instance, ASize)
}

func (t *TToggleBox) ScaleFontToScreen(ASize int32) int32 {
	return ToggleBox_ScaleFontToScreen(t.instance, ASize)
}

func (t *TToggleBox) Scale96ToScreen(ASize int32) int32 {
	return ToggleBox_Scale96ToScreen(t.instance, ASize)
}

func (t *TToggleBox) ScaleScreenTo96(ASize int32) int32 {
	return ToggleBox_ScaleScreenTo96(t.instance, ASize)
}

func (t *TToggleBox) AutoAdjustLayout(AMode TLayoutAdjustmentPolicy, AFromPPI int32, AToPPI int32, AOldFormWidth int32, ANewFormWidth int32) {
	ToggleBox_AutoAdjustLayout(t.instance, AMode, AFromPPI, AToPPI, AOldFormWidth, ANewFormWidth)
}

func (t *TToggleBox) FixDesignFontsPPI(ADesignTimePPI int32) {
	ToggleBox_FixDesignFontsPPI(t.instance, ADesignTimePPI)
}

func (t *TToggleBox) ScaleFontsPPI(AToPPI int32, AProportion float64) {
	ToggleBox_ScaleFontsPPI(t.instance, AToPPI, AProportion)
}

func (t *TToggleBox) AllowGrayed() bool {
	return ToggleBox_GetAllowGrayed(t.instance)
}

func (t *TToggleBox) SetAllowGrayed(value bool) {
	ToggleBox_SetAllowGrayed(t.instance, value)
}

// 获取控件自动调整。
//
// Get Control automatically adjusts.
func (t *TToggleBox) Align() TAlign {
	return ToggleBox_GetAlign(t.instance)
}

// 设置控件自动调整。
//
// Set Control automatically adjusts.
func (t *TToggleBox) SetAlign(value TAlign) {
	ToggleBox_SetAlign(t.instance, value)
}

// 获取四个角位置的锚点。
func (t *TToggleBox) Anchors() TAnchors {
	return ToggleBox_GetAnchors(t.instance)
}

// 设置四个角位置的锚点。
func (t *TToggleBox) SetAnchors(value TAnchors) {
	ToggleBox_SetAnchors(t.instance, value)
}

// 获取自动调整大小。
func (t *TToggleBox) AutoSize() bool {
	return ToggleBox_GetAutoSize(t.instance)
}

// 设置自动调整大小。
func (t *TToggleBox) SetAutoSize(value bool) {
	ToggleBox_SetAutoSize(t.instance, value)
}

// 获取控件标题。
//
// Get the control title.
func (t *TToggleBox) Caption() string {
	return ToggleBox_GetCaption(t.instance)
}

// 设置控件标题。
//
// Set the control title.
func (t *TToggleBox) SetCaption(value string) {
	ToggleBox_SetCaption(t.instance, value)
}

// 获取是否选中。
func (t *TToggleBox) Checked() bool {
	return ToggleBox_GetChecked(t.instance)
}

// 设置是否选中。
func (t *TToggleBox) SetChecked(value bool) {
	ToggleBox_SetChecked(t.instance, value)
}

// 获取颜色。
//
// Get color.
func (t *TToggleBox) Color() TColor {
	return ToggleBox_GetColor(t.instance)
}

// 设置颜色。
//
// Set color.
func (t *TToggleBox) SetColor(value TColor) {
	ToggleBox_SetColor(t.instance, value)
}

// 获取约束控件大小。
func (t *TToggleBox) Constraints() *TSizeConstraints {
	return AsSizeConstraints(ToggleBox_GetConstraints(t.instance))
}

// 设置约束控件大小。
func (t *TToggleBox) SetConstraints(value *TSizeConstraints) {
	ToggleBox_SetConstraints(t.instance, CheckPtr(value))
}

// 获取设置控件双缓冲。
//
// Get Set control double buffering.
func (t *TToggleBox) DoubleBuffered() bool {
	return ToggleBox_GetDoubleBuffered(t.instance)
}

// 设置设置控件双缓冲。
//
// Set Set control double buffering.
func (t *TToggleBox) SetDoubleBuffered(value bool) {
	ToggleBox_SetDoubleBuffered(t.instance, value)
}

// 获取设置控件拖拽时的光标。
//
// Get Set the cursor when the control is dragged.
func (t *TToggleBox) DragCursor() TCursor {
	return ToggleBox_GetDragCursor(t.instance)
}

// 设置设置控件拖拽时的光标。
//
// Set Set the cursor when the control is dragged.
func (t *TToggleBox) SetDragCursor(value TCursor) {
	ToggleBox_SetDragCursor(t.instance, value)
}

// 获取拖拽方式。
//
// Get Drag and drop.
func (t *TToggleBox) DragKind() TDragKind {
	return ToggleBox_GetDragKind(t.instance)
}

// 设置拖拽方式。
//
// Set Drag and drop.
func (t *TToggleBox) SetDragKind(value TDragKind) {
	ToggleBox_SetDragKind(t.instance, value)
}

// 获取拖拽模式。
//
// Get Drag mode.
func (t *TToggleBox) DragMode() TDragMode {
	return ToggleBox_GetDragMode(t.instance)
}

// 设置拖拽模式。
//
// Set Drag mode.
func (t *TToggleBox) SetDragMode(value TDragMode) {
	ToggleBox_SetDragMode(t.instance, value)
}

// 获取控件启用。
//
// Get the control enabled.
func (t *TToggleBox) Enabled() bool {
	return ToggleBox_GetEnabled(t.instance)
}

// 设置控件启用。
//
// Set the control enabled.
func (t *TToggleBox) SetEnabled(value bool) {
	ToggleBox_SetEnabled(t.instance, value)
}

// 获取字体。
//
// Get Font.
func (t *TToggleBox) Font() *TFont {
	return AsFont(ToggleBox_GetFont(t.instance))
}

// 设置字体。
//
// Set Font.
func (t *TToggleBox) SetFont(value *TFont) {
	ToggleBox_SetFont(t.instance, CheckPtr(value))
}

// 获取组件鼠标悬停提示。
//
// Get component mouse hints.
func (t *TToggleBox) Hint() string {
	return ToggleBox_GetHint(t.instance)
}

// 设置组件鼠标悬停提示。
//
// Set component mouse hints.
func (t *TToggleBox) SetHint(value string) {
	ToggleBox_SetHint(t.instance, value)
}

// 设置改变事件。
//
// Set changed event.
func (t *TToggleBox) SetOnChange(fn TNotifyEvent) {
	ToggleBox_SetOnChange(t.instance, fn)
}

// 设置控件单击事件。
//
// Set control click event.
func (t *TToggleBox) SetOnClick(fn TNotifyEvent) {
	ToggleBox_SetOnClick(t.instance, fn)
}

// 设置拖拽下落事件。
//
// Set Drag and drop event.
func (t *TToggleBox) SetOnDragDrop(fn TDragDropEvent) {
	ToggleBox_SetOnDragDrop(t.instance, fn)
}

// 设置拖拽完成事件。
//
// Set Drag and drop completion event.
func (t *TToggleBox) SetOnDragOver(fn TDragOverEvent) {
	ToggleBox_SetOnDragOver(t.instance, fn)
}

// 设置拖拽结束。
//
// Set End of drag.
func (t *TToggleBox) SetOnEndDrag(fn TEndDragEvent) {
	ToggleBox_SetOnEndDrag(t.instance, fn)
}

// 设置焦点进入。
//
// Set Focus entry.
func (t *TToggleBox) SetOnEnter(fn TNotifyEvent) {
	ToggleBox_SetOnEnter(t.instance, fn)
}

// 设置焦点退出。
//
// Set Focus exit.
func (t *TToggleBox) SetOnExit(fn TNotifyEvent) {
	ToggleBox_SetOnExit(t.instance, fn)
}

// 设置鼠标按下事件。
//
// Set Mouse down event.
func (t *TToggleBox) SetOnMouseDown(fn TMouseEvent) {
	ToggleBox_SetOnMouseDown(t.instance, fn)
}

// 设置鼠标进入事件。
//
// Set Mouse entry event.
func (t *TToggleBox) SetOnMouseEnter(fn TNotifyEvent) {
	ToggleBox_SetOnMouseEnter(t.instance, fn)
}

// 设置鼠标离开事件。
//
// Set Mouse leave event.
func (t *TToggleBox) SetOnMouseLeave(fn TNotifyEvent) {
	ToggleBox_SetOnMouseLeave(t.instance, fn)
}

// 设置鼠标移动事件。
func (t *TToggleBox) SetOnMouseMove(fn TMouseMoveEvent) {
	ToggleBox_SetOnMouseMove(t.instance, fn)
}

// 设置鼠标抬起事件。
//
// Set Mouse lift event.
func (t *TToggleBox) SetOnMouseUp(fn TMouseEvent) {
	ToggleBox_SetOnMouseUp(t.instance, fn)
}

// 设置鼠标滚轮事件。
func (t *TToggleBox) SetOnMouseWheel(fn TMouseWheelEvent) {
	ToggleBox_SetOnMouseWheel(t.instance, fn)
}

// 设置鼠标滚轮按下事件。
func (t *TToggleBox) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	ToggleBox_SetOnMouseWheelDown(t.instance, fn)
}

// 设置鼠标滚轮抬起事件。
func (t *TToggleBox) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	ToggleBox_SetOnMouseWheelUp(t.instance, fn)
}

// 获取使用父容器双缓冲。
//
// Get Parent container double buffering.
func (t *TToggleBox) ParentDoubleBuffered() bool {
	return ToggleBox_GetParentDoubleBuffered(t.instance)
}

// 设置使用父容器双缓冲。
//
// Set Parent container double buffering.
func (t *TToggleBox) SetParentDoubleBuffered(value bool) {
	ToggleBox_SetParentDoubleBuffered(t.instance, value)
}

// 获取使用父容器字体。
//
// Get Parent container font.
func (t *TToggleBox) ParentFont() bool {
	return ToggleBox_GetParentFont(t.instance)
}

// 设置使用父容器字体。
//
// Set Parent container font.
func (t *TToggleBox) SetParentFont(value bool) {
	ToggleBox_SetParentFont(t.instance, value)
}

// 获取以父容器的ShowHint属性为准。
func (t *TToggleBox) ParentShowHint() bool {
	return ToggleBox_GetParentShowHint(t.instance)
}

// 设置以父容器的ShowHint属性为准。
func (t *TToggleBox) SetParentShowHint(value bool) {
	ToggleBox_SetParentShowHint(t.instance, value)
}

// 获取右键菜单。
//
// Get Right click menu.
func (t *TToggleBox) PopupMenu() *TPopupMenu {
	return AsPopupMenu(ToggleBox_GetPopupMenu(t.instance))
}

// 设置右键菜单。
//
// Set Right click menu.
func (t *TToggleBox) SetPopupMenu(value IComponent) {
	ToggleBox_SetPopupMenu(t.instance, CheckPtr(value))
}

// 获取显示鼠标悬停提示。
//
// Get Show mouseover tips.
func (t *TToggleBox) ShowHint() bool {
	return ToggleBox_GetShowHint(t.instance)
}

// 设置显示鼠标悬停提示。
//
// Set Show mouseover tips.
func (t *TToggleBox) SetShowHint(value bool) {
	ToggleBox_SetShowHint(t.instance, value)
}

func (t *TToggleBox) State() TCheckBoxState {
	return ToggleBox_GetState(t.instance)
}

func (t *TToggleBox) SetState(value TCheckBoxState) {
	ToggleBox_SetState(t.instance, value)
}

// 获取Tab切换顺序序号。
//
// Get Tab switching sequence number.
func (t *TToggleBox) TabOrder() TTabOrder {
	return ToggleBox_GetTabOrder(t.instance)
}

// 设置Tab切换顺序序号。
//
// Set Tab switching sequence number.
func (t *TToggleBox) SetTabOrder(value TTabOrder) {
	ToggleBox_SetTabOrder(t.instance, value)
}

// 获取Tab可停留。
//
// Get Tab can stay.
func (t *TToggleBox) TabStop() bool {
	return ToggleBox_GetTabStop(t.instance)
}

// 设置Tab可停留。
//
// Set Tab can stay.
func (t *TToggleBox) SetTabStop(value bool) {
	ToggleBox_SetTabStop(t.instance, value)
}

// 获取控件可视。
//
// Get the control visible.
func (t *TToggleBox) Visible() bool {
	return ToggleBox_GetVisible(t.instance)
}

// 设置控件可视。
//
// Set the control visible.
func (t *TToggleBox) SetVisible(value bool) {
	ToggleBox_SetVisible(t.instance, value)
}

// 获取依靠客户端总数。
func (t *TToggleBox) DockClientCount() int32 {
	return ToggleBox_GetDockClientCount(t.instance)
}

// 获取停靠站点。
//
// Get Docking site.
func (t *TToggleBox) DockSite() bool {
	return ToggleBox_GetDockSite(t.instance)
}

// 设置停靠站点。
//
// Set Docking site.
func (t *TToggleBox) SetDockSite(value bool) {
	ToggleBox_SetDockSite(t.instance, value)
}

// 获取鼠标是否在客户端，仅VCL有效。
//
// Get Whether the mouse is on the client, only VCL is valid.
func (t *TToggleBox) MouseInClient() bool {
	return ToggleBox_GetMouseInClient(t.instance)
}

// 获取当前停靠的可视总数。
//
// Get The total number of visible calls currently docked.
func (t *TToggleBox) VisibleDockClientCount() int32 {
	return ToggleBox_GetVisibleDockClientCount(t.instance)
}

// 获取画刷对象。
//
// Get Brush.
func (t *TToggleBox) Brush() *TBrush {
	return AsBrush(ToggleBox_GetBrush(t.instance))
}

// 获取子控件数。
//
// Get Number of child controls.
func (t *TToggleBox) ControlCount() int32 {
	return ToggleBox_GetControlCount(t.instance)
}

// 获取控件句柄。
//
// Get Control handle.
func (t *TToggleBox) Handle() HWND {
	return ToggleBox_GetHandle(t.instance)
}

// 获取父容器句柄。
//
// Get Parent container handle.
func (t *TToggleBox) ParentWindow() HWND {
	return ToggleBox_GetParentWindow(t.instance)
}

// 设置父容器句柄。
//
// Set Parent container handle.
func (t *TToggleBox) SetParentWindow(value HWND) {
	ToggleBox_SetParentWindow(t.instance, value)
}

func (t *TToggleBox) Showing() bool {
	return ToggleBox_GetShowing(t.instance)
}

// 获取使用停靠管理。
func (t *TToggleBox) UseDockManager() bool {
	return ToggleBox_GetUseDockManager(t.instance)
}

// 设置使用停靠管理。
func (t *TToggleBox) SetUseDockManager(value bool) {
	ToggleBox_SetUseDockManager(t.instance, value)
}

func (t *TToggleBox) Action() *TAction {
	return AsAction(ToggleBox_GetAction(t.instance))
}

func (t *TToggleBox) SetAction(value IComponent) {
	ToggleBox_SetAction(t.instance, CheckPtr(value))
}

func (t *TToggleBox) BiDiMode() TBiDiMode {
	return ToggleBox_GetBiDiMode(t.instance)
}

func (t *TToggleBox) SetBiDiMode(value TBiDiMode) {
	ToggleBox_SetBiDiMode(t.instance, value)
}

func (t *TToggleBox) BoundsRect() TRect {
	return ToggleBox_GetBoundsRect(t.instance)
}

func (t *TToggleBox) SetBoundsRect(value TRect) {
	ToggleBox_SetBoundsRect(t.instance, value)
}

// 获取客户区高度。
//
// Get client height.
func (t *TToggleBox) ClientHeight() int32 {
	return ToggleBox_GetClientHeight(t.instance)
}

// 设置客户区高度。
//
// Set client height.
func (t *TToggleBox) SetClientHeight(value int32) {
	ToggleBox_SetClientHeight(t.instance, value)
}

func (t *TToggleBox) ClientOrigin() TPoint {
	return ToggleBox_GetClientOrigin(t.instance)
}

// 获取客户区矩形。
//
// Get client rectangle.
func (t *TToggleBox) ClientRect() TRect {
	return ToggleBox_GetClientRect(t.instance)
}

// 获取客户区宽度。
//
// Get client width.
func (t *TToggleBox) ClientWidth() int32 {
	return ToggleBox_GetClientWidth(t.instance)
}

// 设置客户区宽度。
//
// Set client width.
func (t *TToggleBox) SetClientWidth(value int32) {
	ToggleBox_SetClientWidth(t.instance, value)
}

// 获取控件状态。
//
// Get control state.
func (t *TToggleBox) ControlState() TControlState {
	return ToggleBox_GetControlState(t.instance)
}

// 设置控件状态。
//
// Set control state.
func (t *TToggleBox) SetControlState(value TControlState) {
	ToggleBox_SetControlState(t.instance, value)
}

// 获取控件样式。
//
// Get control style.
func (t *TToggleBox) ControlStyle() TControlStyle {
	return ToggleBox_GetControlStyle(t.instance)
}

// 设置控件样式。
//
// Set control style.
func (t *TToggleBox) SetControlStyle(value TControlStyle) {
	ToggleBox_SetControlStyle(t.instance, value)
}

func (t *TToggleBox) Floating() bool {
	return ToggleBox_GetFloating(t.instance)
}

// 获取控件父容器。
//
// Get control parent container.
func (t *TToggleBox) Parent() *TWinControl {
	return AsWinControl(ToggleBox_GetParent(t.instance))
}

// 设置控件父容器。
//
// Set control parent container.
func (t *TToggleBox) SetParent(value IWinControl) {
	ToggleBox_SetParent(t.instance, CheckPtr(value))
}

// 获取左边位置。
//
// Get Left position.
func (t *TToggleBox) Left() int32 {
	return ToggleBox_GetLeft(t.instance)
}

// 设置左边位置。
//
// Set Left position.
func (t *TToggleBox) SetLeft(value int32) {
	ToggleBox_SetLeft(t.instance, value)
}

// 获取顶边位置。
//
// Get Top position.
func (t *TToggleBox) Top() int32 {
	return ToggleBox_GetTop(t.instance)
}

// 设置顶边位置。
//
// Set Top position.
func (t *TToggleBox) SetTop(value int32) {
	ToggleBox_SetTop(t.instance, value)
}

// 获取宽度。
//
// Get width.
func (t *TToggleBox) Width() int32 {
	return ToggleBox_GetWidth(t.instance)
}

// 设置宽度。
//
// Set width.
func (t *TToggleBox) SetWidth(value int32) {
	ToggleBox_SetWidth(t.instance, value)
}

// 获取高度。
//
// Get height.
func (t *TToggleBox) Height() int32 {
	return ToggleBox_GetHeight(t.instance)
}

// 设置高度。
//
// Set height.
func (t *TToggleBox) SetHeight(value int32) {
	ToggleBox_SetHeight(t.instance, value)
}

// 获取控件光标。
//
// Get control cursor.
func (t *TToggleBox) Cursor() TCursor {
	return ToggleBox_GetCursor(t.instance)
}

// 设置控件光标。
//
// Set control cursor.
func (t *TToggleBox) SetCursor(value TCursor) {
	ToggleBox_SetCursor(t.instance, value)
}

// 获取组件总数。
//
// Get the total number of components.
func (t *TToggleBox) ComponentCount() int32 {
	return ToggleBox_GetComponentCount(t.instance)
}

// 获取组件索引。
//
// Get component index.
func (t *TToggleBox) ComponentIndex() int32 {
	return ToggleBox_GetComponentIndex(t.instance)
}

// 设置组件索引。
//
// Set component index.
func (t *TToggleBox) SetComponentIndex(value int32) {
	ToggleBox_SetComponentIndex(t.instance, value)
}

// 获取组件所有者。
//
// Get component owner.
func (t *TToggleBox) Owner() *TComponent {
	return AsComponent(ToggleBox_GetOwner(t.instance))
}

// 获取组件名称。
//
// Get the component name.
func (t *TToggleBox) Name() string {
	return ToggleBox_GetName(t.instance)
}

// 设置组件名称。
//
// Set the component name.
func (t *TToggleBox) SetName(value string) {
	ToggleBox_SetName(t.instance, value)
}

// 获取对象标记。
//
// Get the control tag.
func (t *TToggleBox) Tag() int {
	return ToggleBox_GetTag(t.instance)
}

// 设置对象标记。
//
// Set the control tag.
func (t *TToggleBox) SetTag(value int) {
	ToggleBox_SetTag(t.instance, value)
}

// 获取左边锚点。
func (t *TToggleBox) AnchorSideLeft() *TAnchorSide {
	return AsAnchorSide(ToggleBox_GetAnchorSideLeft(t.instance))
}

// 设置左边锚点。
func (t *TToggleBox) SetAnchorSideLeft(value *TAnchorSide) {
	ToggleBox_SetAnchorSideLeft(t.instance, CheckPtr(value))
}

// 获取顶边锚点。
func (t *TToggleBox) AnchorSideTop() *TAnchorSide {
	return AsAnchorSide(ToggleBox_GetAnchorSideTop(t.instance))
}

// 设置顶边锚点。
func (t *TToggleBox) SetAnchorSideTop(value *TAnchorSide) {
	ToggleBox_SetAnchorSideTop(t.instance, CheckPtr(value))
}

// 获取右边锚点。
func (t *TToggleBox) AnchorSideRight() *TAnchorSide {
	return AsAnchorSide(ToggleBox_GetAnchorSideRight(t.instance))
}

// 设置右边锚点。
func (t *TToggleBox) SetAnchorSideRight(value *TAnchorSide) {
	ToggleBox_SetAnchorSideRight(t.instance, CheckPtr(value))
}

// 获取底边锚点。
func (t *TToggleBox) AnchorSideBottom() *TAnchorSide {
	return AsAnchorSide(ToggleBox_GetAnchorSideBottom(t.instance))
}

// 设置底边锚点。
func (t *TToggleBox) SetAnchorSideBottom(value *TAnchorSide) {
	ToggleBox_SetAnchorSideBottom(t.instance, CheckPtr(value))
}

func (t *TToggleBox) ChildSizing() *TControlChildSizing {
	return AsControlChildSizing(ToggleBox_GetChildSizing(t.instance))
}

func (t *TToggleBox) SetChildSizing(value *TControlChildSizing) {
	ToggleBox_SetChildSizing(t.instance, CheckPtr(value))
}

// 获取边框间距。
func (t *TToggleBox) BorderSpacing() *TControlBorderSpacing {
	return AsControlBorderSpacing(ToggleBox_GetBorderSpacing(t.instance))
}

// 设置边框间距。
func (t *TToggleBox) SetBorderSpacing(value *TControlBorderSpacing) {
	ToggleBox_SetBorderSpacing(t.instance, CheckPtr(value))
}

// 获取指定索引停靠客户端。
func (t *TToggleBox) DockClients(Index int32) *TControl {
	return AsControl(ToggleBox_GetDockClients(t.instance, Index))
}

// 获取指定索引子控件。
func (t *TToggleBox) Controls(Index int32) *TControl {
	return AsControl(ToggleBox_GetControls(t.instance, Index))
}

// 获取指定索引组件。
//
// Get the specified index component.
func (t *TToggleBox) Components(AIndex int32) *TComponent {
	return AsComponent(ToggleBox_GetComponents(t.instance, AIndex))
}

// 获取锚侧面。
func (t *TToggleBox) AnchorSide(AKind TAnchorKind) *TAnchorSide {
	return AsAnchorSide(ToggleBox_GetAnchorSide(t.instance, AKind))
}
