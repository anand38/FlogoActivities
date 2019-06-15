package main

import "github.com/project-flogo/core/data/coerce"

type Settings struct {
	ASetting string `md:"aSetting"`
}

type Input struct {
	URL string `md:"URL,required"`
	ClientID string `md:"ClientID,required"`
	Topic string `md:"Topic,required"`

}
func (r *Input) FromMap(values map[string]interface{}) error {
	strURLVal, _ := coerce.ToString(values["URL"])
	strClientIDVal, _ := coerce.ToString(values["ClientID"])
	strTopicVal, _ := coerce.ToString(values["Topic"])
	r.URL = strURLVal
	r.ClientID=strClientIDVal
	r.Topic=strTopicVal
	return nil
}
/*
func (r *Input) FromMap(values map[string]interface{}) error {
	strVal, _ := coerce.ToString(values["name"])
	r.AnInput = strVal
	return nil
}
 */

func (r *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"URL": r.URL,"ClientID":r.ClientID,"Topic":r.Topic,
	}
}

type Output struct {
	AnOutput string `md:"anOutput"`
}

func (o *Output) FromMap(values map[string]interface{}) error {
	strVal, _ := coerce.ToString(values["anOutput"])
	o.AnOutput = strVal
	return nil
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"anOutput": o.AnOutput,
	}
}
