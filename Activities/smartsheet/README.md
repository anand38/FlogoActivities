
# SmartSheet GetSheetDetails

This activity allows you to get the SmartSheet Data for your sheet in a json format. It takes SheetID and access token as an input and gives all the data for that sheet in a compact json format, which can 

## Installation

### Flogo CLI

```
flogo install github.com/anand38/FlogoActivities/Activities/smartsheet
```

### Third-party libraries used
- #### GJSON :
GJSON is a Go package that provides a fast and simple way to get values from a json document. It has features such as one line retrieval, dot notation paths, iteration, and parsing json lines.
- #### SJSON :
SJSON is a Go package that provides a very fast and simple way to set a value in a json document. The purpose for this library is to provide efficient json updating in the SurveyMonkey_GetResponse activity.

### Schema

```
{
 "inputs":[
    {
      "name": "Sheet_ID",
      "type": "string"
    },
    {
      "name":"Access_Token",
      "type":"string"
    }
  ],
  "outputs": [
    {
      "name": "Response_Json",
      "type": "string"
    }
  ]
}
```

### Activity Input


| Name | Required | Type | Description |
| ---- | -------- | ---- |------------ |
| Access_Token | True | String | Access Token of your surveymonkey App |
| Sheet_ID  | True | String | ID of the Sheet |


### Activity Output


| Name | Type | Description |
| ---- | ---- | ----------- |
| Response_Json | String | Sheet data in json format |

### Example :
This activity will give the response in a following way,

```
{
	"rows": [{
		"Comments": "",
		"Status": "Complete",
		"% Complete": "1",
		"Assigned To": "",
		"Predecessors": "",
		"Finish": "2018-04-04T16:59:59",
		"Start": "2018-04-04T08:00:00",
		"Duration": "1d",
		"Task Name": "Overview of the Project"
	},
	{
		"Comments": "",
		"Status": "Complete",
		"% Complete": "1",
		"Assigned To": "",
		"Predecessors": "",
		"Finish": "2018-04-06T16:59:59",
		"Start": "2018-04-05T08:00:00",
		"Duration": "2d",
		"Task Name": "Discussion of Business Requirement"
	},
	{
		"Comments": "",
		"Status": "Complete",
		"% Complete": "1",
		"Assigned To": "",
		"Predecessors": "",
		"Finish": "2018-04-10T16:59:59",
		"Start": "2018-04-09T08:00:00",
		"Duration": "2d",
		"Task Name": "Walkthrough of Model"
	},
	{
		"Comments": "",
		"Status": "Complete",
		"% Complete": "1",
		"Assigned To": "",
		"Predecessors": "",
		"Finish": "2018-04-17T16:59:59",
		"Start": "2018-04-11T08:00:00",
		"Duration": "5d",
		"Task Name": "Functionality discussion"
	},
	{
		"Comments": "",
		"Status": "Complete",
		"% Complete": "1",
		"Assigned To": "",
		"Predecessors": "",
		"Finish": "2018-05-21T16:59:59",
		"Start": "2018-05-18T08:00:00",
		"Duration": "2d",
		"Task Name": "Reports generation Module Discussion"
	}]
}
```
