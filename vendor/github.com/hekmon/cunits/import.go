package cunits

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const parseRegex = "^([0-9,]+(\\.[0-9]+)?) ?(([KMGTPEZY]i?)?(B|bit))$"

var sizeMatch = regexp.MustCompile(parseRegex)

// Parse parses un string representation of a number with a suffix
func Parse(sizeSuffix string) (size Bits, err error) {
	// Does it match ?
	match := sizeMatch.FindSubmatch([]byte(sizeSuffix))
	if match == nil {
		err = fmt.Errorf("string does not match the parsing regex: %s", parseRegex)
		return
	}
	if len(match) < 4 {
		err = fmt.Errorf("regex matching did not return enough groups")
		return
	}
	// Extract number
	num, err := strconv.ParseFloat(strings.Replace(string(match[1]), ",", "", -1), 64)
	if err != nil {
		err = fmt.Errorf("extracted number '%s' can't be parsed as float64: %v", string(match[1]), err)
		return
	}
	// Findout the unit
	switch string(match[3]) {
	case "bit":
		size = Bits(num)
	case "B":
		size = ImportInByte(num)
	// Decimal prefix of bits
	case "Kbit":
		size = ImportInKbit(num)
	case "Mbit":
		size = ImportInMbit(num)
	case "Gbit":
		size = ImportInGbit(num)
	case "Tbit":
		size = ImportInTbit(num)
	case "Pbit":
		size = ImportInPbit(num)
	case "Ebit":
		size = ImportInEbit(num)
	case "Zbit":
		size = ImportInZbit(num)
	case "Ybit":
		size = ImportInYbit(num)
	// Binary prefix of bits
	case "Kibit":
		size = ImportInKibit(num)
	case "Mibit":
		size = ImportInMibit(num)
	case "Gibit":
		size = ImportInGibit(num)
	case "Tibit":
		size = ImportInTibit(num)
	case "Pibit":
		size = ImportInPibit(num)
	case "Eibit":
		size = ImportInEibit(num)
	case "Zibit":
		size = ImportInZibit(num)
	case "Yibit":
		size = ImportInYibit(num)
	// Decimal prefix of bytes
	case "KB":
		size = ImportInKB(num)
	case "MB":
		size = ImportInMB(num)
	case "GB":
		size = ImportInGB(num)
	case "TB":
		size = ImportInTB(num)
	case "PB":
		size = ImportInPB(num)
	case "EB":
		size = ImportInEB(num)
	case "ZB":
		size = ImportInZB(num)
	case "YB":
		size = ImportInYB(num)
	// Binary prefix of bytes
	case "KiB":
		size = ImportInKiB(num)
	case "MiB":
		size = ImportInMiB(num)
	case "GiB":
		size = ImportInGiB(num)
	case "TiB":
		size = ImportInTiB(num)
	case "PiB":
		size = ImportInPiB(num)
	case "EiB":
		size = ImportInEiB(num)
	case "ZiB":
		size = ImportInZiB(num)
	case "YiB":
		size = ImportInYiB(num)
	// or not
	default:
		err = fmt.Errorf("extracted unit '%s' is unknown", string(match[3]))
	}
	return
}

// ImportInByte imports a number in byte
func ImportInByte(sizeInByte float64) Bits {
	return Bits(sizeInByte * Byte)
}

/*
	Decimal prefix of bits
*/

// ImportInKbit imports a number in kilobit
func ImportInKbit(sizeInKbit float64) Bits {
	return Bits(sizeInKbit * Kbit)
}

// ImportInMbit imports a number in megabit
func ImportInMbit(sizeInMbit float64) Bits {
	return Bits(sizeInMbit * Mbit)
}

// ImportInGbit imports a number in gigabit
func ImportInGbit(sizeInGbit float64) Bits {
	return Bits(sizeInGbit * Gbit)
}

// ImportInTbit imports a number in terabit
func ImportInTbit(sizeInTbit float64) Bits {
	return Bits(sizeInTbit * Tbit)
}

// ImportInPbit imports a number in petabit
func ImportInPbit(sizeInPbit float64) Bits {
	return Bits(sizeInPbit * Pbit)
}

// ImportInEbit imports a number in exabit
func ImportInEbit(sizeInEbit float64) Bits {
	return Bits(sizeInEbit * Ebit)
}

// ImportInZbit imports a number in zettabit (sizeInZbit better < 1)
func ImportInZbit(sizeInZbit float64) Bits {
	return Bits(sizeInZbit * Zbit)
}

// ImportInYbit imports a number in yottabit (sizeInYbit better < 1)
func ImportInYbit(sizeInYbit float64) Bits {
	return Bits(sizeInYbit * Ybit)
}

/*
	Binary prefix of bits
*/

// ImportInKibit imports a number in kibibit
func ImportInKibit(sizeInKibit float64) Bits {
	return Bits(sizeInKibit * Kibit)
}

// ImportInMibit imports a number in mebibit
func ImportInMibit(sizeInMibit float64) Bits {
	return Bits(sizeInMibit * Mibit)
}

// ImportInGibit imports a number in gibibit
func ImportInGibit(sizeInGibit float64) Bits {
	return Bits(sizeInGibit * Gibit)
}

// ImportInTibit imports a number in tebibit
func ImportInTibit(sizeInTibit float64) Bits {
	return Bits(sizeInTibit * Tibit)
}

// ImportInPibit imports a number in pebibit
func ImportInPibit(sizeInPibit float64) Bits {
	return Bits(sizeInPibit * Pibit)
}

// ImportInEibit imports a number in exbibit
func ImportInEibit(sizeInEibit float64) Bits {
	return Bits(sizeInEibit * Eibit)
}

// ImportInZibit imports a number in zebibit (sizeInZibit better < 1)
func ImportInZibit(sizeInZibit float64) Bits {
	return Bits(sizeInZibit * Zibit)
}

// ImportInYibit imports a number in yobibit (sizeInYibit better < 1)
func ImportInYibit(sizeInYibit float64) Bits {
	return Bits(sizeInYibit * Yibit)
}

/*
	Decimal prefix of bytes
*/

// ImportInKB imports a number in kilobyte
func ImportInKB(sizeInKB float64) Bits {
	return Bits(sizeInKB * KB)
}

// ImportInMB imports a number in megabyte
func ImportInMB(sizeInMB float64) Bits {
	return Bits(sizeInMB * MB)
}

// ImportInGB imports a number in gigabyte
func ImportInGB(sizeInGB float64) Bits {
	return Bits(sizeInGB * GB)
}

// ImportInTB imports a number in terabyte
func ImportInTB(sizeInTB float64) Bits {
	return Bits(sizeInTB * TB)
}

// ImportInPB imports a number in petabyte
func ImportInPB(sizeInPB float64) Bits {
	return Bits(sizeInPB * PB)
}

// ImportInEB imports a number in exabyte
func ImportInEB(sizeInEB float64) Bits {
	return Bits(sizeInEB * EB)
}

// ImportInZB imports a number in zettabyte (sizeInZB better < 1)
func ImportInZB(sizeInZB float64) Bits {
	return Bits(sizeInZB * ZB)
}

// ImportInYB imports a number in yottabyte (sizeInYB better < 1)
func ImportInYB(sizeInYB float64) Bits {
	return Bits(sizeInYB * YB)
}

/*
	Binary prefix of bytes
*/

// ImportInKiB imports a number in kilobyte
func ImportInKiB(sizeInKiB float64) Bits {
	return Bits(sizeInKiB * KiB)
}

// ImportInMiB imports a number in megabyte
func ImportInMiB(sizeInMiB float64) Bits {
	return Bits(sizeInMiB * MiB)
}

// ImportInGiB imports a number in gigabyte
func ImportInGiB(sizeInGiB float64) Bits {
	return Bits(sizeInGiB * GiB)
}

// ImportInTiB imports a number in terabyte
func ImportInTiB(sizeInTiB float64) Bits {
	return Bits(sizeInTiB * TiB)
}

// ImportInPiB imports a number in petabyte
func ImportInPiB(sizeInPiB float64) Bits {
	return Bits(sizeInPiB * PiB)
}

// ImportInEiB imports a number in exabyte (sizeInEiB better < 1)
func ImportInEiB(sizeInEiB float64) Bits {
	return Bits(sizeInEiB * EiB)
}

// ImportInZiB imports a number in zettabyte (sizeInZiB better < 1)
func ImportInZiB(sizeInZiB float64) Bits {
	return Bits(sizeInZiB * ZiB)
}

// ImportInYiB imports a number in yottabyte (sizeInYiB better < 1)
func ImportInYiB(sizeInYiB float64) Bits {
	return Bits(sizeInYiB * YiB)
}
