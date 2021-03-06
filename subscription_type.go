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

type SubscriptionType struct {
	// Name of the subscription
	Name string `json:"name"`
	// Description of the subscription
	Description string `json:"description"`
	// Type defining the subscription e.g.: DSP, Fuel Pass, VIP, etc. (unique value)
	Type string `json:"type"`
	// Status of the subscription type. It can be published, unpublished or draft
	Status string `json:"status"`
	// This attribute group will be added to customer who buys this subscription
	CustomerAttributeGroup string `json:"customerAttributeGroup"`
	// UPCs associated with the subscription
	Gtin13s []string `json:"gtin13s"`
	// Unique ID of subscription type
	Id string `json:"id"`
	// URL of terms and conditions text
	TermsAndConditionsUrl string `json:"termsAndConditionsUrl"`
}

const SubscriptionTypeAvroCRC64Fingerprint = "\x9b\xa9\x86\xf2\xe2\xd5\x1d\x9e"

func NewSubscriptionType() SubscriptionType {
	r := SubscriptionType{}
	r.Gtin13s = make([]string, 0)

	return r
}

func DeserializeSubscriptionType(r io.Reader) (SubscriptionType, error) {
	t := NewSubscriptionType()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeSubscriptionTypeFromSchema(r io.Reader, schema string) (SubscriptionType, error) {
	t := NewSubscriptionType()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeSubscriptionType(r SubscriptionType, w io.Writer) error {
	var err error
	err = vm.WriteString(r.Name, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Description, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Type, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Status, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.CustomerAttributeGroup, w)
	if err != nil {
		return err
	}
	err = writeArrayString(r.Gtin13s, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Id, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.TermsAndConditionsUrl, w)
	if err != nil {
		return err
	}
	return err
}

func (r SubscriptionType) Serialize(w io.Writer) error {
	return writeSubscriptionType(r, w)
}

func (r SubscriptionType) Schema() string {
	return "{\"fields\":[{\"doc\":\"Name of the subscription\",\"name\":\"name\",\"type\":{\"avro.java.string\":\"String\",\"type\":\"string\"}},{\"doc\":\"Description of the subscription\",\"name\":\"description\",\"type\":{\"avro.java.string\":\"String\",\"type\":\"string\"}},{\"doc\":\"Type defining the subscription e.g.: DSP, Fuel Pass, VIP, etc. (unique value)\",\"name\":\"type\",\"type\":{\"avro.java.string\":\"String\",\"type\":\"string\"}},{\"doc\":\"Status of the subscription type. It can be published, unpublished or draft\",\"name\":\"status\",\"type\":{\"avro.java.string\":\"String\",\"type\":\"string\"}},{\"doc\":\"This attribute group will be added to customer who buys this subscription\",\"name\":\"customerAttributeGroup\",\"type\":{\"avro.java.string\":\"String\",\"type\":\"string\"}},{\"doc\":\"UPCs associated with the subscription\",\"name\":\"gtin13s\",\"type\":{\"items\":{\"avro.java.string\":\"String\",\"type\":\"string\"},\"type\":\"array\"}},{\"doc\":\"Unique ID of subscription type\",\"name\":\"id\",\"type\":{\"avro.java.string\":\"String\",\"type\":\"string\"}},{\"doc\":\"URL of terms and conditions text\",\"name\":\"termsAndConditionsUrl\",\"type\":{\"avro.java.string\":\"String\",\"type\":\"string\"}}],\"name\":\"com.kroger.desp.commons.kcp.subscription.subscriptionscore.SubscriptionType\",\"type\":\"record\"}"
}

func (r SubscriptionType) SchemaName() string {
	return "com.kroger.desp.commons.kcp.subscription.subscriptionscore.SubscriptionType"
}

func (_ SubscriptionType) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ SubscriptionType) SetInt(v int32)       { panic("Unsupported operation") }
func (_ SubscriptionType) SetLong(v int64)      { panic("Unsupported operation") }
func (_ SubscriptionType) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ SubscriptionType) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ SubscriptionType) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ SubscriptionType) SetString(v string)   { panic("Unsupported operation") }
func (_ SubscriptionType) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *SubscriptionType) Get(i int) types.Field {
	switch i {
	case 0:
		return &types.String{Target: &r.Name}
	case 1:
		return &types.String{Target: &r.Description}
	case 2:
		return &types.String{Target: &r.Type}
	case 3:
		return &types.String{Target: &r.Status}
	case 4:
		return &types.String{Target: &r.CustomerAttributeGroup}
	case 5:
		r.Gtin13s = make([]string, 0)

		return &ArrayStringWrapper{Target: &r.Gtin13s}
	case 6:
		return &types.String{Target: &r.Id}
	case 7:
		return &types.String{Target: &r.TermsAndConditionsUrl}
	}
	panic("Unknown field index")
}

func (r *SubscriptionType) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *SubscriptionType) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ SubscriptionType) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ SubscriptionType) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ SubscriptionType) Finalize()                        {}

func (_ SubscriptionType) AvroCRC64Fingerprint() []byte {
	return []byte(SubscriptionTypeAvroCRC64Fingerprint)
}

func (r SubscriptionType) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["name"], err = json.Marshal(r.Name)
	if err != nil {
		return nil, err
	}
	output["description"], err = json.Marshal(r.Description)
	if err != nil {
		return nil, err
	}
	output["type"], err = json.Marshal(r.Type)
	if err != nil {
		return nil, err
	}
	output["status"], err = json.Marshal(r.Status)
	if err != nil {
		return nil, err
	}
	output["customerAttributeGroup"], err = json.Marshal(r.CustomerAttributeGroup)
	if err != nil {
		return nil, err
	}
	output["gtin13s"], err = json.Marshal(r.Gtin13s)
	if err != nil {
		return nil, err
	}
	output["id"], err = json.Marshal(r.Id)
	if err != nil {
		return nil, err
	}
	output["termsAndConditionsUrl"], err = json.Marshal(r.TermsAndConditionsUrl)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *SubscriptionType) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["name"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Name); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for name")
	}
	val = func() json.RawMessage {
		if v, ok := fields["description"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Description); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for description")
	}
	val = func() json.RawMessage {
		if v, ok := fields["type"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Type); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for type")
	}
	val = func() json.RawMessage {
		if v, ok := fields["status"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Status); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for status")
	}
	val = func() json.RawMessage {
		if v, ok := fields["customerAttributeGroup"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CustomerAttributeGroup); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for customerAttributeGroup")
	}
	val = func() json.RawMessage {
		if v, ok := fields["gtin13s"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Gtin13s); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for gtin13s")
	}
	val = func() json.RawMessage {
		if v, ok := fields["id"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Id); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for id")
	}
	val = func() json.RawMessage {
		if v, ok := fields["termsAndConditionsUrl"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.TermsAndConditionsUrl); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for termsAndConditionsUrl")
	}
	return nil
}
