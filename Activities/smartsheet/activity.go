package smartsheet

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"net/http"
	"fmt"
	"io/ioutil"
	"github.com/tidwall/gjson"
	"strconv"
	"github.com/tidwall/sjson"
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
					columns:=gjson.Get(string(sheetData),"columns")
					columnLength,_:=strconv.Atoi(gjson.Get(columns.String(),"#").String())
					var col=make([]string,columnLength)
					for t:=0;t<columnLength;t++{
						col[t]=gjson.Get(columns.String(),strconv.Itoa(t)+".title").String()
					}
					rows := gjson.Get(string(sheetData),"rows.#.cells")
					activityOutputTmp:=`{}`
					activityOutput:=""
					for i,_ := range rows.Array(){
						rowcell:= gjson.Get(rows.String(),strconv.Itoa(i)) ///single row
						for tmp:=0;tmp<columnLength;tmp++  {
							activityOutputTmp,_=sjson.Set(activityOutputTmp,"rows."+strconv.Itoa(i)+"."+col[tmp],gjson.Get(rowcell.String(),strconv.Itoa(tmp)+".value").String())
						}
					}
					activityOutput=activityOutputTmp
					fmt.Println(activityOutput)
				}
			}
	return true, nil
	}


