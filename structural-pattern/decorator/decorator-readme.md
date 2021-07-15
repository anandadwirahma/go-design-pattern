# Decorator Pattern

**Decorator Pattern** memungkinkan kamu meletakkan behaviour baru ke suatu object dengan menempatkan object kedalam special wrapper object yang berisi behaviour object tsb.  
Decorator pattern bisa disebut juga dengan wrapper.

_**Problem**_

Bayangkan kamu membuat sebuah library notification yang memungkinkan program lain bisa mengirim notif ke usernya.

Versi awal dari library berdasarkan class `Notifier` yang hanya berisi beberapa kolom, sebuah constructor dan sebuah method `Send`. Method tersebut dapat menerima parameter message dari client dan mengirimkan message tersebut ke list email yang di dikirimkan ke notifier melalu constructor. Aplikasi third-party yang beraksi sebagai client seharusnya membuat dan mengkonfigurasi object notifier sekali, dan kemudian menggunakannya setiap kali terjadi suatu hal penting.

![alt text](https://refactoring.guru/images/patterns/diagrams/decorator/problem1-en-2x.png)

Dalam beberapa kasus, kamu menyadari pengguna library ingin mengirim notif lebih dari email. Beberapa diantaranya ingin menerima SMS tentang critial issues. Diantara lainnya ingin mengirin notif via Facebook, dan Slack.

![alt text](https://refactoring.guru/images/patterns/diagrams/decorator/problem2-2x.png)

Seberapa sulit kah hal ini ?  
Kamu mengextend class `Notifier` dan meletakkan method notification tambahan kedalam sublcass yang baru. Sekarang si client diharuskan meng inisiasi class notification yang diinginkan dan menggunakan untuk semua notification lebih lanjut. 

Tetapi kemudian ada pertanyaan, "Mengapa kamu tidak bisa menggunakan beberapa tipe notifikasi sekaligus ? Semisal terjadi kebakaran, mungkin kamu ingin menginformasikan melalui semua channel."

Kamu mencoba mengatasi masalah ini dengan membuat special subclass yang dikombinasikan dengan beberapa metode notifikasi dalam satu class. Namun, cara ini akan sangat menggembungkan code, tidak hanya code pada library tetapi juga pada client.

![alt text](https://refactoring.guru/images/patterns/diagrams/decorator/problem3-2x.png)

Kamu harus mencari cara lain untuk menyusun class notifikasi supaya codenya tidak membengkak.

_**Solution**_

Meng extend sebuah class merupakan hal pertama yang terpikirkan ketika kamu butuh merubah sebuah behavior suatu object. Namun inheritance memiliki beberapa catatan penting yang perlu kamu perhatikan.

- Inheritance adalah static. Kamu tidak dapat merubah behavior dari existing object saat runtime. Kamu hanya dapat mereplace keseluruhan object dengan object lain yang dibuat dari sublcass yang berbeda.
- Sublcass hanya dapat memiliki satu parent class. Di banyak bahasa pemrograman, inheritance tidak membiarkan sebuah class mewarisi behavior dari multiple class di waktu yang sama.


Salah satu cara untuk mengatasi masalah ini dengan menggunakan `Aggregation` atau `Composition` daripada `Inheritance`. Kedua alternatif ini bekerja hampir dengan cara yang sama: satu object memiliki sebuah reference ke object lain dan mendelegasikan beberapa pekerjaannya, sedangkan inheritance, object itu sendiri yang mampu mengerjakannya, menurunkan behavior dari superclassnya.

Dengan pendekatan yang baru ini kamu dapat dengan mudah mengganti helper object yang dihubungkan dengan yang lain, merubah behavior dari container saat runtime. Sebuah object dapat menggunakan behavior dari berbagai class, mempunyai referensi ke multiple object dan mendelegasikan semua jenis pekerjaannya. Aggregation/Composition adalah kunci utama dibalik beragam design pattern, termasuk Decorator.

![alt text](https://refactoring.guru/images/patterns/diagrams/decorator/solution1-en-2x.png)

"Wrapper" alternatif singkatan dari Decorator pattern yang dengan jelas menggambarkan ide utama dari pattern tersebut. Wrapper adalah sebuah object yang dapat dihubungkan dengan beberapa target object. Wrapper berisi kumpulan method yang sama dengan target dan di delegasikan ke semua request yang menerimanya. Namun, wrapper dapat merubah hasil dengan melakukan sesuatu antara sebelum atau sesudah mengirimkan request ke target.

Kapan sebuah wrapper menjadi decorator ?  
Wrapper mengimplementasikan interface yang sama seperti object yang di bungkusnya. Itu sebabnya dari perspektif client, objek ini identikal. Membuat field wrapper reference menerima objek apapun yang mengikuti interfacenya. Ini akan memungkinkan kamu menutupi object di beberapa wrapper, menambahkan kombinasi behavior dari semua wrapper tersebut.

Pada contoh notifikasi kita, mari kita tinggalkan behavior notifikasi email sederhana di dalam class `Notifier`, tapi ubah semua method notifikasi lainnya menjadi decorator.

![alt text](https://refactoring.guru/images/patterns/diagrams/decorator/solution2-2x.png)

Client code perlu membungkus sebuah object notifier basic menjadi satu set decorator yang sesuai preferensi client. Menghasilkan object yang akan disusun sebagai sebuah stack.

![alt text](https://refactoring.guru/images/patterns/diagrams/decorator/solution3-en-2x.png)

Decorator terakhir di stack akan menjadi object yang sesungguhnya akan dipakai client. Karena semua decorator mengimplement interface yang sama seperti base notifier, sisa kode dari client tidak peduli apakah hal itu bekerja dengan object notifier yang asli atau dengan decorator.

Kita bisa menerapkan pendekatan yang sama ke behavior lainnya seperti memformat pesan atau menyusun list penerima. Client dapat mendekorasi object dengan custom decorator, selama mengikuti interface yang sama seperti yang lainnya.