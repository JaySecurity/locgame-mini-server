package bsonutils

import (
	"fmt"
	"reflect"
	"strings"
	"sync"
	"time"

	"locgame-mini-server/pkg/dto/base"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type tagOptions map[string]struct{}

func (t tagOptions) Has(opt string) bool {
	if _, ok := t[opt]; ok {
		return true
	}
	return false
}

type tagData struct {
	tagName string
	tagOpts tagOptions
}

var parseCache map[string]*tagData

var tagMutex *sync.RWMutex
var fieldsMutex *sync.RWMutex

func parseTag(tag string) (string, tagOptions) {
	if parseCache == nil {
		parseCache = make(map[string]*tagData)
		tagMutex = &sync.RWMutex{}
	}

	tagMutex.RLock()
	data := parseCache[tag]
	tagMutex.RUnlock()
	if data == nil {
		res := strings.Split(tag, ",")
		m := make(tagOptions)
		for i, opt := range res {
			if i == 0 {
				continue
			}
			m[opt] = struct{}{}
		}
		data = &tagData{tagName: res[0], tagOpts: m}
		tagMutex.Lock()
		parseCache[tag] = data
		tagMutex.Unlock()
	}

	return data.tagName, data.tagOpts
}

const bsonKey = "bson"

type BSONOptions struct {
	proj             interface{}
	parentStructName string
	ignoreSubStruct  bool
}

func (o *BSONOptions) SetProjection(proj interface{}) {
	o.proj = proj
}

func (o *BSONOptions) SetParentStructName(name string) {
	o.parentStructName = name
}

func (o *BSONOptions) SetIgnoreSubStructs(value bool) {
	o.ignoreSubStruct = value
}

func ToBSONMap(v interface{}, options ...*BSONOptions) bson.M {
	out := bson.M{}

	value, valueType := getValueAndType(v)
	fields := getFields(valueType)

	for _, field := range fields {
		name := field.Name
		isSubStruct := false
		// isSlice := false
		val := value.FieldByName(name)

		if strings.EqualFold(name, "online") {
			continue
		}
		var finalVal interface{}

		// Identify whether the struct field has tags or not
		tagName, tagOpts := parseTag(field.Tag.Get(bsonKey))
		if tagName != "" {
			name = tagName
		}

		//This code sets boolean value to false so this is a patch to prevent it to do with online field
		if tagOpts.Has("online") {
			continue
		}

		var projSubValue interface{}

		// Decide whether to omit the field if it is empty or not
		if tagOpts.Has("omitempty") {
			if len(options) > 0 && options[0].proj != nil {
				projValue := reflect.ValueOf(options[0].proj)

				for projValue.Kind() == reflect.Ptr {
					projValue = projValue.Elem()
				}

				if projValue.Kind() != reflect.Slice {
					projValue = projValue.FieldByName(field.Name)
				}

				if projValue.IsValid() && !projValue.IsZero() {
					projSubValue = projValue.Interface()
				}
			}

			if val.IsZero() && projSubValue == nil {
				continue
			}

			// Handling edge cases that reflect.value.IsZero doesn't catch
			switch val.Kind() {
			case reflect.Slice:
				if val.Len() == 0 {
					continue
				}
			case reflect.Map:
				if len(val.MapKeys()) == 0 {
					continue
				}
			}
		}

		// Handle omitnil tag
		if tagOpts.Has("omitnil") {
			if val.Kind() == reflect.Ptr && val.IsNil() {
				continue
			}

			if val.Kind() == reflect.Slice && val.IsNil() {
				if val.Len() == 0 {
					continue
				}
			}
		}

		var fullName string
		if len(options) > 0 && options[0].parentStructName != "" {
			fullName += name
		} else {
			fullName = name
			if len(options) == 0 {
				options = append(options, &BSONOptions{})
			}
		}

		options[0].parentStructName = fullName

		// If nested data structures should not be omitted
		finalVal = nestedData(val, options...)

		v := reflect.ValueOf(val.Interface())
		if v.Kind() == reflect.Ptr {
			v = v.Elem()
		}

		switch v.Kind() {
		case reflect.Map, reflect.Struct:
			if v.Type() == reflect.TypeOf(base.ObjectID{}) || v.Type() == reflect.TypeOf(base.Timestamp{}) || len(options) > 0 && options[0].ignoreSubStruct {
				break
			}
			isSubStruct = true
		}

		if isSubStruct {
			outMap, ok := finalVal.(primitive.M)
			if ok {
				for k := range outMap {
					out[fullName+"."+k] = outMap[k]
				}
			}
		} else {
			out[fullName] = finalVal
		}
	}
	if len(out) == 0 {
		return nil
	}
	return out
}

func getValueAndType(v interface{}) (reflect.Value, reflect.Type) {
	value := reflect.ValueOf(v)

	for value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	valueType := value.Type()
	return value, valueType
}

var fieldsCache map[reflect.Type][]reflect.StructField

func getFields(valueType reflect.Type) []reflect.StructField {
	if fieldsCache == nil {
		fieldsCache = make(map[reflect.Type][]reflect.StructField)
		fieldsMutex = &sync.RWMutex{}
	}
	fieldsMutex.RLock()
	fields := fieldsCache[valueType]
	fieldsMutex.RUnlock()

	if fields == nil {
		fields = make([]reflect.StructField, 0)

		for i := 0; i < valueType.NumField(); i++ {
			field := valueType.Field(i)

			// Can't access the value of unexported fields
			if field.PkgPath != "" {
				continue
			}

			fields = append(fields, field)
		}
		fieldsMutex.Lock()
		fieldsCache[valueType] = fields
		fieldsMutex.Unlock()
	}

	return fields
}

func nestedData(val reflect.Value, options ...*BSONOptions) interface{} {
	var finalVal interface{}
	v := reflect.ValueOf(val.Interface())

	// Converting a pointer to a value
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	switch v.Kind() {
	case reflect.Struct:
		if v.Type() == reflect.TypeOf(base.ObjectID{}) {
			id := v.FieldByName("Value").String()
			if id == "" {
				finalVal = nil
			} else {
				finalVal, _ = primitive.ObjectIDFromHex(id)
			}
		} else if v.Type() == reflect.TypeOf(base.Timestamp{}) {
			finalVal = time.Unix(v.FieldByName("Seconds").Int(), 0)
		} else {
			m := ToBSONMap(val.Interface(), options...)

			if len(m) == 0 {
				finalVal = val.Interface()
			} else {
				finalVal = m
			}
		}

	case reflect.Map:
		// Find the type of the value within the map
		mapElem := val.Type()
		switch mapElem.Kind() {
		case reflect.Ptr, reflect.Array, reflect.Map, reflect.Slice, reflect.Chan:
			mapElem = mapElem.Elem()
			if mapElem.Kind() == reflect.Ptr {
				mapElem = mapElem.Elem()
			}
		}

		// If we need to iterate over some form of struct in the map
		// ie. map[string]struct
		if mapElem.Kind() == reflect.Struct || (mapElem.Kind() == reflect.Slice && mapElem.Elem().Kind() == reflect.Struct) {
			m := bson.M{}
			for _, k := range val.MapKeys() {
				opts := new(BSONOptions)
				opts.SetIgnoreSubStructs(true)
				m[fmt.Sprint(k)] = nestedData(val.MapIndex(k), opts)
			}
			finalVal = m
			break
		}

		if mapElem.Kind() == reflect.Int32 || mapElem.Kind() == reflect.Int64 {
			m := bson.M{}
			for _, k := range val.MapKeys() {
				m[fmt.Sprint(k)] = val.MapIndex(k).Int()
			}
			finalVal = m
			break
		}

		if mapElem.Kind() == reflect.Bool {
			m := bson.M{}
			for _, k := range val.MapKeys() {
				m[fmt.Sprint(k)] = val.MapIndex(k).Bool()
			}
			finalVal = m
			break
		}

		finalVal = val.Interface()

	case reflect.Slice, reflect.Array:
		if val.Type().Kind() == reflect.Ptr {
			val = val.Elem()
		}

		// Ensuring there are no structs (which require further iteration) anywhere within the slice/array
		// As long as there are not, we just pass the value of the array/slice
		if val.Type().Elem().Kind() != reflect.Struct && !(val.Type().Elem().Kind() == reflect.Ptr && val.Type().Elem().Elem().Kind() == reflect.Struct) {
			finalVal = val.Interface()
			break
		}

		// If further iteration is needed, then iterate over the slice
		slices := make([]interface{}, val.Len())
		for x := 0; x < val.Len(); x++ {
			opts := new(BSONOptions)
			opts.SetIgnoreSubStructs(true)
			slices[x] = ToBSONMap(val.Index(x).Interface(), opts)
		}
		finalVal = slices

	default:
		finalVal = val.Interface()
	}

	return finalVal
}
