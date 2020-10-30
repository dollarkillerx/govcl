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

type TToolButton struct {
	IControl
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

// 创建一个新的对象。
//
// Create a new object.
func NewToolButton(owner IComponent) *TToolButton {
	t := new(TToolButton)
	t.instance = ToolButton_Create(CheckPtr(owner))
	t.ptr = unsafe.Pointer(t.instance)
	return t
}

// 动态转换一个已存在的对象实例。
//
// Dynamically convert an existing object instance.
func AsToolButton(obj interface{}) *TToolButton {
	instance, ptr := getInstance(obj)
	if instance == 0 {
		return nil
	}
	return &TToolButton{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// 新建一个对象来自已经存在的对象实例指针。
//
// Create a new object from an existing object instance pointer.
// Deprecated: use AsToolButton.
func ToolButtonFromInst(inst uintptr) *TToolButton {
	return AsToolButton(inst)
}

// 新建一个对象来自已经存在的对象实例。
//
// Create a new object from an existing object instance.
// Deprecated: use AsToolButton.
func ToolButtonFromObj(obj IObject) *TToolButton {
	return AsToolButton(obj)
}

// 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
//
// Create a new object from an unsecured address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsToolButton.
func ToolButtonFromUnsafePointer(ptr unsafe.Pointer) *TToolButton {
	return AsToolButton(ptr)
}

// -------------------------- Deprecated end --------------------------
// 释放对象。
//
// Free object.
func (t *TToolButton) Free() {
	if t.instance != 0 {
		ToolButton_Free(t.instance)
		t.instance, t.ptr = 0, nullptr
	}
}

// 返回对象实例指针。
//
// Return object instance pointer.
func (t *TToolButton) Instance() uintptr {
	return t.instance
}

// 获取一个不安全的地址。
//
// Get an unsafe address.
func (t *TToolButton) UnsafeAddr() unsafe.Pointer {
	return t.ptr
}

// 检测地址是否为空。
//
// Check if the address is empty.
func (t *TToolButton) IsValid() bool {
	return t.instance != 0
}

// 检测当前对象是否继承自目标对象。
//
// Checks whether the current object is inherited from the target object.
func (t *TToolButton) Is() TIs {
	return TIs(t.instance)
}

// 动态转换当前对象为目标对象。
//
// Dynamically convert the current object to the target object.
//func (t *TToolButton) As() TAs {
//    return TAs(t.instance)
//}

// 获取类信息指针。
//
// Get class information pointer.
func TToolButtonClass() TClass {
	return ToolButton_StaticClassType()
}

func (t *TToolButton) CheckMenuDropdown() bool {
	return ToolButton_CheckMenuDropdown(t.instance)
}

// 单击。
func (t *TToolButton) Click() {
	ToolButton_Click(t.instance)
}

// 设置组件边界。
//
// Set component boundaries.
func (t *TToolButton) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
	ToolButton_SetBounds(t.instance, ALeft, ATop, AWidth, AHeight)
}

// 将控件置于最前。
//
// Bring the control to the front.
func (t *TToolButton) BringToFront() {
	ToolButton_BringToFront(t.instance)
}

// 将客户端坐标转为绝对的屏幕坐标。
//
// Convert client coordinates to absolute screen coordinates.
func (t *TToolButton) ClientToScreen(Point TPoint) TPoint {
	return ToolButton_ClientToScreen(t.instance, Point)
}

// 将客户端坐标转为父容器坐标。
//
// Convert client coordinates to parent container coordinates.
func (t *TToolButton) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
	return ToolButton_ClientToParent(t.instance, Point, CheckPtr(AParent))
}

// 是否在拖拽中。
//
// Is it in the middle of dragging.
func (t *TToolButton) Dragging() bool {
	return ToolButton_Dragging(t.instance)
}

// 是否有父容器。
//
// Is there a parent container.
func (t *TToolButton) HasParent() bool {
	return ToolButton_HasParent(t.instance)
}

// 隐藏控件。
//
// Hidden control.
func (t *TToolButton) Hide() {
	ToolButton_Hide(t.instance)
}

// 要求重绘。
//
// Redraw.
func (t *TToolButton) Invalidate() {
	ToolButton_Invalidate(t.instance)
}

// 发送一个消息。
//
// Send a message.
func (t *TToolButton) Perform(Msg uint32, WParam uintptr, LParam int) int {
	return ToolButton_Perform(t.instance, Msg, WParam, LParam)
}

// 刷新控件。
//
// Refresh control.
func (t *TToolButton) Refresh() {
	ToolButton_Refresh(t.instance)
}

// 重绘。
//
// Repaint.
func (t *TToolButton) Repaint() {
	ToolButton_Repaint(t.instance)
}

// 将屏幕坐标转为客户端坐标。
//
// Convert screen coordinates to client coordinates.
func (t *TToolButton) ScreenToClient(Point TPoint) TPoint {
	return ToolButton_ScreenToClient(t.instance, Point)
}

// 将父容器坐标转为客户端坐标。
//
// Convert parent container coordinates to client coordinates.
func (t *TToolButton) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
	return ToolButton_ParentToClient(t.instance, Point, CheckPtr(AParent))
}

// 控件至于最后面。
//
// The control is placed at the end.
func (t *TToolButton) SendToBack() {
	ToolButton_SendToBack(t.instance)
}

// 显示控件。
//
// Show control.
func (t *TToolButton) Show() {
	ToolButton_Show(t.instance)
}

// 控件更新。
//
// Update.
func (t *TToolButton) Update() {
	ToolButton_Update(t.instance)
}

// 获取控件的字符，如果有。
//
// Get the characters of the control, if any.
func (t *TToolButton) GetTextBuf(Buffer *string, BufSize int32) int32 {
	return ToolButton_GetTextBuf(t.instance, Buffer, BufSize)
}

// 获取控件的字符长，如果有。
//
// Get the character length of the control, if any.
func (t *TToolButton) GetTextLen() int32 {
	return ToolButton_GetTextLen(t.instance)
}

// 设置控件字符，如果有。
//
// Set control characters, if any.
func (t *TToolButton) SetTextBuf(Buffer string) {
	ToolButton_SetTextBuf(t.instance, Buffer)
}

// 查找指定名称的组件。
//
// Find the component with the specified name.
func (t *TToolButton) FindComponent(AName string) *TComponent {
	return AsComponent(ToolButton_FindComponent(t.instance, AName))
}

// 获取类名路径。
//
// Get the class name path.
func (t *TToolButton) GetNamePath() string {
	return ToolButton_GetNamePath(t.instance)
}

// 复制一个对象，如果对象实现了此方法的话。
//
// Copy an object, if the object implements this method.
func (t *TToolButton) Assign(Source IObject) {
	ToolButton_Assign(t.instance, CheckPtr(Source))
}

// 获取类的类型信息。
//
// Get class type information.
func (t *TToolButton) ClassType() TClass {
	return ToolButton_ClassType(t.instance)
}

// 获取当前对象类名称。
//
// Get the current object class name.
func (t *TToolButton) ClassName() string {
	return ToolButton_ClassName(t.instance)
}

// 获取当前对象实例大小。
//
// Get the current object instance size.
func (t *TToolButton) InstanceSize() int32 {
	return ToolButton_InstanceSize(t.instance)
}

// 判断当前类是否继承自指定类。
//
// Determine whether the current class inherits from the specified class.
func (t *TToolButton) InheritsFrom(AClass TClass) bool {
	return ToolButton_InheritsFrom(t.instance, AClass)
}

// 与一个对象进行比较。
//
// Compare with an object.
func (t *TToolButton) Equals(Obj IObject) bool {
	return ToolButton_Equals(t.instance, CheckPtr(Obj))
}

// 获取类的哈希值。
//
// Get the hash value of the class.
func (t *TToolButton) GetHashCode() int32 {
	return ToolButton_GetHashCode(t.instance)
}

// 文本类信息。
//
// Text information.
func (t *TToolButton) ToString() string {
	return ToolButton_ToString(t.instance)
}

func (t *TToolButton) AnchorToNeighbour(ASide TAnchorKind, ASpace int32, ASibling IControl) {
	ToolButton_AnchorToNeighbour(t.instance, ASide, ASpace, CheckPtr(ASibling))
}

func (t *TToolButton) AnchorParallel(ASide TAnchorKind, ASpace int32, ASibling IControl) {
	ToolButton_AnchorParallel(t.instance, ASide, ASpace, CheckPtr(ASibling))
}

// 置于指定控件的横向中心。
func (t *TToolButton) AnchorHorizontalCenterTo(ASibling IControl) {
	ToolButton_AnchorHorizontalCenterTo(t.instance, CheckPtr(ASibling))
}

// 置于指定控件的纵向中心。
func (t *TToolButton) AnchorVerticalCenterTo(ASibling IControl) {
	ToolButton_AnchorVerticalCenterTo(t.instance, CheckPtr(ASibling))
}

func (t *TToolButton) AnchorSame(ASide TAnchorKind, ASibling IControl) {
	ToolButton_AnchorSame(t.instance, ASide, CheckPtr(ASibling))
}

func (t *TToolButton) AnchorAsAlign(ATheAlign TAlign, ASpace int32) {
	ToolButton_AnchorAsAlign(t.instance, ATheAlign, ASpace)
}

func (t *TToolButton) AnchorClient(ASpace int32) {
	ToolButton_AnchorClient(t.instance, ASpace)
}

func (t *TToolButton) ScaleDesignToForm(ASize int32) int32 {
	return ToolButton_ScaleDesignToForm(t.instance, ASize)
}

func (t *TToolButton) ScaleFormToDesign(ASize int32) int32 {
	return ToolButton_ScaleFormToDesign(t.instance, ASize)
}

func (t *TToolButton) Scale96ToForm(ASize int32) int32 {
	return ToolButton_Scale96ToForm(t.instance, ASize)
}

func (t *TToolButton) ScaleFormTo96(ASize int32) int32 {
	return ToolButton_ScaleFormTo96(t.instance, ASize)
}

func (t *TToolButton) Scale96ToFont(ASize int32) int32 {
	return ToolButton_Scale96ToFont(t.instance, ASize)
}

func (t *TToolButton) ScaleFontTo96(ASize int32) int32 {
	return ToolButton_ScaleFontTo96(t.instance, ASize)
}

func (t *TToolButton) ScaleScreenToFont(ASize int32) int32 {
	return ToolButton_ScaleScreenToFont(t.instance, ASize)
}

func (t *TToolButton) ScaleFontToScreen(ASize int32) int32 {
	return ToolButton_ScaleFontToScreen(t.instance, ASize)
}

func (t *TToolButton) Scale96ToScreen(ASize int32) int32 {
	return ToolButton_Scale96ToScreen(t.instance, ASize)
}

func (t *TToolButton) ScaleScreenTo96(ASize int32) int32 {
	return ToolButton_ScaleScreenTo96(t.instance, ASize)
}

func (t *TToolButton) AutoAdjustLayout(AMode TLayoutAdjustmentPolicy, AFromPPI int32, AToPPI int32, AOldFormWidth int32, ANewFormWidth int32) {
	ToolButton_AutoAdjustLayout(t.instance, AMode, AFromPPI, AToPPI, AOldFormWidth, ANewFormWidth)
}

func (t *TToolButton) FixDesignFontsPPI(ADesignTimePPI int32) {
	ToolButton_FixDesignFontsPPI(t.instance, ADesignTimePPI)
}

func (t *TToolButton) ScaleFontsPPI(AToPPI int32, AProportion float64) {
	ToolButton_ScaleFontsPPI(t.instance, AToPPI, AProportion)
}

func (t *TToolButton) Index() int32 {
	return ToolButton_GetIndex(t.instance)
}

func (t *TToolButton) Action() *TAction {
	return AsAction(ToolButton_GetAction(t.instance))
}

func (t *TToolButton) SetAction(value IComponent) {
	ToolButton_SetAction(t.instance, CheckPtr(value))
}

func (t *TToolButton) AllowAllUp() bool {
	return ToolButton_GetAllowAllUp(t.instance)
}

func (t *TToolButton) SetAllowAllUp(value bool) {
	ToolButton_SetAllowAllUp(t.instance, value)
}

// 获取自动调整大小。
func (t *TToolButton) AutoSize() bool {
	return ToolButton_GetAutoSize(t.instance)
}

// 设置自动调整大小。
func (t *TToolButton) SetAutoSize(value bool) {
	ToolButton_SetAutoSize(t.instance, value)
}

// 获取控件标题。
//
// Get the control title.
func (t *TToolButton) Caption() string {
	return ToolButton_GetCaption(t.instance)
}

// 设置控件标题。
//
// Set the control title.
func (t *TToolButton) SetCaption(value string) {
	ToolButton_SetCaption(t.instance, value)
}

func (t *TToolButton) Down() bool {
	return ToolButton_GetDown(t.instance)
}

func (t *TToolButton) SetDown(value bool) {
	ToolButton_SetDown(t.instance, value)
}

// 获取设置控件拖拽时的光标。
//
// Get Set the cursor when the control is dragged.
func (t *TToolButton) DragCursor() TCursor {
	return ToolButton_GetDragCursor(t.instance)
}

// 设置设置控件拖拽时的光标。
//
// Set Set the cursor when the control is dragged.
func (t *TToolButton) SetDragCursor(value TCursor) {
	ToolButton_SetDragCursor(t.instance, value)
}

// 获取拖拽方式。
//
// Get Drag and drop.
func (t *TToolButton) DragKind() TDragKind {
	return ToolButton_GetDragKind(t.instance)
}

// 设置拖拽方式。
//
// Set Drag and drop.
func (t *TToolButton) SetDragKind(value TDragKind) {
	ToolButton_SetDragKind(t.instance, value)
}

// 获取拖拽模式。
//
// Get Drag mode.
func (t *TToolButton) DragMode() TDragMode {
	return ToolButton_GetDragMode(t.instance)
}

// 设置拖拽模式。
//
// Set Drag mode.
func (t *TToolButton) SetDragMode(value TDragMode) {
	ToolButton_SetDragMode(t.instance, value)
}

func (t *TToolButton) DropdownMenu() *TPopupMenu {
	return AsPopupMenu(ToolButton_GetDropdownMenu(t.instance))
}

func (t *TToolButton) SetDropdownMenu(value IComponent) {
	ToolButton_SetDropdownMenu(t.instance, CheckPtr(value))
}

// 获取控件启用。
//
// Get the control enabled.
func (t *TToolButton) Enabled() bool {
	return ToolButton_GetEnabled(t.instance)
}

// 设置控件启用。
//
// Set the control enabled.
func (t *TToolButton) SetEnabled(value bool) {
	ToolButton_SetEnabled(t.instance, value)
}

func (t *TToolButton) Grouped() bool {
	return ToolButton_GetGrouped(t.instance)
}

func (t *TToolButton) SetGrouped(value bool) {
	ToolButton_SetGrouped(t.instance, value)
}

// 获取高度。
//
// Get height.
func (t *TToolButton) Height() int32 {
	return ToolButton_GetHeight(t.instance)
}

// 设置高度。
//
// Set height.
func (t *TToolButton) SetHeight(value int32) {
	ToolButton_SetHeight(t.instance, value)
}

// 获取图像在images中的索引。
func (t *TToolButton) ImageIndex() int32 {
	return ToolButton_GetImageIndex(t.instance)
}

// 设置图像在images中的索引。
func (t *TToolButton) SetImageIndex(value int32) {
	ToolButton_SetImageIndex(t.instance, value)
}

func (t *TToolButton) Indeterminate() bool {
	return ToolButton_GetIndeterminate(t.instance)
}

func (t *TToolButton) SetIndeterminate(value bool) {
	ToolButton_SetIndeterminate(t.instance, value)
}

func (t *TToolButton) Marked() bool {
	return ToolButton_GetMarked(t.instance)
}

func (t *TToolButton) SetMarked(value bool) {
	ToolButton_SetMarked(t.instance, value)
}

// 获取以父容器的ShowHint属性为准。
func (t *TToolButton) ParentShowHint() bool {
	return ToolButton_GetParentShowHint(t.instance)
}

// 设置以父容器的ShowHint属性为准。
func (t *TToolButton) SetParentShowHint(value bool) {
	ToolButton_SetParentShowHint(t.instance, value)
}

// 获取右键菜单。
//
// Get Right click menu.
func (t *TToolButton) PopupMenu() *TPopupMenu {
	return AsPopupMenu(ToolButton_GetPopupMenu(t.instance))
}

// 设置右键菜单。
//
// Set Right click menu.
func (t *TToolButton) SetPopupMenu(value IComponent) {
	ToolButton_SetPopupMenu(t.instance, CheckPtr(value))
}

func (t *TToolButton) Wrap() bool {
	return ToolButton_GetWrap(t.instance)
}

func (t *TToolButton) SetWrap(value bool) {
	ToolButton_SetWrap(t.instance, value)
}

// 获取显示鼠标悬停提示。
//
// Get Show mouseover tips.
func (t *TToolButton) ShowHint() bool {
	return ToolButton_GetShowHint(t.instance)
}

// 设置显示鼠标悬停提示。
//
// Set Show mouseover tips.
func (t *TToolButton) SetShowHint(value bool) {
	ToolButton_SetShowHint(t.instance, value)
}

func (t *TToolButton) Style() TToolButtonStyle {
	return ToolButton_GetStyle(t.instance)
}

func (t *TToolButton) SetStyle(value TToolButtonStyle) {
	ToolButton_SetStyle(t.instance, value)
}

// 获取控件可视。
//
// Get the control visible.
func (t *TToolButton) Visible() bool {
	return ToolButton_GetVisible(t.instance)
}

// 设置控件可视。
//
// Set the control visible.
func (t *TToolButton) SetVisible(value bool) {
	ToolButton_SetVisible(t.instance, value)
}

// 获取宽度。
//
// Get width.
func (t *TToolButton) Width() int32 {
	return ToolButton_GetWidth(t.instance)
}

// 设置宽度。
//
// Set width.
func (t *TToolButton) SetWidth(value int32) {
	ToolButton_SetWidth(t.instance, value)
}

// 设置控件单击事件。
//
// Set control click event.
func (t *TToolButton) SetOnClick(fn TNotifyEvent) {
	ToolButton_SetOnClick(t.instance, fn)
}

// 设置上下文弹出事件，一般是右键时弹出。
//
// Set Context popup event, usually pop up when right click.
func (t *TToolButton) SetOnContextPopup(fn TContextPopupEvent) {
	ToolButton_SetOnContextPopup(t.instance, fn)
}

// 设置拖拽下落事件。
//
// Set Drag and drop event.
func (t *TToolButton) SetOnDragDrop(fn TDragDropEvent) {
	ToolButton_SetOnDragDrop(t.instance, fn)
}

// 设置拖拽完成事件。
//
// Set Drag and drop completion event.
func (t *TToolButton) SetOnDragOver(fn TDragOverEvent) {
	ToolButton_SetOnDragOver(t.instance, fn)
}

// 设置停靠结束事件。
//
// Set Dock end event.
func (t *TToolButton) SetOnEndDock(fn TEndDragEvent) {
	ToolButton_SetOnEndDock(t.instance, fn)
}

// 设置拖拽结束。
//
// Set End of drag.
func (t *TToolButton) SetOnEndDrag(fn TEndDragEvent) {
	ToolButton_SetOnEndDrag(t.instance, fn)
}

// 设置鼠标按下事件。
//
// Set Mouse down event.
func (t *TToolButton) SetOnMouseDown(fn TMouseEvent) {
	ToolButton_SetOnMouseDown(t.instance, fn)
}

// 设置鼠标进入事件。
//
// Set Mouse entry event.
func (t *TToolButton) SetOnMouseEnter(fn TNotifyEvent) {
	ToolButton_SetOnMouseEnter(t.instance, fn)
}

// 设置鼠标离开事件。
//
// Set Mouse leave event.
func (t *TToolButton) SetOnMouseLeave(fn TNotifyEvent) {
	ToolButton_SetOnMouseLeave(t.instance, fn)
}

// 设置鼠标移动事件。
func (t *TToolButton) SetOnMouseMove(fn TMouseMoveEvent) {
	ToolButton_SetOnMouseMove(t.instance, fn)
}

// 设置鼠标抬起事件。
//
// Set Mouse lift event.
func (t *TToolButton) SetOnMouseUp(fn TMouseEvent) {
	ToolButton_SetOnMouseUp(t.instance, fn)
}

// 设置启动停靠。
func (t *TToolButton) SetOnStartDock(fn TStartDockEvent) {
	ToolButton_SetOnStartDock(t.instance, fn)
}

// 获取控件自动调整。
//
// Get Control automatically adjusts.
func (t *TToolButton) Align() TAlign {
	return ToolButton_GetAlign(t.instance)
}

// 设置控件自动调整。
//
// Set Control automatically adjusts.
func (t *TToolButton) SetAlign(value TAlign) {
	ToolButton_SetAlign(t.instance, value)
}

// 获取四个角位置的锚点。
func (t *TToolButton) Anchors() TAnchors {
	return ToolButton_GetAnchors(t.instance)
}

// 设置四个角位置的锚点。
func (t *TToolButton) SetAnchors(value TAnchors) {
	ToolButton_SetAnchors(t.instance, value)
}

func (t *TToolButton) BiDiMode() TBiDiMode {
	return ToolButton_GetBiDiMode(t.instance)
}

func (t *TToolButton) SetBiDiMode(value TBiDiMode) {
	ToolButton_SetBiDiMode(t.instance, value)
}

func (t *TToolButton) BoundsRect() TRect {
	return ToolButton_GetBoundsRect(t.instance)
}

func (t *TToolButton) SetBoundsRect(value TRect) {
	ToolButton_SetBoundsRect(t.instance, value)
}

// 获取客户区高度。
//
// Get client height.
func (t *TToolButton) ClientHeight() int32 {
	return ToolButton_GetClientHeight(t.instance)
}

// 设置客户区高度。
//
// Set client height.
func (t *TToolButton) SetClientHeight(value int32) {
	ToolButton_SetClientHeight(t.instance, value)
}

func (t *TToolButton) ClientOrigin() TPoint {
	return ToolButton_GetClientOrigin(t.instance)
}

// 获取客户区矩形。
//
// Get client rectangle.
func (t *TToolButton) ClientRect() TRect {
	return ToolButton_GetClientRect(t.instance)
}

// 获取客户区宽度。
//
// Get client width.
func (t *TToolButton) ClientWidth() int32 {
	return ToolButton_GetClientWidth(t.instance)
}

// 设置客户区宽度。
//
// Set client width.
func (t *TToolButton) SetClientWidth(value int32) {
	ToolButton_SetClientWidth(t.instance, value)
}

// 获取约束控件大小。
func (t *TToolButton) Constraints() *TSizeConstraints {
	return AsSizeConstraints(ToolButton_GetConstraints(t.instance))
}

// 设置约束控件大小。
func (t *TToolButton) SetConstraints(value *TSizeConstraints) {
	ToolButton_SetConstraints(t.instance, CheckPtr(value))
}

// 获取控件状态。
//
// Get control state.
func (t *TToolButton) ControlState() TControlState {
	return ToolButton_GetControlState(t.instance)
}

// 设置控件状态。
//
// Set control state.
func (t *TToolButton) SetControlState(value TControlState) {
	ToolButton_SetControlState(t.instance, value)
}

// 获取控件样式。
//
// Get control style.
func (t *TToolButton) ControlStyle() TControlStyle {
	return ToolButton_GetControlStyle(t.instance)
}

// 设置控件样式。
//
// Set control style.
func (t *TToolButton) SetControlStyle(value TControlStyle) {
	ToolButton_SetControlStyle(t.instance, value)
}

func (t *TToolButton) Floating() bool {
	return ToolButton_GetFloating(t.instance)
}

// 获取控件父容器。
//
// Get control parent container.
func (t *TToolButton) Parent() *TWinControl {
	return AsWinControl(ToolButton_GetParent(t.instance))
}

// 设置控件父容器。
//
// Set control parent container.
func (t *TToolButton) SetParent(value IWinControl) {
	ToolButton_SetParent(t.instance, CheckPtr(value))
}

// 获取左边位置。
//
// Get Left position.
func (t *TToolButton) Left() int32 {
	return ToolButton_GetLeft(t.instance)
}

// 设置左边位置。
//
// Set Left position.
func (t *TToolButton) SetLeft(value int32) {
	ToolButton_SetLeft(t.instance, value)
}

// 获取顶边位置。
//
// Get Top position.
func (t *TToolButton) Top() int32 {
	return ToolButton_GetTop(t.instance)
}

// 设置顶边位置。
//
// Set Top position.
func (t *TToolButton) SetTop(value int32) {
	ToolButton_SetTop(t.instance, value)
}

// 获取控件光标。
//
// Get control cursor.
func (t *TToolButton) Cursor() TCursor {
	return ToolButton_GetCursor(t.instance)
}

// 设置控件光标。
//
// Set control cursor.
func (t *TToolButton) SetCursor(value TCursor) {
	ToolButton_SetCursor(t.instance, value)
}

// 获取组件鼠标悬停提示。
//
// Get component mouse hints.
func (t *TToolButton) Hint() string {
	return ToolButton_GetHint(t.instance)
}

// 设置组件鼠标悬停提示。
//
// Set component mouse hints.
func (t *TToolButton) SetHint(value string) {
	ToolButton_SetHint(t.instance, value)
}

// 获取组件总数。
//
// Get the total number of components.
func (t *TToolButton) ComponentCount() int32 {
	return ToolButton_GetComponentCount(t.instance)
}

// 获取组件索引。
//
// Get component index.
func (t *TToolButton) ComponentIndex() int32 {
	return ToolButton_GetComponentIndex(t.instance)
}

// 设置组件索引。
//
// Set component index.
func (t *TToolButton) SetComponentIndex(value int32) {
	ToolButton_SetComponentIndex(t.instance, value)
}

// 获取组件所有者。
//
// Get component owner.
func (t *TToolButton) Owner() *TComponent {
	return AsComponent(ToolButton_GetOwner(t.instance))
}

// 获取组件名称。
//
// Get the component name.
func (t *TToolButton) Name() string {
	return ToolButton_GetName(t.instance)
}

// 设置组件名称。
//
// Set the component name.
func (t *TToolButton) SetName(value string) {
	ToolButton_SetName(t.instance, value)
}

// 获取对象标记。
//
// Get the control tag.
func (t *TToolButton) Tag() int {
	return ToolButton_GetTag(t.instance)
}

// 设置对象标记。
//
// Set the control tag.
func (t *TToolButton) SetTag(value int) {
	ToolButton_SetTag(t.instance, value)
}

// 获取左边锚点。
func (t *TToolButton) AnchorSideLeft() *TAnchorSide {
	return AsAnchorSide(ToolButton_GetAnchorSideLeft(t.instance))
}

// 设置左边锚点。
func (t *TToolButton) SetAnchorSideLeft(value *TAnchorSide) {
	ToolButton_SetAnchorSideLeft(t.instance, CheckPtr(value))
}

// 获取顶边锚点。
func (t *TToolButton) AnchorSideTop() *TAnchorSide {
	return AsAnchorSide(ToolButton_GetAnchorSideTop(t.instance))
}

// 设置顶边锚点。
func (t *TToolButton) SetAnchorSideTop(value *TAnchorSide) {
	ToolButton_SetAnchorSideTop(t.instance, CheckPtr(value))
}

// 获取右边锚点。
func (t *TToolButton) AnchorSideRight() *TAnchorSide {
	return AsAnchorSide(ToolButton_GetAnchorSideRight(t.instance))
}

// 设置右边锚点。
func (t *TToolButton) SetAnchorSideRight(value *TAnchorSide) {
	ToolButton_SetAnchorSideRight(t.instance, CheckPtr(value))
}

// 获取底边锚点。
func (t *TToolButton) AnchorSideBottom() *TAnchorSide {
	return AsAnchorSide(ToolButton_GetAnchorSideBottom(t.instance))
}

// 设置底边锚点。
func (t *TToolButton) SetAnchorSideBottom(value *TAnchorSide) {
	ToolButton_SetAnchorSideBottom(t.instance, CheckPtr(value))
}

// 获取边框间距。
func (t *TToolButton) BorderSpacing() *TControlBorderSpacing {
	return AsControlBorderSpacing(ToolButton_GetBorderSpacing(t.instance))
}

// 设置边框间距。
func (t *TToolButton) SetBorderSpacing(value *TControlBorderSpacing) {
	ToolButton_SetBorderSpacing(t.instance, CheckPtr(value))
}

// 获取指定索引组件。
//
// Get the specified index component.
func (t *TToolButton) Components(AIndex int32) *TComponent {
	return AsComponent(ToolButton_GetComponents(t.instance, AIndex))
}

// 获取锚侧面。
func (t *TToolButton) AnchorSide(AKind TAnchorKind) *TAnchorSide {
	return AsAnchorSide(ToolButton_GetAnchorSide(t.instance, AKind))
}
