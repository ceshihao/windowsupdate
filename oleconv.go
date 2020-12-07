package windowsupdate

import (
	"time"

	"github.com/go-ole/go-ole"
)

func toIDispatchErr(result *ole.VARIANT, err error) (*ole.IDispatch, error) {
	if err != nil {
		return nil, err
	}
	return result.ToIDispatch(), nil
}

func toStringSliceErr(result *ole.VARIANT, err error) ([]string, error) {
	if err != nil {
		return nil, err
	}
	array := result.ToArray()
	if array == nil {
		return nil, nil
	}
	return array.ToStringArray(), nil
}

func toInt64Err(result *ole.VARIANT, err error) (int64, error) {
	if err != nil {
		return 0, err
	}
	return variantToInt64(result), nil
}

func toInt32Err(result *ole.VARIANT, err error) (int32, error) {
	if err != nil {
		return 0, err
	}
	return variantToInt32(result), nil
}

func toFloat64Err(result *ole.VARIANT, err error) (float64, error) {
	if err != nil {
		return 0, err
	}
	return variantToFloat64(result), nil
}

func toFloat32Err(result *ole.VARIANT, err error) (float32, error) {
	if err != nil {
		return 0, err
	}
	return variantToFloat32(result), nil
}

func toStringErr(result *ole.VARIANT, err error) (string, error) {
	if err != nil {
		return "", err
	}
	return variantToString(result), nil
}

func toBoolErr(result *ole.VARIANT, err error) (bool, error) {
	if err != nil {
		return false, err
	}
	return variantToBool(result), nil
}

func toTimeErr(result *ole.VARIANT, err error) (*time.Time, error) {
	if err != nil {
		return nil, err
	}
	return variantToTime(result), nil
}

func variantToInt64(v *ole.VARIANT) int64 {
	return v.Value().(int64)
}

func variantToInt32(v *ole.VARIANT) int32 {
	return v.Value().(int32)
}

func variantToFloat64(v *ole.VARIANT) float64 {
	return v.Value().(float64)
}

func variantToFloat32(v *ole.VARIANT) float32 {
	return v.Value().(float32)
}

func variantToString(v *ole.VARIANT) string {
	return v.Value().(string)
}

func variantToBool(v *ole.VARIANT) bool {
	return v.Value().(bool)
}

func variantToTime(v *ole.VARIANT) *time.Time {
	value := v.Value()
	if value == nil {
		return nil
	}
	valueTime := value.(time.Time)
	return &valueTime
}
