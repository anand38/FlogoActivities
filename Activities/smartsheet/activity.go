package smartsheet

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"net/http"
	"fmt"
	"io/ioutil"
	"github.com/anand38/FlogoActivities/smartsheetcode"
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

	//accessToken:="y3ht4k57vq57974zvtsst362pn"
	accessToken:=context.GetInput("Access_Token").(string)
	sheetId:=context.GetInput("Sheet_ID").(string)
	activityOutput:=""


			sheetUrl:="https://api.smartsheet.com/2.0/sheets/"+sheetId
			{
				req,_:=http.NewRequest("GET",sheetUrl,nil)
				req.Header.Set("Authorization","Bearer "+accessToken)
				req.Header.Set("Content-Type","application/json")
				cl := &http.Client{}
				success_resp,err_resp := cl.Do(req)
				if err_resp !=nil{
					fmt.Print("Error Occurred: ",err_resp)
				}else {
					sheetData,_:=ioutil.ReadAll(success_resp.Body)
					activityOutput=smartsheetcode.SetSheetDetails(string(sheetData))
					fmt.Println(activityOutput)
				}
			}
	context.SetOutput("Response_Json", activityOutput)
	return true, nil
	}


