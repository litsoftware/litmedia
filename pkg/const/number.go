package _const

const (
	_max = ^uint(0)
	_log = _max>>8&1 + _max>>16&1 + _max>>32&1 + _max>>64&1

	UintBytes = 1 << _log
	UintBits  = IntBytes * 8
	IntBytes  = UintBytes
	IntBits   = UintBits
	MaxUint   = (1 << IntBits) - 1
	MinUint   = 0
	MaxInt    = MaxUint >> 1
	MinInt    = -MaxInt - 1
)

const (
	Uint8Bytes = 1
	Uint8Bits  = Uint8Bytes * 8
	Int8Bytes  = Uint8Bytes
	Int8Bits   = Uint8Bits
	MaxUint8   = (1 << Uint8Bits) - 1
	MinUint8   = 0
	MaxInt8    = MaxUint8 >> 1
	MinInt8    = -MaxInt8 - 1
)

const (
	Uint16Bytes = 2
	Uint16Bits  = Uint16Bytes * 8
	Int16Bytes  = Uint16Bytes
	Int16Bits   = Uint16Bits
	MaxUint16   = (1 << Uint16Bits) - 1
	MinUint16   = 0
	MaxInt16    = MaxUint16 >> 1
	MinInt16    = -MaxInt16 - 1
)

const (
	Uint32Bytes = 4
	Uint32Bits  = Uint32Bytes * 8
	Int32Bytes  = Uint32Bytes
	Int32Bits   = Uint32Bits
	MaxUint32   = (1 << Uint32Bits) - 1
	MinUint32   = 0
	MaxInt32    = MaxUint32 >> 1
	MinInt32    = -MaxInt32 - 1
)

const (
	Uint64Bytes = 8
	Uint64Bits  = Uint64Bytes * 8
	Int64Bytes  = Uint64Bytes
	Int64Bits   = Uint64Bits
	MaxUint64   = (1 << Uint64Bits) - 1
	MinUint64   = 0
	MaxInt64    = MaxUint64 >> 1
	MinInt64    = -MaxInt64 - 1
)
