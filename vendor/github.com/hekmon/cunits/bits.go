package cunits

import "fmt"

// Bits represent a size in bits
type Bits uint64

// String allows direct reprensetation of Bit by calling GetHumanSizeRepresentation()
func (size Bits) String() string {
	return size.GetHumanSizeRepresentation()
}

// GetHumanSizeRepresentation returns the size in a human readable binary prefix of bytes format
func (size Bits) GetHumanSizeRepresentation() string {
	// if size >= YiB {
	// 	return size.YiBString()
	// }
	// if size >= ZiB {
	// 	return size.ZiBString()
	// }
	if size >= EiB {
		return size.EiBString()
	}
	if size >= PiB {
		return size.PiBString()
	}
	if size >= TiB {
		return size.TiBString()
	}
	if size >= GiB {
		return size.GiBString()
	}
	if size >= MiB {
		return size.MiBString()
	}
	if size >= KiB {
		return size.KiBString()
	}
	return size.ByteString()
}

// GetHumanSpeedRepresentation returns the size in a human readable decimal prefix of bits format
func (size Bits) GetHumanSpeedRepresentation() string {
	// if size >= Ybit {
	// 	return size.YbitString()
	// }
	// if size >= Zbit {
	// 	return size.ZbitString()
	// }
	if size >= Ebit {
		return size.EbitString()
	}
	if size >= Pbit {
		return size.PbitString()
	}
	if size >= Tbit {
		return size.TbitString()
	}
	if size >= Gbit {
		return size.GbitString()
	}
	if size >= Mbit {
		return size.MbitString()
	}
	if size >= Kbit {
		return size.KbitString()
	}
	return size.BitString()
}

/*
	Base forms
*/

// BitString returns the size in bit with unit suffix
func (size Bits) BitString() string {
	return fmt.Sprintf("%d b", size)
}

// Byte returns the size in byte
func (size Bits) Byte() float64 {
	return float64(size) / 8
}

// ByteString returns the size in byte with unit suffix
func (size Bits) ByteString() string {
	return fmt.Sprintf("%.2f B", size.Byte())
}

/*
	Decimal prefix of Bit
*/

// Kbit returns the size in kilobit
func (size Bits) Kbit() float64 {
	return float64(size) / Kbit
}

// KbitString returns the size in kilobit with unit suffix
func (size Bits) KbitString() string {
	return fmt.Sprintf("%.2f Kb", size.Kbit())
}

// Mbit returns the size in megabit
func (size Bits) Mbit() float64 {
	return float64(size) / Mbit
}

// MbitString returns the size in megabit with unit suffix
func (size Bits) MbitString() string {
	return fmt.Sprintf("%.2f Mb", size.Mbit())
}

// Gbit returns the size in gigabit
func (size Bits) Gbit() float64 {
	return float64(size) / Gbit
}

// GbitString returns the size in gigabit with unit suffix
func (size Bits) GbitString() string {
	return fmt.Sprintf("%.2f Gb", size.Gbit())
}

// Tbit returns the size in terabit
func (size Bits) Tbit() float64 {
	return float64(size) / Tbit
}

// TbitString returns the size in terabit with unit suffix
func (size Bits) TbitString() string {
	return fmt.Sprintf("%.2f Tb", size.Tbit())
}

// Pbit returns the size in petabit
func (size Bits) Pbit() float64 {
	return float64(size) / Pbit
}

// PbitString returns the size in petabit with unit suffix
func (size Bits) PbitString() string {
	return fmt.Sprintf("%.2f Pb", size.Pbit())
}

// Ebit returns the size in exabit
func (size Bits) Ebit() float64 {
	return float64(size) / Ebit
}

// EbitString returns the size in exabit with unit suffix
func (size Bits) EbitString() string {
	return fmt.Sprintf("%.2f Eb", size.Ebit())
}

// Zbit returns the size in zettabit
func (size Bits) Zbit() float64 {
	return float64(size) / Zbit
}

// ZbitString returns the size in zettabit with unit suffix (carefull with sub zeros !)
func (size Bits) ZbitString() string {
	return fmt.Sprintf("%f Zb", size.Zbit())
}

// Ybit returns the size in yottabit
func (size Bits) Ybit() float64 {
	return float64(size) / Ybit
}

// YbitString returns the size in yottabit with unit suffix (carefull with sub zeros !)
func (size Bits) YbitString() string {
	return fmt.Sprintf("%f Yb", size.Ybit())
}

/*
	Binary prefix of Bit
*/

// Kibit returns the size in kibibit
func (size Bits) Kibit() float64 {
	return float64(size) / Kibit
}

// KibitString returns the size in kibibit with unit suffix
func (size Bits) KibitString() string {
	return fmt.Sprintf("%.2f Kib", size.Kibit())
}

// Mibit returns the size in mebibit
func (size Bits) Mibit() float64 {
	return float64(size) / Mibit
}

// MibitString returns the size in mebibit with unit suffix
func (size Bits) MibitString() string {
	return fmt.Sprintf("%.2f Mib", size.Mibit())
}

// Gibit returns the size in gibibit
func (size Bits) Gibit() float64 {
	return float64(size) / Gibit
}

// GibitString returns the size in gibibit with unit suffix
func (size Bits) GibitString() string {
	return fmt.Sprintf("%.2f Gib", size.Gibit())
}

// Tibit returns the size in tebibit
func (size Bits) Tibit() float64 {
	return float64(size) / Tibit
}

// TibitString returns the size in tebibit with unit suffix
func (size Bits) TibitString() string {
	return fmt.Sprintf("%.2f Tib", size.Tibit())
}

// Pibit returns the size in pebibit
func (size Bits) Pibit() float64 {
	return float64(size) / Pibit
}

// PibitString returns the size in pebibit with unit suffix
func (size Bits) PibitString() string {
	return fmt.Sprintf("%.2f Pib", size.Pibit())
}

// Eibit returns the size in exbibit
func (size Bits) Eibit() float64 {
	return float64(size) / Eibit
}

// EibitString returns the size in exbibit with unit suffix
func (size Bits) EibitString() string {
	return fmt.Sprintf("%.2f Eib", size.Eibit())
}

// Zibit returns the size in zebibit
func (size Bits) Zibit() float64 {
	return float64(size) / Zibit
}

// ZibitString returns the size in zebibit with unit suffix (carefull with sub zeros !)
func (size Bits) ZibitString() string {
	return fmt.Sprintf("%f Zib", size.Zibit())
}

// Yibit returns the size in yobibit
func (size Bits) Yibit() float64 {
	return float64(size) / Yibit
}

// YibitString returns the size in yobibit with unit suffix (carefull with sub zeros !)
func (size Bits) YibitString() string {
	return fmt.Sprintf("%f Yib", size.Yibit())
}

/*
	Decimal prefix of bytes
*/

// KB returns the size in kilobyte
func (size Bits) KB() float64 {
	return float64(size) / KB
}

// KBString returns the size in kilobyte with unit suffix
func (size Bits) KBString() string {
	return fmt.Sprintf("%.2f KB", size.KB())
}

// MB returns the size in megabyte
func (size Bits) MB() float64 {
	return float64(size) / MB
}

// MBString returns the size in megabyte with unit suffix
func (size Bits) MBString() string {
	return fmt.Sprintf("%.2f MB", size.MB())
}

// GB returns the size in gigabyte
func (size Bits) GB() float64 {
	return float64(size) / GB
}

// GBString returns the size in gigabyte with unit suffix
func (size Bits) GBString() string {
	return fmt.Sprintf("%.2f GB", size.GB())
}

// TB returns the size in terabyte
func (size Bits) TB() float64 {
	return float64(size) / TB
}

// TBString returns the size in terabyte with unit suffix
func (size Bits) TBString() string {
	return fmt.Sprintf("%.2f TB", size.TB())
}

// PB returns the size in petabyte
func (size Bits) PB() float64 {
	return float64(size) / PB
}

// PBString returns the size in petabyte with unit suffix
func (size Bits) PBString() string {
	return fmt.Sprintf("%.2f PB", size.PB())
}

// EB returns the size in exabyte
func (size Bits) EB() float64 {
	return float64(size) / EB
}

// EBString returns the size in exabyte with unit suffix
func (size Bits) EBString() string {
	return fmt.Sprintf("%.2f EB", size.EB())
}

// ZB returns the size in zettabyte
func (size Bits) ZB() float64 {
	return float64(size) / ZB
}

// ZBString returns the size in zettabyte with unit suffix (carefull with sub zeros !)
func (size Bits) ZBString() string {
	return fmt.Sprintf("%f ZB", size.ZB())
}

// YB returns the size in yottabyte
func (size Bits) YB() float64 {
	return float64(size) / YB
}

// YBString returns the size in yottabyte with unit suffix (carefull with sub zeros !)
func (size Bits) YBString() string {
	return fmt.Sprintf("%f YB", size.YB())
}

/*
	Binary prefix of bytes
*/

// KiB returns the size in kibibyte
func (size Bits) KiB() float64 {
	return float64(size) / KiB
}

// KiBString returns the size in kibibyte with unit suffix
func (size Bits) KiBString() string {
	return fmt.Sprintf("%.2f KiB", size.KiB())
}

// MiB returns the size in mebibyte
func (size Bits) MiB() float64 {
	return float64(size) / MiB
}

// MiBString returns the size in mebibyte with unit suffix
func (size Bits) MiBString() string {
	return fmt.Sprintf("%.2f MiB", size.MiB())
}

// GiB returns the size in gibibyte
func (size Bits) GiB() float64 {
	return float64(size) / GiB
}

// GiBString returns the size in gibibyte with unit suffix
func (size Bits) GiBString() string {
	return fmt.Sprintf("%.2f GiB", size.GiB())
}

// TiB returns the size in tebibyte
func (size Bits) TiB() float64 {
	return float64(size) / TiB
}

// TiBString returns the size in tebibyte wit unit suffix
func (size Bits) TiBString() string {
	return fmt.Sprintf("%.2f TiB", size.TiB())
}

// PiB returns the size in pebibyte
func (size Bits) PiB() float64 {
	return float64(size) / PiB
}

// PiBString returns the size in pebibyte with unit suffix
func (size Bits) PiBString() string {
	return fmt.Sprintf("%.2f PiB", size.PiB())
}

// EiB returns the size in exbibyte
func (size Bits) EiB() float64 {
	return float64(size) / EiB
}

// EiBString returns the size in exbibyte with unit suffix (carefull with sub zeros !)
func (size Bits) EiBString() string {
	return fmt.Sprintf("%f EiB", size.EiB())
}

// ZiB returns the size in zebibyte
func (size Bits) ZiB() float64 {
	return float64(size) / ZiB
}

// ZiBString returns the size in zebibyte with unit suffix (carefull with sub zeros !)
func (size Bits) ZiBString() string {
	return fmt.Sprintf("%f ZiB", size.ZiB())
}

// YiB returns the size in yobibyte
func (size Bits) YiB() float64 {
	return float64(size) / YiB
}

// YiBString returns the size in yobibyte with unit suffix (carefull with sub zeros !)
func (size Bits) YiBString() string {
	return fmt.Sprintf("%f YiB", size.YiB())
}
