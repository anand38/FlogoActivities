package smartsheetcode

import (
	"github.com/tidwall/gjson"
	"strconv"
	"github.com/tidwall/sjson"
)

func setSurveyDetails(sheetData string)string{
	activityOutput:=""
	columns:=gjson.Get(string(sheetData),"columns")
	columnLength,_:=strconv.Atoi(gjson.Get(columns.String(),"#").String())
	var col=make([]string,columnLength)
	for t:=0;t<columnLength;t++{
		col[t]=gjson.Get(columns.String(),strconv.Itoa(t)+".title").String()
	}
	rows := gjson.Get(string(sheetData),"rows.#.cells")
	activityOutputTmp:=`{}`

	for i,_ := range rows.Array(){
		rowcell:= gjson.Get(rows.String(),strconv.Itoa(i)) ///single row
		for tmp:=0;tmp<columnLength;tmp++  {
			activityOutputTmp,_=sjson.Set(activityOutputTmp,"rows."+strconv.Itoa(i)+"."+col[tmp],gjson.Get(rowcell.String(),strconv.Itoa(tmp)+".value").String())
		}
	}
	activityOutput=activityOutputTmp
	return activityOutput
}
