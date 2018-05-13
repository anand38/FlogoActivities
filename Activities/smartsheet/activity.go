package smartsheet

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/anand38/FlogoActivities/smartsheetcode"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

var activityLog = logger.GetLogger("activity-flogo-SmartSheet-getSheetDetails")


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


	//variables initialized here
	accessToken:=context.GetInput("Access_Token").(string)
	sheetID:=context.GetInput("Sheet_ID").(string)

	result,err:=smartsheetcode.GetSheetDetails(sheetID,accessToken)
	if err!=nil {
		activityLog.Errorf("Error occurred:"+err.Error())
		return false,err
	}

	//setting the output here
	context.SetOutput("Response_Json", result)

	return true, nil
	}


