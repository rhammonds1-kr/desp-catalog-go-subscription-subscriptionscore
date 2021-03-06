// Code generated by github.com/actgardner/gogen-avro/v8. DO NOT EDIT.
/*
 * SOURCES:
 *     DateTime.avsc
 *     Enrollment.avsc
 *     EnrollmentCancelledEvent.avsc
 *     EnrollmentCreatedEvent.avsc
 *     EnrollmentExpirationEvent.avsc
 *     EnrollmentExpiredEvent.avsc
 *     EnrollmentRenewalEvent.avsc
 *     EventHeader.avsc
 *     SubscriptionType.avsc
 *     SubscriptionTypeCreatedEvent.avsc
 *     SubscriptionTypeUpdatedEvent.avsc
 */
package subscription_subscriptionscore

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v9/compiler"
	"github.com/actgardner/gogen-avro/v9/vm"
	"github.com/actgardner/gogen-avro/v9/vm/types"
)

var _ = fmt.Printf

type SubscriptionTypeCreatedEvent struct {
	// Header field for this event
	EventHeader EventHeader `json:"eventHeader"`
	// Subscription entity created
	Subscription SubscriptionType `json:"subscription"`
}

const SubscriptionTypeCreatedEventAvroCRC64Fingerprint = "\xef\a\xa0\x97!s\xa5\x02"

func NewSubscriptionTypeCreatedEvent() SubscriptionTypeCreatedEvent {
	r := SubscriptionTypeCreatedEvent{}
	r.EventHeader = NewEventHeader()

	r.Subscription = NewSubscriptionType()

	return r
}

func DeserializeSubscriptionTypeCreatedEvent(r io.Reader) (SubscriptionTypeCreatedEvent, error) {
	t := NewSubscriptionTypeCreatedEvent()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeSubscriptionTypeCreatedEventFromSchema(r io.Reader, schema string) (SubscriptionTypeCreatedEvent, error) {
	t := NewSubscriptionTypeCreatedEvent()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeSubscriptionTypeCreatedEvent(r SubscriptionTypeCreatedEvent, w io.Writer) error {
	var err error
	err = writeEventHeader(r.EventHeader, w)
	if err != nil {
		return err
	}
	err = writeSubscriptionType(r.Subscription, w)
	if err != nil {
		return err
	}
	return err
}

func (r SubscriptionTypeCreatedEvent) Serialize(w io.Writer) error {
	return writeSubscriptionTypeCreatedEvent(r, w)
}

func (r SubscriptionTypeCreatedEvent) Schema() string {
	return "{\"fields\":[{\"doc\":\"Header field for this event\",\"name\":\"eventHeader\",\"type\":{\"doc\":\"The below fields include header information and should be included on every event in the DESP. Inspired by: https://github.com/cloudevents/spec/blob/v0.2/spec.md\",\"fields\":[{\"doc\":\"A unique identifier of the event - for example, a randomly generated GUID\",\"name\":\"id\",\"type\":{\"avro.java.string\":\"String\",\"type\":\"string\"}},{\"doc\":\"Time the event occurred in milliseconds since epoch, UTC timezone.\",\"name\":\"time\",\"type\":\"long\"},{\"doc\":\"Type of occurrence which has happened. Reference the domain.event registered in schema-registry.\",\"name\":\"type\",\"type\":{\"avro.java.string\":\"String\",\"type\":\"string\"}},{\"doc\":\"Service that produced the event. Future: reference to producer registry.\",\"name\":\"source\",\"type\":{\"avro.java.string\":\"String\",\"type\":\"string\"}}],\"name\":\"EventHeader\",\"namespace\":\"com.kroger.desp.commons.kcp.subscription.subscriptionscore\",\"type\":\"record\"}},{\"doc\":\"Subscription entity created\",\"name\":\"subscription\",\"type\":{\"fields\":[{\"doc\":\"Name of the subscription\",\"name\":\"name\",\"type\":{\"avro.java.string\":\"String\",\"type\":\"string\"}},{\"doc\":\"Description of the subscription\",\"name\":\"description\",\"type\":{\"avro.java.string\":\"String\",\"type\":\"string\"}},{\"doc\":\"Type defining the subscription e.g.: DSP, Fuel Pass, VIP, etc. (unique value)\",\"name\":\"type\",\"type\":{\"avro.java.string\":\"String\",\"type\":\"string\"}},{\"doc\":\"Status of the subscription type. It can be published, unpublished or draft\",\"name\":\"status\",\"type\":{\"avro.java.string\":\"String\",\"type\":\"string\"}},{\"doc\":\"This attribute group will be added to customer who buys this subscription\",\"name\":\"customerAttributeGroup\",\"type\":{\"avro.java.string\":\"String\",\"type\":\"string\"}},{\"doc\":\"UPCs associated with the subscription\",\"name\":\"gtin13s\",\"type\":{\"items\":{\"avro.java.string\":\"String\",\"type\":\"string\"},\"type\":\"array\"}},{\"doc\":\"Unique ID of subscription type\",\"name\":\"id\",\"type\":{\"avro.java.string\":\"String\",\"type\":\"string\"}},{\"doc\":\"URL of terms and conditions text\",\"name\":\"termsAndConditionsUrl\",\"type\":{\"avro.java.string\":\"String\",\"type\":\"string\"}}],\"name\":\"SubscriptionType\",\"namespace\":\"com.kroger.desp.commons.kcp.subscription.subscriptionscore\",\"type\":\"record\"}}],\"name\":\"com.kroger.desp.events.kcp.subscription.subscriptionscore.SubscriptionTypeCreatedEvent\",\"type\":\"record\"}"
}

func (r SubscriptionTypeCreatedEvent) SchemaName() string {
	return "com.kroger.desp.events.kcp.subscription.subscriptionscore.SubscriptionTypeCreatedEvent"
}

func (_ SubscriptionTypeCreatedEvent) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ SubscriptionTypeCreatedEvent) SetInt(v int32)       { panic("Unsupported operation") }
func (_ SubscriptionTypeCreatedEvent) SetLong(v int64)      { panic("Unsupported operation") }
func (_ SubscriptionTypeCreatedEvent) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ SubscriptionTypeCreatedEvent) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ SubscriptionTypeCreatedEvent) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ SubscriptionTypeCreatedEvent) SetString(v string)   { panic("Unsupported operation") }
func (_ SubscriptionTypeCreatedEvent) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *SubscriptionTypeCreatedEvent) Get(i int) types.Field {
	switch i {
	case 0:
		r.EventHeader = NewEventHeader()

		return &types.Record{Target: &r.EventHeader}
	case 1:
		r.Subscription = NewSubscriptionType()

		return &types.Record{Target: &r.Subscription}
	}
	panic("Unknown field index")
}

func (r *SubscriptionTypeCreatedEvent) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *SubscriptionTypeCreatedEvent) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ SubscriptionTypeCreatedEvent) AppendMap(key string) types.Field {
	panic("Unsupported operation")
}
func (_ SubscriptionTypeCreatedEvent) AppendArray() types.Field { panic("Unsupported operation") }
func (_ SubscriptionTypeCreatedEvent) Finalize()                {}

func (_ SubscriptionTypeCreatedEvent) AvroCRC64Fingerprint() []byte {
	return []byte(SubscriptionTypeCreatedEventAvroCRC64Fingerprint)
}

func (r SubscriptionTypeCreatedEvent) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["eventHeader"], err = json.Marshal(r.EventHeader)
	if err != nil {
		return nil, err
	}
	output["subscription"], err = json.Marshal(r.Subscription)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *SubscriptionTypeCreatedEvent) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["eventHeader"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.EventHeader); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for eventHeader")
	}
	val = func() json.RawMessage {
		if v, ok := fields["subscription"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Subscription); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for subscription")
	}
	return nil
}
