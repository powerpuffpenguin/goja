try{
var b=make()

 b[0]=255
 b[0]++
 b[0]++
    print(b)
    b=makei64()
    
    print(b)
    console.log(b)
    print(b.Add(1))
    console.log(b.Add(1))
    // console.log(b.toString())
    // console.log('1',b.Add(1))
    // console.log('-1',b.Add(-1))
    
    // make("async")

    // var array=makeAll()
    // for (var i = 0; i < array.length; i++) {
    //     const element = array[i];
    //     console.log(i,typeof element,element)
    // }
    
}catch(e){
    console.log(e)
    console.log(e instanceof GoError)
}