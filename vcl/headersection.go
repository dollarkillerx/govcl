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

type THeaderSection struct {
	IObject
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

// 创建一个新的对象。
//
// Create a new object.
func NewHeaderSection(AOwner *TCollection) *THeaderSection {
	h := new(THeaderSection)
	h.instance = HeaderSection_Create(CheckPtr(AOwner))
	h.ptr = unsafe.Pointer(h.instance)
	setFinalizer(h, (*THeaderSection).Free)
	return h
}

// 动态转换一个已存在的对象实例。
//
// Dynamically convert an existing object instance.
func AsHeaderSection(obj interface{}) *THeaderSection {
	instance, ptr := getInstance(obj)
	if instance == 0 {
		return nil
	}
	return &THeaderSection{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// 新建一个对象来自已经存在的对象实例指针。
//
// Create a new object from an existing object instance pointer.
// Deprecated: use AsHeaderSection.
func HeaderSectionFromInst(inst uintptr) *THeaderSection {
	return AsHeaderSection(inst)
}

// 新建一个对象来自已经存在的对象实例。
//
// Create a new object from an existing object instance.
// Deprecated: use AsHeaderSection.
func HeaderSectionFromObj(obj IObject) *THeaderSection {
	return AsHeaderSection(obj)
}

// 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
//
// Create a new object from an unsecured address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsHeaderSection.
func HeaderSectionFromUnsafePointer(ptr unsafe.Pointer) *THeaderSection {
	return AsHeaderSection(ptr)
}

// -------------------------- Deprecated end --------------------------
// 释放对象。
//
// Free object.
func (h *THeaderSection) Free() {
	if h.instance != 0 {
		HeaderSection_Free(h.instance)
		h.instance, h.ptr = 0, nullptr
	}
}

// 返回对象实例指针。
//
// Return object instance pointer.
func (h *THeaderSection) Instance() uintptr {
	return h.instance
}

// 获取一个不安全的地址。
//
// Get an unsafe address.
func (h *THeaderSection) UnsafeAddr() unsafe.Pointer {
	return h.ptr
}

// 检测地址是否为空。
//
// Check if the address is empty.
func (h *THeaderSection) IsValid() bool {
	return h.instance != 0
}

// 检测当前对象是否继承自目标对象。
//
// Checks whether the current object is inherited from the target object.
func (h *THeaderSection) Is() TIs {
	return TIs(h.instance)
}

// 动态转换当前对象为目标对象。
//
// Dynamically convert the current object to the target object.
//func (h *THeaderSection) As() TAs {
//    return TAs(h.instance)
//}

// 获取类信息指针。
//
// Get class information pointer.
func THeaderSectionClass() TClass {
	return HeaderSection_StaticClassType()
}

// 复制一个对象，如果对象实现了此方法的话。
//
// Copy an object, if the object implements this method.
func (h *THeaderSection) Assign(Source IObject) {
	HeaderSection_Assign(h.instance, CheckPtr(Source))
}

// 获取类名路径。
//
// Get the class name path.
func (h *THeaderSection) GetNamePath() string {
	return HeaderSection_GetNamePath(h.instance)
}

// 获取类的类型信息。
//
// Get class type information.
func (h *THeaderSection) ClassType() TClass {
	return HeaderSection_ClassType(h.instance)
}

// 获取当前对象类名称。
//
// Get the current object class name.
func (h *THeaderSection) ClassName() string {
	return HeaderSection_ClassName(h.instance)
}

// 获取当前对象实例大小。
//
// Get the current object instance size.
func (h *THeaderSection) InstanceSize() int32 {
	return HeaderSection_InstanceSize(h.instance)
}

// 判断当前类是否继承自指定类。
//
// Determine whether the current class inherits from the specified class.
func (h *THeaderSection) InheritsFrom(AClass TClass) bool {
	return HeaderSection_InheritsFrom(h.instance, AClass)
}

// 与一个对象进行比较。
//
// Compare with an object.
func (h *THeaderSection) Equals(Obj IObject) bool {
	return HeaderSection_Equals(h.instance, CheckPtr(Obj))
}

// 获取类的哈希值。
//
// Get the hash value of the class.
func (h *THeaderSection) GetHashCode() int32 {
	return HeaderSection_GetHashCode(h.instance)
}

// 文本类信息。
//
// Text information.
func (h *THeaderSection) ToString() string {
	return HeaderSection_ToString(h.instance)
}

// 获取左边位置。
//
// Get Left position.
func (h *THeaderSection) Left() int32 {
	return HeaderSection_GetLeft(h.instance)
}

func (h *THeaderSection) Right() int32 {
	return HeaderSection_GetRight(h.instance)
}

// 获取文字对齐。
//
// Get Text alignment.
func (h *THeaderSection) Alignment() TAlignment {
	return HeaderSection_GetAlignment(h.instance)
}

// 设置文字对齐。
//
// Set Text alignment.
func (h *THeaderSection) SetAlignment(value TAlignment) {
	HeaderSection_SetAlignment(h.instance, value)
}

// 获取图像在images中的索引。
func (h *THeaderSection) ImageIndex() int32 {
	return HeaderSection_GetImageIndex(h.instance)
}

// 设置图像在images中的索引。
func (h *THeaderSection) SetImageIndex(value int32) {
	HeaderSection_SetImageIndex(h.instance, value)
}

func (h *THeaderSection) MaxWidth() int32 {
	return HeaderSection_GetMaxWidth(h.instance)
}

func (h *THeaderSection) SetMaxWidth(value int32) {
	HeaderSection_SetMaxWidth(h.instance, value)
}

func (h *THeaderSection) MinWidth() int32 {
	return HeaderSection_GetMinWidth(h.instance)
}

func (h *THeaderSection) SetMinWidth(value int32) {
	HeaderSection_SetMinWidth(h.instance, value)
}

// 获取文本。
func (h *THeaderSection) Text() string {
	return HeaderSection_GetText(h.instance)
}

// 设置文本。
func (h *THeaderSection) SetText(value string) {
	HeaderSection_SetText(h.instance, value)
}

// 获取宽度。
//
// Get width.
func (h *THeaderSection) Width() int32 {
	return HeaderSection_GetWidth(h.instance)
}

// 设置宽度。
//
// Set width.
func (h *THeaderSection) SetWidth(value int32) {
	HeaderSection_SetWidth(h.instance, value)
}

func (h *THeaderSection) Collection() *TCollection {
	return AsCollection(HeaderSection_GetCollection(h.instance))
}

func (h *THeaderSection) SetCollection(value *TCollection) {
	HeaderSection_SetCollection(h.instance, CheckPtr(value))
}

func (h *THeaderSection) Index() int32 {
	return HeaderSection_GetIndex(h.instance)
}

func (h *THeaderSection) SetIndex(value int32) {
	HeaderSection_SetIndex(h.instance, value)
}

func (h *THeaderSection) DisplayName() string {
	return HeaderSection_GetDisplayName(h.instance)
}

func (h *THeaderSection) SetDisplayName(value string) {
	HeaderSection_SetDisplayName(h.instance, value)
}
