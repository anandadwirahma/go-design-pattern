# Composite Pattern

**Composite Pattern** memungkinkan kamu menyusun object seperti structure pohon dan kemudian bekerja dengan struktur ini seolah olah mereka adalah objek individu.

_**Problem**_

Menggunakan composite pattern masuk akal hanya ketika core model aplikasimu dapat direpresntasikan seperti pohon.

Contoh, misalkan kamu punya 2 object: `Products` dan `Boxes`. Sebuah `Box` dapat berisi beberapa `Products` serta beberapa kotak yang lebih kecil. Dan di box kecil ini dapat berisi beberapa `Products` atau bahkan box yang lebih kecil lagi, dst.

Kamu memutuskan untuk membuat sistem pemesanan yang menggunakan class ini. Pesanan dapat berisi simple product tanpa pembungkus, serta kotak yang diisi dengan barang dan kotak lain. Bagaimana kamu menghitung total harga dari pesanan tersebut ?

![alt text](https://refactoring.guru/images/patterns/diagrams/composite/problem-en-2x.png)

Kamu dapat mencoba pendekatan: membuka semua box, ambil semua product, dan hitung totalnya. Hal ini yang bisa dilakukan di kehidupan nyata. Tapi di program ini, tidak sederhana untuk menjalankan perulangannya. Kamu harus tahu class dari `Products` dan `Boxes`, nesting level dari boxes and detail nasty lainnya sebelumnya. Semuanya membuat direct approact yang terlalu canggun atau tidak bahkan mungkin.