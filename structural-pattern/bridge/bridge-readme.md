#Bridge Pattern

**Bridge Patterne** memungkinkan kamu memecah sebuah class yang besar kedalam 2 abstract dan implementasi yang terpisah yang dapat di develope secara independen satu sama lain. 

![alt text](https://refactoring.guru/images/patterns/diagrams/bridge/problem-en-2x.png)

Pada ilustrasi diatas, seandainya kita memiliki class `Shape` dengan subclass `Circle` & `Square`. Kemudian kamu ingin mengextend class agar mempunyai warna. Kamu berencana akan membuat subclass shape `Red` dan `Blue`. Namun karena kamu sudah memiliki 2 subclass, kamu jadi akan membuat 4 kombinasi class misalnya `BlueCircle` dan `RedSquare`.

Menambahkan tipe shape dan color baru pada hirarki akan membuat tumbuh secara exponensial. Sebagai contoh, ketika menambahkan sebuah triangle shape kamu akan membutuhkan 2 subclass (satu utk setiap warna). Kemudian kamu ingin menambahkan satu warna baru, maka akan butuh membuat 3 subclass (satu untuk setiap tipe shape). 

_**Solution**_

Masalah ini terjadi karena kita mencoba mengextend class shape dalam 2 dimensi independen. `Shape` dan `Color`. Hal ini merupakan issue yang umum pada class inheritance.

Bridge pattern mencoba memecahkan masalah ini dengan merubah cara dari inheritance ke object composit, sehingga original class akan mereferensikan ke object dari hirarki yang baru, daripada mempunyai semua state dan behaviors dalam satu class.

![alt text](https://refactoring.guru/images/patterns/diagrams/bridge/solution-en-2x.png)

Dengan pendekatan ini, kita dapat meng extract color code ke dalam class `Color` dengan 2 sublcass: `Red` & `Blue`. Kemudian class `Shape` mendapat referensi field yang menunjuk ke salah satu object `Color`. Sekarang shape dapat mendelegasikan pekerjaan color apapun ke object color yang terhubung. Referensi itu akan bertindak sebagai `bridge` antara `Shpae` dan `Color` class. Sekarang, menambahkan warna baru tidak perlu merubah hirarki shape, dan sebaliknya.