// *********************************************************************************************************************
// ***                                        CONFIDENTIAL --- CUSTOM STUDIOS                                        ***
// *********************************************************************************************************************
// * Auth: ColeCai                                                                                                     *
// * Date: 2023/10/25 22:09:20                                                                                         *
// * Proj: work                                                                                                        *
// * Pack: padding                                                                                                     *
// * File: header_test.go                                                                                              *
// *-------------------------------------------------------------------------------------------------------------------*
// * Overviews:                                                                                                        *
// *-------------------------------------------------------------------------------------------------------------------*
package padding

import (
	"strings"
	"testing"
)

// *********************************************************************************************************************
// * SUMMARY:                                                                                                          *
// * 	- inputs:                                                                                                      *
// *        - title string:
// * 	- output:                                                                                                      *
// * WARNING:                                                                                                          *
// * HISTORY:                                                                                                          *
// *    -create: 2023/10/25 22:34:51 ColeCai.                                                                          *
// *    -update: 2023/10/25 22:36:46 ColeCai.                                                                          *
// *********************************************************************************************************************

func GenStr(title, content string) string {
	pre := "// * "
	pre = pre + title + " "
	pre += content
	l := 120 - len(pre) - 1
	pre += strings.Repeat(" ", l)
	pre += "*"
	return pre
}

func TestHeader(t *testing.T) {
	auht := "COleCai"
	date := "2023/10/25 22:09:20"
	proj := "work"
	pack := "padding"
	file := "headerxxxxxxxx_test.go"

	l1 := "// " + strings.Repeat("*", 117)
	tpk := strings.Repeat(" ", 50)
	l2 := "// " + "***" + tpk + "G O L A N D" + tpk + "***"
	l3 := l1
	l4 := GenStr("Auth:", auht)
	l5 := GenStr("Date:", date)
	l6 := GenStr("Proj:", proj)
	l7 := GenStr("Pack:", pack)
	l8 := GenStr("File:", file)
	l9 := "// *" + strings.Repeat("-", 115) + "*"
	l10 := GenStr("Overviews:", "")
	l11 := l9
	l12 := GenStr("Functions:", "")
	l13 := "// *" + strings.Repeat(" -", 57) + " *"

	t.Log(l1)
	t.Log(l2)
	t.Log(l3)
	t.Log(l4)
	t.Log(l5)
	t.Log(l6)
	t.Log(l7)
	t.Log(l8)
	t.Log(l9)
	t.Log(l10)
	t.Log(l11)
	t.Log(l12)
	t.Log(l13)
}
