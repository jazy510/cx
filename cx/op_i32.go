package base

import (
	"fmt"
	"strconv"
	"math"
	"math/rand"
	"github.com/skycoin/skycoin/src/cipher/encoder"
)

func op_i32_i32 (expr *CXExpression, fp int) {
	inp1, out1 := expr.Inputs[0], expr.Outputs[0]
	out1Offset := GetFinalOffset(fp, out1)

	switch out1.Type {
	case TYPE_STR:
		WriteObject(out1Offset, encoder.Serialize(strconv.Itoa(int(ReadI32(fp, inp1)))))
	case TYPE_BYTE:
		WriteMemory(out1Offset, FromByte(byte(ReadI32(fp, inp1))))
	case TYPE_I32:
		WriteMemory(out1Offset, FromI32(ReadI32(fp, inp1)))
	case TYPE_I64:
		WriteMemory(out1Offset, FromI64(int64(ReadI32(fp, inp1))))
	case TYPE_F32:
		WriteMemory(out1Offset, FromF32(float32(ReadI32(fp, inp1))))
	case TYPE_F64:
		WriteMemory(out1Offset, FromF64(float64(ReadI32(fp, inp1))))
	}
}

func op_i32_print(expr *CXExpression, fp int) {
	inp1 := expr.Inputs[0]
	fmt.Println(ReadI32(fp, inp1))
}

// op_i32_add. The add built-in function returns the add of two numbers
func op_i32_add(expr *CXExpression, fp int) {
	inp1, inp2, out1 := expr.Inputs[0], expr.Inputs[1], expr.Outputs[0]
	outB1 := FromI32(ReadI32(fp, inp1) + ReadI32(fp, inp2))
	WriteMemory(GetFinalOffset(fp, out1), outB1)
}

// op_i32_sub. The sub built-in function returns the substract of two numbers
func op_i32_sub(expr *CXExpression, fp int) {
	inp1, inp2, out1 := expr.Inputs[0], expr.Inputs[1], expr.Outputs[0]
	outB1 := FromI32(ReadI32(fp, inp1) - ReadI32(fp, inp2))
	WriteMemory(GetFinalOffset(fp, out1), outB1)
}

// op_i32_sub. The mul built-in function returns the multiplication of two numbers
func op_i32_mul(expr *CXExpression, fp int) {
	inp1, inp2, out1 := expr.Inputs[0], expr.Inputs[1], expr.Outputs[0]
	outB1 := FromI32(ReadI32(fp, inp1) * ReadI32(fp, inp2))
	WriteMemory(GetFinalOffset(fp, out1), outB1)
}

// op_i32_sub. The div built-in function returns the divides two numbers
func op_i32_div(expr *CXExpression, fp int) {
	inp1, inp2, out1 := expr.Inputs[0], expr.Inputs[1], expr.Outputs[0]
	outB1 := FromI32(ReadI32(fp, inp1) / ReadI32(fp, inp2))
	WriteMemory(GetFinalOffset(fp, out1), outB1)
}

// op_i32_abs. The div built-in function returns the absolute number of the number
func op_i32_abs(expr *CXExpression, fp int) {
	inp1, out1 := expr.Inputs[0], expr.Outputs[0]
	outB1 := FromI32(int32(math.Abs(float64(ReadI32(fp, inp1)))))
	WriteMemory(GetFinalOffset(fp, out1), outB1)
}

// op_i32_pow. The div built-in function returns x**n for n>0 otherwise 1
func op_i32_pow(expr *CXExpression, fp int) {
	inp1, inp2, out1 := expr.Inputs[0], expr.Inputs[1], expr.Outputs[0]
	outB1 := FromI32(int32(math.Pow(float64(ReadI32(fp, inp1)), float64(ReadI32(fp, inp2)))))
	WriteMemory(GetFinalOffset(fp, out1), outB1)
}

// op_i32_gt. The gt built-in function returns true if x number is greater than a y number
func op_i32_gt(expr *CXExpression, fp int) {
	inp1, inp2, out1 := expr.Inputs[0], expr.Inputs[1], expr.Outputs[0]
	outB1 := FromBool(ReadI32(fp, inp1) > ReadI32(fp, inp2))
	WriteMemory(GetFinalOffset(fp, out1), outB1)
}

// op_i32_gteq. The gteq built-in function returns true if x number is greater or
// equal than a y number
func op_i32_gteq(expr *CXExpression, fp int) {
	inp1, inp2, out1 := expr.Inputs[0], expr.Inputs[1], expr.Outputs[0]
	outB1 := FromBool(ReadI32(fp, inp1) >= ReadI32(fp, inp2))
	WriteMemory(GetFinalOffset(fp, out1), outB1)
}

// op_i32_lt. The lt built-in function returns true if x number is less then
func op_i32_lt(expr *CXExpression, fp int) {
	inp1, inp2, out1 := expr.Inputs[0], expr.Inputs[1], expr.Outputs[0]
	outB1 := FromBool(ReadI32(fp, inp1) < ReadI32(fp, inp2))
	WriteMemory(GetFinalOffset(fp, out1), outB1)
}

// op_i32_lteq. The lteq built-in function returns true if x number is less or
// equal than a y number
func op_i32_lteq(expr *CXExpression, fp int) {
	inp1, inp2, out1 := expr.Inputs[0], expr.Inputs[1], expr.Outputs[0]
	outB1 := FromBool(ReadI32(fp, inp1) <= ReadI32(fp, inp2))
	WriteMemory(GetFinalOffset(fp, out1), outB1)
}

// op_i32_eq. The eq built-in function returns true if x number is equal to the y number
func op_i32_eq(expr *CXExpression, fp int) {
	inp1, inp2, out1 := expr.Inputs[0], expr.Inputs[1], expr.Outputs[0]
	outB1 := FromBool(ReadI32(fp, inp1) == ReadI32(fp, inp2))
	WriteMemory(GetFinalOffset(fp, out1), outB1)
}

// op_i32_uneq. The uneq built-in function returns true if x number is diferent to the y number
func op_i32_uneq(expr *CXExpression, fp int) {
	inp1, inp2, out1 := expr.Inputs[0], expr.Inputs[1], expr.Outputs[0]
	outB1 := FromBool(ReadI32(fp, inp1) != ReadI32(fp, inp2))
	WriteMemory(GetFinalOffset(fp, out1), outB1)
}

func op_i32_mod(expr *CXExpression, fp int) {
	inp1, inp2, out1 := expr.Inputs[0], expr.Inputs[1], expr.Outputs[0]
	outB1 := FromI32(ReadI32(fp, inp1) % ReadI32(fp, inp2))
	WriteMemory(GetFinalOffset(fp, out1), outB1)
}

func op_i32_rand(expr *CXExpression, fp int) {
	inp1, inp2, out1 := expr.Inputs[0], expr.Inputs[1], expr.Outputs[0]

	minimum := ReadI32(fp, inp1)
	maximum := ReadI32(fp, inp2)

	outB1 := FromI32(int32(rand.Intn(int(maximum-minimum)) + int(minimum)))

	WriteMemory(GetFinalOffset(fp, out1), outB1)
}

func op_i32_bitand(expr *CXExpression, fp int) {
	inp1, inp2, out1 := expr.Inputs[0], expr.Inputs[1], expr.Outputs[0]
	outB1 := FromI32(ReadI32(fp, inp1) & ReadI32(fp, inp2))
	WriteMemory(GetFinalOffset(fp, out1), outB1)
}

func op_i32_bitor(expr *CXExpression, fp int) {
	inp1, inp2, out1 := expr.Inputs[0], expr.Inputs[1], expr.Outputs[0]
	outB1 := FromI32(ReadI32(fp, inp1) | ReadI32(fp, inp2))
	WriteMemory(GetFinalOffset(fp, out1), outB1)
}

func op_i32_bitxor(expr *CXExpression, fp int) {
	inp1, inp2, out1 := expr.Inputs[0], expr.Inputs[1], expr.Outputs[0]
	outB1 := FromI32(ReadI32(fp, inp1) ^ ReadI32(fp, inp2))
	WriteMemory(GetFinalOffset(fp, out1), outB1)
}

func op_i32_bitclear(expr *CXExpression, fp int) {
	inp1, inp2, out1 := expr.Inputs[0], expr.Inputs[1], expr.Outputs[0]
	outB1 := FromI32(ReadI32(fp, inp1) &^ ReadI32(fp, inp2))
	WriteMemory(GetFinalOffset(fp, out1), outB1)
}

func op_i32_bitshl(expr *CXExpression, fp int) {
	inp1, inp2, out1 := expr.Inputs[0], expr.Inputs[1], expr.Outputs[0]
	outB1 := FromI32(int32(uint32(ReadI32(fp, inp1)) << uint32(ReadI32(fp, inp2))))
	WriteMemory(GetFinalOffset(fp, out1), outB1)
}

func op_i32_bitshr(expr *CXExpression, fp int) {
	inp1, inp2, out1 := expr.Inputs[0], expr.Inputs[1], expr.Outputs[0]
	outB1 := FromI32(int32(uint32(ReadI32(fp, inp1)) >> uint32(ReadI32(fp, inp2))))
	WriteMemory(GetFinalOffset(fp, out1), outB1)
}

// op_i32_sqrt. The sqrt built-in function returns the square root of x number
func op_i32_sqrt(expr *CXExpression, fp int) {
	inp1, out1 := expr.Inputs[0], expr.Outputs[0]
	outB1 := FromI32(int32(math.Sqrt(float64(ReadI32(fp, inp1)))))
	WriteMemory(GetFinalOffset(fp, out1), outB1)
}

// op_i32_log. The log built-in function returns the natural logarithm of x number
func op_i32_log(expr *CXExpression, fp int) {
	inp1, out1 := expr.Inputs[0], expr.Outputs[0]
	outB1 := FromI32(int32(math.Log(float64(ReadI32(fp, inp1)))))
	WriteMemory(GetFinalOffset(fp, out1), outB1)
}

// op_i32_log2. The log2 built-in function returns the natural logarithm based 2 of x number
func op_i32_log2(expr *CXExpression, fp int) {
	inp1, out1 := expr.Inputs[0], expr.Outputs[0]
	outB1 := FromI32(int32(math.Log2(float64(ReadI32(fp, inp1)))))
	WriteMemory(GetFinalOffset(fp, out1), outB1)
}

// op_i32_log10. The log10 built-in function returns the natural logarithm based 2 of x number
func op_i32_log10(expr *CXExpression, fp int) {
	inp1, out1 := expr.Inputs[0], expr.Outputs[0]
	outB1 := FromI32(int32(math.Log10(float64(ReadI32(fp, inp1)))))
	WriteMemory(GetFinalOffset(fp, out1), outB1)
}

// op_i32_max. The max built-in function returns the max value between x and y numbers
func op_i32_max(expr *CXExpression, fp int) {
	inp1, inp2, out1 := expr.Inputs[0], expr.Inputs[1], expr.Outputs[0]
	outB1 := FromI32(int32(math.Max(float64(ReadI32(fp, inp1)), float64(ReadI32(fp, inp2)))))
	WriteMemory(GetFinalOffset(fp, out1), outB1)
}

// op_i32_min. The min built-in function returns the min value between x and y numbers
func op_i32_min(expr *CXExpression, fp int) {
	inp1, inp2, out1 := expr.Inputs[0], expr.Inputs[1], expr.Outputs[0]
	outB1 := FromI32(int32(math.Min(float64(ReadI32(fp, inp1)), float64(ReadI32(fp, inp2)))))
	WriteMemory(GetFinalOffset(fp, out1), outB1)
}
