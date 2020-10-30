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

type TWinControl struct {
	IWinControl
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

// 创建一个新的对象。
//
// Create a new object.
func NewWinControl(owner IComponent) *TWinControl {
	w := new(TWinControl)
	w.instance = WinControl_Create(CheckPtr(owner))
	w.ptr = unsafe.Pointer(w.instance)
	return w
}

// 动态转换一个已存在的对象实例。
//
// Dynamically convert an existing object instance.
func AsWinControl(obj interface{}) *TWinControl {
	instance, ptr := getInstance(obj)
	if instance == 0 {
		return nil
	}
	return &TWinControl{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// 新建一个对象来自已经存在的对象实例指针。
//
// Create a new object from an existing object instance pointer.
// Deprecated: use AsWinControl.
func WinControlFromInst(inst uintptr) *TWinControl {
	return AsWinControl(inst)
}

// 新建一个对象来自已经存在的对象实例。
//
// Create a new object from an existing object instance.
// Deprecated: use AsWinControl.
func WinControlFromObj(obj IObject) *TWinControl {
	return AsWinControl(obj)
}

// 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
//
// Create a new object from an unsecured address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsWinControl.
func WinControlFromUnsafePointer(ptr unsafe.Pointer) *TWinControl {
	return AsWinControl(ptr)
}

// -------------------------- Deprecated end --------------------------
// 释放对象。
//
// Free object.
func (w *TWinControl) Free() {
	if w.instance != 0 {
		WinControl_Free(w.instance)
		w.instance, w.ptr = 0, nullptr
	}
}

// 返回对象实例指针。
//
// Return object instance pointer.
func (w *TWinControl) Instance() uintptr {
	return w.instance
}

// 获取一个不安全的地址。
//
// Get an unsafe address.
func (w *TWinControl) UnsafeAddr() unsafe.Pointer {
	return w.ptr
}

// 检测地址是否为空。
//
// Check if the address is empty.
func (w *TWinControl) IsValid() bool {
	return w.instance != 0
}

// 检测当前对象是否继承自目标对象。
//
// Checks whether the current object is inherited from the target object.
func (w *TWinControl) Is() TIs {
	return TIs(w.instance)
}

// 动态转换当前对象为目标对象。
//
// Dynamically convert the current object to the target object.
//func (w *TWinControl) As() TAs {
//    return TAs(w.instance)
//}

// 获取类信息指针。
//
// Get class information pointer.
func TWinControlClass() TClass {
	return WinControl_StaticClassType()
}

// 是否可以获得焦点。
func (w *TWinControl) CanFocus() bool {
	return WinControl_CanFocus(w.instance)
}

// 返回是否包含指定控件。
//
// it's contain a specified control.
func (w *TWinControl) ContainsControl(Control IControl) bool {
	return WinControl_ContainsControl(w.instance, CheckPtr(Control))
}

// 返回指定坐标及相关属性位置控件。
//
// Returns the specified coordinate and the relevant attribute position control..
func (w *TWinControl) ControlAtPos(Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) *TControl {
	return AsControl(WinControl_ControlAtPos(w.instance, Pos, AllowDisabled, AllowWinControls, AllLevels))
}

// 禁用控件的对齐。
//
// Disable control alignment.
func (w *TWinControl) DisableAlign() {
	WinControl_DisableAlign(w.instance)
}

// 启用控件对齐。
//
// Enabled control alignment.
func (w *TWinControl) EnableAlign() {
	WinControl_EnableAlign(w.instance)
}

// 查找子控件。
//
// Find sub controls.
func (w *TWinControl) FindChildControl(ControlName string) *TControl {
	return AsControl(WinControl_FindChildControl(w.instance, ControlName))
}

func (w *TWinControl) FlipChildren(AllLevels bool) {
	WinControl_FlipChildren(w.instance, AllLevels)
}

// 返回是否获取焦点。
//
// Return to get focus.
func (w *TWinControl) Focused() bool {
	return WinControl_Focused(w.instance)
}

// 句柄是否已经分配。
//
// Is the handle already allocated.
func (w *TWinControl) HandleAllocated() bool {
	return WinControl_HandleAllocated(w.instance)
}

// 插入一个控件。
//
// Insert a control.
func (w *TWinControl) InsertControl(AControl IControl) {
	WinControl_InsertControl(w.instance, CheckPtr(AControl))
}

// 要求重绘。
//
// Redraw.
func (w *TWinControl) Invalidate() {
	WinControl_Invalidate(w.instance)
}

// 绘画至指定DC。
//
// Painting to the specified DC.
func (w *TWinControl) PaintTo(DC HDC, X int32, Y int32) {
	WinControl_PaintTo(w.instance, DC, X, Y)
}

// 移除一个控件。
//
// Remove a control.
func (w *TWinControl) RemoveControl(AControl IControl) {
	WinControl_RemoveControl(w.instance, CheckPtr(AControl))
}

// 重新对齐。
//
// Realign.
func (w *TWinControl) Realign() {
	WinControl_Realign(w.instance)
}

// 重绘。
//
// Repaint.
func (w *TWinControl) Repaint() {
	WinControl_Repaint(w.instance)
}

// 按比例缩放。
//
// Scale by.
func (w *TWinControl) ScaleBy(M int32, D int32) {
	WinControl_ScaleBy(w.instance, M, D)
}

// 滚动至指定位置。
//
// Scroll by.
func (w *TWinControl) ScrollBy(DeltaX int32, DeltaY int32) {
	WinControl_ScrollBy(w.instance, DeltaX, DeltaY)
}

// 设置组件边界。
//
// Set component boundaries.
func (w *TWinControl) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
	WinControl_SetBounds(w.instance, ALeft, ATop, AWidth, AHeight)
}

// 设置控件焦点。
//
// Set control focus.
func (w *TWinControl) SetFocus() {
	WinControl_SetFocus(w.instance)
}

// 控件更新。
//
// Update.
func (w *TWinControl) Update() {
	WinControl_Update(w.instance)
}

// 将控件置于最前。
//
// Bring the control to the front.
func (w *TWinControl) BringToFront() {
	WinControl_BringToFront(w.instance)
}

// 将客户端坐标转为绝对的屏幕坐标。
//
// Convert client coordinates to absolute screen coordinates.
func (w *TWinControl) ClientToScreen(Point TPoint) TPoint {
	return WinControl_ClientToScreen(w.instance, Point)
}

// 将客户端坐标转为父容器坐标。
//
// Convert client coordinates to parent container coordinates.
func (w *TWinControl) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
	return WinControl_ClientToParent(w.instance, Point, CheckPtr(AParent))
}

// 是否在拖拽中。
//
// Is it in the middle of dragging.
func (w *TWinControl) Dragging() bool {
	return WinControl_Dragging(w.instance)
}

// 是否有父容器。
//
// Is there a parent container.
func (w *TWinControl) HasParent() bool {
	return WinControl_HasParent(w.instance)
}

// 隐藏控件。
//
// Hidden control.
func (w *TWinControl) Hide() {
	WinControl_Hide(w.instance)
}

// 发送一个消息。
//
// Send a message.
func (w *TWinControl) Perform(Msg uint32, WParam uintptr, LParam int) int {
	return WinControl_Perform(w.instance, Msg, WParam, LParam)
}

// 刷新控件。
//
// Refresh control.
func (w *TWinControl) Refresh() {
	WinControl_Refresh(w.instance)
}

// 将屏幕坐标转为客户端坐标。
//
// Convert screen coordinates to client coordinates.
func (w *TWinControl) ScreenToClient(Point TPoint) TPoint {
	return WinControl_ScreenToClient(w.instance, Point)
}

// 将父容器坐标转为客户端坐标。
//
// Convert parent container coordinates to client coordinates.
func (w *TWinControl) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
	return WinControl_ParentToClient(w.instance, Point, CheckPtr(AParent))
}

// 控件至于最后面。
//
// The control is placed at the end.
func (w *TWinControl) SendToBack() {
	WinControl_SendToBack(w.instance)
}

// 显示控件。
//
// Show control.
func (w *TWinControl) Show() {
	WinControl_Show(w.instance)
}

// 获取控件的字符，如果有。
//
// Get the characters of the control, if any.
func (w *TWinControl) GetTextBuf(Buffer *string, BufSize int32) int32 {
	return WinControl_GetTextBuf(w.instance, Buffer, BufSize)
}

// 获取控件的字符长，如果有。
//
// Get the character length of the control, if any.
func (w *TWinControl) GetTextLen() int32 {
	return WinControl_GetTextLen(w.instance)
}

// 设置控件字符，如果有。
//
// Set control characters, if any.
func (w *TWinControl) SetTextBuf(Buffer string) {
	WinControl_SetTextBuf(w.instance, Buffer)
}

// 查找指定名称的组件。
//
// Find the component with the specified name.
func (w *TWinControl) FindComponent(AName string) *TComponent {
	return AsComponent(WinControl_FindComponent(w.instance, AName))
}

// 获取类名路径。
//
// Get the class name path.
func (w *TWinControl) GetNamePath() string {
	return WinControl_GetNamePath(w.instance)
}

// 复制一个对象，如果对象实现了此方法的话。
//
// Copy an object, if the object implements this method.
func (w *TWinControl) Assign(Source IObject) {
	WinControl_Assign(w.instance, CheckPtr(Source))
}

// 获取类的类型信息。
//
// Get class type information.
func (w *TWinControl) ClassType() TClass {
	return WinControl_ClassType(w.instance)
}

// 获取当前对象类名称。
//
// Get the current object class name.
func (w *TWinControl) ClassName() string {
	return WinControl_ClassName(w.instance)
}

// 获取当前对象实例大小。
//
// Get the current object instance size.
func (w *TWinControl) InstanceSize() int32 {
	return WinControl_InstanceSize(w.instance)
}

// 判断当前类是否继承自指定类。
//
// Determine whether the current class inherits from the specified class.
func (w *TWinControl) InheritsFrom(AClass TClass) bool {
	return WinControl_InheritsFrom(w.instance, AClass)
}

// 与一个对象进行比较。
//
// Compare with an object.
func (w *TWinControl) Equals(Obj IObject) bool {
	return WinControl_Equals(w.instance, CheckPtr(Obj))
}

// 获取类的哈希值。
//
// Get the hash value of the class.
func (w *TWinControl) GetHashCode() int32 {
	return WinControl_GetHashCode(w.instance)
}

// 文本类信息。
//
// Text information.
func (w *TWinControl) ToString() string {
	return WinControl_ToString(w.instance)
}

func (w *TWinControl) AnchorToNeighbour(ASide TAnchorKind, ASpace int32, ASibling IControl) {
	WinControl_AnchorToNeighbour(w.instance, ASide, ASpace, CheckPtr(ASibling))
}

func (w *TWinControl) AnchorParallel(ASide TAnchorKind, ASpace int32, ASibling IControl) {
	WinControl_AnchorParallel(w.instance, ASide, ASpace, CheckPtr(ASibling))
}

// 置于指定控件的横向中心。
func (w *TWinControl) AnchorHorizontalCenterTo(ASibling IControl) {
	WinControl_AnchorHorizontalCenterTo(w.instance, CheckPtr(ASibling))
}

// 置于指定控件的纵向中心。
func (w *TWinControl) AnchorVerticalCenterTo(ASibling IControl) {
	WinControl_AnchorVerticalCenterTo(w.instance, CheckPtr(ASibling))
}

func (w *TWinControl) AnchorSame(ASide TAnchorKind, ASibling IControl) {
	WinControl_AnchorSame(w.instance, ASide, CheckPtr(ASibling))
}

func (w *TWinControl) AnchorAsAlign(ATheAlign TAlign, ASpace int32) {
	WinControl_AnchorAsAlign(w.instance, ATheAlign, ASpace)
}

func (w *TWinControl) AnchorClient(ASpace int32) {
	WinControl_AnchorClient(w.instance, ASpace)
}

func (w *TWinControl) ScaleDesignToForm(ASize int32) int32 {
	return WinControl_ScaleDesignToForm(w.instance, ASize)
}

func (w *TWinControl) ScaleFormToDesign(ASize int32) int32 {
	return WinControl_ScaleFormToDesign(w.instance, ASize)
}

func (w *TWinControl) Scale96ToForm(ASize int32) int32 {
	return WinControl_Scale96ToForm(w.instance, ASize)
}

func (w *TWinControl) ScaleFormTo96(ASize int32) int32 {
	return WinControl_ScaleFormTo96(w.instance, ASize)
}

func (w *TWinControl) Scale96ToFont(ASize int32) int32 {
	return WinControl_Scale96ToFont(w.instance, ASize)
}

func (w *TWinControl) ScaleFontTo96(ASize int32) int32 {
	return WinControl_ScaleFontTo96(w.instance, ASize)
}

func (w *TWinControl) ScaleScreenToFont(ASize int32) int32 {
	return WinControl_ScaleScreenToFont(w.instance, ASize)
}

func (w *TWinControl) ScaleFontToScreen(ASize int32) int32 {
	return WinControl_ScaleFontToScreen(w.instance, ASize)
}

func (w *TWinControl) Scale96ToScreen(ASize int32) int32 {
	return WinControl_Scale96ToScreen(w.instance, ASize)
}

func (w *TWinControl) ScaleScreenTo96(ASize int32) int32 {
	return WinControl_ScaleScreenTo96(w.instance, ASize)
}

func (w *TWinControl) AutoAdjustLayout(AMode TLayoutAdjustmentPolicy, AFromPPI int32, AToPPI int32, AOldFormWidth int32, ANewFormWidth int32) {
	WinControl_AutoAdjustLayout(w.instance, AMode, AFromPPI, AToPPI, AOldFormWidth, ANewFormWidth)
}

func (w *TWinControl) FixDesignFontsPPI(ADesignTimePPI int32) {
	WinControl_FixDesignFontsPPI(w.instance, ADesignTimePPI)
}

func (w *TWinControl) ScaleFontsPPI(AToPPI int32, AProportion float64) {
	WinControl_ScaleFontsPPI(w.instance, AToPPI, AProportion)
}

// 获取依靠客户端总数。
func (w *TWinControl) DockClientCount() int32 {
	return WinControl_GetDockClientCount(w.instance)
}

// 获取停靠站点。
//
// Get Docking site.
func (w *TWinControl) DockSite() bool {
	return WinControl_GetDockSite(w.instance)
}

// 设置停靠站点。
//
// Set Docking site.
func (w *TWinControl) SetDockSite(value bool) {
	WinControl_SetDockSite(w.instance, value)
}

// 获取设置控件双缓冲。
//
// Get Set control double buffering.
func (w *TWinControl) DoubleBuffered() bool {
	return WinControl_GetDoubleBuffered(w.instance)
}

// 设置设置控件双缓冲。
//
// Set Set control double buffering.
func (w *TWinControl) SetDoubleBuffered(value bool) {
	WinControl_SetDoubleBuffered(w.instance, value)
}

// 获取鼠标是否在客户端，仅VCL有效。
//
// Get Whether the mouse is on the client, only VCL is valid.
func (w *TWinControl) MouseInClient() bool {
	return WinControl_GetMouseInClient(w.instance)
}

// 获取当前停靠的可视总数。
//
// Get The total number of visible calls currently docked.
func (w *TWinControl) VisibleDockClientCount() int32 {
	return WinControl_GetVisibleDockClientCount(w.instance)
}

// 获取画刷对象。
//
// Get Brush.
func (w *TWinControl) Brush() *TBrush {
	return AsBrush(WinControl_GetBrush(w.instance))
}

// 获取子控件数。
//
// Get Number of child controls.
func (w *TWinControl) ControlCount() int32 {
	return WinControl_GetControlCount(w.instance)
}

// 获取控件句柄。
//
// Get Control handle.
func (w *TWinControl) Handle() HWND {
	return WinControl_GetHandle(w.instance)
}

// 获取使用父容器双缓冲。
//
// Get Parent container double buffering.
func (w *TWinControl) ParentDoubleBuffered() bool {
	return WinControl_GetParentDoubleBuffered(w.instance)
}

// 设置使用父容器双缓冲。
//
// Set Parent container double buffering.
func (w *TWinControl) SetParentDoubleBuffered(value bool) {
	WinControl_SetParentDoubleBuffered(w.instance, value)
}

// 获取父容器句柄。
//
// Get Parent container handle.
func (w *TWinControl) ParentWindow() HWND {
	return WinControl_GetParentWindow(w.instance)
}

// 设置父容器句柄。
//
// Set Parent container handle.
func (w *TWinControl) SetParentWindow(value HWND) {
	WinControl_SetParentWindow(w.instance, value)
}

func (w *TWinControl) Showing() bool {
	return WinControl_GetShowing(w.instance)
}

// 获取Tab切换顺序序号。
//
// Get Tab switching sequence number.
func (w *TWinControl) TabOrder() TTabOrder {
	return WinControl_GetTabOrder(w.instance)
}

// 设置Tab切换顺序序号。
//
// Set Tab switching sequence number.
func (w *TWinControl) SetTabOrder(value TTabOrder) {
	WinControl_SetTabOrder(w.instance, value)
}

// 获取Tab可停留。
//
// Get Tab can stay.
func (w *TWinControl) TabStop() bool {
	return WinControl_GetTabStop(w.instance)
}

// 设置Tab可停留。
//
// Set Tab can stay.
func (w *TWinControl) SetTabStop(value bool) {
	WinControl_SetTabStop(w.instance, value)
}

// 获取使用停靠管理。
func (w *TWinControl) UseDockManager() bool {
	return WinControl_GetUseDockManager(w.instance)
}

// 设置使用停靠管理。
func (w *TWinControl) SetUseDockManager(value bool) {
	WinControl_SetUseDockManager(w.instance, value)
}

// 获取控件启用。
//
// Get the control enabled.
func (w *TWinControl) Enabled() bool {
	return WinControl_GetEnabled(w.instance)
}

// 设置控件启用。
//
// Set the control enabled.
func (w *TWinControl) SetEnabled(value bool) {
	WinControl_SetEnabled(w.instance, value)
}

func (w *TWinControl) Action() *TAction {
	return AsAction(WinControl_GetAction(w.instance))
}

func (w *TWinControl) SetAction(value IComponent) {
	WinControl_SetAction(w.instance, CheckPtr(value))
}

// 获取控件自动调整。
//
// Get Control automatically adjusts.
func (w *TWinControl) Align() TAlign {
	return WinControl_GetAlign(w.instance)
}

// 设置控件自动调整。
//
// Set Control automatically adjusts.
func (w *TWinControl) SetAlign(value TAlign) {
	WinControl_SetAlign(w.instance, value)
}

// 获取四个角位置的锚点。
func (w *TWinControl) Anchors() TAnchors {
	return WinControl_GetAnchors(w.instance)
}

// 设置四个角位置的锚点。
func (w *TWinControl) SetAnchors(value TAnchors) {
	WinControl_SetAnchors(w.instance, value)
}

func (w *TWinControl) BiDiMode() TBiDiMode {
	return WinControl_GetBiDiMode(w.instance)
}

func (w *TWinControl) SetBiDiMode(value TBiDiMode) {
	WinControl_SetBiDiMode(w.instance, value)
}

func (w *TWinControl) BoundsRect() TRect {
	return WinControl_GetBoundsRect(w.instance)
}

func (w *TWinControl) SetBoundsRect(value TRect) {
	WinControl_SetBoundsRect(w.instance, value)
}

// 获取客户区高度。
//
// Get client height.
func (w *TWinControl) ClientHeight() int32 {
	return WinControl_GetClientHeight(w.instance)
}

// 设置客户区高度。
//
// Set client height.
func (w *TWinControl) SetClientHeight(value int32) {
	WinControl_SetClientHeight(w.instance, value)
}

func (w *TWinControl) ClientOrigin() TPoint {
	return WinControl_GetClientOrigin(w.instance)
}

// 获取客户区矩形。
//
// Get client rectangle.
func (w *TWinControl) ClientRect() TRect {
	return WinControl_GetClientRect(w.instance)
}

// 获取客户区宽度。
//
// Get client width.
func (w *TWinControl) ClientWidth() int32 {
	return WinControl_GetClientWidth(w.instance)
}

// 设置客户区宽度。
//
// Set client width.
func (w *TWinControl) SetClientWidth(value int32) {
	WinControl_SetClientWidth(w.instance, value)
}

// 获取约束控件大小。
func (w *TWinControl) Constraints() *TSizeConstraints {
	return AsSizeConstraints(WinControl_GetConstraints(w.instance))
}

// 设置约束控件大小。
func (w *TWinControl) SetConstraints(value *TSizeConstraints) {
	WinControl_SetConstraints(w.instance, CheckPtr(value))
}

// 获取控件状态。
//
// Get control state.
func (w *TWinControl) ControlState() TControlState {
	return WinControl_GetControlState(w.instance)
}

// 设置控件状态。
//
// Set control state.
func (w *TWinControl) SetControlState(value TControlState) {
	WinControl_SetControlState(w.instance, value)
}

// 获取控件样式。
//
// Get control style.
func (w *TWinControl) ControlStyle() TControlStyle {
	return WinControl_GetControlStyle(w.instance)
}

// 设置控件样式。
//
// Set control style.
func (w *TWinControl) SetControlStyle(value TControlStyle) {
	WinControl_SetControlStyle(w.instance, value)
}

func (w *TWinControl) Floating() bool {
	return WinControl_GetFloating(w.instance)
}

// 获取显示鼠标悬停提示。
//
// Get Show mouseover tips.
func (w *TWinControl) ShowHint() bool {
	return WinControl_GetShowHint(w.instance)
}

// 设置显示鼠标悬停提示。
//
// Set Show mouseover tips.
func (w *TWinControl) SetShowHint(value bool) {
	WinControl_SetShowHint(w.instance, value)
}

// 获取控件可视。
//
// Get the control visible.
func (w *TWinControl) Visible() bool {
	return WinControl_GetVisible(w.instance)
}

// 设置控件可视。
//
// Set the control visible.
func (w *TWinControl) SetVisible(value bool) {
	WinControl_SetVisible(w.instance, value)
}

// 获取控件父容器。
//
// Get control parent container.
func (w *TWinControl) Parent() *TWinControl {
	return AsWinControl(WinControl_GetParent(w.instance))
}

// 设置控件父容器。
//
// Set control parent container.
func (w *TWinControl) SetParent(value IWinControl) {
	WinControl_SetParent(w.instance, CheckPtr(value))
}

// 获取左边位置。
//
// Get Left position.
func (w *TWinControl) Left() int32 {
	return WinControl_GetLeft(w.instance)
}

// 设置左边位置。
//
// Set Left position.
func (w *TWinControl) SetLeft(value int32) {
	WinControl_SetLeft(w.instance, value)
}

// 获取顶边位置。
//
// Get Top position.
func (w *TWinControl) Top() int32 {
	return WinControl_GetTop(w.instance)
}

// 设置顶边位置。
//
// Set Top position.
func (w *TWinControl) SetTop(value int32) {
	WinControl_SetTop(w.instance, value)
}

// 获取宽度。
//
// Get width.
func (w *TWinControl) Width() int32 {
	return WinControl_GetWidth(w.instance)
}

// 设置宽度。
//
// Set width.
func (w *TWinControl) SetWidth(value int32) {
	WinControl_SetWidth(w.instance, value)
}

// 获取高度。
//
// Get height.
func (w *TWinControl) Height() int32 {
	return WinControl_GetHeight(w.instance)
}

// 设置高度。
//
// Set height.
func (w *TWinControl) SetHeight(value int32) {
	WinControl_SetHeight(w.instance, value)
}

// 获取控件光标。
//
// Get control cursor.
func (w *TWinControl) Cursor() TCursor {
	return WinControl_GetCursor(w.instance)
}

// 设置控件光标。
//
// Set control cursor.
func (w *TWinControl) SetCursor(value TCursor) {
	WinControl_SetCursor(w.instance, value)
}

// 获取组件鼠标悬停提示。
//
// Get component mouse hints.
func (w *TWinControl) Hint() string {
	return WinControl_GetHint(w.instance)
}

// 设置组件鼠标悬停提示。
//
// Set component mouse hints.
func (w *TWinControl) SetHint(value string) {
	WinControl_SetHint(w.instance, value)
}

// 获取组件总数。
//
// Get the total number of components.
func (w *TWinControl) ComponentCount() int32 {
	return WinControl_GetComponentCount(w.instance)
}

// 获取组件索引。
//
// Get component index.
func (w *TWinControl) ComponentIndex() int32 {
	return WinControl_GetComponentIndex(w.instance)
}

// 设置组件索引。
//
// Set component index.
func (w *TWinControl) SetComponentIndex(value int32) {
	WinControl_SetComponentIndex(w.instance, value)
}

// 获取组件所有者。
//
// Get component owner.
func (w *TWinControl) Owner() *TComponent {
	return AsComponent(WinControl_GetOwner(w.instance))
}

// 获取组件名称。
//
// Get the component name.
func (w *TWinControl) Name() string {
	return WinControl_GetName(w.instance)
}

// 设置组件名称。
//
// Set the component name.
func (w *TWinControl) SetName(value string) {
	WinControl_SetName(w.instance, value)
}

// 获取对象标记。
//
// Get the control tag.
func (w *TWinControl) Tag() int {
	return WinControl_GetTag(w.instance)
}

// 设置对象标记。
//
// Set the control tag.
func (w *TWinControl) SetTag(value int) {
	WinControl_SetTag(w.instance, value)
}

// 获取左边锚点。
func (w *TWinControl) AnchorSideLeft() *TAnchorSide {
	return AsAnchorSide(WinControl_GetAnchorSideLeft(w.instance))
}

// 设置左边锚点。
func (w *TWinControl) SetAnchorSideLeft(value *TAnchorSide) {
	WinControl_SetAnchorSideLeft(w.instance, CheckPtr(value))
}

// 获取顶边锚点。
func (w *TWinControl) AnchorSideTop() *TAnchorSide {
	return AsAnchorSide(WinControl_GetAnchorSideTop(w.instance))
}

// 设置顶边锚点。
func (w *TWinControl) SetAnchorSideTop(value *TAnchorSide) {
	WinControl_SetAnchorSideTop(w.instance, CheckPtr(value))
}

// 获取右边锚点。
func (w *TWinControl) AnchorSideRight() *TAnchorSide {
	return AsAnchorSide(WinControl_GetAnchorSideRight(w.instance))
}

// 设置右边锚点。
func (w *TWinControl) SetAnchorSideRight(value *TAnchorSide) {
	WinControl_SetAnchorSideRight(w.instance, CheckPtr(value))
}

// 获取底边锚点。
func (w *TWinControl) AnchorSideBottom() *TAnchorSide {
	return AsAnchorSide(WinControl_GetAnchorSideBottom(w.instance))
}

// 设置底边锚点。
func (w *TWinControl) SetAnchorSideBottom(value *TAnchorSide) {
	WinControl_SetAnchorSideBottom(w.instance, CheckPtr(value))
}

func (w *TWinControl) ChildSizing() *TControlChildSizing {
	return AsControlChildSizing(WinControl_GetChildSizing(w.instance))
}

func (w *TWinControl) SetChildSizing(value *TControlChildSizing) {
	WinControl_SetChildSizing(w.instance, CheckPtr(value))
}

// 获取边框间距。
func (w *TWinControl) BorderSpacing() *TControlBorderSpacing {
	return AsControlBorderSpacing(WinControl_GetBorderSpacing(w.instance))
}

// 设置边框间距。
func (w *TWinControl) SetBorderSpacing(value *TControlBorderSpacing) {
	WinControl_SetBorderSpacing(w.instance, CheckPtr(value))
}

// 获取指定索引停靠客户端。
func (w *TWinControl) DockClients(Index int32) *TControl {
	return AsControl(WinControl_GetDockClients(w.instance, Index))
}

// 获取指定索引子控件。
func (w *TWinControl) Controls(Index int32) *TControl {
	return AsControl(WinControl_GetControls(w.instance, Index))
}

// 获取指定索引组件。
//
// Get the specified index component.
func (w *TWinControl) Components(AIndex int32) *TComponent {
	return AsComponent(WinControl_GetComponents(w.instance, AIndex))
}

// 获取锚侧面。
func (w *TWinControl) AnchorSide(AKind TAnchorKind) *TAnchorSide {
	return AsAnchorSide(WinControl_GetAnchorSide(w.instance, AKind))
}
