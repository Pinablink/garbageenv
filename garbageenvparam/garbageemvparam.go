package garbageenvparam

import (
	"garbageenv/garbagehelp"
	"os"
	"strconv"

	uuid "github.com/satori/go.uuid"
)

const (
	URL_SERVICE_SMS        string = "<SUA CONFIG AQUI>"
	USUARIO_SERVICE_SMS           = "<SUA CONFIG AQUI>"
	TOKEN_PERM_SERVICE            = "<SUA CONFIG AQUI>"
	NUM_TEL_SERVICE_TARGET        = "<SUA CONFIG AQUI>"
)

type GarbageEnvParam struct {
	GarbageEnvParamUUID uuid.UUID
	ParamClearPath      string
	ParamBeforeHour     int
	ParamBeforeMin      int
	GarbageEnvMsg       func()
}

//
func NewGarbageEnvParam() *GarbageEnvParam {
	u1, _ := uuid.NewV4()
	return &GarbageEnvParam{GarbageEnvParamUUID: u1, ParamBeforeHour: 0, ParamBeforeMin: 0}
}

//
func (ref *GarbageEnvParam) ValidArgs() bool {
	var retTeste bool = false

	if len(os.Args) == 2 && os.Args[1] == "-h" {

		ref.GarbageEnvMsg = showSShowHelp
		retTeste = false

	} else if len(os.Args) == 5 {

		var paramControlOK bool = validMainParam()

		if !paramControlOK {

			ref.GarbageEnvMsg = showWShowHelp
			retTeste = false

		} else {

			retTeste = ref.validValueHour()

		}

	} else if len(os.Args) == 7 {

		var strParamBeforeMin string = os.Args[5]
		var paramControlOK bool = (validMainParam() && strParamBeforeMin == "-bm")

		if !paramControlOK {
			ref.GarbageEnvMsg = showWShowHelp
			retTeste = false
		} else {

			retTeste = ref.validValueMin()

		}

	} else {
		ref.GarbageEnvMsg = showWShowHelp
		retTeste = false
	}

	return retTeste
}

//
func (ref *GarbageEnvParam) validValueHour() bool {
	var strValPathClear string = os.Args[2]
	var strValBeforeHour string = os.Args[4]

	valBH, err1 := strconv.Atoi(strValBeforeHour)

	if err1 != nil {
		return false
	}

	ref.ParamClearPath = strValPathClear
	ref.ParamBeforeHour = valBH

	return true
}

//
func (ref *GarbageEnvParam) validValueMin() bool {
	var strValPathClear string = os.Args[2]
	var strValBeforeMin string = os.Args[6]

	valBM, err1 := strconv.Atoi(strValBeforeMin)

	if err1 != nil {
		return false
	}

	ref.ParamClearPath = strValPathClear
	ref.ParamBeforeMin = valBM

	return true
}

//
func validMainParam() bool {
	var strParamPathClear string = os.Args[1]
	var strParamBeforeHour string = os.Args[3]

	return (strParamPathClear == "-p" && strParamBeforeHour == "-bh")
}

//
func showWShowHelp() {
	garbagehelp.WarningShowHelp()
}

//
func showSShowHelp() {
	garbagehelp.Showhelp()
}
