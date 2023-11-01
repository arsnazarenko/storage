package serialization

import "fmt"

type Serialisable interface {
    Serialize() ([]byte, error)
    Deserialize(buf []byte) error
}

func SerializeUint(v uint64) ([]byte, error) {
    return nil, nil
}

func SerializeInt(v int64) ([]byte, error) {
    return nil, nil
}

func SerializeString(v string) ([]byte, error) {
    return nil, nil
}

func SerializeFloat(v float64) ([]byte, error) {
    return nil, nil
}

func SerializeStringList(l []string) ([]byte, error) {
    return nil, nil
}

func SerializeStringMap(v map[string]string) ([]byte, error) {
    return nil, nil
}

func SerializeObject(v Serialisable) ([]byte, error) {
    return nil, nil
}

func SerializeObjectList[T Serialisable](l []T) ([]byte, error) {
    return nil, nil
}

func GetType(buf []byte) (ValueType, error) {
    if buf == nil || len(buf) == 0 {
        return 0, fmt.Errorf("Buffer can be nil or empty\n")
    }
    t := uint8(buf[0]); 
    if t < uint8(UintType) || t >= uint8(MaxType) {        
        return 0, fmt.Errorf("Non-existent data type: %d\n", t)
    }
    return ValueType(t), nil
}

func DeserializeTo(src []byte, dst interface{}) error {
    if dst == nil {
        return fmt.Errorf("Destination address caanot be nil")
    }
    t, err := GetType(src)
    if err != nil {
        return err
    }    
    switch t {
    case UintType: 
    if ptr, ok := dst.(*uint64); ok {
                    
    }
    case IntType: 
    if ptr, ok := dst.(*int64); ok {
           
    }
    case FloatType: 
    if ptr, ok := dst.(*float64); ok {
                
    }

    case StringType: 
    if ptr, ok := dst.(*string); ok {
                
    }

    case StringListType: 
    if ptr, ok := dst.(*[]string); ok {
                
    }

    case StringMapType: 
    if ptr, ok := dst.(*map[string]string); ok {
                
    }
    default:
        panic("Unsupported type!")
    }
    return nil
}

