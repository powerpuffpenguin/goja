try {
    var b = make()
    print(b)
    printType(b)
    fmt(b)
    b = b.slice(0, 2)
    print(b)
    printType(b)
    fmt(b)
    // b = makeb()
    // print(b)
    // printType(b)

} catch (e) {
    print(e)
    print(e instanceof GoError)
}