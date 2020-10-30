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

type TMaskEdit struct {
	IWinControl
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

// 创建一个新的对象。
//
// Create a new object.
func NewMaskEdit(owner IComponent) *TMaskEdit {
	m := new(TMaskEdit)
	m.instance = MaskEdit_Create(CheckPtr(owner))
	m.ptr = unsafe.Pointer(m.instance)
	return m
}

// 动态转换一个已存在的对象实例。
//
// Dynamically convert an existing object instance.
func AsMaskEdit(obj interface{}) *TMaskEdit {
	instance, ptr := getInstance(obj)
	if instance == 0 {
		return nil
	}
	return &TMaskEdit{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// 新建一个对象来自已经存在的对象实例指针。
//
// Create a new object from an existing object instance pointer.
// Deprecated: use AsMaskEdit.
func MaskEditFromInst(inst uintptr) *TMaskEdit {
	return AsMaskEdit(inst)
}

// 新建一个对象来自已经存在的对象实例。
//
// Create a new object from an existing object instance.
// Deprecated: use AsMaskEdit.
func MaskEditFromObj(obj IObject) *TMaskEdit {
	return AsMaskEdit(obj)
}

// 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
//
// Create a new object from an unsecured address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsMaskEdit.
func MaskEditFromUnsafePointer(ptr unsafe.Pointer) *TMaskEdit {
	return AsMaskEdit(ptr)
}

// -------------------------- Deprecated end --------------------------
// 释放对象。
//
// Free object.
func (m *TMaskEdit) Free() {
	if m.instance != 0 {
		MaskEdit_Free(m.instance)
		m.instance, m.ptr = 0, nullptr
	}
}

// 返回对象实例指针。
//
// Return object instance pointer.
func (m *TMaskEdit) Instance() uintptr {
	return m.instance
}

// 获取一个不安全的地址。
//
// Get an unsafe address.
func (m *TMaskEdit) UnsafeAddr() unsafe.Pointer {
	return m.ptr
}

// 检测地址是否为空。
//
// Check if the address is empty.
func (m *TMaskEdit) IsValid() bool {
	return m.instance != 0
}

// 检测当前对象是否继承自目标对象。
//
// Checks whether the current object is inherited from the target object.
func (m *TMaskEdit) Is() TIs {
	return TIs(m.instance)
}

// 动态转换当前对象为目标对象。
//
// Dynamically convert the current object to the target object.
//func (m *TMaskEdit) As() TAs {
//    return TAs(m.instance)
//}

// 获取类信息指针。
//
// Get class information pointer.
func TMaskEditClass() TClass {
	return MaskEdit_StaticClassType()
}

func (m *TMaskEdit) ValidateEdit() {
	MaskEdit_ValidateEdit(m.instance)
}

// 清除。
func (m *TMaskEdit) Clear() {
	MaskEdit_Clear(m.instance)
}

// 获取控件的字符长，如果有。
//
// Get the character length of the control, if any.
func (m *TMaskEdit) GetTextLen() int32 {
	return MaskEdit_GetTextLen(m.instance)
}

// 清除选择。
func (m *TMaskEdit) ClearSelection() {
	MaskEdit_ClearSelection(m.instance)
}

// 复制到粘贴板。
func (m *TMaskEdit) CopyToClipboard() {
	MaskEdit_CopyToClipboard(m.instance)
}

// 剪切到粘贴板。
func (m *TMaskEdit) CutToClipboard() {
	MaskEdit_CutToClipboard(m.instance)
}

// 从剪切板粘贴。
func (m *TMaskEdit) PasteFromClipboard() {
	MaskEdit_PasteFromClipboard(m.instance)
}

// 撤销上一次操作。
func (m *TMaskEdit) Undo() {
	MaskEdit_Undo(m.instance)
}

// 全选。
func (m *TMaskEdit) SelectAll() {
	MaskEdit_SelectAll(m.instance)
}

// 是否可以获得焦点。
func (m *TMaskEdit) CanFocus() bool {
	return MaskEdit_CanFocus(m.instance)
}

// 返回是否包含指定控件。
//
// it's contain a specified control.
func (m *TMaskEdit) ContainsControl(Control IControl) bool {
	return MaskEdit_ContainsControl(m.instance, CheckPtr(Control))
}

// 返回指定坐标及相关属性位置控件。
//
// Returns the specified coordinate and the relevant attribute position control..
func (m *TMaskEdit) ControlAtPos(Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) *TControl {
	return AsControl(MaskEdit_ControlAtPos(m.instance, Pos, AllowDisabled, AllowWinControls, AllLevels))
}

// 禁用控件的对齐。
//
// Disable control alignment.
func (m *TMaskEdit) DisableAlign() {
	MaskEdit_DisableAlign(m.instance)
}

// 启用控件对齐。
//
// Enabled control alignment.
func (m *TMaskEdit) EnableAlign() {
	MaskEdit_EnableAlign(m.instance)
}

// 查找子控件。
//
// Find sub controls.
func (m *TMaskEdit) FindChildControl(ControlName string) *TControl {
	return AsControl(MaskEdit_FindChildControl(m.instance, ControlName))
}

func (m *TMaskEdit) FlipChildren(AllLevels bool) {
	MaskEdit_FlipChildren(m.instance, AllLevels)
}

// 返回是否获取焦点。
//
// Return to get focus.
func (m *TMaskEdit) Focused() bool {
	return MaskEdit_Focused(m.instance)
}

// 句柄是否已经分配。
//
// Is the handle already allocated.
func (m *TMaskEdit) HandleAllocated() bool {
	return MaskEdit_HandleAllocated(m.instance)
}

// 插入一个控件。
//
// Insert a control.
func (m *TMaskEdit) InsertControl(AControl IControl) {
	MaskEdit_InsertControl(m.instance, CheckPtr(AControl))
}

// 要求重绘。
//
// Redraw.
func (m *TMaskEdit) Invalidate() {
	MaskEdit_Invalidate(m.instance)
}

// 绘画至指定DC。
//
// Painting to the specified DC.
func (m *TMaskEdit) PaintTo(DC HDC, X int32, Y int32) {
	MaskEdit_PaintTo(m.instance, DC, X, Y)
}

// 移除一个控件。
//
// Remove a control.
func (m *TMaskEdit) RemoveControl(AControl IControl) {
	MaskEdit_RemoveControl(m.instance, CheckPtr(AControl))
}

// 重新对齐。
//
// Realign.
func (m *TMaskEdit) Realign() {
	MaskEdit_Realign(m.instance)
}

// 重绘。
//
// Repaint.
func (m *TMaskEdit) Repaint() {
	MaskEdit_Repaint(m.instance)
}

// 按比例缩放。
//
// Scale by.
func (m *TMaskEdit) ScaleBy(M int32, D int32) {
	MaskEdit_ScaleBy(m.instance, M, D)
}

// 滚动至指定位置。
//
// Scroll by.
func (m *TMaskEdit) ScrollBy(DeltaX int32, DeltaY int32) {
	MaskEdit_ScrollBy(m.instance, DeltaX, DeltaY)
}

// 设置组件边界。
//
// Set component boundaries.
func (m *TMaskEdit) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
	MaskEdit_SetBounds(m.instance, ALeft, ATop, AWidth, AHeight)
}

// 设置控件焦点。
//
// Set control focus.
func (m *TMaskEdit) SetFocus() {
	MaskEdit_SetFocus(m.instance)
}

// 控件更新。
//
// Update.
func (m *TMaskEdit) Update() {
	MaskEdit_Update(m.instance)
}

// 将控件置于最前。
//
// Bring the control to the front.
func (m *TMaskEdit) BringToFront() {
	MaskEdit_BringToFront(m.instance)
}

// 将客户端坐标转为绝对的屏幕坐标。
//
// Convert client coordinates to absolute screen coordinates.
func (m *TMaskEdit) ClientToScreen(Point TPoint) TPoint {
	return MaskEdit_ClientToScreen(m.instance, Point)
}

// 将客户端坐标转为父容器坐标。
//
// Convert client coordinates to parent container coordinates.
func (m *TMaskEdit) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
	return MaskEdit_ClientToParent(m.instance, Point, CheckPtr(AParent))
}

// 是否在拖拽中。
//
// Is it in the middle of dragging.
func (m *TMaskEdit) Dragging() bool {
	return MaskEdit_Dragging(m.instance)
}

// 是否有父容器。
//
// Is there a parent container.
func (m *TMaskEdit) HasParent() bool {
	return MaskEdit_HasParent(m.instance)
}

// 隐藏控件。
//
// Hidden control.
func (m *TMaskEdit) Hide() {
	MaskEdit_Hide(m.instance)
}

// 发送一个消息。
//
// Send a message.
func (m *TMaskEdit) Perform(Msg uint32, WParam uintptr, LParam int) int {
	return MaskEdit_Perform(m.instance, Msg, WParam, LParam)
}

// 刷新控件。
//
// Refresh control.
func (m *TMaskEdit) Refresh() {
	MaskEdit_Refresh(m.instance)
}

// 将屏幕坐标转为客户端坐标。
//
// Convert screen coordinates to client coordinates.
func (m *TMaskEdit) ScreenToClient(Point TPoint) TPoint {
	return MaskEdit_ScreenToClient(m.instance, Point)
}

// 将父容器坐标转为客户端坐标。
//
// Convert parent container coordinates to client coordinates.
func (m *TMaskEdit) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
	return MaskEdit_ParentToClient(m.instance, Point, CheckPtr(AParent))
}

// 控件至于最后面。
//
// The control is placed at the end.
func (m *TMaskEdit) SendToBack() {
	MaskEdit_SendToBack(m.instance)
}

// 显示控件。
//
// Show control.
func (m *TMaskEdit) Show() {
	MaskEdit_Show(m.instance)
}

// 获取控件的字符，如果有。
//
// Get the characters of the control, if any.
func (m *TMaskEdit) GetTextBuf(Buffer *string, BufSize int32) int32 {
	return MaskEdit_GetTextBuf(m.instance, Buffer, BufSize)
}

// 设置控件字符，如果有。
//
// Set control characters, if any.
func (m *TMaskEdit) SetTextBuf(Buffer string) {
	MaskEdit_SetTextBuf(m.instance, Buffer)
}

// 查找指定名称的组件。
//
// Find the component with the specified name.
func (m *TMaskEdit) FindComponent(AName string) *TComponent {
	return AsComponent(MaskEdit_FindComponent(m.instance, AName))
}

// 获取类名路径。
//
// Get the class name path.
func (m *TMaskEdit) GetNamePath() string {
	return MaskEdit_GetNamePath(m.instance)
}

// 复制一个对象，如果对象实现了此方法的话。
//
// Copy an object, if the object implements this method.
func (m *TMaskEdit) Assign(Source IObject) {
	MaskEdit_Assign(m.instance, CheckPtr(Source))
}

// 获取类的类型信息。
//
// Get class type information.
func (m *TMaskEdit) ClassType() TClass {
	return MaskEdit_ClassType(m.instance)
}

// 获取当前对象类名称。
//
// Get the current object class name.
func (m *TMaskEdit) ClassName() string {
	return MaskEdit_ClassName(m.instance)
}

// 获取当前对象实例大小。
//
// Get the current object instance size.
func (m *TMaskEdit) InstanceSize() int32 {
	return MaskEdit_InstanceSize(m.instance)
}

// 判断当前类是否继承自指定类。
//
// Determine whether the current class inherits from the specified class.
func (m *TMaskEdit) InheritsFrom(AClass TClass) bool {
	return MaskEdit_InheritsFrom(m.instance, AClass)
}

// 与一个对象进行比较。
//
// Compare with an object.
func (m *TMaskEdit) Equals(Obj IObject) bool {
	return MaskEdit_Equals(m.instance, CheckPtr(Obj))
}

// 获取类的哈希值。
//
// Get the hash value of the class.
func (m *TMaskEdit) GetHashCode() int32 {
	return MaskEdit_GetHashCode(m.instance)
}

// 文本类信息。
//
// Text information.
func (m *TMaskEdit) ToString() string {
	return MaskEdit_ToString(m.instance)
}

func (m *TMaskEdit) AnchorToNeighbour(ASide TAnchorKind, ASpace int32, ASibling IControl) {
	MaskEdit_AnchorToNeighbour(m.instance, ASide, ASpace, CheckPtr(ASibling))
}

func (m *TMaskEdit) AnchorParallel(ASide TAnchorKind, ASpace int32, ASibling IControl) {
	MaskEdit_AnchorParallel(m.instance, ASide, ASpace, CheckPtr(ASibling))
}

// 置于指定控件的横向中心。
func (m *TMaskEdit) AnchorHorizontalCenterTo(ASibling IControl) {
	MaskEdit_AnchorHorizontalCenterTo(m.instance, CheckPtr(ASibling))
}

// 置于指定控件的纵向中心。
func (m *TMaskEdit) AnchorVerticalCenterTo(ASibling IControl) {
	MaskEdit_AnchorVerticalCenterTo(m.instance, CheckPtr(ASibling))
}

func (m *TMaskEdit) AnchorSame(ASide TAnchorKind, ASibling IControl) {
	MaskEdit_AnchorSame(m.instance, ASide, CheckPtr(ASibling))
}

func (m *TMaskEdit) AnchorAsAlign(ATheAlign TAlign, ASpace int32) {
	MaskEdit_AnchorAsAlign(m.instance, ATheAlign, ASpace)
}

func (m *TMaskEdit) AnchorClient(ASpace int32) {
	MaskEdit_AnchorClient(m.instance, ASpace)
}

func (m *TMaskEdit) ScaleDesignToForm(ASize int32) int32 {
	return MaskEdit_ScaleDesignToForm(m.instance, ASize)
}

func (m *TMaskEdit) ScaleFormToDesign(ASize int32) int32 {
	return MaskEdit_ScaleFormToDesign(m.instance, ASize)
}

func (m *TMaskEdit) Scale96ToForm(ASize int32) int32 {
	return MaskEdit_Scale96ToForm(m.instance, ASize)
}

func (m *TMaskEdit) ScaleFormTo96(ASize int32) int32 {
	return MaskEdit_ScaleFormTo96(m.instance, ASize)
}

func (m *TMaskEdit) Scale96ToFont(ASize int32) int32 {
	return MaskEdit_Scale96ToFont(m.instance, ASize)
}

func (m *TMaskEdit) ScaleFontTo96(ASize int32) int32 {
	return MaskEdit_ScaleFontTo96(m.instance, ASize)
}

func (m *TMaskEdit) ScaleScreenToFont(ASize int32) int32 {
	return MaskEdit_ScaleScreenToFont(m.instance, ASize)
}

func (m *TMaskEdit) ScaleFontToScreen(ASize int32) int32 {
	return MaskEdit_ScaleFontToScreen(m.instance, ASize)
}

func (m *TMaskEdit) Scale96ToScreen(ASize int32) int32 {
	return MaskEdit_Scale96ToScreen(m.instance, ASize)
}

func (m *TMaskEdit) ScaleScreenTo96(ASize int32) int32 {
	return MaskEdit_ScaleScreenTo96(m.instance, ASize)
}

func (m *TMaskEdit) AutoAdjustLayout(AMode TLayoutAdjustmentPolicy, AFromPPI int32, AToPPI int32, AOldFormWidth int32, ANewFormWidth int32) {
	MaskEdit_AutoAdjustLayout(m.instance, AMode, AFromPPI, AToPPI, AOldFormWidth, ANewFormWidth)
}

func (m *TMaskEdit) FixDesignFontsPPI(ADesignTimePPI int32) {
	MaskEdit_FixDesignFontsPPI(m.instance, ADesignTimePPI)
}

func (m *TMaskEdit) ScaleFontsPPI(AToPPI int32, AProportion float64) {
	MaskEdit_ScaleFontsPPI(m.instance, AToPPI, AProportion)
}

// 获取控件自动调整。
//
// Get Control automatically adjusts.
func (m *TMaskEdit) Align() TAlign {
	return MaskEdit_GetAlign(m.instance)
}

// 设置控件自动调整。
//
// Set Control automatically adjusts.
func (m *TMaskEdit) SetAlign(value TAlign) {
	MaskEdit_SetAlign(m.instance, value)
}

// 获取文字对齐。
//
// Get Text alignment.
func (m *TMaskEdit) Alignment() TAlignment {
	return MaskEdit_GetAlignment(m.instance)
}

// 设置文字对齐。
//
// Set Text alignment.
func (m *TMaskEdit) SetAlignment(value TAlignment) {
	MaskEdit_SetAlignment(m.instance, value)
}

// 获取四个角位置的锚点。
func (m *TMaskEdit) Anchors() TAnchors {
	return MaskEdit_GetAnchors(m.instance)
}

// 设置四个角位置的锚点。
func (m *TMaskEdit) SetAnchors(value TAnchors) {
	MaskEdit_SetAnchors(m.instance, value)
}

// 获取自动选择。
func (m *TMaskEdit) AutoSelect() bool {
	return MaskEdit_GetAutoSelect(m.instance)
}

// 设置自动选择。
func (m *TMaskEdit) SetAutoSelect(value bool) {
	MaskEdit_SetAutoSelect(m.instance, value)
}

// 获取自动调整大小。
func (m *TMaskEdit) AutoSize() bool {
	return MaskEdit_GetAutoSize(m.instance)
}

// 设置自动调整大小。
func (m *TMaskEdit) SetAutoSize(value bool) {
	MaskEdit_SetAutoSize(m.instance, value)
}

func (m *TMaskEdit) BiDiMode() TBiDiMode {
	return MaskEdit_GetBiDiMode(m.instance)
}

func (m *TMaskEdit) SetBiDiMode(value TBiDiMode) {
	MaskEdit_SetBiDiMode(m.instance, value)
}

// 获取窗口边框样式。比如：无边框，单一边框等。
func (m *TMaskEdit) BorderStyle() TBorderStyle {
	return MaskEdit_GetBorderStyle(m.instance)
}

// 设置窗口边框样式。比如：无边框，单一边框等。
func (m *TMaskEdit) SetBorderStyle(value TBorderStyle) {
	MaskEdit_SetBorderStyle(m.instance, value)
}

func (m *TMaskEdit) CharCase() TEditCharCase {
	return MaskEdit_GetCharCase(m.instance)
}

func (m *TMaskEdit) SetCharCase(value TEditCharCase) {
	MaskEdit_SetCharCase(m.instance, value)
}

// 获取颜色。
//
// Get color.
func (m *TMaskEdit) Color() TColor {
	return MaskEdit_GetColor(m.instance)
}

// 设置颜色。
//
// Set color.
func (m *TMaskEdit) SetColor(value TColor) {
	MaskEdit_SetColor(m.instance, value)
}

// 获取约束控件大小。
func (m *TMaskEdit) Constraints() *TSizeConstraints {
	return AsSizeConstraints(MaskEdit_GetConstraints(m.instance))
}

// 设置约束控件大小。
func (m *TMaskEdit) SetConstraints(value *TSizeConstraints) {
	MaskEdit_SetConstraints(m.instance, CheckPtr(value))
}

// 获取设置控件双缓冲。
//
// Get Set control double buffering.
func (m *TMaskEdit) DoubleBuffered() bool {
	return MaskEdit_GetDoubleBuffered(m.instance)
}

// 设置设置控件双缓冲。
//
// Set Set control double buffering.
func (m *TMaskEdit) SetDoubleBuffered(value bool) {
	MaskEdit_SetDoubleBuffered(m.instance, value)
}

// 获取设置控件拖拽时的光标。
//
// Get Set the cursor when the control is dragged.
func (m *TMaskEdit) DragCursor() TCursor {
	return MaskEdit_GetDragCursor(m.instance)
}

// 设置设置控件拖拽时的光标。
//
// Set Set the cursor when the control is dragged.
func (m *TMaskEdit) SetDragCursor(value TCursor) {
	MaskEdit_SetDragCursor(m.instance, value)
}

// 获取拖拽方式。
//
// Get Drag and drop.
func (m *TMaskEdit) DragKind() TDragKind {
	return MaskEdit_GetDragKind(m.instance)
}

// 设置拖拽方式。
//
// Set Drag and drop.
func (m *TMaskEdit) SetDragKind(value TDragKind) {
	MaskEdit_SetDragKind(m.instance, value)
}

// 获取拖拽模式。
//
// Get Drag mode.
func (m *TMaskEdit) DragMode() TDragMode {
	return MaskEdit_GetDragMode(m.instance)
}

// 设置拖拽模式。
//
// Set Drag mode.
func (m *TMaskEdit) SetDragMode(value TDragMode) {
	MaskEdit_SetDragMode(m.instance, value)
}

// 获取控件启用。
//
// Get the control enabled.
func (m *TMaskEdit) Enabled() bool {
	return MaskEdit_GetEnabled(m.instance)
}

// 设置控件启用。
//
// Set the control enabled.
func (m *TMaskEdit) SetEnabled(value bool) {
	MaskEdit_SetEnabled(m.instance, value)
}

// 获取字体。
//
// Get Font.
func (m *TMaskEdit) Font() *TFont {
	return AsFont(MaskEdit_GetFont(m.instance))
}

// 设置字体。
//
// Set Font.
func (m *TMaskEdit) SetFont(value *TFont) {
	MaskEdit_SetFont(m.instance, CheckPtr(value))
}

// 获取最大长度。
func (m *TMaskEdit) MaxLength() int32 {
	return MaskEdit_GetMaxLength(m.instance)
}

// 设置最大长度。
func (m *TMaskEdit) SetMaxLength(value int32) {
	MaskEdit_SetMaxLength(m.instance, value)
}

// 获取使用父容器颜色。
//
// Get parent color.
func (m *TMaskEdit) ParentColor() bool {
	return MaskEdit_GetParentColor(m.instance)
}

// 设置使用父容器颜色。
//
// Set parent color.
func (m *TMaskEdit) SetParentColor(value bool) {
	MaskEdit_SetParentColor(m.instance, value)
}

// 获取使用父容器双缓冲。
//
// Get Parent container double buffering.
func (m *TMaskEdit) ParentDoubleBuffered() bool {
	return MaskEdit_GetParentDoubleBuffered(m.instance)
}

// 设置使用父容器双缓冲。
//
// Set Parent container double buffering.
func (m *TMaskEdit) SetParentDoubleBuffered(value bool) {
	MaskEdit_SetParentDoubleBuffered(m.instance, value)
}

// 获取使用父容器字体。
//
// Get Parent container font.
func (m *TMaskEdit) ParentFont() bool {
	return MaskEdit_GetParentFont(m.instance)
}

// 设置使用父容器字体。
//
// Set Parent container font.
func (m *TMaskEdit) SetParentFont(value bool) {
	MaskEdit_SetParentFont(m.instance, value)
}

// 获取以父容器的ShowHint属性为准。
func (m *TMaskEdit) ParentShowHint() bool {
	return MaskEdit_GetParentShowHint(m.instance)
}

// 设置以父容器的ShowHint属性为准。
func (m *TMaskEdit) SetParentShowHint(value bool) {
	MaskEdit_SetParentShowHint(m.instance, value)
}

// 获取密码掩码字符。
func (m *TMaskEdit) PasswordChar() uint16 {
	return MaskEdit_GetPasswordChar(m.instance)
}

// 设置密码掩码字符。
func (m *TMaskEdit) SetPasswordChar(value uint16) {
	MaskEdit_SetPasswordChar(m.instance, value)
}

// 获取右键菜单。
//
// Get Right click menu.
func (m *TMaskEdit) PopupMenu() *TPopupMenu {
	return AsPopupMenu(MaskEdit_GetPopupMenu(m.instance))
}

// 设置右键菜单。
//
// Set Right click menu.
func (m *TMaskEdit) SetPopupMenu(value IComponent) {
	MaskEdit_SetPopupMenu(m.instance, CheckPtr(value))
}

// 获取只读。
func (m *TMaskEdit) ReadOnly() bool {
	return MaskEdit_GetReadOnly(m.instance)
}

// 设置只读。
func (m *TMaskEdit) SetReadOnly(value bool) {
	MaskEdit_SetReadOnly(m.instance, value)
}

// 获取显示鼠标悬停提示。
//
// Get Show mouseover tips.
func (m *TMaskEdit) ShowHint() bool {
	return MaskEdit_GetShowHint(m.instance)
}

// 设置显示鼠标悬停提示。
//
// Set Show mouseover tips.
func (m *TMaskEdit) SetShowHint(value bool) {
	MaskEdit_SetShowHint(m.instance, value)
}

// 获取Tab切换顺序序号。
//
// Get Tab switching sequence number.
func (m *TMaskEdit) TabOrder() TTabOrder {
	return MaskEdit_GetTabOrder(m.instance)
}

// 设置Tab切换顺序序号。
//
// Set Tab switching sequence number.
func (m *TMaskEdit) SetTabOrder(value TTabOrder) {
	MaskEdit_SetTabOrder(m.instance, value)
}

// 获取Tab可停留。
//
// Get Tab can stay.
func (m *TMaskEdit) TabStop() bool {
	return MaskEdit_GetTabStop(m.instance)
}

// 设置Tab可停留。
//
// Set Tab can stay.
func (m *TMaskEdit) SetTabStop(value bool) {
	MaskEdit_SetTabStop(m.instance, value)
}

// 获取文本。
func (m *TMaskEdit) Text() string {
	strLen := m.GetTextLen()
	if strLen != 0 {
		var buffStr string
		m.GetTextBuf(&buffStr, strLen+1)
		return buffStr
	}
	return ""
}

// 设置文本。
func (m *TMaskEdit) SetText(value string) {
	MaskEdit_SetText(m.instance, value)
}

// 获取提示文本。
func (m *TMaskEdit) TextHint() string {
	return MaskEdit_GetTextHint(m.instance)
}

// 设置提示文本。
func (m *TMaskEdit) SetTextHint(value string) {
	MaskEdit_SetTextHint(m.instance, value)
}

// 获取控件可视。
//
// Get the control visible.
func (m *TMaskEdit) Visible() bool {
	return MaskEdit_GetVisible(m.instance)
}

// 设置控件可视。
//
// Set the control visible.
func (m *TMaskEdit) SetVisible(value bool) {
	MaskEdit_SetVisible(m.instance, value)
}

// 设置改变事件。
//
// Set changed event.
func (m *TMaskEdit) SetOnChange(fn TNotifyEvent) {
	MaskEdit_SetOnChange(m.instance, fn)
}

// 设置控件单击事件。
//
// Set control click event.
func (m *TMaskEdit) SetOnClick(fn TNotifyEvent) {
	MaskEdit_SetOnClick(m.instance, fn)
}

// 设置双击事件。
func (m *TMaskEdit) SetOnDblClick(fn TNotifyEvent) {
	MaskEdit_SetOnDblClick(m.instance, fn)
}

// 设置拖拽下落事件。
//
// Set Drag and drop event.
func (m *TMaskEdit) SetOnDragDrop(fn TDragDropEvent) {
	MaskEdit_SetOnDragDrop(m.instance, fn)
}

// 设置拖拽完成事件。
//
// Set Drag and drop completion event.
func (m *TMaskEdit) SetOnDragOver(fn TDragOverEvent) {
	MaskEdit_SetOnDragOver(m.instance, fn)
}

// 设置停靠结束事件。
//
// Set Dock end event.
func (m *TMaskEdit) SetOnEndDock(fn TEndDragEvent) {
	MaskEdit_SetOnEndDock(m.instance, fn)
}

// 设置拖拽结束。
//
// Set End of drag.
func (m *TMaskEdit) SetOnEndDrag(fn TEndDragEvent) {
	MaskEdit_SetOnEndDrag(m.instance, fn)
}

// 设置焦点进入。
//
// Set Focus entry.
func (m *TMaskEdit) SetOnEnter(fn TNotifyEvent) {
	MaskEdit_SetOnEnter(m.instance, fn)
}

// 设置焦点退出。
//
// Set Focus exit.
func (m *TMaskEdit) SetOnExit(fn TNotifyEvent) {
	MaskEdit_SetOnExit(m.instance, fn)
}

// 设置键盘按键按下事件。
//
// Set Keyboard button press event.
func (m *TMaskEdit) SetOnKeyDown(fn TKeyEvent) {
	MaskEdit_SetOnKeyDown(m.instance, fn)
}

// 设置键键下事件。
func (m *TMaskEdit) SetOnKeyPress(fn TKeyPressEvent) {
	MaskEdit_SetOnKeyPress(m.instance, fn)
}

// 设置键盘按键抬起事件。
//
// Set Keyboard button lift event.
func (m *TMaskEdit) SetOnKeyUp(fn TKeyEvent) {
	MaskEdit_SetOnKeyUp(m.instance, fn)
}

// 设置鼠标按下事件。
//
// Set Mouse down event.
func (m *TMaskEdit) SetOnMouseDown(fn TMouseEvent) {
	MaskEdit_SetOnMouseDown(m.instance, fn)
}

// 设置鼠标进入事件。
//
// Set Mouse entry event.
func (m *TMaskEdit) SetOnMouseEnter(fn TNotifyEvent) {
	MaskEdit_SetOnMouseEnter(m.instance, fn)
}

// 设置鼠标离开事件。
//
// Set Mouse leave event.
func (m *TMaskEdit) SetOnMouseLeave(fn TNotifyEvent) {
	MaskEdit_SetOnMouseLeave(m.instance, fn)
}

// 设置鼠标移动事件。
func (m *TMaskEdit) SetOnMouseMove(fn TMouseMoveEvent) {
	MaskEdit_SetOnMouseMove(m.instance, fn)
}

// 设置鼠标抬起事件。
//
// Set Mouse lift event.
func (m *TMaskEdit) SetOnMouseUp(fn TMouseEvent) {
	MaskEdit_SetOnMouseUp(m.instance, fn)
}

// 设置启动停靠。
func (m *TMaskEdit) SetOnStartDock(fn TStartDockEvent) {
	MaskEdit_SetOnStartDock(m.instance, fn)
}

func (m *TMaskEdit) IsMasked() bool {
	return MaskEdit_GetIsMasked(m.instance)
}

func (m *TMaskEdit) EditText() string {
	return MaskEdit_GetEditText(m.instance)
}

func (m *TMaskEdit) SetEditText(value string) {
	MaskEdit_SetEditText(m.instance, value)
}

// 获取能否撤销。
func (m *TMaskEdit) CanUndo() bool {
	return MaskEdit_GetCanUndo(m.instance)
}

// 获取修改。
//
// Get modified.
func (m *TMaskEdit) Modified() bool {
	return MaskEdit_GetModified(m.instance)
}

// 设置修改。
//
// Set modified.
func (m *TMaskEdit) SetModified(value bool) {
	MaskEdit_SetModified(m.instance, value)
}

// 获取选择的长度。
func (m *TMaskEdit) SelLength() int32 {
	return MaskEdit_GetSelLength(m.instance)
}

// 设置选择的长度。
func (m *TMaskEdit) SetSelLength(value int32) {
	MaskEdit_SetSelLength(m.instance, value)
}

// 获取选择的启始位置。
func (m *TMaskEdit) SelStart() int32 {
	return MaskEdit_GetSelStart(m.instance)
}

// 设置选择的启始位置。
func (m *TMaskEdit) SetSelStart(value int32) {
	MaskEdit_SetSelStart(m.instance, value)
}

// 获取选择的文本。
func (m *TMaskEdit) SelText() string {
	return MaskEdit_GetSelText(m.instance)
}

// 设置选择的文本。
func (m *TMaskEdit) SetSelText(value string) {
	MaskEdit_SetSelText(m.instance, value)
}

// 获取依靠客户端总数。
func (m *TMaskEdit) DockClientCount() int32 {
	return MaskEdit_GetDockClientCount(m.instance)
}

// 获取停靠站点。
//
// Get Docking site.
func (m *TMaskEdit) DockSite() bool {
	return MaskEdit_GetDockSite(m.instance)
}

// 设置停靠站点。
//
// Set Docking site.
func (m *TMaskEdit) SetDockSite(value bool) {
	MaskEdit_SetDockSite(m.instance, value)
}

// 获取鼠标是否在客户端，仅VCL有效。
//
// Get Whether the mouse is on the client, only VCL is valid.
func (m *TMaskEdit) MouseInClient() bool {
	return MaskEdit_GetMouseInClient(m.instance)
}

// 获取当前停靠的可视总数。
//
// Get The total number of visible calls currently docked.
func (m *TMaskEdit) VisibleDockClientCount() int32 {
	return MaskEdit_GetVisibleDockClientCount(m.instance)
}

// 获取画刷对象。
//
// Get Brush.
func (m *TMaskEdit) Brush() *TBrush {
	return AsBrush(MaskEdit_GetBrush(m.instance))
}

// 获取子控件数。
//
// Get Number of child controls.
func (m *TMaskEdit) ControlCount() int32 {
	return MaskEdit_GetControlCount(m.instance)
}

// 获取控件句柄。
//
// Get Control handle.
func (m *TMaskEdit) Handle() HWND {
	return MaskEdit_GetHandle(m.instance)
}

// 获取父容器句柄。
//
// Get Parent container handle.
func (m *TMaskEdit) ParentWindow() HWND {
	return MaskEdit_GetParentWindow(m.instance)
}

// 设置父容器句柄。
//
// Set Parent container handle.
func (m *TMaskEdit) SetParentWindow(value HWND) {
	MaskEdit_SetParentWindow(m.instance, value)
}

func (m *TMaskEdit) Showing() bool {
	return MaskEdit_GetShowing(m.instance)
}

// 获取使用停靠管理。
func (m *TMaskEdit) UseDockManager() bool {
	return MaskEdit_GetUseDockManager(m.instance)
}

// 设置使用停靠管理。
func (m *TMaskEdit) SetUseDockManager(value bool) {
	MaskEdit_SetUseDockManager(m.instance, value)
}

func (m *TMaskEdit) Action() *TAction {
	return AsAction(MaskEdit_GetAction(m.instance))
}

func (m *TMaskEdit) SetAction(value IComponent) {
	MaskEdit_SetAction(m.instance, CheckPtr(value))
}

func (m *TMaskEdit) BoundsRect() TRect {
	return MaskEdit_GetBoundsRect(m.instance)
}

func (m *TMaskEdit) SetBoundsRect(value TRect) {
	MaskEdit_SetBoundsRect(m.instance, value)
}

// 获取客户区高度。
//
// Get client height.
func (m *TMaskEdit) ClientHeight() int32 {
	return MaskEdit_GetClientHeight(m.instance)
}

// 设置客户区高度。
//
// Set client height.
func (m *TMaskEdit) SetClientHeight(value int32) {
	MaskEdit_SetClientHeight(m.instance, value)
}

func (m *TMaskEdit) ClientOrigin() TPoint {
	return MaskEdit_GetClientOrigin(m.instance)
}

// 获取客户区矩形。
//
// Get client rectangle.
func (m *TMaskEdit) ClientRect() TRect {
	return MaskEdit_GetClientRect(m.instance)
}

// 获取客户区宽度。
//
// Get client width.
func (m *TMaskEdit) ClientWidth() int32 {
	return MaskEdit_GetClientWidth(m.instance)
}

// 设置客户区宽度。
//
// Set client width.
func (m *TMaskEdit) SetClientWidth(value int32) {
	MaskEdit_SetClientWidth(m.instance, value)
}

// 获取控件状态。
//
// Get control state.
func (m *TMaskEdit) ControlState() TControlState {
	return MaskEdit_GetControlState(m.instance)
}

// 设置控件状态。
//
// Set control state.
func (m *TMaskEdit) SetControlState(value TControlState) {
	MaskEdit_SetControlState(m.instance, value)
}

// 获取控件样式。
//
// Get control style.
func (m *TMaskEdit) ControlStyle() TControlStyle {
	return MaskEdit_GetControlStyle(m.instance)
}

// 设置控件样式。
//
// Set control style.
func (m *TMaskEdit) SetControlStyle(value TControlStyle) {
	MaskEdit_SetControlStyle(m.instance, value)
}

func (m *TMaskEdit) Floating() bool {
	return MaskEdit_GetFloating(m.instance)
}

// 获取控件父容器。
//
// Get control parent container.
func (m *TMaskEdit) Parent() *TWinControl {
	return AsWinControl(MaskEdit_GetParent(m.instance))
}

// 设置控件父容器。
//
// Set control parent container.
func (m *TMaskEdit) SetParent(value IWinControl) {
	MaskEdit_SetParent(m.instance, CheckPtr(value))
}

// 获取左边位置。
//
// Get Left position.
func (m *TMaskEdit) Left() int32 {
	return MaskEdit_GetLeft(m.instance)
}

// 设置左边位置。
//
// Set Left position.
func (m *TMaskEdit) SetLeft(value int32) {
	MaskEdit_SetLeft(m.instance, value)
}

// 获取顶边位置。
//
// Get Top position.
func (m *TMaskEdit) Top() int32 {
	return MaskEdit_GetTop(m.instance)
}

// 设置顶边位置。
//
// Set Top position.
func (m *TMaskEdit) SetTop(value int32) {
	MaskEdit_SetTop(m.instance, value)
}

// 获取宽度。
//
// Get width.
func (m *TMaskEdit) Width() int32 {
	return MaskEdit_GetWidth(m.instance)
}

// 设置宽度。
//
// Set width.
func (m *TMaskEdit) SetWidth(value int32) {
	MaskEdit_SetWidth(m.instance, value)
}

// 获取高度。
//
// Get height.
func (m *TMaskEdit) Height() int32 {
	return MaskEdit_GetHeight(m.instance)
}

// 设置高度。
//
// Set height.
func (m *TMaskEdit) SetHeight(value int32) {
	MaskEdit_SetHeight(m.instance, value)
}

// 获取控件光标。
//
// Get control cursor.
func (m *TMaskEdit) Cursor() TCursor {
	return MaskEdit_GetCursor(m.instance)
}

// 设置控件光标。
//
// Set control cursor.
func (m *TMaskEdit) SetCursor(value TCursor) {
	MaskEdit_SetCursor(m.instance, value)
}

// 获取组件鼠标悬停提示。
//
// Get component mouse hints.
func (m *TMaskEdit) Hint() string {
	return MaskEdit_GetHint(m.instance)
}

// 设置组件鼠标悬停提示。
//
// Set component mouse hints.
func (m *TMaskEdit) SetHint(value string) {
	MaskEdit_SetHint(m.instance, value)
}

// 获取组件总数。
//
// Get the total number of components.
func (m *TMaskEdit) ComponentCount() int32 {
	return MaskEdit_GetComponentCount(m.instance)
}

// 获取组件索引。
//
// Get component index.
func (m *TMaskEdit) ComponentIndex() int32 {
	return MaskEdit_GetComponentIndex(m.instance)
}

// 设置组件索引。
//
// Set component index.
func (m *TMaskEdit) SetComponentIndex(value int32) {
	MaskEdit_SetComponentIndex(m.instance, value)
}

// 获取组件所有者。
//
// Get component owner.
func (m *TMaskEdit) Owner() *TComponent {
	return AsComponent(MaskEdit_GetOwner(m.instance))
}

// 获取组件名称。
//
// Get the component name.
func (m *TMaskEdit) Name() string {
	return MaskEdit_GetName(m.instance)
}

// 设置组件名称。
//
// Set the component name.
func (m *TMaskEdit) SetName(value string) {
	MaskEdit_SetName(m.instance, value)
}

// 获取对象标记。
//
// Get the control tag.
func (m *TMaskEdit) Tag() int {
	return MaskEdit_GetTag(m.instance)
}

// 设置对象标记。
//
// Set the control tag.
func (m *TMaskEdit) SetTag(value int) {
	MaskEdit_SetTag(m.instance, value)
}

// 获取左边锚点。
func (m *TMaskEdit) AnchorSideLeft() *TAnchorSide {
	return AsAnchorSide(MaskEdit_GetAnchorSideLeft(m.instance))
}

// 设置左边锚点。
func (m *TMaskEdit) SetAnchorSideLeft(value *TAnchorSide) {
	MaskEdit_SetAnchorSideLeft(m.instance, CheckPtr(value))
}

// 获取顶边锚点。
func (m *TMaskEdit) AnchorSideTop() *TAnchorSide {
	return AsAnchorSide(MaskEdit_GetAnchorSideTop(m.instance))
}

// 设置顶边锚点。
func (m *TMaskEdit) SetAnchorSideTop(value *TAnchorSide) {
	MaskEdit_SetAnchorSideTop(m.instance, CheckPtr(value))
}

// 获取右边锚点。
func (m *TMaskEdit) AnchorSideRight() *TAnchorSide {
	return AsAnchorSide(MaskEdit_GetAnchorSideRight(m.instance))
}

// 设置右边锚点。
func (m *TMaskEdit) SetAnchorSideRight(value *TAnchorSide) {
	MaskEdit_SetAnchorSideRight(m.instance, CheckPtr(value))
}

// 获取底边锚点。
func (m *TMaskEdit) AnchorSideBottom() *TAnchorSide {
	return AsAnchorSide(MaskEdit_GetAnchorSideBottom(m.instance))
}

// 设置底边锚点。
func (m *TMaskEdit) SetAnchorSideBottom(value *TAnchorSide) {
	MaskEdit_SetAnchorSideBottom(m.instance, CheckPtr(value))
}

func (m *TMaskEdit) ChildSizing() *TControlChildSizing {
	return AsControlChildSizing(MaskEdit_GetChildSizing(m.instance))
}

func (m *TMaskEdit) SetChildSizing(value *TControlChildSizing) {
	MaskEdit_SetChildSizing(m.instance, CheckPtr(value))
}

// 获取边框间距。
func (m *TMaskEdit) BorderSpacing() *TControlBorderSpacing {
	return AsControlBorderSpacing(MaskEdit_GetBorderSpacing(m.instance))
}

// 设置边框间距。
func (m *TMaskEdit) SetBorderSpacing(value *TControlBorderSpacing) {
	MaskEdit_SetBorderSpacing(m.instance, CheckPtr(value))
}

// 获取指定索引停靠客户端。
func (m *TMaskEdit) DockClients(Index int32) *TControl {
	return AsControl(MaskEdit_GetDockClients(m.instance, Index))
}

// 获取指定索引子控件。
func (m *TMaskEdit) Controls(Index int32) *TControl {
	return AsControl(MaskEdit_GetControls(m.instance, Index))
}

// 获取指定索引组件。
//
// Get the specified index component.
func (m *TMaskEdit) Components(AIndex int32) *TComponent {
	return AsComponent(MaskEdit_GetComponents(m.instance, AIndex))
}

// 获取锚侧面。
func (m *TMaskEdit) AnchorSide(AKind TAnchorKind) *TAnchorSide {
	return AsAnchorSide(MaskEdit_GetAnchorSide(m.instance, AKind))
}
