#Abstract Factory

<p align="justify">
<b>Abstract Factory</b> merupakan creational design pattern yang memungkinkan kamu untuk membuat families object yang berelasi tanpa harus menentukan concrete classnya.
</p>

_**Problem**_

<p align="justify">
Seandainya kamu membuat simulator toko furniture. Kode kamu berisi class yang mewakili:

1. Families dari product yang berelasi. Contoh: `chair` + `sofa` + `coffeeTable`
2. Beberapa variasi dari familiynya. Contoh: product `chair` + `sofa` + `coffeeTable` memiliki beberapa variasi, antara lain: `Modern`, `Victorian`, `ArtDeco`

</p>

Kamu membutuhkan cara untuk membuat individual furniture objects sehingga hasilnya cocok dengan object lain dalam satu family yang sama. Misalkan kamu ingin membuat satu set `chair` + `sofa` + `coffeeTable` dengan model `ArtDeco`.
![alt text](https://refactoring.guru/images/patterns/diagrams/abstract-factory/problem-en-2x.png?id=7de667bc24583c3ac03c)

_**Solution**_

Hal pertama yang disarankan oleh Abstract Factory Pattern secara eksplisit mendeklarasikan interfaces untuk setiap product yang berbeda dari product familiy (e.g., chair, sofa or coffee table).  

Kemudian kamu bisa membuat semua variasi dari product mengikuti interfacenya.  
Contoh: 
- Semua variasi dari kursi bisa mengimplement interface `Chair`.
- Semua variasi coffee table bisa mengimplement interface `CoffeeTable`.
- dsb.
  
![alt text](https://refactoring.guru/images/patterns/diagrams/abstract-factory/solution1-2x.png?id=eacec286153e058db925)

Selanjutnya buat sebuah interface `Abstract Factory` yang berisi method untuk membuat semua product yang merupakan bagian dari product family. (Contoh: `createChair`, `createSofa` and `createCoffeeTable`). Method ini harus mengembalikan `Abstract Product` yang diwakili oleh interface yang kita ekstrak sebelumnya. Contoh: `Chair`, `Sofa`, `CoffeeTable`, dsb.

![alt text](https://refactoring.guru/images/patterns/diagrams/abstract-factory/solution2-2x.png?id=b21557081f75ac7b4110)

Setiap variasi dari produk family, dibuatkan factory class terpisah berdasarkan interface `AbstractFactory`. Sebuah factory adalah class yang mengembalikan produk dari jenis tertentu. Contoh: `IKEAFurnitureFactory` hanya dapat membuat `IKEAChair`, `IKEASofa` and `IKEACoffeeTable` objects.