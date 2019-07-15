package cunits

// Byte represent a byte
const Byte = 8

// Decimal prefix of bits
const (
	// Kbit represents a kilobit
	Kbit = 1000
	// Mbit represents a megabit
	Mbit = 1000000
	// Gbit represents a gigabit
	Gbit = 1000000000
	// Tbit represents a terabit
	Tbit = 1000000000000
	// Pbit represents a petabit
	Pbit = 1000000000000000
	// Ebit represents an exabit
	Ebit = 1000000000000000000
	// Zbit represents a zettabit (overflows int64)
	Zbit = 1000000000000000000000
	// Ybit represents a yottabit (overflows int64)
	Ybit = 1000000000000000000000000
)

// Binary prefix of bits
const (
	// Kibit represents a kibibit
	Kibit = 1 << 10
	// Mibit represents a mebibit
	Mibit = 1 << 20
	// Gibit represents a gibibit
	Gibit = 1 << 30
	// Tibit represents a tebibit
	Tibit = 1 << 40
	// Pibit represents a pebibit
	Pibit = 1 << 50
	// Eibit represents an exbibit
	Eibit = 1 << 60
	// Zibit represents a zebibit (overflows int64)
	Zibit = 1 << 70
	// Yibit represents a yobibit (overflows int64)
	Yibit = 1 << 80
)

// Decimal prefix of bytes
const (
	// KB represents a kilobyte
	KB = Kbit * Byte
	// MB represents a megabyte
	MB = Mbit * Byte
	// GB represents a gigabyte
	GB = Gbit * Byte
	// TB represents a terabyte
	TB = Tbit * Byte
	// PB represents a petabyte
	PB = Pbit * Byte
	// EB represents an exabyte
	EB = Ebit * Byte
	// ZB represents a zettabyte (overflows int64)
	ZB = Zbit * Byte
	// YB represents a yottabyte (overflows int64)
	YB = Ybit * Byte
)

// Binary prefix of bytes
const (
	// KiB represents a kibibyte
	KiB = Kibit * Byte
	// MiB represents a mebibyte
	MiB = Mibit * Byte
	// GiB represents a gibibyte
	GiB = Gibit * Byte
	// TiB represents a tebibyte
	TiB = Tibit * Byte
	// PiB represents a pebibyte
	PiB = Pibit * Byte
	// EiB represents an exbibyte (overflows int64)
	EiB = Eibit * Byte
	// ZiB represents a zebibyte (overflows int64)
	ZiB = Zbit * Byte
	// YiB represents a yobibyte (overflows int64)
	YiB = Ybit * Byte
)
