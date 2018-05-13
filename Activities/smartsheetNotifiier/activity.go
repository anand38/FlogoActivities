package smartsheetNotifiier

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"fmt"
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

	accessToken:=context.GetInput("Access_Token").(string)
	sheetID:=context.GetInput("Sheet_ID").(string)
	status:=context.GetInput("Status").(string)
	percentComplete:=context.GetInput("Percent_Complete").(string)
	subject:=context.GetInput("Sheet_ID").(string)
	message:=context.GetInput("Message").(string)

	fmt.Print("Accesstoken",accessToken)
	fmt.Println("Sheetid",sheetID)
	fmt.Println(status)
	fmt.Println(percentComplete)
	fmt.Println("subject",subject)
	fmt.Println("Message",message)

	if status==""{
		fmt.Println("status is empty")
	}
	if percentComplete=="" {
		fmt.Println("percent complete is empty")
	}
	return true, nil
}
