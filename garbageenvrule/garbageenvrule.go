package garbageenvrule

import (
	"errors"
	"fmt"
	"garbageenv/garbageenvreport"
	"io/fs"
	"io/ioutil"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/Pinablink/mTools/mtime"
	uuid "github.com/satori/go.uuid"
)

var mapDirCls map[string]time.Time

//
type GarbageEnvItem struct {
	nameDir     string
	verRemove   bool
	filesRemove []fs.FileInfo
}

//
type GarbageEnvRule struct {
	GarbageEnvRuleUUID uuid.UUID
	filePathInfo       string
	beforeHourn        int
	beforeMinn         int
	garbageRemoveList  []GarbageEnvItem
}

//
func NewGarbageEnvRule(strFilePath string, beforeHour, beforeMin int) *GarbageEnvRule {
	u1, _ := uuid.NewV4()
	refGarbageRemList := make([]GarbageEnvItem, 0)
	return &GarbageEnvRule{GarbageEnvRuleUUID: u1,
		filePathInfo:      strFilePath,
		beforeHourn:       beforeHour,
		beforeMinn:        beforeMin,
		garbageRemoveList: refGarbageRemList}
}

//
func (ref *GarbageEnvRule) Scan() error {

	mapDirCls = make(map[string]time.Time)
	err := ref.scanVer(ref.filePathInfo, "")

	if err != nil {
		return err
	}

	return nil
}

//
func (ref *GarbageEnvRule) RemoveResource(report garbageenvreport.GarbageEnvReport) (garbageenvreport.GarbageEnvReport, error) {

	for _, dataRem := range ref.garbageRemoveList {

		report.QTTotalProcessed = (report.QTTotalProcessed + len(dataRem.filesRemove))

		for _, filesRem := range dataRem.filesRemove {

			strRefFile := dataRem.nameDir + barOS() + filesRem.Name()

			error := os.Remove(strRefFile)
			report.QTRemovedFile = (report.QTRemovedFile + 1)

			if error != nil {
				return report, errors.New("ERROR_REMOVE")
			}

		}

		if dataRem.verRemove {
			report.QTTotalProcessed = (report.QTTotalProcessed + 1)
			strVal := strSplitOS(dataRem.nameDir)
			value, ok := mapDirCls[strVal]

			if ok {
				refCalcTime := createReferPreviusTime(value, ref.beforeHourn, ref.beforeMinn)
				prevOk, perror := refCalcTime.Calc()

				if perror != nil {
					return report, errors.New("ERROR_REMOVE_DIR")
				} else {

					if prevOk {

						remError := os.Remove(dataRem.nameDir)
						report.QTRemovedDir = (report.QTRemovedDir + 1)

						if remError != nil {
							return report, errors.New("ERROR_REMOVE_DIR")
						}

					}

				}

			}

		}

	}

	return report, nil
}

//
func barOS() string {
	strCurr := runtime.GOOS

	if strCurr == "windows" {
		return "\\"
	} else if strCurr == "linux" {
		return "/"
	}

	return ""
}

//
func strSplitOS(strTarget string) string {
	var retStr string = ""
	strCurr := runtime.GOOS

	if strCurr == "windows" {
		strColl := strings.Split(strTarget, "\\")
		lenColl := len(strColl)

		if lenColl > 0 {
			retStr = strColl[lenColl-1]
		}

	} else if strCurr == "linux" {
		strColl := strings.Split(strTarget, "/")
		lenColl := len(strColl)

		if lenColl > 0 {
			retStr = strColl[lenColl-1]
		}
	}

	return retStr
}

//
func dirTarget(strRefSeed, strRefChid string) string {

	var strPath string

	if len(strRefChid) > 0 {
		var strBarOS string = barOS()
		var strBasechild string = "%s%s%s"
		strPath = fmt.Sprintf(strBasechild, strRefSeed, strBarOS, strRefChid)

	} else {
		strPath = strRefSeed
	}

	return strPath
}

//
func (ref *GarbageEnvRule) scanVer(strSeed, strPathDir string) error {
	var refReturn GarbageEnvItem = GarbageEnvItem{nameDir: strPathDir}
	var listFileInfo []fs.FileInfo = make([]fs.FileInfo, 0)
	strTarget := dirTarget(strSeed, strPathDir)
	iFilesInfo, err := ioutil.ReadDir(strTarget)

	if err != nil {
		return err
	}

	for _, fileInfo := range iFilesInfo {

		strNameResource := fileInfo.Name()
		// infSize := fileInfo.Size()
		dtTimeCreated := fileInfo.ModTime()

		if fileInfo.IsDir() {

			mapDirCls[fileInfo.Name()] = fileInfo.ModTime()
			err = ref.scanVer(strTarget, strNameResource)

			if err != nil {
				return err
			}

		} else {

			refCalcTime := createReferPreviusTime(dtTimeCreated, ref.beforeHourn, ref.beforeMinn)

			prev, err := refCalcTime.Calc()

			if err != nil {
				return errors.New("ERR-CALC-TIME")
			} else {

				if prev {
					listFileInfo = append(listFileInfo, fileInfo)
				}

			}

		}

	}

	refReturn.filesRemove = listFileInfo
	refReturn.verRemove = !(strTarget == ref.filePathInfo)
	refReturn.nameDir = strTarget
	ref.garbageRemoveList = append(ref.garbageRemoveList, refReturn)

	return err
}

func createReferPreviusTime(refDtCreateDir time.Time, refBeforeTime ...int) mtime.Previustime {
	var valBeforeHour int = refBeforeTime[0]
	var valBeforeMin int = refBeforeTime[1]
	var calcTimer mtime.Previustime = *mtime.NewPreviustime(refDtCreateDir)

	calcTimer.DefHourIn(time.Now())
	calcTimer.DefHourInterval(getHour(valBeforeHour))

	if valBeforeMin != 0 {
		calcTimer.DefMinuteInterval(valBeforeMin)
	}

	return calcTimer
}

func getHour(refHourNum int) mtime.HOUR {
	var refHour mtime.HOUR

	switch refHourNum {

	case 0:
		refHour = mtime.NOT_HOUR
	case 1:
		refHour = mtime.HOUR_1
	case 2:
		refHour = mtime.HOUR_2
	case 3:
		refHour = mtime.HOUR_3
	case 4:
		refHour = mtime.HOUR_4
	case 5:
		refHour = mtime.HOUR_5
	case 6:
		refHour = mtime.HOUR_6
	case 7:
		refHour = mtime.HOUR_7
	case 8:
		refHour = mtime.HOUR_8
	case 9:
		refHour = mtime.HOUR_9
	case 10:
		refHour = mtime.HOUR_10
	case 11:
		refHour = mtime.HOUR_11
	case 12:
		refHour = mtime.HOUR_12
	case 13:
		refHour = mtime.HOUR_13
	case 14:
		refHour = mtime.HOUR_14
	case 15:
		refHour = mtime.HOUR_15
	case 16:
		refHour = mtime.HOUR_16
	case 17:
		refHour = mtime.HOUR_17
	case 18:
		refHour = mtime.HOUR_18
	case 19:
		refHour = mtime.HOUR_19
	case 20:
		refHour = mtime.HOUR_20
	case 21:
		refHour = mtime.HOUR_21
	case 22:
		refHour = mtime.HOUR_22
	case 23:
		refHour = mtime.HOUR_23
	case 24:
		refHour = mtime.HOUR_24
	}

	return refHour
}
