package serialization

type ValueType uint8

const (
    UintType ValueType = iota + 1
    IntType
    FloatType
    StringType
    StringListType
    StringMapType
    ObjectType
    MaxType /* Max type value */
)

