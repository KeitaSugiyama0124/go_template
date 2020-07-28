package logs

import (
	"fmt"
	"runtime"
)

func PutString( hierarchy int , text string ) {
    fmt.Println( "----------------------------" )
    fmt.Println( runtime.Caller( hierarchy ) )
    fmt.Println( text )
    fmt.Println( "----------------------------" )
}
func PutError( hierarchy int , err error ) {
    if err != nil {
        fmt.Println( "----------------------------" )
        fmt.Println( runtime.Caller(hierarchy) )
        fmt.Println( err.Error() )
        fmt.Println( "----------------------------" )
        panic(err.Error())
    }
}