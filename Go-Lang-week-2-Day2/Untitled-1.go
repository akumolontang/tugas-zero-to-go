package main

import "fmt"

// POINTER
// func main (){
//   var numberA int = 4
//   fmt.Println(numberA)

//   var numberB *int = &numberA

//   fmt.Println(numberB)
//   fmt.Println(*numberB)



// STRUCT
// type Person struct {
//   name string
//   age int
//   height int
//   gender bool
// }
// func main () {
//   yusqie:= Person{
//     name: "Yusqie Mafaza",
//     age: 23,
//     height: 170,
//     gender: false,
//   }
//   fmt.Println(yusqie)


// NGEPRINT NAMA DOANG
// type Person struct {
//   name string
//   age int
//   height int
//   gender bool
// }
// func main () {
//   yusqie:= Person{
//     name: "Yusqie Mafaza",
//     age: 23,
//     height: 170,
//     gender: false,
//   }
//   fmt.Println(yusqie.name)


// STRUCT CARA LAIN
  // type Person struct {
  //   name string
  //   age int
  //   height int
  //   gender bool
  // }
  // func main () {
  // var kania = Person{"Kania Nur", 27, 160, true}
  // fmt.Println(kania)


// STRUCT CARA LAIN
// type Person struct {
  //   name string
  //   age int
  //   height int
  //   gender bool
  // }
  // func main () {
  // salis := new(Person)
  // salis.name = "Salis"
  // salis.age = 22
  // salis.height = 160
  // salis.gender =  true

  // fmt.Println(*salis)



// METHOD
// type Person struct {
//     name string
//     age int
//     height int
//     weight int
//     gender bool
// }

// func (P Person) PanggilNama() string {
//   return "Halo, nama saya" + P.name
// }

// func (P Person) CariTahunLahir() int {
//   return 2019 - P.age
// }

// func (P Person) UbahBerat (turun int) int {
//   return P.weight - turun
// }

// func main () {
//   yusqie:= Person{
//     name: "Yusqie Mafaza",
//     age: 23,
//     height: 170,
//     weight: 50,
//     gender: false,
//   }

// fmt.Println(yusqie)
// fmt.Println(yusqie.PanggilNama())
// fmt.Println("saya lahir di tahun ",yusqie.CariTahunLahir())
// fmt.Println("saya telah diet, BB saya menjadi ", yusqie.UbahBerat(10))



// CONTOH 2 METHOD
// type Person struct {
//   name string
//   age int
// }

// func (P Person) GetName() string {     //kalau di definisiin (string) harus di return
//   return P.name + " amazing!"
// }

// func (P *Person) IncreaseAge() {
//   P.age = P.age + 1
// }

// func main () {
//   b:= Person{"John", 50}
//   fmt.Printf("%v\n", b)
//   fmt.Println(b.GetName())

//   b.IncreaseAge()
//   fmt.Println(b.age)
// }



// INTERFACE
// type person struct {
//   name string
//   age int
// }


// func main () {
//   var secret interface{}
    
//   secret = "ethan hunt"
//   fmt.Println(secret)

//   secret = []string{"apple", "mango", "banana"}
//   fmt.Println(secret)

//   secret = 12.4
//   fmt.Println(secret)

//   secret = person{
//     name: "John",
//     age: 20,
//   }

//   fmt.Println(secret)
// }



// CONTOH INTERFACE 2
func main () {
  var numberA, numberB interface{}
  var numberB interface{}                 // kalau gamau digabungin bisa juga di define terpisah
  numberA = 10
  numberB = 20
  fmt.Println(numberA.(int) + numberB.(int))      //ini namanya casting (tipe data interface tapi dibuat jadi int)
}