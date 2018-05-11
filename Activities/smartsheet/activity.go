package smartsheet

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"net/http"
	"io/ioutil"
	"github.com/anand38/FlogoActivities/smartsheetcode"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"errors"
	"github.com/tidwall/gjson"
)

// MyActivity is a stub for your Activity implementation
type MyActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &MyActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *MyActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *MyActivity) Eval(context activity.Context) (done bool, err error)  {

	// do eval
	logger.Debug("Started execution....")
	accessToken:=context.GetInput("Access_Token").(string)
	sheetId:=context.GetInput("Sheet_ID").(string)
	activityOutput:=""
	errReturn:=""

			sheetUrl:="https://api.smartsheet.com/2.0/sheets/"+sheetId
			{
				req,_:=http.NewRequest("GET",sheetUrl,nil)
				req.Header.Set("Authorization","Bearer "+accessToken)
				req.Header.Set("Content-Type","application/json")
				cl := &http.Client{}
				success_resp,err_resp := cl.Do(req)
				if err_resp !=nil{
					errReturn="The HTTP request failed while getting sheet details..."
					//fmt.Print("Error Occurred: ",err_resp.Error())
					logger.Debug("Some error occurred")
					return false,errors.New(errReturn)
				}else {
					sheetData,_:=ioutil.ReadAll(success_resp.Body)
					logger.Debug(sheetData)
					//fmt.Println(string(sheetData))
					errCode:=gjson.Get(string(sheetData),"errorCode")
					if(errCode.Exists()){
						errMessage:=gjson.Get(string(sheetData),"message")
						logger.Debug(errMessage)
						//fmt.Println(errMessage)
					}
					activityOutput=smartsheetcode.SetSheetDetails(string(sheetData))
					//fmt.Println(activityOutput)
					logger.Debug(activityOutput)
				}
			}

	context.SetOutput("Response_Json", activityOutput)
	return true, nil
	}


