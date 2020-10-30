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

type TTreeView struct {
	IWinControl
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

// 创建一个新的对象。
//
// Create a new object.
func NewTreeView(owner IComponent) *TTreeView {
	t := new(TTreeView)
	t.instance = TreeView_Create(CheckPtr(owner))
	t.ptr = unsafe.Pointer(t.instance)
	return t
}

// 动态转换一个已存在的对象实例。
//
// Dynamically convert an existing object instance.
func AsTreeView(obj interface{}) *TTreeView {
	instance, ptr := getInstance(obj)
	if instance == 0 {
		return nil
	}
	return &TTreeView{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// 新建一个对象来自已经存在的对象实例指针。
//
// Create a new object from an existing object instance pointer.
// Deprecated: use AsTreeView.
func TreeViewFromInst(inst uintptr) *TTreeView {
	return AsTreeView(inst)
}

// 新建一个对象来自已经存在的对象实例。
//
// Create a new object from an existing object instance.
// Deprecated: use AsTreeView.
func TreeViewFromObj(obj IObject) *TTreeView {
	return AsTreeView(obj)
}

// 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
//
// Create a new object from an unsecured address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsTreeView.
func TreeViewFromUnsafePointer(ptr unsafe.Pointer) *TTreeView {
	return AsTreeView(ptr)
}

// -------------------------- Deprecated end --------------------------
// 释放对象。
//
// Free object.
func (t *TTreeView) Free() {
	if t.instance != 0 {
		TreeView_Free(t.instance)
		t.instance, t.ptr = 0, nullptr
	}
}

// 返回对象实例指针。
//
// Return object instance pointer.
func (t *TTreeView) Instance() uintptr {
	return t.instance
}

// 获取一个不安全的地址。
//
// Get an unsafe address.
func (t *TTreeView) UnsafeAddr() unsafe.Pointer {
	return t.ptr
}

// 检测地址是否为空。
//
// Check if the address is empty.
func (t *TTreeView) IsValid() bool {
	return t.instance != 0
}

// 检测当前对象是否继承自目标对象。
//
// Checks whether the current object is inherited from the target object.
func (t *TTreeView) Is() TIs {
	return TIs(t.instance)
}

// 动态转换当前对象为目标对象。
//
// Dynamically convert the current object to the target object.
//func (t *TTreeView) As() TAs {
//    return TAs(t.instance)
//}

// 获取类信息指针。
//
// Get class information pointer.
func TTreeViewClass() TClass {
	return TreeView_StaticClassType()
}

// 按字母排序，所有都参数无效，仅仅用来兼容Delphi的。
//
// Sorted alphabetically, all parameters are invalid, Only used to be compatible with Delphi.
func (t *TTreeView) AlphaSort(ARecurse bool) bool {
	return TreeView_AlphaSort(t.instance, ARecurse)
}

func (t *TTreeView) FullCollapse() {
	TreeView_FullCollapse(t.instance)
}

func (t *TTreeView) FullExpand() {
	TreeView_FullExpand(t.instance)
}

func (t *TTreeView) GetHitTestInfoAt(X int32, Y int32) THitTests {
	return TreeView_GetHitTestInfoAt(t.instance, X, Y)
}

func (t *TTreeView) GetNodeAt(X int32, Y int32) *TTreeNode {
	return AsTreeNode(TreeView_GetNodeAt(t.instance, X, Y))
}

func (t *TTreeView) IsEditing() bool {
	return TreeView_IsEditing(t.instance)
}

// 从文件加载。
func (t *TTreeView) LoadFromFile(FileName string) {
	TreeView_LoadFromFile(t.instance, FileName)
}

// 文件流加载。
func (t *TTreeView) LoadFromStream(Stream IStream) {
	TreeView_LoadFromStream(t.instance, CheckPtr(Stream))
}

// 保存至文件。
func (t *TTreeView) SaveToFile(FileName string) {
	TreeView_SaveToFile(t.instance, FileName)
}

// 保存至流。
func (t *TTreeView) SaveToStream(Stream IStream) {
	TreeView_SaveToStream(t.instance, CheckPtr(Stream))
}

// 清除选择。
func (t *TTreeView) ClearSelection(KeepPrimary bool) {
	TreeView_ClearSelection(t.instance, KeepPrimary)
}

// 自定义排序，所有都参数无效，仅仅用来兼容Delphi的。
//
// Custom sorting, All parameters are invalid, Only used to be compatible with Delphi.
func (t *TTreeView) CustomSort(SortProc PFNTVCOMPARE, Data int, ARecurse bool) bool {
	return TreeView_CustomSort(t.instance, SortProc, Data, ARecurse)
}

// 是否可以获得焦点。
func (t *TTreeView) CanFocus() bool {
	return TreeView_CanFocus(t.instance)
}

// 返回是否包含指定控件。
//
// it's contain a specified control.
func (t *TTreeView) ContainsControl(Control IControl) bool {
	return TreeView_ContainsControl(t.instance, CheckPtr(Control))
}

// 返回指定坐标及相关属性位置控件。
//
// Returns the specified coordinate and the relevant attribute position control..
func (t *TTreeView) ControlAtPos(Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) *TControl {
	return AsControl(TreeView_ControlAtPos(t.instance, Pos, AllowDisabled, AllowWinControls, AllLevels))
}

// 禁用控件的对齐。
//
// Disable control alignment.
func (t *TTreeView) DisableAlign() {
	TreeView_DisableAlign(t.instance)
}

// 启用控件对齐。
//
// Enabled control alignment.
func (t *TTreeView) EnableAlign() {
	TreeView_EnableAlign(t.instance)
}

// 查找子控件。
//
// Find sub controls.
func (t *TTreeView) FindChildControl(ControlName string) *TControl {
	return AsControl(TreeView_FindChildControl(t.instance, ControlName))
}

func (t *TTreeView) FlipChildren(AllLevels bool) {
	TreeView_FlipChildren(t.instance, AllLevels)
}

// 返回是否获取焦点。
//
// Return to get focus.
func (t *TTreeView) Focused() bool {
	return TreeView_Focused(t.instance)
}

// 句柄是否已经分配。
//
// Is the handle already allocated.
func (t *TTreeView) HandleAllocated() bool {
	return TreeView_HandleAllocated(t.instance)
}

// 插入一个控件。
//
// Insert a control.
func (t *TTreeView) InsertControl(AControl IControl) {
	TreeView_InsertControl(t.instance, CheckPtr(AControl))
}

// 要求重绘。
//
// Redraw.
func (t *TTreeView) Invalidate() {
	TreeView_Invalidate(t.instance)
}

// 绘画至指定DC。
//
// Painting to the specified DC.
func (t *TTreeView) PaintTo(DC HDC, X int32, Y int32) {
	TreeView_PaintTo(t.instance, DC, X, Y)
}

// 移除一个控件。
//
// Remove a control.
func (t *TTreeView) RemoveControl(AControl IControl) {
	TreeView_RemoveControl(t.instance, CheckPtr(AControl))
}

// 重新对齐。
//
// Realign.
func (t *TTreeView) Realign() {
	TreeView_Realign(t.instance)
}

// 重绘。
//
// Repaint.
func (t *TTreeView) Repaint() {
	TreeView_Repaint(t.instance)
}

// 按比例缩放。
//
// Scale by.
func (t *TTreeView) ScaleBy(M int32, D int32) {
	TreeView_ScaleBy(t.instance, M, D)
}

// 滚动至指定位置。
//
// Scroll by.
func (t *TTreeView) ScrollBy(DeltaX int32, DeltaY int32) {
	TreeView_ScrollBy(t.instance, DeltaX, DeltaY)
}

// 设置组件边界。
//
// Set component boundaries.
func (t *TTreeView) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
	TreeView_SetBounds(t.instance, ALeft, ATop, AWidth, AHeight)
}

// 设置控件焦点。
//
// Set control focus.
func (t *TTreeView) SetFocus() {
	TreeView_SetFocus(t.instance)
}

// 控件更新。
//
// Update.
func (t *TTreeView) Update() {
	TreeView_Update(t.instance)
}

// 将控件置于最前。
//
// Bring the control to the front.
func (t *TTreeView) BringToFront() {
	TreeView_BringToFront(t.instance)
}

// 将客户端坐标转为绝对的屏幕坐标。
//
// Convert client coordinates to absolute screen coordinates.
func (t *TTreeView) ClientToScreen(Point TPoint) TPoint {
	return TreeView_ClientToScreen(t.instance, Point)
}

// 将客户端坐标转为父容器坐标。
//
// Convert client coordinates to parent container coordinates.
func (t *TTreeView) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
	return TreeView_ClientToParent(t.instance, Point, CheckPtr(AParent))
}

// 是否在拖拽中。
//
// Is it in the middle of dragging.
func (t *TTreeView) Dragging() bool {
	return TreeView_Dragging(t.instance)
}

// 是否有父容器。
//
// Is there a parent container.
func (t *TTreeView) HasParent() bool {
	return TreeView_HasParent(t.instance)
}

// 隐藏控件。
//
// Hidden control.
func (t *TTreeView) Hide() {
	TreeView_Hide(t.instance)
}

// 发送一个消息。
//
// Send a message.
func (t *TTreeView) Perform(Msg uint32, WParam uintptr, LParam int) int {
	return TreeView_Perform(t.instance, Msg, WParam, LParam)
}

// 刷新控件。
//
// Refresh control.
func (t *TTreeView) Refresh() {
	TreeView_Refresh(t.instance)
}

// 将屏幕坐标转为客户端坐标。
//
// Convert screen coordinates to client coordinates.
func (t *TTreeView) ScreenToClient(Point TPoint) TPoint {
	return TreeView_ScreenToClient(t.instance, Point)
}

// 将父容器坐标转为客户端坐标。
//
// Convert parent container coordinates to client coordinates.
func (t *TTreeView) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
	return TreeView_ParentToClient(t.instance, Point, CheckPtr(AParent))
}

// 控件至于最后面。
//
// The control is placed at the end.
func (t *TTreeView) SendToBack() {
	TreeView_SendToBack(t.instance)
}

// 显示控件。
//
// Show control.
func (t *TTreeView) Show() {
	TreeView_Show(t.instance)
}

// 获取控件的字符，如果有。
//
// Get the characters of the control, if any.
func (t *TTreeView) GetTextBuf(Buffer *string, BufSize int32) int32 {
	return TreeView_GetTextBuf(t.instance, Buffer, BufSize)
}

// 获取控件的字符长，如果有。
//
// Get the character length of the control, if any.
func (t *TTreeView) GetTextLen() int32 {
	return TreeView_GetTextLen(t.instance)
}

// 设置控件字符，如果有。
//
// Set control characters, if any.
func (t *TTreeView) SetTextBuf(Buffer string) {
	TreeView_SetTextBuf(t.instance, Buffer)
}

// 查找指定名称的组件。
//
// Find the component with the specified name.
func (t *TTreeView) FindComponent(AName string) *TComponent {
	return AsComponent(TreeView_FindComponent(t.instance, AName))
}

// 获取类名路径。
//
// Get the class name path.
func (t *TTreeView) GetNamePath() string {
	return TreeView_GetNamePath(t.instance)
}

// 复制一个对象，如果对象实现了此方法的话。
//
// Copy an object, if the object implements this method.
func (t *TTreeView) Assign(Source IObject) {
	TreeView_Assign(t.instance, CheckPtr(Source))
}

// 获取类的类型信息。
//
// Get class type information.
func (t *TTreeView) ClassType() TClass {
	return TreeView_ClassType(t.instance)
}

// 获取当前对象类名称。
//
// Get the current object class name.
func (t *TTreeView) ClassName() string {
	return TreeView_ClassName(t.instance)
}

// 获取当前对象实例大小。
//
// Get the current object instance size.
func (t *TTreeView) InstanceSize() int32 {
	return TreeView_InstanceSize(t.instance)
}

// 判断当前类是否继承自指定类。
//
// Determine whether the current class inherits from the specified class.
func (t *TTreeView) InheritsFrom(AClass TClass) bool {
	return TreeView_InheritsFrom(t.instance, AClass)
}

// 与一个对象进行比较。
//
// Compare with an object.
func (t *TTreeView) Equals(Obj IObject) bool {
	return TreeView_Equals(t.instance, CheckPtr(Obj))
}

// 获取类的哈希值。
//
// Get the hash value of the class.
func (t *TTreeView) GetHashCode() int32 {
	return TreeView_GetHashCode(t.instance)
}

// 文本类信息。
//
// Text information.
func (t *TTreeView) ToString() string {
	return TreeView_ToString(t.instance)
}

func (t *TTreeView) AnchorToNeighbour(ASide TAnchorKind, ASpace int32, ASibling IControl) {
	TreeView_AnchorToNeighbour(t.instance, ASide, ASpace, CheckPtr(ASibling))
}

func (t *TTreeView) AnchorParallel(ASide TAnchorKind, ASpace int32, ASibling IControl) {
	TreeView_AnchorParallel(t.instance, ASide, ASpace, CheckPtr(ASibling))
}

// 置于指定控件的横向中心。
func (t *TTreeView) AnchorHorizontalCenterTo(ASibling IControl) {
	TreeView_AnchorHorizontalCenterTo(t.instance, CheckPtr(ASibling))
}

// 置于指定控件的纵向中心。
func (t *TTreeView) AnchorVerticalCenterTo(ASibling IControl) {
	TreeView_AnchorVerticalCenterTo(t.instance, CheckPtr(ASibling))
}

func (t *TTreeView) AnchorSame(ASide TAnchorKind, ASibling IControl) {
	TreeView_AnchorSame(t.instance, ASide, CheckPtr(ASibling))
}

func (t *TTreeView) AnchorAsAlign(ATheAlign TAlign, ASpace int32) {
	TreeView_AnchorAsAlign(t.instance, ATheAlign, ASpace)
}

func (t *TTreeView) AnchorClient(ASpace int32) {
	TreeView_AnchorClient(t.instance, ASpace)
}

func (t *TTreeView) ScaleDesignToForm(ASize int32) int32 {
	return TreeView_ScaleDesignToForm(t.instance, ASize)
}

func (t *TTreeView) ScaleFormToDesign(ASize int32) int32 {
	return TreeView_ScaleFormToDesign(t.instance, ASize)
}

func (t *TTreeView) Scale96ToForm(ASize int32) int32 {
	return TreeView_Scale96ToForm(t.instance, ASize)
}

func (t *TTreeView) ScaleFormTo96(ASize int32) int32 {
	return TreeView_ScaleFormTo96(t.instance, ASize)
}

func (t *TTreeView) Scale96ToFont(ASize int32) int32 {
	return TreeView_Scale96ToFont(t.instance, ASize)
}

func (t *TTreeView) ScaleFontTo96(ASize int32) int32 {
	return TreeView_ScaleFontTo96(t.instance, ASize)
}

func (t *TTreeView) ScaleScreenToFont(ASize int32) int32 {
	return TreeView_ScaleScreenToFont(t.instance, ASize)
}

func (t *TTreeView) ScaleFontToScreen(ASize int32) int32 {
	return TreeView_ScaleFontToScreen(t.instance, ASize)
}

func (t *TTreeView) Scale96ToScreen(ASize int32) int32 {
	return TreeView_Scale96ToScreen(t.instance, ASize)
}

func (t *TTreeView) ScaleScreenTo96(ASize int32) int32 {
	return TreeView_ScaleScreenTo96(t.instance, ASize)
}

func (t *TTreeView) AutoAdjustLayout(AMode TLayoutAdjustmentPolicy, AFromPPI int32, AToPPI int32, AOldFormWidth int32, ANewFormWidth int32) {
	TreeView_AutoAdjustLayout(t.instance, AMode, AFromPPI, AToPPI, AOldFormWidth, ANewFormWidth)
}

func (t *TTreeView) FixDesignFontsPPI(ADesignTimePPI int32) {
	TreeView_FixDesignFontsPPI(t.instance, ADesignTimePPI)
}

func (t *TTreeView) ScaleFontsPPI(AToPPI int32, AProportion float64) {
	TreeView_ScaleFontsPPI(t.instance, AToPPI, AProportion)
}

func (t *TTreeView) DefaultItemHeight() int32 {
	return TreeView_GetDefaultItemHeight(t.instance)
}

func (t *TTreeView) SetDefaultItemHeight(value int32) {
	TreeView_SetDefaultItemHeight(t.instance, value)
}

func (t *TTreeView) ExpandSignColor() TColor {
	return TreeView_GetExpandSignColor(t.instance)
}

func (t *TTreeView) SetExpandSignColor(value TColor) {
	TreeView_SetExpandSignColor(t.instance, value)
}

func (t *TTreeView) ExpandSignSize() int32 {
	return TreeView_GetExpandSignSize(t.instance)
}

func (t *TTreeView) SetExpandSignSize(value int32) {
	TreeView_SetExpandSignSize(t.instance, value)
}

func (t *TTreeView) ExpandSignType() TTreeViewExpandSignType {
	return TreeView_GetExpandSignType(t.instance)
}

func (t *TTreeView) SetExpandSignType(value TTreeViewExpandSignType) {
	TreeView_SetExpandSignType(t.instance, value)
}

func (t *TTreeView) HotTrackColor() TColor {
	return TreeView_GetHotTrackColor(t.instance)
}

func (t *TTreeView) SetHotTrackColor(value TColor) {
	TreeView_SetHotTrackColor(t.instance, value)
}

func (t *TTreeView) ImagesWidth() int32 {
	return TreeView_GetImagesWidth(t.instance)
}

func (t *TTreeView) SetImagesWidth(value int32) {
	TreeView_SetImagesWidth(t.instance, value)
}

func (t *TTreeView) Options() TTreeViewOptions {
	return TreeView_GetOptions(t.instance)
}

func (t *TTreeView) SetOptions(value TTreeViewOptions) {
	TreeView_SetOptions(t.instance, value)
}

func (t *TTreeView) ScrollBars() TScrollStyle {
	return TreeView_GetScrollBars(t.instance)
}

func (t *TTreeView) SetScrollBars(value TScrollStyle) {
	TreeView_SetScrollBars(t.instance, value)
}

func (t *TTreeView) SelectionColor() TColor {
	return TreeView_GetSelectionColor(t.instance)
}

func (t *TTreeView) SetSelectionColor(value TColor) {
	TreeView_SetSelectionColor(t.instance, value)
}

func (t *TTreeView) SelectionFontColor() TColor {
	return TreeView_GetSelectionFontColor(t.instance)
}

func (t *TTreeView) SetSelectionFontColor(value TColor) {
	TreeView_SetSelectionFontColor(t.instance, value)
}

func (t *TTreeView) SelectionFontColorUsed() bool {
	return TreeView_GetSelectionFontColorUsed(t.instance)
}

func (t *TTreeView) SetSelectionFontColorUsed(value bool) {
	TreeView_SetSelectionFontColorUsed(t.instance, value)
}

func (t *TTreeView) SeparatorColor() TColor {
	return TreeView_GetSeparatorColor(t.instance)
}

func (t *TTreeView) SetSeparatorColor(value TColor) {
	TreeView_SetSeparatorColor(t.instance, value)
}

func (t *TTreeView) StateImagesWidth() int32 {
	return TreeView_GetStateImagesWidth(t.instance)
}

func (t *TTreeView) SetStateImagesWidth(value int32) {
	TreeView_SetStateImagesWidth(t.instance, value)
}

func (t *TTreeView) ToolTips() bool {
	return TreeView_GetToolTips(t.instance)
}

func (t *TTreeView) SetToolTips(value bool) {
	TreeView_SetToolTips(t.instance, value)
}

func (t *TTreeView) TreeLineColor() TColor {
	return TreeView_GetTreeLineColor(t.instance)
}

func (t *TTreeView) SetTreeLineColor(value TColor) {
	TreeView_SetTreeLineColor(t.instance, value)
}

func (t *TTreeView) TreeLinePenStyle() TPenStyle {
	return TreeView_GetTreeLinePenStyle(t.instance)
}

func (t *TTreeView) SetTreeLinePenStyle(value TPenStyle) {
	TreeView_SetTreeLinePenStyle(t.instance, value)
}

// 获取控件自动调整。
//
// Get Control automatically adjusts.
func (t *TTreeView) Align() TAlign {
	return TreeView_GetAlign(t.instance)
}

// 设置控件自动调整。
//
// Set Control automatically adjusts.
func (t *TTreeView) SetAlign(value TAlign) {
	TreeView_SetAlign(t.instance, value)
}

// 获取四个角位置的锚点。
func (t *TTreeView) Anchors() TAnchors {
	return TreeView_GetAnchors(t.instance)
}

// 设置四个角位置的锚点。
func (t *TTreeView) SetAnchors(value TAnchors) {
	TreeView_SetAnchors(t.instance, value)
}

func (t *TTreeView) AutoExpand() bool {
	return TreeView_GetAutoExpand(t.instance)
}

func (t *TTreeView) SetAutoExpand(value bool) {
	TreeView_SetAutoExpand(t.instance, value)
}

func (t *TTreeView) BiDiMode() TBiDiMode {
	return TreeView_GetBiDiMode(t.instance)
}

func (t *TTreeView) SetBiDiMode(value TBiDiMode) {
	TreeView_SetBiDiMode(t.instance, value)
}

// 获取窗口边框样式。比如：无边框，单一边框等。
func (t *TTreeView) BorderStyle() TBorderStyle {
	return TreeView_GetBorderStyle(t.instance)
}

// 设置窗口边框样式。比如：无边框，单一边框等。
func (t *TTreeView) SetBorderStyle(value TBorderStyle) {
	TreeView_SetBorderStyle(t.instance, value)
}

// 获取边框的宽度。
func (t *TTreeView) BorderWidth() int32 {
	return TreeView_GetBorderWidth(t.instance)
}

// 设置边框的宽度。
func (t *TTreeView) SetBorderWidth(value int32) {
	TreeView_SetBorderWidth(t.instance, value)
}

// 获取颜色。
//
// Get color.
func (t *TTreeView) Color() TColor {
	return TreeView_GetColor(t.instance)
}

// 设置颜色。
//
// Set color.
func (t *TTreeView) SetColor(value TColor) {
	TreeView_SetColor(t.instance, value)
}

// 获取约束控件大小。
func (t *TTreeView) Constraints() *TSizeConstraints {
	return AsSizeConstraints(TreeView_GetConstraints(t.instance))
}

// 设置约束控件大小。
func (t *TTreeView) SetConstraints(value *TSizeConstraints) {
	TreeView_SetConstraints(t.instance, CheckPtr(value))
}

// 获取设置控件双缓冲。
//
// Get Set control double buffering.
func (t *TTreeView) DoubleBuffered() bool {
	return TreeView_GetDoubleBuffered(t.instance)
}

// 设置设置控件双缓冲。
//
// Set Set control double buffering.
func (t *TTreeView) SetDoubleBuffered(value bool) {
	TreeView_SetDoubleBuffered(t.instance, value)
}

// 获取拖拽方式。
//
// Get Drag and drop.
func (t *TTreeView) DragKind() TDragKind {
	return TreeView_GetDragKind(t.instance)
}

// 设置拖拽方式。
//
// Set Drag and drop.
func (t *TTreeView) SetDragKind(value TDragKind) {
	TreeView_SetDragKind(t.instance, value)
}

// 获取设置控件拖拽时的光标。
//
// Get Set the cursor when the control is dragged.
func (t *TTreeView) DragCursor() TCursor {
	return TreeView_GetDragCursor(t.instance)
}

// 设置设置控件拖拽时的光标。
//
// Set Set the cursor when the control is dragged.
func (t *TTreeView) SetDragCursor(value TCursor) {
	TreeView_SetDragCursor(t.instance, value)
}

// 获取拖拽模式。
//
// Get Drag mode.
func (t *TTreeView) DragMode() TDragMode {
	return TreeView_GetDragMode(t.instance)
}

// 设置拖拽模式。
//
// Set Drag mode.
func (t *TTreeView) SetDragMode(value TDragMode) {
	TreeView_SetDragMode(t.instance, value)
}

// 获取控件启用。
//
// Get the control enabled.
func (t *TTreeView) Enabled() bool {
	return TreeView_GetEnabled(t.instance)
}

// 设置控件启用。
//
// Set the control enabled.
func (t *TTreeView) SetEnabled(value bool) {
	TreeView_SetEnabled(t.instance, value)
}

// 获取字体。
//
// Get Font.
func (t *TTreeView) Font() *TFont {
	return AsFont(TreeView_GetFont(t.instance))
}

// 设置字体。
//
// Set Font.
func (t *TTreeView) SetFont(value *TFont) {
	TreeView_SetFont(t.instance, CheckPtr(value))
}

// 获取隐藏选择。
func (t *TTreeView) HideSelection() bool {
	return TreeView_GetHideSelection(t.instance)
}

// 设置隐藏选择。
func (t *TTreeView) SetHideSelection(value bool) {
	TreeView_SetHideSelection(t.instance, value)
}

func (t *TTreeView) HotTrack() bool {
	return TreeView_GetHotTrack(t.instance)
}

func (t *TTreeView) SetHotTrack(value bool) {
	TreeView_SetHotTrack(t.instance, value)
}

// 获取图标索引列表对象。
func (t *TTreeView) Images() *TImageList {
	return AsImageList(TreeView_GetImages(t.instance))
}

// 设置图标索引列表对象。
func (t *TTreeView) SetImages(value IComponent) {
	TreeView_SetImages(t.instance, CheckPtr(value))
}

func (t *TTreeView) Indent() int32 {
	return TreeView_GetIndent(t.instance)
}

func (t *TTreeView) SetIndent(value int32) {
	TreeView_SetIndent(t.instance, value)
}

func (t *TTreeView) MultiSelect() bool {
	return TreeView_GetMultiSelect(t.instance)
}

func (t *TTreeView) SetMultiSelect(value bool) {
	TreeView_SetMultiSelect(t.instance, value)
}

func (t *TTreeView) MultiSelectStyle() TMultiSelectStyle {
	return TreeView_GetMultiSelectStyle(t.instance)
}

func (t *TTreeView) SetMultiSelectStyle(value TMultiSelectStyle) {
	TreeView_SetMultiSelectStyle(t.instance, value)
}

// 获取使用父容器颜色。
//
// Get parent color.
func (t *TTreeView) ParentColor() bool {
	return TreeView_GetParentColor(t.instance)
}

// 设置使用父容器颜色。
//
// Set parent color.
func (t *TTreeView) SetParentColor(value bool) {
	TreeView_SetParentColor(t.instance, value)
}

// 获取使用父容器双缓冲。
//
// Get Parent container double buffering.
func (t *TTreeView) ParentDoubleBuffered() bool {
	return TreeView_GetParentDoubleBuffered(t.instance)
}

// 设置使用父容器双缓冲。
//
// Set Parent container double buffering.
func (t *TTreeView) SetParentDoubleBuffered(value bool) {
	TreeView_SetParentDoubleBuffered(t.instance, value)
}

// 获取使用父容器字体。
//
// Get Parent container font.
func (t *TTreeView) ParentFont() bool {
	return TreeView_GetParentFont(t.instance)
}

// 设置使用父容器字体。
//
// Set Parent container font.
func (t *TTreeView) SetParentFont(value bool) {
	TreeView_SetParentFont(t.instance, value)
}

// 获取以父容器的ShowHint属性为准。
func (t *TTreeView) ParentShowHint() bool {
	return TreeView_GetParentShowHint(t.instance)
}

// 设置以父容器的ShowHint属性为准。
func (t *TTreeView) SetParentShowHint(value bool) {
	TreeView_SetParentShowHint(t.instance, value)
}

// 获取右键菜单。
//
// Get Right click menu.
func (t *TTreeView) PopupMenu() *TPopupMenu {
	return AsPopupMenu(TreeView_GetPopupMenu(t.instance))
}

// 设置右键菜单。
//
// Set Right click menu.
func (t *TTreeView) SetPopupMenu(value IComponent) {
	TreeView_SetPopupMenu(t.instance, CheckPtr(value))
}

// 获取只读。
func (t *TTreeView) ReadOnly() bool {
	return TreeView_GetReadOnly(t.instance)
}

// 设置只读。
func (t *TTreeView) SetReadOnly(value bool) {
	TreeView_SetReadOnly(t.instance, value)
}

func (t *TTreeView) RightClickSelect() bool {
	return TreeView_GetRightClickSelect(t.instance)
}

func (t *TTreeView) SetRightClickSelect(value bool) {
	TreeView_SetRightClickSelect(t.instance, value)
}

func (t *TTreeView) RowSelect() bool {
	return TreeView_GetRowSelect(t.instance)
}

func (t *TTreeView) SetRowSelect(value bool) {
	TreeView_SetRowSelect(t.instance, value)
}

func (t *TTreeView) ShowButtons() bool {
	return TreeView_GetShowButtons(t.instance)
}

func (t *TTreeView) SetShowButtons(value bool) {
	TreeView_SetShowButtons(t.instance, value)
}

// 获取显示鼠标悬停提示。
//
// Get Show mouseover tips.
func (t *TTreeView) ShowHint() bool {
	return TreeView_GetShowHint(t.instance)
}

// 设置显示鼠标悬停提示。
//
// Set Show mouseover tips.
func (t *TTreeView) SetShowHint(value bool) {
	TreeView_SetShowHint(t.instance, value)
}

func (t *TTreeView) ShowLines() bool {
	return TreeView_GetShowLines(t.instance)
}

func (t *TTreeView) SetShowLines(value bool) {
	TreeView_SetShowLines(t.instance, value)
}

func (t *TTreeView) ShowRoot() bool {
	return TreeView_GetShowRoot(t.instance)
}

func (t *TTreeView) SetShowRoot(value bool) {
	TreeView_SetShowRoot(t.instance, value)
}

func (t *TTreeView) SortType() TSortType {
	return TreeView_GetSortType(t.instance)
}

func (t *TTreeView) SetSortType(value TSortType) {
	TreeView_SetSortType(t.instance, value)
}

func (t *TTreeView) StateImages() *TImageList {
	return AsImageList(TreeView_GetStateImages(t.instance))
}

func (t *TTreeView) SetStateImages(value IComponent) {
	TreeView_SetStateImages(t.instance, CheckPtr(value))
}

// 获取Tab切换顺序序号。
//
// Get Tab switching sequence number.
func (t *TTreeView) TabOrder() TTabOrder {
	return TreeView_GetTabOrder(t.instance)
}

// 设置Tab切换顺序序号。
//
// Set Tab switching sequence number.
func (t *TTreeView) SetTabOrder(value TTabOrder) {
	TreeView_SetTabOrder(t.instance, value)
}

// 获取Tab可停留。
//
// Get Tab can stay.
func (t *TTreeView) TabStop() bool {
	return TreeView_GetTabStop(t.instance)
}

// 设置Tab可停留。
//
// Set Tab can stay.
func (t *TTreeView) SetTabStop(value bool) {
	TreeView_SetTabStop(t.instance, value)
}

// 获取控件可视。
//
// Get the control visible.
func (t *TTreeView) Visible() bool {
	return TreeView_GetVisible(t.instance)
}

// 设置控件可视。
//
// Set the control visible.
func (t *TTreeView) SetVisible(value bool) {
	TreeView_SetVisible(t.instance, value)
}

func (t *TTreeView) SetOnAddition(fn TTVExpandedEvent) {
	TreeView_SetOnAddition(t.instance, fn)
}

func (t *TTreeView) SetOnAdvancedCustomDraw(fn TTVAdvancedCustomDrawEvent) {
	TreeView_SetOnAdvancedCustomDraw(t.instance, fn)
}

func (t *TTreeView) SetOnAdvancedCustomDrawItem(fn TTVAdvancedCustomDrawItemEvent) {
	TreeView_SetOnAdvancedCustomDrawItem(t.instance, fn)
}

// 设置改变事件。
//
// Set changed event.
func (t *TTreeView) SetOnChange(fn TTVChangedEvent) {
	TreeView_SetOnChange(t.instance, fn)
}

func (t *TTreeView) SetOnChanging(fn TTVChangingEvent) {
	TreeView_SetOnChanging(t.instance, fn)
}

// 设置控件单击事件。
//
// Set control click event.
func (t *TTreeView) SetOnClick(fn TNotifyEvent) {
	TreeView_SetOnClick(t.instance, fn)
}

func (t *TTreeView) SetOnCollapsed(fn TTVExpandedEvent) {
	TreeView_SetOnCollapsed(t.instance, fn)
}

func (t *TTreeView) SetOnCollapsing(fn TTVCollapsingEvent) {
	TreeView_SetOnCollapsing(t.instance, fn)
}

func (t *TTreeView) SetOnCompare(fn TTVCompareEvent) {
	TreeView_SetOnCompare(t.instance, fn)
}

// 设置上下文弹出事件，一般是右键时弹出。
//
// Set Context popup event, usually pop up when right click.
func (t *TTreeView) SetOnContextPopup(fn TContextPopupEvent) {
	TreeView_SetOnContextPopup(t.instance, fn)
}

func (t *TTreeView) SetOnCustomDraw(fn TTVCustomDrawEvent) {
	TreeView_SetOnCustomDraw(t.instance, fn)
}

func (t *TTreeView) SetOnCustomDrawItem(fn TTVCustomDrawItemEvent) {
	TreeView_SetOnCustomDrawItem(t.instance, fn)
}

// 设置双击事件。
func (t *TTreeView) SetOnDblClick(fn TNotifyEvent) {
	TreeView_SetOnDblClick(t.instance, fn)
}

func (t *TTreeView) SetOnDeletion(fn TTVExpandedEvent) {
	TreeView_SetOnDeletion(t.instance, fn)
}

// 设置拖拽下落事件。
//
// Set Drag and drop event.
func (t *TTreeView) SetOnDragDrop(fn TDragDropEvent) {
	TreeView_SetOnDragDrop(t.instance, fn)
}

// 设置拖拽完成事件。
//
// Set Drag and drop completion event.
func (t *TTreeView) SetOnDragOver(fn TDragOverEvent) {
	TreeView_SetOnDragOver(t.instance, fn)
}

func (t *TTreeView) SetOnEdited(fn TTVEditedEvent) {
	TreeView_SetOnEdited(t.instance, fn)
}

func (t *TTreeView) SetOnEditing(fn TTVEditingEvent) {
	TreeView_SetOnEditing(t.instance, fn)
}

// 设置拖拽结束。
//
// Set End of drag.
func (t *TTreeView) SetOnEndDrag(fn TEndDragEvent) {
	TreeView_SetOnEndDrag(t.instance, fn)
}

// 设置焦点进入。
//
// Set Focus entry.
func (t *TTreeView) SetOnEnter(fn TNotifyEvent) {
	TreeView_SetOnEnter(t.instance, fn)
}

// 设置焦点退出。
//
// Set Focus exit.
func (t *TTreeView) SetOnExit(fn TNotifyEvent) {
	TreeView_SetOnExit(t.instance, fn)
}

func (t *TTreeView) SetOnExpanding(fn TTVExpandingEvent) {
	TreeView_SetOnExpanding(t.instance, fn)
}

func (t *TTreeView) SetOnExpanded(fn TTVExpandedEvent) {
	TreeView_SetOnExpanded(t.instance, fn)
}

func (t *TTreeView) SetOnGetSelectedIndex(fn TTVExpandedEvent) {
	TreeView_SetOnGetSelectedIndex(t.instance, fn)
}

// 设置键盘按键按下事件。
//
// Set Keyboard button press event.
func (t *TTreeView) SetOnKeyDown(fn TKeyEvent) {
	TreeView_SetOnKeyDown(t.instance, fn)
}

// 设置键键下事件。
func (t *TTreeView) SetOnKeyPress(fn TKeyPressEvent) {
	TreeView_SetOnKeyPress(t.instance, fn)
}

// 设置键盘按键抬起事件。
//
// Set Keyboard button lift event.
func (t *TTreeView) SetOnKeyUp(fn TKeyEvent) {
	TreeView_SetOnKeyUp(t.instance, fn)
}

// 设置鼠标按下事件。
//
// Set Mouse down event.
func (t *TTreeView) SetOnMouseDown(fn TMouseEvent) {
	TreeView_SetOnMouseDown(t.instance, fn)
}

// 设置鼠标进入事件。
//
// Set Mouse entry event.
func (t *TTreeView) SetOnMouseEnter(fn TNotifyEvent) {
	TreeView_SetOnMouseEnter(t.instance, fn)
}

// 设置鼠标离开事件。
//
// Set Mouse leave event.
func (t *TTreeView) SetOnMouseLeave(fn TNotifyEvent) {
	TreeView_SetOnMouseLeave(t.instance, fn)
}

// 设置鼠标移动事件。
func (t *TTreeView) SetOnMouseMove(fn TMouseMoveEvent) {
	TreeView_SetOnMouseMove(t.instance, fn)
}

// 设置鼠标抬起事件。
//
// Set Mouse lift event.
func (t *TTreeView) SetOnMouseUp(fn TMouseEvent) {
	TreeView_SetOnMouseUp(t.instance, fn)
}

func (t *TTreeView) Items() *TTreeNodes {
	return AsTreeNodes(TreeView_GetItems(t.instance))
}

func (t *TTreeView) SetItems(value *TTreeNodes) {
	TreeView_SetItems(t.instance, CheckPtr(value))
}

// 获取画布。
func (t *TTreeView) Canvas() *TCanvas {
	return AsCanvas(TreeView_GetCanvas(t.instance))
}

func (t *TTreeView) DropTarget() *TTreeNode {
	return AsTreeNode(TreeView_GetDropTarget(t.instance))
}

func (t *TTreeView) SetDropTarget(value *TTreeNode) {
	TreeView_SetDropTarget(t.instance, CheckPtr(value))
}

func (t *TTreeView) Selected() *TTreeNode {
	return AsTreeNode(TreeView_GetSelected(t.instance))
}

func (t *TTreeView) SetSelected(value *TTreeNode) {
	TreeView_SetSelected(t.instance, CheckPtr(value))
}

func (t *TTreeView) TopItem() *TTreeNode {
	return AsTreeNode(TreeView_GetTopItem(t.instance))
}

func (t *TTreeView) SetTopItem(value *TTreeNode) {
	TreeView_SetTopItem(t.instance, CheckPtr(value))
}

func (t *TTreeView) SelectionCount() uint32 {
	return TreeView_GetSelectionCount(t.instance)
}

// 获取依靠客户端总数。
func (t *TTreeView) DockClientCount() int32 {
	return TreeView_GetDockClientCount(t.instance)
}

// 获取停靠站点。
//
// Get Docking site.
func (t *TTreeView) DockSite() bool {
	return TreeView_GetDockSite(t.instance)
}

// 设置停靠站点。
//
// Set Docking site.
func (t *TTreeView) SetDockSite(value bool) {
	TreeView_SetDockSite(t.instance, value)
}

// 获取鼠标是否在客户端，仅VCL有效。
//
// Get Whether the mouse is on the client, only VCL is valid.
func (t *TTreeView) MouseInClient() bool {
	return TreeView_GetMouseInClient(t.instance)
}

// 获取当前停靠的可视总数。
//
// Get The total number of visible calls currently docked.
func (t *TTreeView) VisibleDockClientCount() int32 {
	return TreeView_GetVisibleDockClientCount(t.instance)
}

// 获取画刷对象。
//
// Get Brush.
func (t *TTreeView) Brush() *TBrush {
	return AsBrush(TreeView_GetBrush(t.instance))
}

// 获取子控件数。
//
// Get Number of child controls.
func (t *TTreeView) ControlCount() int32 {
	return TreeView_GetControlCount(t.instance)
}

// 获取控件句柄。
//
// Get Control handle.
func (t *TTreeView) Handle() HWND {
	return TreeView_GetHandle(t.instance)
}

// 获取父容器句柄。
//
// Get Parent container handle.
func (t *TTreeView) ParentWindow() HWND {
	return TreeView_GetParentWindow(t.instance)
}

// 设置父容器句柄。
//
// Set Parent container handle.
func (t *TTreeView) SetParentWindow(value HWND) {
	TreeView_SetParentWindow(t.instance, value)
}

func (t *TTreeView) Showing() bool {
	return TreeView_GetShowing(t.instance)
}

// 获取使用停靠管理。
func (t *TTreeView) UseDockManager() bool {
	return TreeView_GetUseDockManager(t.instance)
}

// 设置使用停靠管理。
func (t *TTreeView) SetUseDockManager(value bool) {
	TreeView_SetUseDockManager(t.instance, value)
}

func (t *TTreeView) Action() *TAction {
	return AsAction(TreeView_GetAction(t.instance))
}

func (t *TTreeView) SetAction(value IComponent) {
	TreeView_SetAction(t.instance, CheckPtr(value))
}

func (t *TTreeView) BoundsRect() TRect {
	return TreeView_GetBoundsRect(t.instance)
}

func (t *TTreeView) SetBoundsRect(value TRect) {
	TreeView_SetBoundsRect(t.instance, value)
}

// 获取客户区高度。
//
// Get client height.
func (t *TTreeView) ClientHeight() int32 {
	return TreeView_GetClientHeight(t.instance)
}

// 设置客户区高度。
//
// Set client height.
func (t *TTreeView) SetClientHeight(value int32) {
	TreeView_SetClientHeight(t.instance, value)
}

func (t *TTreeView) ClientOrigin() TPoint {
	return TreeView_GetClientOrigin(t.instance)
}

// 获取客户区矩形。
//
// Get client rectangle.
func (t *TTreeView) ClientRect() TRect {
	return TreeView_GetClientRect(t.instance)
}

// 获取客户区宽度。
//
// Get client width.
func (t *TTreeView) ClientWidth() int32 {
	return TreeView_GetClientWidth(t.instance)
}

// 设置客户区宽度。
//
// Set client width.
func (t *TTreeView) SetClientWidth(value int32) {
	TreeView_SetClientWidth(t.instance, value)
}

// 获取控件状态。
//
// Get control state.
func (t *TTreeView) ControlState() TControlState {
	return TreeView_GetControlState(t.instance)
}

// 设置控件状态。
//
// Set control state.
func (t *TTreeView) SetControlState(value TControlState) {
	TreeView_SetControlState(t.instance, value)
}

// 获取控件样式。
//
// Get control style.
func (t *TTreeView) ControlStyle() TControlStyle {
	return TreeView_GetControlStyle(t.instance)
}

// 设置控件样式。
//
// Set control style.
func (t *TTreeView) SetControlStyle(value TControlStyle) {
	TreeView_SetControlStyle(t.instance, value)
}

func (t *TTreeView) Floating() bool {
	return TreeView_GetFloating(t.instance)
}

// 获取控件父容器。
//
// Get control parent container.
func (t *TTreeView) Parent() *TWinControl {
	return AsWinControl(TreeView_GetParent(t.instance))
}

// 设置控件父容器。
//
// Set control parent container.
func (t *TTreeView) SetParent(value IWinControl) {
	TreeView_SetParent(t.instance, CheckPtr(value))
}

// 获取左边位置。
//
// Get Left position.
func (t *TTreeView) Left() int32 {
	return TreeView_GetLeft(t.instance)
}

// 设置左边位置。
//
// Set Left position.
func (t *TTreeView) SetLeft(value int32) {
	TreeView_SetLeft(t.instance, value)
}

// 获取顶边位置。
//
// Get Top position.
func (t *TTreeView) Top() int32 {
	return TreeView_GetTop(t.instance)
}

// 设置顶边位置。
//
// Set Top position.
func (t *TTreeView) SetTop(value int32) {
	TreeView_SetTop(t.instance, value)
}

// 获取宽度。
//
// Get width.
func (t *TTreeView) Width() int32 {
	return TreeView_GetWidth(t.instance)
}

// 设置宽度。
//
// Set width.
func (t *TTreeView) SetWidth(value int32) {
	TreeView_SetWidth(t.instance, value)
}

// 获取高度。
//
// Get height.
func (t *TTreeView) Height() int32 {
	return TreeView_GetHeight(t.instance)
}

// 设置高度。
//
// Set height.
func (t *TTreeView) SetHeight(value int32) {
	TreeView_SetHeight(t.instance, value)
}

// 获取控件光标。
//
// Get control cursor.
func (t *TTreeView) Cursor() TCursor {
	return TreeView_GetCursor(t.instance)
}

// 设置控件光标。
//
// Set control cursor.
func (t *TTreeView) SetCursor(value TCursor) {
	TreeView_SetCursor(t.instance, value)
}

// 获取组件鼠标悬停提示。
//
// Get component mouse hints.
func (t *TTreeView) Hint() string {
	return TreeView_GetHint(t.instance)
}

// 设置组件鼠标悬停提示。
//
// Set component mouse hints.
func (t *TTreeView) SetHint(value string) {
	TreeView_SetHint(t.instance, value)
}

// 获取组件总数。
//
// Get the total number of components.
func (t *TTreeView) ComponentCount() int32 {
	return TreeView_GetComponentCount(t.instance)
}

// 获取组件索引。
//
// Get component index.
func (t *TTreeView) ComponentIndex() int32 {
	return TreeView_GetComponentIndex(t.instance)
}

// 设置组件索引。
//
// Set component index.
func (t *TTreeView) SetComponentIndex(value int32) {
	TreeView_SetComponentIndex(t.instance, value)
}

// 获取组件所有者。
//
// Get component owner.
func (t *TTreeView) Owner() *TComponent {
	return AsComponent(TreeView_GetOwner(t.instance))
}

// 获取组件名称。
//
// Get the component name.
func (t *TTreeView) Name() string {
	return TreeView_GetName(t.instance)
}

// 设置组件名称。
//
// Set the component name.
func (t *TTreeView) SetName(value string) {
	TreeView_SetName(t.instance, value)
}

// 获取对象标记。
//
// Get the control tag.
func (t *TTreeView) Tag() int {
	return TreeView_GetTag(t.instance)
}

// 设置对象标记。
//
// Set the control tag.
func (t *TTreeView) SetTag(value int) {
	TreeView_SetTag(t.instance, value)
}

// 获取左边锚点。
func (t *TTreeView) AnchorSideLeft() *TAnchorSide {
	return AsAnchorSide(TreeView_GetAnchorSideLeft(t.instance))
}

// 设置左边锚点。
func (t *TTreeView) SetAnchorSideLeft(value *TAnchorSide) {
	TreeView_SetAnchorSideLeft(t.instance, CheckPtr(value))
}

// 获取顶边锚点。
func (t *TTreeView) AnchorSideTop() *TAnchorSide {
	return AsAnchorSide(TreeView_GetAnchorSideTop(t.instance))
}

// 设置顶边锚点。
func (t *TTreeView) SetAnchorSideTop(value *TAnchorSide) {
	TreeView_SetAnchorSideTop(t.instance, CheckPtr(value))
}

// 获取右边锚点。
func (t *TTreeView) AnchorSideRight() *TAnchorSide {
	return AsAnchorSide(TreeView_GetAnchorSideRight(t.instance))
}

// 设置右边锚点。
func (t *TTreeView) SetAnchorSideRight(value *TAnchorSide) {
	TreeView_SetAnchorSideRight(t.instance, CheckPtr(value))
}

// 获取底边锚点。
func (t *TTreeView) AnchorSideBottom() *TAnchorSide {
	return AsAnchorSide(TreeView_GetAnchorSideBottom(t.instance))
}

// 设置底边锚点。
func (t *TTreeView) SetAnchorSideBottom(value *TAnchorSide) {
	TreeView_SetAnchorSideBottom(t.instance, CheckPtr(value))
}

func (t *TTreeView) ChildSizing() *TControlChildSizing {
	return AsControlChildSizing(TreeView_GetChildSizing(t.instance))
}

func (t *TTreeView) SetChildSizing(value *TControlChildSizing) {
	TreeView_SetChildSizing(t.instance, CheckPtr(value))
}

// 获取边框间距。
func (t *TTreeView) BorderSpacing() *TControlBorderSpacing {
	return AsControlBorderSpacing(TreeView_GetBorderSpacing(t.instance))
}

// 设置边框间距。
func (t *TTreeView) SetBorderSpacing(value *TControlBorderSpacing) {
	TreeView_SetBorderSpacing(t.instance, CheckPtr(value))
}

func (t *TTreeView) Selections(Index int32) *TTreeNode {
	return AsTreeNode(TreeView_GetSelections(t.instance, Index))
}

// 获取指定索引停靠客户端。
func (t *TTreeView) DockClients(Index int32) *TControl {
	return AsControl(TreeView_GetDockClients(t.instance, Index))
}

// 获取指定索引子控件。
func (t *TTreeView) Controls(Index int32) *TControl {
	return AsControl(TreeView_GetControls(t.instance, Index))
}

// 获取指定索引组件。
//
// Get the specified index component.
func (t *TTreeView) Components(AIndex int32) *TComponent {
	return AsComponent(TreeView_GetComponents(t.instance, AIndex))
}

// 获取锚侧面。
func (t *TTreeView) AnchorSide(AKind TAnchorKind) *TAnchorSide {
	return AsAnchorSide(TreeView_GetAnchorSide(t.instance, AKind))
}
