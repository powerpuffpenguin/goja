class Native {
    private constructor()
    private readonly __Native: Native
}
class GoNumber extends Native {
    private constructor()
    private readonly __GoNumber: GoNumber
}
type NumberLike = number | string | null | undefined | GoNumber

interface GoReadChannel<T> extends Native {
    private readonly __GoReadChannel: GoReadChannel<T>
}
interface GoWriteChannel<T> extends Native {
    private readonly __WriteChannel: GoWriteChannel<T>
}
interface GoChannel<T> extends Native implements GoReadChannel<T>, GoWriteChannel<T> {
    private readonly __GoReadChannel: GoReadChannel<T>
    private readonly __GoWriteChannel: GoWriteChannel<T>
    private readonly __GoChannel: GoChannel<T>
}
function goRecv<T>(ch: GoReadChannel<T>): [T, boolean]
function goTryRecv<T>(ch: GoReadChannel<T>): [T, boolean]
function goSend<T>(ch: GoWriteChannel<T>, x: any)
function goTrySend<T>(ch: GoWriteChannel<T>, x: any): boolean
function goClose<T>(ch: GoWriteChannel<T>)
function goRecv<T>(ch: GoReadChannel<T>, scheduler: Scheduler): Promise<[T, boolean]>
function goTryRecv<T>(ch: GoReadChannel<T>, scheduler: Scheduler): Promise<[T, boolean]>
function goSend<T>(ch: GoWriteChannel<T>, x: any, scheduler: Scheduler): Promise<undefined>
function goTrySend<T>(ch: GoWriteChannel<T>, x: any, scheduler: Scheduler): Promise<boolean>
function goClose<T>(ch: GoWriteChannel<T>, scheduler: Scheduler): Promise<undefined>
class SelectDir extends Native {
    private readonly __SelectDir: SelectDir
    private constructor() { }
}
const SelectSend: SelectDir
const SelectRecv: SelectDir
const SelectDefault: SelectDir
function NewSendCase(ch: GoWriteChannel, val: any): SelectCase
function NewRecvCase(ch: GoReadChannel): SelectCase
function NewSendCase(ch: GoWriteChannel, val: any, scheduler: Scheduler): Promise<SelectCase>
function NewRecvCase(ch: GoReadChannel, scheduler: Scheduler): Promise<SelectCase>
class SelectCase extends Native {
    private readonly __SelectCase: SelectCase
    private constructor() { }
}
const DefaultCase: SelectCase
function goSelect(...cases: Array<SelectCase>): [GoInt, any, boolean]
function goSelect(...cases: Array<SelectCase>, scheduler: Scheduler): Promise<[GoInt, any, boolean]>

class Scheduler extends Native {
    private readonly __Scheduler: Scheduler
    private constructor() { }
}
const DefaultScheduler: Scheduler
class GoRune extends Native {
    private readonly __GoRune: GoRune
    private constructor()
}
class GoErrorNative extends Native {
    private readonly __GoErrorNative: GoErrorNative
    private constructor()
}
class GoError extends Native {
    private readonly __GoError: GoError
    private constructor()
    name: string
    message: string
    stack?: string
    value: GoErrorNative
}
class Completer<T> {
    constructor()
    get promise(): Promise<T>
    resolve(value?: T | PromiseLike<T>)
    reject(reason?: any)
    toString(): string
}
class GoMap extends Native {
    private readonly __GoMap: GoMap
    private constructor()
}
class GoSlice extends Native {
    private readonly __GoSlice: GoSlice
    private constructor()
}
function goLen(v: GoMap | GoSlice): GoInt
function goHasKey(v: GoMap, key: any): boolean
function isGoSlice(v: GoSlice): boolean
function isGoMap(v: GoMap): boolean

const MaxInt64: GoInt64
const MaxInt32: GoInt32
const MaxInt16: GoInt16
const MaxInt8: GoInt8
const MaxUint64: GoUint64
const MaxUint32: GoUint32
const MaxUint16: GoUint16
const MaxUint8: GoUint8
const MaxFloat64: GoFloat64
const MaxFloat32: GoFloat32
const MinInt64: GoInt64
const MinInt32: GoInt32
const MinInt16: GoInt16
const MinInt8: GoInt8

class GoInt extends GoNumber {
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
function NewInt(val: NumberLike): GoInt
function NewInt(val: string, base: number | string): GoInt

class GoInt64 extends GoNumber {
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
function NewInt64(val: NumberLike): GoInt64
function NewInt64(val: string, base: number | string): GoInt64

class GoInt32 extends GoNumber {
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
function NewInt32(val: NumberLike): GoInt32
function NewInt32(val: string, base: number | string): GoInt32

class GoInt16 extends GoNumber {
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
function NewInt16(val: NumberLike): GoInt16
function NewInt16(val: string, base: number | string): GoInt16

class GoInt8 extends GoNumber {
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
function NewInt8(val: NumberLike): GoInt8
function NewInt8(val: string, base: number | string): GoInt8

class GoUint extends GoNumber {
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
function NewUint(val: NumberLike): GoUint
function NewUint(val: string, base: number | string): GoUint

class GoUint64 extends GoNumber {
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
function NewUint64(val: NumberLike): GoUint64
function NewUint64(val: string, base: number | string): GoUint64

class GoUint32 extends GoNumber {
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
function NewUint32(val: NumberLike): GoUint32
function NewUint32(val: string, base: number | string): GoUint32

class GoUint16 extends GoNumber {
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
function NewUint16(val: NumberLike): GoUint16
function NewUint16(val: string, base: number | string): GoUint16

class GoUint8 extends GoNumber {
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
function NewUint8(val: NumberLike): GoUint8
function NewUint8(val: string, base: number | string): GoUint8

class GoFloat64 extends GoNumber {
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
function NewFloat64(val: NumberLike): GoFloat64
function NewFloat64(val: string): GoFloat64

class GoFloat32 extends GoNumber {
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
function NewFloat32(val: NumberLike): GoFloat32
function NewFloat32(val: string): GoFloat32

class GoIntArray extends GoSlice {
    private readonly __GoIntArray: GoIntArray
    private constructor()
    String(): string
    Len(): GoInt
    Swap(i: NumberLike, j: NumberLike)
    Less(i: NumberLike, j: NumberLike): boolean
    Cap(): GoInt
    Copy(src: GoIntArray): GoInt
    Slice(start: NumberLike): GoIntArray
    SliceEnd(start: NumberLike, end: NumberLike): GoIntArray
    Append(...data: Array<NumberLike>): GoIntArray
    Get(index: NumberLike): GoInt
    Set(index: NumberLike, val: NumberLike)
    Join(sep: string): string
    Asc()
    Desc()
}
function NewIntArray(): GoIntArray
function NewIntArray(len: NumberLike): GoIntArray
function NewIntArray(len: NumberLike, cap: NumberLike): GoIntArray
class GoInt64Array extends GoSlice {
    private readonly __GoInt64Array: GoInt64Array
    private constructor()
    String(): string
    Len(): GoInt
    Swap(i: NumberLike, j: NumberLike)
    Less(i: NumberLike, j: NumberLike): boolean
    Cap(): GoInt
    Copy(src: GoInt64Array): GoInt
    Slice(start: NumberLike): GoInt64Array
    SliceEnd(start: NumberLike, end: NumberLike): GoInt64Array
    Append(...data: Array<NumberLike>): GoInt64Array
    Get(index: NumberLike): GoInt64
    Set(index: NumberLike, val: NumberLike)
    Join(sep: string): string
    Asc()
    Desc()
}
function NewInt64Array(): GoInt64Array
function NewInt64Array(len: NumberLike): GoInt64Array
function NewInt64Array(len: NumberLike, cap: NumberLike): GoInt64Array
class GoInt32Array extends GoSlice {
    private readonly __GoInt32Array: GoInt32Array
    private constructor()
    String(): string
    Len(): GoInt
    Swap(i: NumberLike, j: NumberLike)
    Less(i: NumberLike, j: NumberLike): boolean
    Cap(): GoInt
    Copy(src: GoInt32Array): GoInt
    Slice(start: NumberLike): GoInt32Array
    SliceEnd(start: NumberLike, end: NumberLike): GoInt32Array
    Append(...data: Array<NumberLike>): GoInt32Array
    Get(index: NumberLike): GoInt32
    Set(index: NumberLike, val: NumberLike)
    Join(sep: string): string
    Asc()
    Desc()
}
function NewInt32Array(): GoInt32Array
function NewInt32Array(len: NumberLike): GoInt32Array
function NewInt32Array(len: NumberLike, cap: NumberLike): GoInt32Array
class GoInt16Array extends GoSlice {
    private readonly __GoInt16Array: GoInt16Array
    private constructor()
    String(): string
    Len(): GoInt
    Swap(i: NumberLike, j: NumberLike)
    Less(i: NumberLike, j: NumberLike): boolean
    Cap(): GoInt
    Copy(src: GoInt16Array): GoInt
    Slice(start: NumberLike): GoInt16Array
    SliceEnd(start: NumberLike, end: NumberLike): GoInt16Array
    Append(...data: Array<NumberLike>): GoInt16Array
    Get(index: NumberLike): GoInt16
    Set(index: NumberLike, val: NumberLike)
    Join(sep: string): string
    Asc()
    Desc()
}
function NewInt16Array(): GoInt16Array
function NewInt16Array(len: NumberLike): GoInt16Array
function NewInt16Array(len: NumberLike, cap: NumberLike): GoInt16Array
class GoInt8Array extends GoSlice {
    private readonly __GoInt8Array: GoInt8Array
    private constructor()
    String(): string
    Len(): GoInt
    Swap(i: NumberLike, j: NumberLike)
    Less(i: NumberLike, j: NumberLike): boolean
    Cap(): GoInt
    Copy(src: GoInt8Array): GoInt
    Slice(start: NumberLike): GoInt8Array
    SliceEnd(start: NumberLike, end: NumberLike): GoInt8Array
    Append(...data: Array<NumberLike>): GoInt8Array
    Get(index: NumberLike): GoInt8
    Set(index: NumberLike, val: NumberLike)
    Join(sep: string): string
    Asc()
    Desc()
}
function NewInt8Array(): GoInt8Array
function NewInt8Array(len: NumberLike): GoInt8Array
function NewInt8Array(len: NumberLike, cap: NumberLike): GoInt8Array
class GoUintArray extends GoSlice {
    private readonly __GoUintArray: GoUintArray
    private constructor()
    String(): string
    Len(): GoInt
    Swap(i: NumberLike, j: NumberLike)
    Less(i: NumberLike, j: NumberLike): boolean
    Cap(): GoInt
    Copy(src: GoUintArray): GoInt
    Slice(start: NumberLike): GoUintArray
    SliceEnd(start: NumberLike, end: NumberLike): GoUintArray
    Append(...data: Array<NumberLike>): GoUintArray
    Get(index: NumberLike): GoUint
    Set(index: NumberLike, val: NumberLike)
    Join(sep: string): string
    Asc()
    Desc()
}
function NewUintArray(): GoUintArray
function NewUintArray(len: NumberLike): GoUintArray
function NewUintArray(len: NumberLike, cap: NumberLike): GoUintArray
class GoUint64Array extends GoSlice {
    private readonly __GoUint64Array: GoUint64Array
    private constructor()
    String(): string
    Len(): GoInt
    Swap(i: NumberLike, j: NumberLike)
    Less(i: NumberLike, j: NumberLike): boolean
    Cap(): GoInt
    Copy(src: GoUint64Array): GoInt
    Slice(start: NumberLike): GoUint64Array
    SliceEnd(start: NumberLike, end: NumberLike): GoUint64Array
    Append(...data: Array<NumberLike>): GoUint64Array
    Get(index: NumberLike): GoUint64
    Set(index: NumberLike, val: NumberLike)
    Join(sep: string): string
    Asc()
    Desc()
}
function NewUint64Array(): GoUint64Array
function NewUint64Array(len: NumberLike): GoUint64Array
function NewUint64Array(len: NumberLike, cap: NumberLike): GoUint64Array
class GoUint32Array extends GoSlice {
    private readonly __GoUint32Array: GoUint32Array
    private constructor()
    String(): string
    Len(): GoInt
    Swap(i: NumberLike, j: NumberLike)
    Less(i: NumberLike, j: NumberLike): boolean
    Cap(): GoInt
    Copy(src: GoUint32Array): GoInt
    Slice(start: NumberLike): GoUint32Array
    SliceEnd(start: NumberLike, end: NumberLike): GoUint32Array
    Append(...data: Array<NumberLike>): GoUint32Array
    Get(index: NumberLike): GoUint32
    Set(index: NumberLike, val: NumberLike)
    Join(sep: string): string
    Asc()
    Desc()
}
function NewUint32Array(): GoUint32Array
function NewUint32Array(len: NumberLike): GoUint32Array
function NewUint32Array(len: NumberLike, cap: NumberLike): GoUint32Array
class GoUint16Array extends GoSlice {
    private readonly __GoUint16Array: GoUint16Array
    private constructor()
    String(): string
    Len(): GoInt
    Swap(i: NumberLike, j: NumberLike)
    Less(i: NumberLike, j: NumberLike): boolean
    Cap(): GoInt
    Copy(src: GoUint16Array): GoInt
    Slice(start: NumberLike): GoUint16Array
    SliceEnd(start: NumberLike, end: NumberLike): GoUint16Array
    Append(...data: Array<NumberLike>): GoUint16Array
    Get(index: NumberLike): GoUint16
    Set(index: NumberLike, val: NumberLike)
    Join(sep: string): string
    Asc()
    Desc()
}
function NewUint16Array(): GoUint16Array
function NewUint16Array(len: NumberLike): GoUint16Array
function NewUint16Array(len: NumberLike, cap: NumberLike): GoUint16Array
class GoUint8Array extends GoSlice {
    private readonly __GoUint8Array: GoUint8Array
    private constructor()
    String(): string
    Len(): GoInt
    Swap(i: NumberLike, j: NumberLike)
    Less(i: NumberLike, j: NumberLike): boolean
    Cap(): GoInt
    Copy(src: GoUint8Array): GoInt
    Slice(start: NumberLike): GoUint8Array
    SliceEnd(start: NumberLike, end: NumberLike): GoUint8Array
    Append(...data: Array<NumberLike>): GoUint8Array
    Get(index: NumberLike): GoUint8
    Set(index: NumberLike, val: NumberLike)
    Join(sep: string): string
    Asc()
    Desc()
}
function NewUint8Array(): GoUint8Array
function NewUint8Array(len: NumberLike): GoUint8Array
function NewUint8Array(len: NumberLike, cap: NumberLike): GoUint8Array
class GoFloat64Array extends GoSlice {
    private readonly __GoFloat64Array: GoFloat64Array
    private constructor()
    String(): string
    Len(): GoInt
    Swap(i: NumberLike, j: NumberLike)
    Less(i: NumberLike, j: NumberLike): boolean
    Cap(): GoInt
    Copy(src: GoFloat64Array): GoInt
    Slice(start: NumberLike): GoFloat64Array
    SliceEnd(start: NumberLike, end: NumberLike): GoFloat64Array
    Append(...data: Array<NumberLike>): GoFloat64Array
    Get(index: NumberLike): GoFloat64
    Set(index: NumberLike, val: NumberLike)
    Join(sep: string): string
    Asc()
    Desc()
}
function NewFloat64Array(): GoFloat64Array
function NewFloat64Array(len: NumberLike): GoFloat64Array
function NewFloat64Array(len: NumberLike, cap: NumberLike): GoFloat64Array
class GoFloat32Array extends GoSlice {
    private readonly __GoFloat32Array: GoFloat32Array
    private constructor()
    String(): string
    Len(): GoInt
    Swap(i: NumberLike, j: NumberLike)
    Less(i: NumberLike, j: NumberLike): boolean
    Cap(): GoInt
    Copy(src: GoFloat32Array): GoInt
    Slice(start: NumberLike): GoFloat32Array
    SliceEnd(start: NumberLike, end: NumberLike): GoFloat32Array
    Append(...data: Array<NumberLike>): GoFloat32Array
    Get(index: NumberLike): GoFloat32
    Set(index: NumberLike, val: NumberLike)
    Join(sep: string): string
    Asc()
    Desc()
}
function NewFloat32Array(): GoFloat32Array
function NewFloat32Array(len: NumberLike): GoFloat32Array
function NewFloat32Array(len: NumberLike, cap: NumberLike): GoFloat32Array