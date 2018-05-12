package smartsheet

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
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
	
	accessToken:=context.GetInput("Access_Token").(string)
	sheetId:=context.GetInput("Sheet_ID").(string)
	activityOutput:=""

	activityOutput,err=smartsheetcode.GetSheetDetails(sheetId,accessToken)
	if err!=nil {
		return false,err
	}else{
		context.SetOutput("Response_Json", activityOutput)
	}
	return true, nil
	}


