//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package vcl

import (
	. "github.com/dollarkillerx/govcl/vcl/api"
)

func (p *TPrinter) SetPrinter(aName string) {
	Printer_SetPrinter(p.instance, aName)
}
