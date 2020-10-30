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

type TPopupMenu struct {
	IComponent
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

// 创建一个新的对象。
//
// Create a new object.
func NewPopupMenu(owner IComponent) *TPopupMenu {
	p := new(TPopupMenu)
	p.instance = PopupMenu_Create(CheckPtr(owner))
	p.ptr = unsafe.Pointer(p.instance)
	return p
}

// 动态转换一个已存在的对象实例。
//
// Dynamically convert an existing object instance.
func AsPopupMenu(obj interface{}) *TPopupMenu {
	instance, ptr := getInstance(obj)
	if instance == 0 {
		return nil
	}
	return &TPopupMenu{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// 新建一个对象来自已经存在的对象实例指针。
//
// Create a new object from an existing object instance pointer.
// Deprecated: use AsPopupMenu.
func PopupMenuFromInst(inst uintptr) *TPopupMenu {
	return AsPopupMenu(inst)
}

// 新建一个对象来自已经存在的对象实例。
//
// Create a new object from an existing object instance.
// Deprecated: use AsPopupMenu.
func PopupMenuFromObj(obj IObject) *TPopupMenu {
	return AsPopupMenu(obj)
}

// 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
//
// Create a new object from an unsecured address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsPopupMenu.
func PopupMenuFromUnsafePointer(ptr unsafe.Pointer) *TPopupMenu {
	return AsPopupMenu(ptr)
}

// -------------------------- Deprecated end --------------------------
// 释放对象。
//
// Free object.
func (p *TPopupMenu) Free() {
	if p.instance != 0 {
		PopupMenu_Free(p.instance)
		p.instance, p.ptr = 0, nullptr
	}
}

// 返回对象实例指针。
//
// Return object instance pointer.
func (p *TPopupMenu) Instance() uintptr {
	return p.instance
}

// 获取一个不安全的地址。
//
// Get an unsafe address.
func (p *TPopupMenu) UnsafeAddr() unsafe.Pointer {
	return p.ptr
}

// 检测地址是否为空。
//
// Check if the address is empty.
func (p *TPopupMenu) IsValid() bool {
	return p.instance != 0
}

// 检测当前对象是否继承自目标对象。
//
// Checks whether the current object is inherited from the target object.
func (p *TPopupMenu) Is() TIs {
	return TIs(p.instance)
}

// 动态转换当前对象为目标对象。
//
// Dynamically convert the current object to the target object.
//func (p *TPopupMenu) As() TAs {
//    return TAs(p.instance)
//}

// 获取类信息指针。
//
// Get class information pointer.
func TPopupMenuClass() TClass {
	return PopupMenu_StaticClassType()
}

func (p *TPopupMenu) CloseMenu() {
	PopupMenu_CloseMenu(p.instance)
}

func (p *TPopupMenu) Popup(X int32, Y int32) {
	PopupMenu_Popup(p.instance, X, Y)
}

// 查找指定名称的组件。
//
// Find the component with the specified name.
func (p *TPopupMenu) FindComponent(AName string) *TComponent {
	return AsComponent(PopupMenu_FindComponent(p.instance, AName))
}

// 获取类名路径。
//
// Get the class name path.
func (p *TPopupMenu) GetNamePath() string {
	return PopupMenu_GetNamePath(p.instance)
}

// 是否有父容器。
//
// Is there a parent container.
func (p *TPopupMenu) HasParent() bool {
	return PopupMenu_HasParent(p.instance)
}

// 复制一个对象，如果对象实现了此方法的话。
//
// Copy an object, if the object implements this method.
func (p *TPopupMenu) Assign(Source IObject) {
	PopupMenu_Assign(p.instance, CheckPtr(Source))
}

// 获取类的类型信息。
//
// Get class type information.
func (p *TPopupMenu) ClassType() TClass {
	return PopupMenu_ClassType(p.instance)
}

// 获取当前对象类名称。
//
// Get the current object class name.
func (p *TPopupMenu) ClassName() string {
	return PopupMenu_ClassName(p.instance)
}

// 获取当前对象实例大小。
//
// Get the current object instance size.
func (p *TPopupMenu) InstanceSize() int32 {
	return PopupMenu_InstanceSize(p.instance)
}

// 判断当前类是否继承自指定类。
//
// Determine whether the current class inherits from the specified class.
func (p *TPopupMenu) InheritsFrom(AClass TClass) bool {
	return PopupMenu_InheritsFrom(p.instance, AClass)
}

// 与一个对象进行比较。
//
// Compare with an object.
func (p *TPopupMenu) Equals(Obj IObject) bool {
	return PopupMenu_Equals(p.instance, CheckPtr(Obj))
}

// 获取类的哈希值。
//
// Get the hash value of the class.
func (p *TPopupMenu) GetHashCode() int32 {
	return PopupMenu_GetHashCode(p.instance)
}

// 文本类信息。
//
// Text information.
func (p *TPopupMenu) ToString() string {
	return PopupMenu_ToString(p.instance)
}

func (p *TPopupMenu) ImagesWidth() int32 {
	return PopupMenu_GetImagesWidth(p.instance)
}

func (p *TPopupMenu) SetImagesWidth(value int32) {
	PopupMenu_SetImagesWidth(p.instance, value)
}

func (p *TPopupMenu) PopupComponent() *TComponent {
	return AsComponent(PopupMenu_GetPopupComponent(p.instance))
}

func (p *TPopupMenu) SetPopupComponent(value IComponent) {
	PopupMenu_SetPopupComponent(p.instance, CheckPtr(value))
}

func (p *TPopupMenu) PopupPoint() TPoint {
	return PopupMenu_GetPopupPoint(p.instance)
}

// 获取文字对齐。
//
// Get Text alignment.
func (p *TPopupMenu) Alignment() TPopupAlignment {
	return PopupMenu_GetAlignment(p.instance)
}

// 设置文字对齐。
//
// Set Text alignment.
func (p *TPopupMenu) SetAlignment(value TPopupAlignment) {
	PopupMenu_SetAlignment(p.instance, value)
}

func (p *TPopupMenu) BiDiMode() TBiDiMode {
	return PopupMenu_GetBiDiMode(p.instance)
}

func (p *TPopupMenu) SetBiDiMode(value TBiDiMode) {
	PopupMenu_SetBiDiMode(p.instance, value)
}

// 获取图标索引列表对象。
func (p *TPopupMenu) Images() *TImageList {
	return AsImageList(PopupMenu_GetImages(p.instance))
}

// 设置图标索引列表对象。
func (p *TPopupMenu) SetImages(value IComponent) {
	PopupMenu_SetImages(p.instance, CheckPtr(value))
}

func (p *TPopupMenu) OwnerDraw() bool {
	return PopupMenu_GetOwnerDraw(p.instance)
}

func (p *TPopupMenu) SetOwnerDraw(value bool) {
	PopupMenu_SetOwnerDraw(p.instance, value)
}

func (p *TPopupMenu) SetOnPopup(fn TNotifyEvent) {
	PopupMenu_SetOnPopup(p.instance, fn)
}

// 获取控件句柄。
//
// Get Control handle.
func (p *TPopupMenu) Handle() HMENU {
	return PopupMenu_GetHandle(p.instance)
}

func (p *TPopupMenu) Items() *TMenuItem {
	return AsMenuItem(PopupMenu_GetItems(p.instance))
}

// 获取组件总数。
//
// Get the total number of components.
func (p *TPopupMenu) ComponentCount() int32 {
	return PopupMenu_GetComponentCount(p.instance)
}

// 获取组件索引。
//
// Get component index.
func (p *TPopupMenu) ComponentIndex() int32 {
	return PopupMenu_GetComponentIndex(p.instance)
}

// 设置组件索引。
//
// Set component index.
func (p *TPopupMenu) SetComponentIndex(value int32) {
	PopupMenu_SetComponentIndex(p.instance, value)
}

// 获取组件所有者。
//
// Get component owner.
func (p *TPopupMenu) Owner() *TComponent {
	return AsComponent(PopupMenu_GetOwner(p.instance))
}

// 获取组件名称。
//
// Get the component name.
func (p *TPopupMenu) Name() string {
	return PopupMenu_GetName(p.instance)
}

// 设置组件名称。
//
// Set the component name.
func (p *TPopupMenu) SetName(value string) {
	PopupMenu_SetName(p.instance, value)
}

// 获取对象标记。
//
// Get the control tag.
func (p *TPopupMenu) Tag() int {
	return PopupMenu_GetTag(p.instance)
}

// 设置对象标记。
//
// Set the control tag.
func (p *TPopupMenu) SetTag(value int) {
	PopupMenu_SetTag(p.instance, value)
}

// 获取指定索引组件。
//
// Get the specified index component.
func (p *TPopupMenu) Components(AIndex int32) *TComponent {
	return AsComponent(PopupMenu_GetComponents(p.instance, AIndex))
}
