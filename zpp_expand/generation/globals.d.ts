class Native {
    private constructor()
    private readonly __Native: Native
}
type NumberLike = number | string | null | undefined | GoInt | GoInt64 | GoInt32 | GoInt16 | GoInt8 | GoUint | GoUint64 | GoUint32 | GoUint16 | GoUint8 | GoFloat64 | GoFloat32

class GoInt extends Native {
    private readonly __GoInt: GoInt
    private constructor()
    String(): string
    ToNumber(): number
    ToInt(): GoInt
    ToInt64(): GoInt64
    ToInt32(): GoInt32
    ToInt16(): GoInt16
    ToInt8(): GoInt8
    ToUint(): GoUint
    ToUint64(): GoUint64
    ToUint32(): GoUint32
    ToUint16(): GoUint16
    ToUint8(): GoUint8
    ToFloat64(): GoFloat64
    ToFloat32(): GoFloat32
    ABS(): GoInt
    Negate(): GoInt
    Add(...value: Array<NumberLike>): GoInt
    Sub(...value: Array<NumberLike>): GoInt
    Mul(...value: Array<NumberLike>): GoInt
    Div(...value: Array<NumberLike>): GoInt
    Mod(...value: Array<NumberLike>): GoInt
    And(...value: Array<NumberLike>): GoInt
    AndNot(...value: Array<NumberLike>): GoInt
    Not(): GoInt
    Or(...value: Array<NumberLike>): GoInt
    Xor(...value: Array<NumberLike>): GoInt
    ShiftLeft(...value: Array<NumberLike>): GoInt
    ShiftRight(...value: Array<NumberLike>): GoInt
    Compare(val: NumberLike): number
    Max(...value: Array<NumberLike>): GoInt
    Min(...value: Array<NumberLike>): GoInt
}
function NewInt(val: number): GoInt
function NewInt(val: string, base: number | string): GoInt

class GoInt64 extends Native {
    private readonly __GoInt64: GoInt64
    private constructor()
    String(): string
    ToNumber(): number
    ToInt(): GoInt
    ToInt64(): GoInt64
    ToInt32(): GoInt32
    ToInt16(): GoInt16
    ToInt8(): GoInt8
    ToUint(): GoUint
    ToUint64(): GoUint64
    ToUint32(): GoUint32
    ToUint16(): GoUint16
    ToUint8(): GoUint8
    ToFloat64(): GoFloat64
    ToFloat32(): GoFloat32
    ABS(): GoInt64
    Negate(): GoInt64
    Add(...value: Array<NumberLike>): GoInt64
    Sub(...value: Array<NumberLike>): GoInt64
    Mul(...value: Array<NumberLike>): GoInt64
    Div(...value: Array<NumberLike>): GoInt64
    Mod(...value: Array<NumberLike>): GoInt64
    And(...value: Array<NumberLike>): GoInt64
    AndNot(...value: Array<NumberLike>): GoInt64
    Not(): GoInt64
    Or(...value: Array<NumberLike>): GoInt64
    Xor(...value: Array<NumberLike>): GoInt64
    ShiftLeft(...value: Array<NumberLike>): GoInt64
    ShiftRight(...value: Array<NumberLike>): GoInt64
    Compare(val: NumberLike): number
    Max(...value: Array<NumberLike>): GoInt64
    Min(...value: Array<NumberLike>): GoInt64
}
function NewInt64(val: number): GoInt64
function NewInt64(val: string, base: number | string): GoInt64

class GoInt32 extends Native {
    private readonly __GoInt32: GoInt32
    private constructor()
    String(): string
    ToNumber(): number
    ToInt(): GoInt
    ToInt64(): GoInt64
    ToInt32(): GoInt32
    ToInt16(): GoInt16
    ToInt8(): GoInt8
    ToUint(): GoUint
    ToUint64(): GoUint64
    ToUint32(): GoUint32
    ToUint16(): GoUint16
    ToUint8(): GoUint8
    ToFloat64(): GoFloat64
    ToFloat32(): GoFloat32
    ABS(): GoInt32
    Negate(): GoInt32
    Add(...value: Array<NumberLike>): GoInt32
    Sub(...value: Array<NumberLike>): GoInt32
    Mul(...value: Array<NumberLike>): GoInt32
    Div(...value: Array<NumberLike>): GoInt32
    Mod(...value: Array<NumberLike>): GoInt32
    And(...value: Array<NumberLike>): GoInt32
    AndNot(...value: Array<NumberLike>): GoInt32
    Not(): GoInt32
    Or(...value: Array<NumberLike>): GoInt32
    Xor(...value: Array<NumberLike>): GoInt32
    ShiftLeft(...value: Array<NumberLike>): GoInt32
    ShiftRight(...value: Array<NumberLike>): GoInt32
    Compare(val: NumberLike): number
    Max(...value: Array<NumberLike>): GoInt32
    Min(...value: Array<NumberLike>): GoInt32
}
function NewInt32(val: number): GoInt32
function NewInt32(val: string, base: number | string): GoInt32

class GoInt16 extends Native {
    private readonly __GoInt16: GoInt16
    private constructor()
    String(): string
    ToNumber(): number
    ToInt(): GoInt
    ToInt64(): GoInt64
    ToInt32(): GoInt32
    ToInt16(): GoInt16
    ToInt8(): GoInt8
    ToUint(): GoUint
    ToUint64(): GoUint64
    ToUint32(): GoUint32
    ToUint16(): GoUint16
    ToUint8(): GoUint8
    ToFloat64(): GoFloat64
    ToFloat32(): GoFloat32
    ABS(): GoInt16
    Negate(): GoInt16
    Add(...value: Array<NumberLike>): GoInt16
    Sub(...value: Array<NumberLike>): GoInt16
    Mul(...value: Array<NumberLike>): GoInt16
    Div(...value: Array<NumberLike>): GoInt16
    Mod(...value: Array<NumberLike>): GoInt16
    And(...value: Array<NumberLike>): GoInt16
    AndNot(...value: Array<NumberLike>): GoInt16
    Not(): GoInt16
    Or(...value: Array<NumberLike>): GoInt16
    Xor(...value: Array<NumberLike>): GoInt16
    ShiftLeft(...value: Array<NumberLike>): GoInt16
    ShiftRight(...value: Array<NumberLike>): GoInt16
    Compare(val: NumberLike): number
    Max(...value: Array<NumberLike>): GoInt16
    Min(...value: Array<NumberLike>): GoInt16
}
function NewInt16(val: number): GoInt16
function NewInt16(val: string, base: number | string): GoInt16

class GoInt8 extends Native {
    private readonly __GoInt8: GoInt8
    private constructor()
    String(): string
    ToNumber(): number
    ToInt(): GoInt
    ToInt64(): GoInt64
    ToInt32(): GoInt32
    ToInt16(): GoInt16
    ToInt8(): GoInt8
    ToUint(): GoUint
    ToUint64(): GoUint64
    ToUint32(): GoUint32
    ToUint16(): GoUint16
    ToUint8(): GoUint8
    ToFloat64(): GoFloat64
    ToFloat32(): GoFloat32
    ABS(): GoInt8
    Negate(): GoInt8
    Add(...value: Array<NumberLike>): GoInt8
    Sub(...value: Array<NumberLike>): GoInt8
    Mul(...value: Array<NumberLike>): GoInt8
    Div(...value: Array<NumberLike>): GoInt8
    Mod(...value: Array<NumberLike>): GoInt8
    And(...value: Array<NumberLike>): GoInt8
    AndNot(...value: Array<NumberLike>): GoInt8
    Not(): GoInt8
    Or(...value: Array<NumberLike>): GoInt8
    Xor(...value: Array<NumberLike>): GoInt8
    ShiftLeft(...value: Array<NumberLike>): GoInt8
    ShiftRight(...value: Array<NumberLike>): GoInt8
    Compare(val: NumberLike): number
    Max(...value: Array<NumberLike>): GoInt8
    Min(...value: Array<NumberLike>): GoInt8
}
function NewInt8(val: number): GoInt8
function NewInt8(val: string, base: number | string): GoInt8

class GoUint extends Native {
    private readonly __GoUint: GoUint
    private constructor()
    String(): string
    ToNumber(): number
    ToInt(): GoInt
    ToInt64(): GoInt64
    ToInt32(): GoInt32
    ToInt16(): GoInt16
    ToInt8(): GoInt8
    ToUint(): GoUint
    ToUint64(): GoUint64
    ToUint32(): GoUint32
    ToUint16(): GoUint16
    ToUint8(): GoUint8
    ToFloat64(): GoFloat64
    ToFloat32(): GoFloat32
    Add(...value: Array<NumberLike>): GoUint
    Sub(...value: Array<NumberLike>): GoUint
    Mul(...value: Array<NumberLike>): GoUint
    Div(...value: Array<NumberLike>): GoUint
    Mod(...value: Array<NumberLike>): GoUint
    And(...value: Array<NumberLike>): GoUint
    AndNot(...value: Array<NumberLike>): GoUint
    Not(): GoUint
    Or(...value: Array<NumberLike>): GoUint
    Xor(...value: Array<NumberLike>): GoUint
    ShiftLeft(...value: Array<NumberLike>): GoUint
    ShiftRight(...value: Array<NumberLike>): GoUint
    Compare(val: NumberLike): number
    Max(...value: Array<NumberLike>): GoUint
    Min(...value: Array<NumberLike>): GoUint
}
function NewUint(val: number): GoUint
function NewUint(val: string, base: number | string): GoUint

class GoUint64 extends Native {
    private readonly __GoUint64: GoUint64
    private constructor()
    String(): string
    ToNumber(): number
    ToInt(): GoInt
    ToInt64(): GoInt64
    ToInt32(): GoInt32
    ToInt16(): GoInt16
    ToInt8(): GoInt8
    ToUint(): GoUint
    ToUint64(): GoUint64
    ToUint32(): GoUint32
    ToUint16(): GoUint16
    ToUint8(): GoUint8
    ToFloat64(): GoFloat64
    ToFloat32(): GoFloat32
    Add(...value: Array<NumberLike>): GoUint64
    Sub(...value: Array<NumberLike>): GoUint64
    Mul(...value: Array<NumberLike>): GoUint64
    Div(...value: Array<NumberLike>): GoUint64
    Mod(...value: Array<NumberLike>): GoUint64
    And(...value: Array<NumberLike>): GoUint64
    AndNot(...value: Array<NumberLike>): GoUint64
    Not(): GoUint64
    Or(...value: Array<NumberLike>): GoUint64
    Xor(...value: Array<NumberLike>): GoUint64
    ShiftLeft(...value: Array<NumberLike>): GoUint64
    ShiftRight(...value: Array<NumberLike>): GoUint64
    Compare(val: NumberLike): number
    Max(...value: Array<NumberLike>): GoUint64
    Min(...value: Array<NumberLike>): GoUint64
}
function NewUint64(val: number): GoUint64
function NewUint64(val: string, base: number | string): GoUint64

class GoUint32 extends Native {
    private readonly __GoUint32: GoUint32
    private constructor()
    String(): string
    ToNumber(): number
    ToInt(): GoInt
    ToInt64(): GoInt64
    ToInt32(): GoInt32
    ToInt16(): GoInt16
    ToInt8(): GoInt8
    ToUint(): GoUint
    ToUint64(): GoUint64
    ToUint32(): GoUint32
    ToUint16(): GoUint16
    ToUint8(): GoUint8
    ToFloat64(): GoFloat64
    ToFloat32(): GoFloat32
    Add(...value: Array<NumberLike>): GoUint32
    Sub(...value: Array<NumberLike>): GoUint32
    Mul(...value: Array<NumberLike>): GoUint32
    Div(...value: Array<NumberLike>): GoUint32
    Mod(...value: Array<NumberLike>): GoUint32
    And(...value: Array<NumberLike>): GoUint32
    AndNot(...value: Array<NumberLike>): GoUint32
    Not(): GoUint32
    Or(...value: Array<NumberLike>): GoUint32
    Xor(...value: Array<NumberLike>): GoUint32
    ShiftLeft(...value: Array<NumberLike>): GoUint32
    ShiftRight(...value: Array<NumberLike>): GoUint32
    Compare(val: NumberLike): number
    Max(...value: Array<NumberLike>): GoUint32
    Min(...value: Array<NumberLike>): GoUint32
}
function NewUint32(val: number): GoUint32
function NewUint32(val: string, base: number | string): GoUint32

class GoUint16 extends Native {
    private readonly __GoUint16: GoUint16
    private constructor()
    String(): string
    ToNumber(): number
    ToInt(): GoInt
    ToInt64(): GoInt64
    ToInt32(): GoInt32
    ToInt16(): GoInt16
    ToInt8(): GoInt8
    ToUint(): GoUint
    ToUint64(): GoUint64
    ToUint32(): GoUint32
    ToUint16(): GoUint16
    ToUint8(): GoUint8
    ToFloat64(): GoFloat64
    ToFloat32(): GoFloat32
    Add(...value: Array<NumberLike>): GoUint16
    Sub(...value: Array<NumberLike>): GoUint16
    Mul(...value: Array<NumberLike>): GoUint16
    Div(...value: Array<NumberLike>): GoUint16
    Mod(...value: Array<NumberLike>): GoUint16
    And(...value: Array<NumberLike>): GoUint16
    AndNot(...value: Array<NumberLike>): GoUint16
    Not(): GoUint16
    Or(...value: Array<NumberLike>): GoUint16
    Xor(...value: Array<NumberLike>): GoUint16
    ShiftLeft(...value: Array<NumberLike>): GoUint16
    ShiftRight(...value: Array<NumberLike>): GoUint16
    Compare(val: NumberLike): number
    Max(...value: Array<NumberLike>): GoUint16
    Min(...value: Array<NumberLike>): GoUint16
}
function NewUint16(val: number): GoUint16
function NewUint16(val: string, base: number | string): GoUint16

class GoUint8 extends Native {
    private readonly __GoUint8: GoUint8
    private constructor()
    String(): string
    ToNumber(): number
    ToInt(): GoInt
    ToInt64(): GoInt64
    ToInt32(): GoInt32
    ToInt16(): GoInt16
    ToInt8(): GoInt8
    ToUint(): GoUint
    ToUint64(): GoUint64
    ToUint32(): GoUint32
    ToUint16(): GoUint16
    ToUint8(): GoUint8
    ToFloat64(): GoFloat64
    ToFloat32(): GoFloat32
    Add(...value: Array<NumberLike>): GoUint8
    Sub(...value: Array<NumberLike>): GoUint8
    Mul(...value: Array<NumberLike>): GoUint8
    Div(...value: Array<NumberLike>): GoUint8
    Mod(...value: Array<NumberLike>): GoUint8
    And(...value: Array<NumberLike>): GoUint8
    AndNot(...value: Array<NumberLike>): GoUint8
    Not(): GoUint8
    Or(...value: Array<NumberLike>): GoUint8
    Xor(...value: Array<NumberLike>): GoUint8
    ShiftLeft(...value: Array<NumberLike>): GoUint8
    ShiftRight(...value: Array<NumberLike>): GoUint8
    Compare(val: NumberLike): number
    Max(...value: Array<NumberLike>): GoUint8
    Min(...value: Array<NumberLike>): GoUint8
}
function NewUint8(val: number): GoUint8
function NewUint8(val: string, base: number | string): GoUint8

class GoFloat64 extends Native {
    private readonly __GoFloat64: GoFloat64
    private constructor()
    String(): string
    ToNumber(): number
    ToInt(): GoInt
    ToInt64(): GoInt64
    ToInt32(): GoInt32
    ToInt16(): GoInt16
    ToInt8(): GoInt8
    ToUint(): GoUint
    ToUint64(): GoUint64
    ToUint32(): GoUint32
    ToUint16(): GoUint16
    ToUint8(): GoUint8
    ToFloat64(): GoFloat64
    ToFloat32(): GoFloat32
    ABS(): GoFloat64
    Negate(): GoFloat64
    Add(...value: Array<NumberLike>): GoFloat64
    Sub(...value: Array<NumberLike>): GoFloat64
    Mul(...value: Array<NumberLike>): GoFloat64
    Div(...value: Array<NumberLike>): GoFloat64
    Sqrt(): GoFloat64
    Compare(val: NumberLike): number
    Max(...value: Array<NumberLike>): GoFloat64
    Min(...value: Array<NumberLike>): GoFloat64
}
function NewFloat64(val: number): GoFloat64
function NewFloat64(val: string): GoFloat64

class GoFloat32 extends Native {
    private readonly __GoFloat32: GoFloat32
    private constructor()
    String(): string
    ToNumber(): number
    ToInt(): GoInt
    ToInt64(): GoInt64
    ToInt32(): GoInt32
    ToInt16(): GoInt16
    ToInt8(): GoInt8
    ToUint(): GoUint
    ToUint64(): GoUint64
    ToUint32(): GoUint32
    ToUint16(): GoUint16
    ToUint8(): GoUint8
    ToFloat64(): GoFloat64
    ToFloat32(): GoFloat32
    ABS(): GoFloat32
    Negate(): GoFloat32
    Add(...value: Array<NumberLike>): GoFloat32
    Sub(...value: Array<NumberLike>): GoFloat32
    Mul(...value: Array<NumberLike>): GoFloat32
    Div(...value: Array<NumberLike>): GoFloat32
    Sqrt(): GoFloat32
    Compare(val: NumberLike): number
    Max(...value: Array<NumberLike>): GoFloat32
    Min(...value: Array<NumberLike>): GoFloat32
}
function NewFloat32(val: number): GoFloat32
function NewFloat32(val: string): GoFloat32
