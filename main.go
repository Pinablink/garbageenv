package main

import (
	"fmt"
	"garbageenv/garbageenvparam"
	"garbageenv/garbageenvreport"
	"garbageenv/garbageenvrule"
	"strconv"

	"github.com/Pinablink/warning"
)

func main() {

	var garbageParams *garbageenvparam.GarbageEnvParam = garbageenvparam.NewGarbageEnvParam()

	ok := garbageParams.ValidArgs()

	if ok {

		var refGarbageRule *garbageenvrule.GarbageEnvRule = garbageenvrule.NewGarbageEnvRule(
			garbageParams.ParamClearPath, garbageParams.ParamBeforeHour, garbageParams.ParamBeforeMin)
		err := refGarbageRule.Scan()

		if err != nil {
			var msg string

			if err.Error() == "ERR-CALC-TIME" {
				msg = "GarbageEnv: Ocorreu erro no proc de Scan " + "ERR-CALC-TIME"
			} else {
				msg = "GarbageEnv: Ocorreu erro no proc de Scan"
			}

			sendMsgWarning(msg, true)
			panic(err)
		}

		var report = garbageenvreport.GarbageEnvReport{QTRemovedFile: 0,
			QTRemovedDir: 0, QTTotalProcessed: 0,
		}

		report, err = refGarbageRule.RemoveResource(report)
		var msg string
		var strMessage string = "%s ;QTFILREM - %s; QTDIREM - %s; QTTOTALPROC - %s"
		var errorProc bool = false

		if err != nil {

			msg = "GarbageEnv: Erro RemoveResource"
			msg = fmt.Sprintf(strMessage, msg, strconv.Itoa(report.QTRemovedFile), strconv.Itoa(report.QTRemovedDir), strconv.Itoa(report.QTTotalProcessed))
			errorProc = true

		} else {

			qtRemovedFile := report.QTRemovedFile
			qtRemovedDir := report.QTRemovedDir
			qtTotalProcessed := report.QTTotalProcessed

			if qtRemovedFile > 0 || qtRemovedDir > 0 {
				msg = "GarbageEnv: Sucesso na Remoção de Recursos"
			} else {
				msg = "GarbageEnv: Não houve Remoção de Recursos"
			}

			msg = fmt.Sprintf(strMessage, msg, strconv.Itoa(qtRemovedFile), strconv.Itoa(qtRemovedDir), strconv.Itoa(qtTotalProcessed))

		}
		sendMsgWarning(msg, errorProc)

	} else {
		//garbageParams.GarbageEnvMsg()
		// Esta funcionalidade necessita de uma conta que disponibilize serviço de SMS
		msg := "GarbageEnv: Problemas na Parametrização"
		sendMsgWarning(msg, true)
	}

}

func sendMsgWarning(refMsg string, isErr bool) {
	var refWarningMsg *warning.Warning = warning.NewWarning()
	refWarningMsg.ConfigSMSWarning(garbageenvparam.URL_SERVICE_SMS, garbageenvparam.USUARIO_SERVICE_SMS, garbageenvparam.TOKEN_PERM_SERVICE, garbageenvparam.NUM_TEL_SERVICE_TARGET)
	refWarningMsg.SendSMSMessage(refMsg, isErr)
}
