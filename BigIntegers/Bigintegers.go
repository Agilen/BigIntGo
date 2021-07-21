package bigintegers

import (
	"fmt"
	"strconv"
	"strings"
)

func ReadDec(str string) []uint64 {
	str = strings.TrimSpace(str)
	numStr := strings.Split(str, "")
	var num []uint64
	for i := len(numStr) - 1; i >= 0; i-- {
		n, _ := strconv.ParseUint(numStr[i], 10, 64)
		num = append(num, n)
	}
	res := []uint64{0}
	for i := 0; i < len(num); i++ {
		res = LongAdd(LongMul(Pow(10, i), []uint64{num[i]}), res)
	}
	res = DelNull(res)

	return res
}
func ReadHex(str string) []uint64 {
	str = strings.TrimSpace(str)
	Lenght := 0
	if len(str) < 16 {
		Digit := make([]uint64, 1)
		n, _ := strconv.ParseUint(str, 16, 64)
		Digit[0] = uint64(n)
		return Digit
	}
	if len(str)%16 != 0 {
		Lenght = len(str)/16 + 1
	} else {
		Lenght = len(str) / 16
	}
	Digit := make([]uint64, Lenght)
	for i := 0; i < Lenght-1; i++ {
		s := str[len(str)-(i+1)*16 : len(str)-i*16]
		n, _ := strconv.ParseUint(s, 16, 64)
		Digit[i] = uint64(n)
	}
	if len(str)%16 != 0 {
		s := str[0 : len(str)%16]
		n, _ := strconv.ParseUint(s, 16, 64)
		Digit[Lenght-1] = uint64(n)
	} else if len(str)%16 == 0 {
		s := str[0:16]
		n, _ := strconv.ParseUint(s, 16, 64)
		Digit[len(Digit)-1] = uint64(n)
	}

	return Digit
}
func ReadBin(str string) []uint64 {
	str = strings.TrimSpace(str)
	Lenght := 0
	if len(str) < 64 {
		Digit := make([]uint64, 1)
		n, _ := strconv.ParseUint(str, 2, 64)
		Digit[0] = uint64(n)
		return Digit
	}
	if len(str)%64 != 0 {
		Lenght = len(str)/64 + 1
	} else {
		Lenght = len(str) / 64
	}
	Digit := make([]uint64, Lenght)
	for i := 0; i < Lenght-1; i++ {
		s := str[len(str)-(i+1)*64 : len(str)-i*64]
		n, _ := strconv.ParseUint(s, 2, 64)
		Digit[i] = uint64(n)
	}
	if len(str)%64 != 0 {
		s := str[0 : len(str)%64]
		n, _ := strconv.ParseUint(s, 2, 64)
		Digit[Lenght-1] = uint64(n)
	} else if len(str)%64 == 0 {
		s := str[0:64]
		n, _ := strconv.ParseUint(s, 2, 64)
		Digit[len(Digit)-1] = uint64(n)
	}

	return Digit
}
func DelLeadZero(a string) string {
	if string(a[0]) == "0" {
		for string(a[0]) == "0" {
			a = a[1:]
		}
	}
	return a
}
func DelNull(a []uint64) []uint64 {
	k := 0
	for i := len(a) - 1; i >= 0; i-- {
		if a[i] != 0 {
			k = i
			break
		}
	}
	a = a[0 : k+1]
	return a
}
func ToHex(a []uint64) string {
	result := ""
	digit := ""
	buf := ""
	k := len(a) - 1
	for i := 0; i < len(a); i++ {
		digit = fmt.Sprintf("%X", a[k])
		if len(digit) < 16 {
			for i := 0; i < 16-len(digit); i++ {
				buf += "0"
			}
		}
		buf += digit
		result += buf
		buf = ""
		k--
	}

	return result
}
func ToDec(a []uint64) string {
	ten := []uint64{10}
	output := ""
	for LongCmp(a, ten) == 1 {
		mod := LongMod(a, ten)
		output = strconv.Itoa(int(mod[0])) + output
		a = LongDiv(a, ten)
	}
	output = strconv.Itoa(int(a[0])) + output
	return output
}
func ToUInt32(a []uint64) []uint32 {
	str := ToHex(a)
	Lenght := 0
	if len(str) < 8 {
		Digit := make([]uint32, 1)
		n, _ := strconv.ParseUint(str, 16, 64)
		Digit[0] = uint32(n)
		return Digit
	}
	if len(str)%8 != 0 {
		Lenght = len(str)/8 + 1
	} else {
		Lenght = len(str) / 8
	}
	Digit := make([]uint32, Lenght)
	for i := 0; i < Lenght-1; i++ {
		s := str[len(str)-(i+1)*8 : len(str)-i*8]
		n, _ := strconv.ParseUint(s, 16, 64)
		Digit[i] = uint32(n)
	}
	if len(str)%8 != 0 {
		s := str[0 : len(str)%8]
		n, _ := strconv.ParseUint(s, 16, 64)
		Digit[Lenght-1] = uint32(n)
	} else if len(str)%8 == 0 {
		s := str[0:8]
		n, _ := strconv.ParseUint(s, 16, 64)
		Digit[len(Digit)-1] = uint32(n)
	}

	return Digit
}
func ToUInt64(a []uint32) []uint64 {
	result := ""
	digit := ""
	buf := ""
	k := len(a) - 1
	for i := 0; i < len(a); i++ {
		digit = fmt.Sprintf("%X", a[k])
		if len(digit) < 8 {
			for i := 0; i < 8-len(digit); i++ {
				buf += "0"
			}
		}
		buf += digit
		result += buf
		buf = ""
		k--
	}

	return ReadHex(result)
}
func ToBin(a []uint64) string {
	result := ""
	digit := ""
	buf := ""
	k := len(a) - 1
	for i := 0; i < len(a); i++ {
		digit = strconv.FormatUint((a[k]), 2)
		if len(digit) < 64 {
			for i := 0; i < 64-len(digit); i++ {
				buf += "0"
			}
		}
		buf += digit
		result += buf
		buf = ""
		k--
	}

	return result
}
func ToBinDigit(a uint64) string {
	result := ""
	digit := ""
	buf := ""

	digit = strconv.FormatUint((a), 2)
	if len(digit) < 64 {
		for i := 0; i < 64-len(digit); i++ {
			buf += "0"
		}
	}
	buf += digit
	result += buf
	buf = ""

	return result
}

func Pow(a, k int) []uint64 {
	if k == 0 {
		return []uint64{1}
	}
	n := []uint64{uint64(a)}
	res := []uint64{uint64(a)}
	for i := 0; i < k-1; i++ {
		res = LongMul(res, n)
		res = DelNull(res)
	}
	return res
}
func LongCmp(a, b []uint64) int {
	a, b = SameSize(a, b)

	for i := len(a) - 1; i >= 0; i-- {
		if a[i] > b[i] {
			return 1 //>
		} else if a[i] < b[i] {
			return -1 //<
		}
	}
	return 0 //=
}
func LongAdd(a, b []uint64) []uint64 {
	a, b = SameSize(a, b)
	C := make([]uint64, len(a))
	carry := uint64(0)
	for i := 0; i < len(a); i++ {
		temp := a[i] + b[i] + carry
		C[i] = temp & 0xffffffffffffffff
		carry = isCarryExist(a[i], b[i])

	}

	return DelNull(C)
}
func LongSub(a, b []uint64) []uint64 {
	a, b = SameSize(a, b)
	c := make([]uint64, len(a))
	borrow := uint64(0)
	for i := 0; i < len(a); i++ {
		c[i] = a[i] - b[i] - (borrow)
		if b[i] != 0 && b[i]+(borrow) == 0 {
			borrow = 1
		} else if a[i] >= b[i]+(borrow) {
			borrow = 0
		} else {
			borrow = 1
		}
	}
	return DelNull(c)
}
func LongMulOneDigit(a []uint32, b uint32) []uint32 {
	c := make([]uint32, len(a)+1)
	carry := uint64(0)
	for i := 0; i < len(a); i++ {
		temp := (uint64(a[i])*uint64(b) + carry)
		c[i] = uint32(temp & 0xffffffff)
		carry = temp >> 32
	}
	c[len(a)] = uint32(carry)
	return c
}
func LongMul(a, b []uint64) []uint64 {
	a, b = SameSize(a, b)
	A := ToUInt32(a)
	B := ToUInt32(b)
	c := make([]uint64, len(a)*2)
	for i := 0; i < len(B); i++ {
		temp := LongMulOneDigit(A, B[i])
		temp64 := ToUInt64(temp)

		temps := LongShiftLeft(temp64, i*32)

		c = LongAdd(c, temps)
	}
	return DelNull(c)
}
func LongShiftLeft(a []uint64, n int) []uint64 {

	if n <= -1 {
		return []uint64{1}
	}
	if n == 0 {
		return a
	}
	bin := ToBin(a)

	for i := 0; i < n; i++ {
		bin += "0"
	}

	return ReadBin(bin)
}
func LongDivMod(a, b []uint64) ([]uint64, []uint64) {
	k := BitLength(b)
	r := a
	t := 0
	var c []uint64
	var q []uint64
	i := 0
	for LongCmp(r, b) == 1 || LongCmp(r, b) == 0 {

		t = BitLength(r)
		c = LongShiftLeft(b, t-k)
		for LongCmp(r, c) == -1 {

			if LongCmp(r, c) == -1 {
				t = t - 1
				c = LongShiftLeft(b, t-k)
			}

		}

		r = LongSub(r, c)
		q = LongAdd(q, LongShiftLeft([]uint64{2}, t-k-1))

		i++

	}
	return DelNull(q), DelNull(r)
}
func LongDiv(a, b []uint64) []uint64 {
	k := BitLength(b)
	r := a
	t := 0
	var c []uint64
	var q []uint64

	for LongCmp(r, b) == 1 || LongCmp(r, b) == 0 {

		t = BitLength(r)
		c = LongShiftLeft(b, t-k)
		for LongCmp(r, c) == -1 {

			if LongCmp(r, c) == -1 {
				t = t - 1
				c = LongShiftLeft(b, t-k)
			}
		}
		r = LongSub(r, c)
		q = LongAdd(q, LongShiftLeft([]uint64{2}, t-k-1))
	}
	return DelNull(q)
}
func LongMod(a, b []uint64) []uint64 {
	k := BitLength(b)
	r := a
	t := 0
	var c []uint64
	for LongCmp(r, b) == 1 || LongCmp(r, b) == 0 {

		t = BitLength(r)
		c = LongShiftLeft(b, t-k)
		for LongCmp(r, c) == -1 {
			if LongCmp(r, c) == -1 {
				t = t - 1
				c = LongShiftLeft(b, t-k)
			}
		}
		r = LongSub(r, c)
	}
	return DelNull(r)
}
func LongModPowerBarrett(a, b, n []uint64) []uint64 {
	zero := []uint64{0}

	if LongCmp(b, zero) == 0 {
		return []uint64{0}
	}
	B := DelLeadZero(ToBin(b))
	m := len(B)
	k := BitLength(n)
	mu := LongDiv(LongShiftLeft([]uint64{1}, 2*k), n)
	c := []uint64{1}

	for i := m - 1; i >= 0; i-- {
		if string(B[i]) == "1" {
			buf := LongMul(a, c)
			c = BarrettReduction(buf, n, mu)
		}
		a = BarrettReduction(LongMul(a, a), n, mu)
	}
	return DelNull(c)
}
func BarrettReduction(x, n, mu []uint64) []uint64 {

	if LongCmp(x, n) == 0 || LongCmp(x, n) == -1 {
		return x
	}
	k := BitLength(n)

	q := KillDigits(x, k-1)
	q = LongMul(q, mu)
	q = KillDigits(q, k+1)
	r := LongMul(q, n)
	t := LongShiftLeft([]uint64{1}, k+1)
	r1 := LongMod(x, t)
	r2 := LongMod(r, t)
	if LongCmp(r1, r2) == 0 || LongCmp(r1, r2) == 1 {
		r = LongSub(r1, r2)
	} else {
		r = LongSub(LongAdd(t, r1), r2)
	}
	for LongCmp(r, n) == 0 || LongCmp(r, n) == 1 {
		r = LongSub(r, n)
	}
	return r
}

func KillDigits(a []uint64, k int) []uint64 {
	zero := []uint64{0}

	if LongCmp(a, zero) == 0 {
		return a
	}

	bit := ToBin(a)
	bit = bit[0 : len(bit)-k]
	a = ReadBin(bit)
	return a
}
func isCarryExist(a, b uint64) uint64 {
	A := ToBinDigit(a)
	B := ToBinDigit(b)

	if string(A[0]) == "1" && string(B[0]) == "1" {
		return 1
	} else if string(A[0]) == "0" && string(B[0]) == "0" {
		return 0
	} else {
		for i := 1; i < len(A); i++ {
			if string(A[i]) == "1" && string(B[i]) == "1" {
				return 1
			} else if string(A[i]) == "0" && string(B[i]) == "0" {
				return 0
			}
		}
		return 0
	}
}
func SameSize(a, b []uint64) ([]uint64, []uint64) {
	leng := 0
	if len(a) == len(b) {
		return a, b
	}
	if len(a) > len(b) {
		leng = len(a)
	} else if len(a) < len(b) {
		leng = len(b)
	}
	A := make([]uint64, leng)
	B := make([]uint64, leng)
	for i := 0; i < len(a); i++ {
		A[i] = a[i]
	}
	for i := 0; i < len(b); i++ {
		B[i] = b[i]
	}
	return A, B
}
func BitLength(a []uint64) int {
	bit := DelLeadZero(ToBin(a))
	return len(bit)
}
func IsEvenNumber(a []uint64) bool {
	A := ToBinDigit(a[0])
	b := true
	if string(A[len(A)-1]) == "0" {
		return b
	} else {
		return !b
	}
}
