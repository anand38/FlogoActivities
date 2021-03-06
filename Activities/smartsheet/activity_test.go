package smartsheet

import (
	"io/ioutil"
	"testing"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-contrib/action/flow/test"
	"github.com/stretchr/testify/assert"
)

var activityMetadata *activity.Metadata

func getActivityMetadata() *activity.Metadata {

	if activityMetadata == nil {
		jsonMetadataBytes, err := ioutil.ReadFile("activity.json")
		if err != nil{
			panic("No Json Metadata found for activity.json path")
		}

		activityMetadata = activity.NewMetadata(string(jsonMetadataBytes))
	}

	return activityMetadata
}

func TestCreate(t *testing.T) {

	act := NewActivity(getActivityMetadata())

	if act == nil {
		t.Error("Activity Not Created")
		t.Fail()
		return
	}
}

func TestEval(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()


	//testCase 1
	//setup attrs
	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())
	//setup attrs
	tc.SetInput("Sheet_ID","681368645592964")
	tc.SetInput("Access_Token","nfg0w0cdq7mbzdnru6w2uvvh24")
	act.Eval(tc)
	result1 := tc.GetOutput("Response_Json")

	//Reading expected data from file
	file, err := ioutil.ReadFile("D:\\SheetData.txt")

	if err!=nil {
		assert.Equal(t,"",result1)
	}else{
		assert.Equal(t, string(file), result1)
	}

}
