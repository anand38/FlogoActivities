package smartsheetNotifiier

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/anand38/FlogoActivities/smartsheetcode"

	"fmt"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"github.com/tidwall/gjson"
	"strconv"
	"github.com/tidwall/sjson"
	"strings"
	"net/http"
	"bytes"
	"time"
	"io/ioutil"
)
var activityLog = logger.GetLogger("activity-flogo-SmartSheet-Notifier")

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
	subject:=context.GetInput("Subject").(string)
	message:=context.GetInput("Message").(string)
	output:=`{}`


	sheetdata, err := smartsheetcode.GetSheetDetails(sheetID, accessToken)
	if err != nil {
		activityLog.Errorf(err.Error())
		fmt.Println("some Error occurred")
		return false, nil
	}

	output, _ = sjson.Set(output, "formatDetails.paperSize", "A4")
	output, _ = sjson.Set(output, "format", "PDF")
	output, _ = sjson.Set(output, "ccMe", false)
	output, _ = sjson.Set(output, "message", message)
	output, _ = sjson.Set(output, "subject", subject)


	if status !="" {

		rows := gjson.Get(sheetdata, "rows")
		for i := range rows.Array() {
			if gjson.Get(rows.String(), strconv.Itoa(i)+".Status").String() == status {
				output, _ = sjson.Set(output, "sendTo."+strconv.Itoa(i)+".email", gjson.Get(rows.String(), strconv.Itoa(i)+".Assigned To").String())
			}
		}
		fmt.Println("string1: ",output)

	}else if percentComplete!="" {
		rows := gjson.Get(sheetdata, "rows")
		fmt.Println(rows.String())
		count:=0
		for i := range rows.Array() {
			perCompjson,_:=strconv.ParseFloat(gjson.Get(rows.String(), strconv.Itoa(i)+".% Complete").String(),64)
			strSlice := strings.Split(percentComplete, "%")
			perCompdata,_:=strconv.Atoi(strSlice[0])
			perCompjson=perCompjson*100
			if perCompjson <= float64(perCompdata) {
				fmt.Println("herer.......")
				fmt.Println("Percompdata:",perCompdata)
				fmt.Println("Percomjson:",perCompjson)
				output, _ = sjson.Set(output, "sendTo."+strconv.Itoa(count)+".email", gjson.Get(rows.String(), strconv.Itoa(count)+".Assigned To").String())
				count++
			}

		}
		fmt.Println("string2: ",output)
	}else{
		activityLog.Errorf("Either status or Percent complete needs to be selected")
		return false,nil
	}


		//Sending mail via Smartsheet API
		mailurl:="https://api.smartsheet.com/2.0/sheets/"+sheetID+"/emails"
		req,_:=http.NewRequest("POST",mailurl,bytes.NewBuffer([]byte(output)))
		req.Header.Set("Authorization","Bearer "+accessToken)
		req.Header.Set("Content-Type","application/json")

		cl := &http.Client{
			Timeout: time.Second * 30,
		}
		successResp,errResp := cl.Do(req)  //call to Smartsheet API
		if errResp !=nil{

			return false,nil
		}
		// Close http connection
		defer successResp.Body.Close()
		mailResp,_:=ioutil.ReadAll(successResp.Body)
		fmt.Println("Output:" ,string(mailResp))
		context.SetOutput("Response_Json",string(mailResp))
	return true, nil
}
