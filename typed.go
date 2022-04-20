package typed

import "fmt"
// function TypeOf return a string with the type of the variable provided.
// take's 1 arg.
func TypeOf(s any)string{
     return fmt.Sprintf("%T", s)
}
// function IsSameType return's a boolean of wheter or not 2 variables are of the same type
func IsSameType(first any,second any)bool{
     if(TypeOf(first) == TypeOf(second)){
       return true
     }
     return false

}
