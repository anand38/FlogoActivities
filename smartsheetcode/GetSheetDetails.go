package smartsheetcode

import (
	"github.com/tidwall/gjson"
	"net/http"
	"io/ioutil"
	"github.com/tidwall/sjson"
	"errors"
	"strconv"
	"time"
)

//GetSheetDetails accepts sheetID, accessToken and returns data of that sheetID
func GetSheetDetails(sheetID string,accessToken string)(string,error){

	errReturn:=""
	activityOutputTmp:=`{}`

	sheetURL:="https://api.smartsheet.com/2.0/sheets/"+sheetID
	{
		req,_:=http.NewRequest("GET",sheetURL,nil)
		req.Header.Set("Authorization","Bearer "+accessToken)
		req.Header.Set("Content-Type","application/json")
		cl := &http.Client{
			Timeout: time.Second * 30,
		}
		successResp,errResp := cl.Do(req)  //call to Smartsheet API
		if errResp !=nil{
			//set error return
			errReturn="the HTTP request failed while getting sheet details:"+errResp.Error()
			return "",errors.New(errReturn)
		}
		// Close http connection
		defer successResp.Body.Close()

		sheetData,_:=ioutil.ReadAll(successResp.Body)
		errCode:=gjson.Get(string(sheetData),"errorCode")
		if errCode.Exists(){
			errMessage:=gjson.Get(string(sheetData),"message").String()
			return "",errors.New(errMessage)
		}

		columns:=gjson.Get(string(sheetData),"columns")
		columnLength,_:=strconv.Atoi(gjson.Get(columns.String(),"#").String())
		var col=make([]string,columnLength)
		for t:=0;t<columnLength;t++{
			col[t]=gjson.Get(columns.String(),strconv.Itoa(t)+".title").String()
		}
		rows := gjson.Get(string(sheetData),"rows.#.cells")

		//creating json for output
		for i:= range rows.Array(){
			rowcell:= gjson.Get(rows.String(),strconv.Itoa(i)) //single row
			for tmp:=0;tmp<columnLength;tmp++  {
				activityOutputTmp,_=sjson.Set(activityOutputTmp,"rows."+strconv.Itoa(i)+"."+col[tmp],gjson.Get(rowcell.String(),strconv.Itoa(tmp)+".value").String())
			}
		}


	}
	return activityOutputTmp,nil
}
