package subscriber

import (
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/support/connection"
)

type Settings struct {
	Connection connection.Manager `md:"connection,required"`
}

type HandlerSettings struct {
	Topic            string `md:"topic,required"`
	Subscription     string `md:"subscriptionName,required"`
	SubscriptionType string `md:"subscriptionType"`
	InitialPosition  string `md:"initialPosition"`
	DLQMaxDeliveries int    `md:"dlqMaxDeliveries"`
	DLQTopic         string `md:"dlqTopic"`
}

type Output struct {
	Properties map[string]string `md:"properties"`
	Payload    interface{}       `md:"payload"`
	Topic      string            `md:"topic"`
}

func (o *Output) FromMap(values map[string]interface{}) error {
	var err error
	o.Payload, err = coerce.ToObject(values["payload"])
	if err != nil {
		return err
	}
	o.Properties, err = coerce.ToParams(values["properties"])
	if err != nil {
		return err
	}
	o.Topic, err = coerce.ToString(values["topic"])
	if err != nil {
		return err
	}
	return nil
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"payload":    o.Payload,
		"properties": o.Properties,
		"topic":      o.Topic,
	}
}
