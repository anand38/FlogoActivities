package smartsheetcode

import (
	"github.com/tidwall/gjson"
	"net/http"
	"io/ioutil"
	"github.com/tidwall/sjson"
	"errors"
	"strconv"
)

//GetSheetDetails accepts sheedID,accessToken and returns sheet Details
func GetSheetDetails(sheetID string,accessToken string)(string,error){
	// ActivityLog is the default logger for the Log Activity

	errReturn:=""
	activityOutputTmp:=`{}`

	sheetURL:="https://api.smartsheet.com/2.0/sheets/"+sheetID
	{
		req,_:=http.NewRequest("GET",sheetURL,nil)
		req.Header.Set("Authorization","Bearer "+accessToken)
		req.Header.Set("Content-Type","application/json")
		cl := &http.Client{}
		successResp,errResp := cl.Do(req)
		if errResp !=nil{
			errReturn="the HTTP request failed while getting sheet details"
			//fmt.Print("Error Occurred: ",err_resp.Error())
			return "",errors.New(errReturn)
		}
		sheetData,_:=ioutil.ReadAll(successResp.Body)
		//fmt.Println(string(sheetData))
		errCode:=gjson.Get(string(sheetData),"errorCode")
		if(errCode.Exists()){
			errMessage:=gjson.Get(string(sheetData),"message").String()
			//fmt.Println(errMessage)
			return "",errors.New(errMessage)
		}


		columns:=gjson.Get(string(sheetData),"columns")
		columnLength,_:=strconv.Atoi(gjson.Get(columns.String(),"#").String())
		var col=make([]string,columnLength)
		for t:=0;t<columnLength;t++{
			col[t]=gjson.Get(columns.String(),strconv.Itoa(t)+".title").String()
		}
		rows := gjson.Get(string(sheetData),"rows.#.cells")

		for i:= range rows.Array(){
			rowcell:= gjson.Get(rows.String(),strconv.Itoa(i)) ///single row
			for tmp:=0;tmp<columnLength;tmp++  {
				activityOutputTmp,_=sjson.Set(activityOutputTmp,"rows."+strconv.Itoa(i)+"."+col[tmp],gjson.Get(rowcell.String(),strconv.Itoa(tmp)+".value").String())
			}
		}


	}
	return activityOutputTmp,nil
}
