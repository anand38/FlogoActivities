package smartsheetcode

import (
	"github.com/tidwall/gjson"
	"net/http"
	"io/ioutil"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"errors"
	"fmt"
)


func GetSheetDetails(sheetId string,accessToken string)(string,error){
	// ActivityLog is the default logger for the Log Activity
	var activityLog = logger.GetLogger("activity-flogo-SmartSheet-getSheetDetails")

	errReturn:=""
	data:=""
	sheetUrl:="https://api.smartsheet.com/2.0/sheets/"+sheetId
	{
		req,_:=http.NewRequest("GET",sheetUrl,nil)
		req.Header.Set("Authorization","Bearer "+accessToken)
		req.Header.Set("Content-Type","application/json")
		cl := &http.Client{}
		success_resp,err_resp := cl.Do(req)
		if err_resp !=nil{
			errReturn="the HTTP request failed while getting sheet details"
			//fmt.Print("Error Occurred: ",err_resp.Error())
			activityLog.Errorf("some error occurred while trying to fetch sheet details...")
			return "",errors.New(errReturn)
		}
		sheetData,_:=ioutil.ReadAll(success_resp.Body)
		logger.Debug(sheetData)
		//fmt.Println(string(sheetData))
		errCode:=gjson.Get(string(sheetData),"errorCode")
		if(errCode.Exists()){
			errMessage:=gjson.Get(string(sheetData),"message").String()
			activityLog.Errorf(errMessage)
			fmt.Println(errMessage)
			return "",errors.New(errMessage)
		}
		data=string(sheetData)
	}
	return data,nil
}

