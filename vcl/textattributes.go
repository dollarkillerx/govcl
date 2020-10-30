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

type TTextAttributes struct {
	IObject
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

// 动态转换一个已存在的对象实例。
//
// Dynamically convert an existing object instance.
func AsTextAttributes(obj interface{}) *TTextAttributes {
	instance, ptr := getInstance(obj)
	if instance == 0 {
		return nil
	}
	return &TTextAttributes{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// 新建一个对象来自已经存在的对象实例指针。
//
// Create a new object from an existing object instance pointer.
// Deprecated: use AsTextAttributes.
func TextAttributesFromInst(inst uintptr) *TTextAttributes {
	return AsTextAttributes(inst)
}

// 新建一个对象来自已经存在的对象实例。
//
// Create a new object from an existing object instance.
// Deprecated: use AsTextAttributes.
func TextAttributesFromObj(obj IObject) *TTextAttributes {
	return AsTextAttributes(obj)
}

// 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
//
// Create a new object from an unsecured address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsTextAttributes.
func TextAttributesFromUnsafePointer(ptr unsafe.Pointer) *TTextAttributes {
	return AsTextAttributes(ptr)
}

// -------------------------- Deprecated end --------------------------
// 返回对象实例指针。
//
// Return object instance pointer.
func (t *TTextAttributes) Instance() uintptr {
	return t.instance
}

// 获取一个不安全的地址。
//
// Get an unsafe address.
func (t *TTextAttributes) UnsafeAddr() unsafe.Pointer {
	return t.ptr
}

// 检测地址是否为空。
//
// Check if the address is empty.
func (t *TTextAttributes) IsValid() bool {
	return t.instance != 0
}

// 检测当前对象是否继承自目标对象。
//
// Checks whether the current object is inherited from the target object.
func (t *TTextAttributes) Is() TIs {
	return TIs(t.instance)
}

// 动态转换当前对象为目标对象。
//
// Dynamically convert the current object to the target object.
//func (t *TTextAttributes) As() TAs {
//    return TAs(t.instance)
//}

// 获取类信息指针。
//
// Get class information pointer.
func TTextAttributesClass() TClass {
	return TextAttributes_StaticClassType()
}

// 复制一个对象，如果对象实现了此方法的话。
//
// Copy an object, if the object implements this method.
func (t *TTextAttributes) Assign(Source IObject) {
	TextAttributes_Assign(t.instance, CheckPtr(Source))
}

// 获取类名路径。
//
// Get the class name path.
func (t *TTextAttributes) GetNamePath() string {
	return TextAttributes_GetNamePath(t.instance)
}

// 获取类的类型信息。
//
// Get class type information.
func (t *TTextAttributes) ClassType() TClass {
	return TextAttributes_ClassType(t.instance)
}

// 获取当前对象类名称。
//
// Get the current object class name.
func (t *TTextAttributes) ClassName() string {
	return TextAttributes_ClassName(t.instance)
}

// 获取当前对象实例大小。
//
// Get the current object instance size.
func (t *TTextAttributes) InstanceSize() int32 {
	return TextAttributes_InstanceSize(t.instance)
}

// 判断当前类是否继承自指定类。
//
// Determine whether the current class inherits from the specified class.
func (t *TTextAttributes) InheritsFrom(AClass TClass) bool {
	return TextAttributes_InheritsFrom(t.instance, AClass)
}

// 与一个对象进行比较。
//
// Compare with an object.
func (t *TTextAttributes) Equals(Obj IObject) bool {
	return TextAttributes_Equals(t.instance, CheckPtr(Obj))
}

// 获取类的哈希值。
//
// Get the hash value of the class.
func (t *TTextAttributes) GetHashCode() int32 {
	return TextAttributes_GetHashCode(t.instance)
}

// 文本类信息。
//
// Text information.
func (t *TTextAttributes) ToString() string {
	return TextAttributes_ToString(t.instance)
}

func (t *TTextAttributes) Charset() TFontCharset {
	return TextAttributes_GetCharset(t.instance)
}

func (t *TTextAttributes) SetCharset(value TFontCharset) {
	TextAttributes_SetCharset(t.instance, value)
}

// 获取颜色。
//
// Get color.
func (t *TTextAttributes) Color() TColor {
	return TextAttributes_GetColor(t.instance)
}

// 设置颜色。
//
// Set color.
func (t *TTextAttributes) SetColor(value TColor) {
	TextAttributes_SetColor(t.instance, value)
}

// 获取组件名称。
//
// Get the component name.
func (t *TTextAttributes) Name() string {
	return TextAttributes_GetName(t.instance)
}

// 设置组件名称。
//
// Set the component name.
func (t *TTextAttributes) SetName(value string) {
	TextAttributes_SetName(t.instance, value)
}

func (t *TTextAttributes) Pitch() TFontPitch {
	return TextAttributes_GetPitch(t.instance)
}

func (t *TTextAttributes) SetPitch(value TFontPitch) {
	TextAttributes_SetPitch(t.instance, value)
}

func (t *TTextAttributes) Size() int32 {
	return TextAttributes_GetSize(t.instance)
}

func (t *TTextAttributes) SetSize(value int32) {
	TextAttributes_SetSize(t.instance, value)
}

func (t *TTextAttributes) Style() TFontStyles {
	return TextAttributes_GetStyle(t.instance)
}

func (t *TTextAttributes) SetStyle(value TFontStyles) {
	TextAttributes_SetStyle(t.instance, value)
}

// 获取高度。
//
// Get height.
func (t *TTextAttributes) Height() int32 {
	return TextAttributes_GetHeight(t.instance)
}

// 设置高度。
//
// Set height.
func (t *TTextAttributes) SetHeight(value int32) {
	TextAttributes_SetHeight(t.instance, value)
}
