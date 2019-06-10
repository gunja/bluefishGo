package main

import (
    "fmt"
    "blowfish"
)

func main() {
    fmt.Println("Hello")
    // secret key 32 to 448 bits ( 4 - 56 octets )
    key := []int8 { 1, 2, -3, -4, 5, -6, -7 }
    //fmt.Println("len of key = ", len( key))
    var P, S []uint32
    P, S = blowfish.Blowfish_init( key, len(key))
    var a,b uint32 = 0x12345678, 0xFEDCBA12
    c, d := blowfish.Encipher( a, b, P, S)
    fmt.Printf("c = 0x%X  d=0x%X\n", c, d)
    //c= 0x10B39D0  d=0x122E8362
    var C, D uint32 = 0x10B39D0, 0x122E8362
    A, B := blowfish.Decipher( C, D, P, S)
    fmt.Printf("A = 0x%X  B=0x%X\n", A, B)
}

