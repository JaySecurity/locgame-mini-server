package bsonutils

import (
	"reflect"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"locgame-mini-server/pkg/dto/base"
)

var (
	// Protobuf Timestamp type.
	timestampType = reflect.TypeOf(base.Timestamp{})

	// Time type.
	timeType = reflect.TypeOf(time.Time{})

	// ObjectId type.
	objectIDType          = reflect.TypeOf(base.ObjectID{})
	objectIDPrimitiveType = reflect.TypeOf(primitive.ObjectID{})

	// Eth address type.
	ethAddressType = reflect.TypeOf(common.Address{})

	// Codecs.
	timestampCodecRef  = &timestampCodec{}
	objectIDCodecRef   = &objectIDCodec{}
	ethAddressCodecRef = &ethAddressCodec{}
)

// timestampCodec is codec for Protobuf Timestamp.
type timestampCodec struct {
}

// EncodeValue encodes Protobuf Timestamp value to BSON value.
func (e *timestampCodec) EncodeValue(ectx bsoncodec.EncodeContext, vw bsonrw.ValueWriter, val reflect.Value) error {
	v := val.Addr().Interface().(*base.Timestamp)
	t := time.Unix(v.Seconds, 0)
	enc, err := ectx.LookupEncoder(timeType)
	if err != nil {
		return err
	}
	return enc.EncodeValue(ectx, vw, reflect.ValueOf(t.In(time.UTC)))
}

// DecodeValue decodes BSON value to Timestamp value.
func (e *timestampCodec) DecodeValue(ectx bsoncodec.DecodeContext, vr bsonrw.ValueReader, val reflect.Value) error {
	enc, err := ectx.LookupDecoder(timeType)
	if err != nil {
		return err
	}
	var t time.Time
	if err = enc.DecodeValue(ectx, vr, reflect.ValueOf(&t).Elem()); err != nil {
		return err
	}
	val.Set(reflect.ValueOf(base.Timestamp{Seconds: t.UTC().Unix()}))
	return nil
}

// objectIDCodec is codec for Protobuf ObjectId.
type objectIDCodec struct {
}

// EncodeValue encodes Protobuf ObjectId value to BSON value.
func (e *objectIDCodec) EncodeValue(ectx bsoncodec.EncodeContext, vw bsonrw.ValueWriter, val reflect.Value) error {
	v := val.Addr().Interface().(*base.ObjectID)
	// Create primitive.ObjectId from string
	id, err := primitive.ObjectIDFromHex(v.Value)
	if err != nil {
		return err
	}
	enc, err := ectx.LookupEncoder(objectIDPrimitiveType)
	if err != nil {
		return err
	}
	return enc.EncodeValue(ectx, vw, reflect.ValueOf(id))
}

// DecodeValue decodes BSON value to ObjectId value.
func (e *objectIDCodec) DecodeValue(ectx bsoncodec.DecodeContext, vr bsonrw.ValueReader, val reflect.Value) error {
	enc, err := ectx.LookupDecoder(objectIDPrimitiveType)
	if err != nil {
		return err
	}
	var id primitive.ObjectID
	if err = enc.DecodeValue(ectx, vr, reflect.ValueOf(&id).Elem()); err != nil {
		return err
	}
	val.Set(reflect.ValueOf(base.ObjectID{Value: id.Hex()}))
	return nil
}

// ethAddressCodec is codec for Ethereum address.
type ethAddressCodec struct{}

// EncodeValue encodes Protobuf ObjectId value to BSON value.
func (e *ethAddressCodec) EncodeValue(_ bsoncodec.EncodeContext, vw bsonrw.ValueWriter, val reflect.Value) error {
	v := val.Interface().(common.Address)
	return vw.WriteString(strings.ToLower(v.String()))
}

// DecodeValue decodes BSON value to ObjectId value.
func (e *ethAddressCodec) DecodeValue(_ bsoncodec.DecodeContext, vr bsonrw.ValueReader, val reflect.Value) error {
	value, err := vr.ReadString()
	if err != nil {
		return err
	}
	val.Set(reflect.ValueOf(common.HexToAddress(value)))
	return nil
}

// Register registers Google protocol buffers types codecs.
func Register(rb *bsoncodec.RegistryBuilder) *bsoncodec.RegistryBuilder {
	return rb.
		RegisterCodec(timestampType, timestampCodecRef).
		RegisterCodec(objectIDType, objectIDCodecRef).
		RegisterCodec(ethAddressType, ethAddressCodecRef)
}
