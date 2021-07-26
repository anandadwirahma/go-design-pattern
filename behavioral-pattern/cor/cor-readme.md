# Chain of Responsibility Pattern

**Chain of Responsibility Pattern** memungkinkan kamu menereskan request di sepanjang rantai handlers. Setelah menerima request, setiap handler memutuskan untuk memproses requestnya atau meneruskan ke handler selanjutnya di dalam rantai. 

![alt text](https://refactoring.guru/images/patterns/content/chain-of-responsibility/chain-of-responsibility-2x.png)

_**Problem**_

Seandainya kamu mengerjakan sistem pemesanan online. Kamu ingin membatasi akses system sehingga hanya user yang ter authentifikasi yang dapat melakukan pemesanan. Juga, user yang mempunyai permission harus memiliki full akses ke semua pemesanan.

Setelah sedikit perencanaan, kamu menyadari pengecekan ini harus dilakukan secara sequential. Aplikasi dapat mencoba authentifikasi sebuah user ke sistem kapanpun dia menerima request yang berisi user credential. Namun, jika credential salah dan authentifikasi gagal, tidak ada alasan untuk melanjutkan pengecekan yang lain.

![alt text](https://refactoring.guru/images/patterns/diagrams/chain-of-responsibility/problem1-en-2x.png)

Beberapa bulan ke depan, kamu mengimplementasikan beberapa pengecekan sequential. Seperti:

- Salah satu dari kolegamu memberi saran bahwa tidak aman menyimpan data raw langsung ke sistem pemesanan. Sehingga kamu menambahkkan step validasi ekstra untuk membersihkan data di request.
- Kemudian, seseorang memperhatikan sistem ini rentan terhadap peretasan brute force password. Untuk meniadakan ini, kamu dengan cepat menambahkan pengecekan pengulangan failed request yang datang dari ip yang sama.
- Seseorang juga memberi saran kepadamu untuk dapat mempercepat sistem dengan mengembalikan cached results pada request berulang yang berisi data yang sama. Oleh karena itu, kamu menambahkan pengecekan lain yang memungkinkan request melewati sistem hanya jika tidak ada cached yang cocok.

![alt text](https://refactoring.guru/images/patterns/diagrams/chain-of-responsibility/problem2-en-2x.png)

Pengecekan code, yang sudah terlihat berantakan, menjadi semakin membengkak setiap kamu menambahkan fitur baru. Merubah satu pengecekan terkadang berpengaruh ke yang lain. Paling buruknya, ketika kamu mencoba me reuse pengecekannya untuk melindungi component yang lain dari sistemmu, kamu harus menduplikasi beberapa codemu karena comopnent ini membutuhkan beberapa pengecekan, tapi tidak semua dari mereka.

System ini menjadi sangat sulit dipahami dan mahal perawatannya.

_**Solution**_

Seperti beberapa behavioral design pattern lainnya, `COR` bergantung pada transformasi behavior tertentu ke dalam stand-alone object yang disebut handlers. Pada case kita, setiap pengecekan harus di extract ke classnya sendiri dengan sebuah single method yang melalkukan pengecekan tersebut. Request beserta datanya, di teruskan ke method ini sebagai sebuah argument.

Pattern ini memberimu saran untuk menghubungkan setiap handler kedalam sebuah rantai. Setiap handler yang ditautkan memiliki sebuah field untuk menyimpan reference ke handler selanjutnya di dalam sebuah rantai. Selain memproses request, handler meneruskan request lebih jauh ke sepanjang rantainya. Request berjalan di sepanjang rantai sampai semua handler mempunyai kesempatan untuk memprosessnya.

Disini bagian terbaiknya: sebuah handler dapat memutuskan untuk tidak meneruskan request lebih jauh ke bawah rantai dan secara efektif menghentikan pemrosesan selanjutnya.

Pada contoh kita sistem pemesanan, sebuah handler melakukan pemrosesan dan memutuskan apakah request diteruskan ke proses selanjutnya. Asumsikan requestnya berisi data yang benar, semua handler dapat menjalankan perilaku utama mereka, apakah pemeriksaan authentikasi atau caching.

