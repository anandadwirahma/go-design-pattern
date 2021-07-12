#Factory Method

**Factory Method** merupakan creational design pattern yang menyediakan sebuah interface untuk membuat object di superclass,
tetapi memperbolehkan subclassnya untuk merubah tipe dari object yang akan dibuat.

_**Problem**_

Seandainya kita membuat sebuah aplikasi management logistik. Pada versi pertama dari apps tersebut hanya dapat menghandle transportasi menggunakan truck.
Jadi semua logic codingan kita ada di dalam `Class Truck`.

Setelah perusahaan tersebut menjadi popular, Setiap hari menerima request dari transportasi laut untuk memasukan transportasi laut ke dalam apps.

Jadi, Bagaimana cara kita mengimplementasikan hal tersebut supaya lebih rapi ?

_**Solution**_

Factory method pattern menyarankan kamu untuk menggantikan pemanggilan constructor langsung (menggunakan method `new`) dengan memanggil sebuah _`factory method`_.  
Object akan tetap dibuat via operator `new`, tetapi dipanggil dari dalam `factory method`.  
Object yang di return oleh `factory method` sering disebut dengan _products_.

![alt text](https://refactoring.guru/images/patterns/diagrams/factory-method/solution1-2x.png)

simulasi pada bagian products, kurang lebih sebagai berikut:

![alt text](https://refactoring.guru/images/patterns/diagrams/factory-method/solution2-en-2x.png)

1. Kita akan buat sebuah interface Transport.
   Transport memiliki behaviour/kemampuan/fitur/method bernama `deliver()`.
   

2. Transport mempunyai implementasi bernama `Truck` & `Ship`, keduanya menurunkan method yang sama yaitu `deliver()`.
   Tetapi terdapat perbedaan implementasi method `deliver()` pada masing masing subclassnya.
   
   Misal:
    - method `deliver()` pada subclass `Truck` memiliki sistem navigasi darat yaitu google maps
    - method `deliver()` pada subclass `Ship` memiliki sistem navigasi laut yaitu kompas
    
Jadi, kedua subclass tersebut memiliki implementasi yang berbeda, tetapi memiliki signature yang sama. Signature disini yaitu method `deliver()` yang dimana diturunkan dari class yang sama yaitu `Transport`.