// <?php
// class SomeObject {
//     public $classVar;
//     public function __construct( $classVar ) {
//         $this->classVar = $classVar;
//     }
// }

    
// $object    = new SomeObject( "Hello, world." );
// $reference = $object;
// $reference->classVar = "Look! I can access object!";
// echo $object->classVar;    
// echo $reference->classVar; 

package main

import ("fmt")

type SomeStruct struct{
    Field string
}

// foo nằm trong tập các method của SomeStruct
// (s *SomeStruct) là một receiver mà con trỏ SomeStruct trỏ đến

func (s *SomeStruct) foo(field string) {    
    s.Field = field
}

func main() {
    someStruct := new(SomeStruct)
    
    someStruct.foo("a")
    fmt.Println(someStruct.Field)  // "a"
    someStruct.foo("b")
    fmt.Println(someStruct.Field)  // "b"
}

