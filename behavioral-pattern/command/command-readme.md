# Command Pattern

**Command Pattern** merupakan behavioral design pattern yang mengubah sebuah request menjadi sebuah stand-alone object yang berisi semua informasi requestnya. Transformasi ini memungkinkan kamu mengirimkan request sebagai sebuah argument method, delay or queue request execution, dan support operasi undoable.

_**Problem**_

Seandanya kamu mengerjakan sebuah aplikasi text-editor. Task kamu saat ini adalah membuat sebuah toolbar dengan banyak tombol untuk berbagai operasi editor. Kamu membuat class button sangat rapi yang dapat digunakan untuk banyak button pada toolbar, serta untuk generic buttin di berbagai dialog.

![alt text](https://refactoring.guru/images/patterns/diagrams/command/problem1-2x.png)

Meskipun semua button terlihat sama, mereka semua seharusnya melakukan hal yang berbeda. Dimana kamu akan meletakkan code untuk beragam click handler dari tombol ini ?  
Solusi sederhana dengan membuat banyak subclass untuk setiap tempat dimana tombol digunakan. Subclass ini akan berisi code yang harus dieksekusi pada click button.

![alt text](https://refactoring.guru/images/patterns/diagrams/command/problem2-2x.png)

Tak lama kemudian, kamu menyadari pendekatan ini sangat cacat. Pertama, kamu mempunyai subclass yang sangat banyak, dan itu tidak masalah jika kamu tidak mengambil resiko merusak/merubah code di subclass ini setiap kamu memodifikasi class base buttonnya. Sederhananya, code GUI menjadi depend pada volatile code dari bisnis logic.

![alt text](https://refactoring.guru/images/patterns/diagrams/command/problem3-en-2x.png)

Dan disini bagian terburuknya. Beberapa operasi, seperti mengcopy atau mempaste text, akan perlu memanggil dari beberapa tempat. Contohnya, user dapat click tombol "Copy" kecil pada toolbar, atau copy sesuatu via context menu, atau hanya hit `ctrl+c` pada keyboard.

Mulanya, ketika aplikasi kita hanya mempunyai toolbar, tidak masalah menempatkan implementasi dari berbagai operasi kedalam button subclcass. Dengan kata lain, mempunyai code untuk mengcopy text kedalam `CopyButton` subclass masih baik2 saja. Tetapi kemudian, ketika kamu mengimplement context menu, shortcuts, dan yang lainnya, kamu harus menduplikasi operasi codenya di banyak class atau membuat menu yang bergantung pada button, yang merupakan pilihan yang lebih buruk.

_**Solution**_

Design software yang baik sering didasarkan pada pirnsip dari pemisahan concern, yang biasanya menghasilkan pemecahan app menjadi beberapa layer. Contoh paling umum: sebuah layer graphic user interface dan layer lainnya untuk bisnis logic. Layer GUI bertanggung jawab untuk merender gambar yang bagus pada layar, menangkap beberapa input dan menampilkan hasil dari apa yang pengguna dan app lakukan. Namun, ketika melakukan sesuatu yang penting, seperti menghitung lintasa bulan atau menyusun laporan tahunan, GUI layer mendelegasikan tugasnya ke layer yang mendasari bisnis logci.

Di code mungkin terlihat sperti ini: object GUI memanggil sebuah method dari object bisnis logic dan mengirimkannya beberapa argument. Proses ini biasanya di digambarkan sebagai satu objek mengirimkan request lain.

![alt text](https://refactoring.guru/images/patterns/diagrams/command/solution1-en-2x.png)

Command pattern menyarankan GUI objek tidak boleh mengirimkan request secara langsung. Malahan, kamu harus mengextract semua detail request, seperti object yang dipanggil, nama dari method, dan list argumentnya menjadi sebuah `Command Class` yang terpisah dengan sebuah singe method yang mentrigger requestnya.

Command objek menyediakan tautan antara berbagai GUI dan object bisnis logic. Sekarang, objek GUI tidak perlu tahu bisnis logic apa yang akan diterima dan bagaimana memprosesnya. Objek GUI hanya mentrigger command, yang menangani semua detailnya.

![alt text](https://refactoring.guru/images/patterns/diagrams/command/solution2-en-2x.png)

Langkah selanjutnya adalah membuat command yang kamu miliki mengimplementasikan interface yang sama. Biasanya hanya sebuah single eksekusi method yang tanpa menggunakan parameter. Interface ini memungkinkan kamu menggunakan berbagai command dengan request pengirim yang sama, tanpa menduplikasi concrete class dari commandnya. Bonusnya, kamu dapat mengubah object command yang terhubung ke sender, secara efektif merubah behavior sender ketika runtime.

Kamu harus memperhatikan satu bagian puzzle yang hilang, yaitu request parameter. Object GUI mungkin telah menyediakan bisnis layer objek dengan beberapa parameter. Karena `command` execution method tidak mempunyain parameter, bagaimana kita mengirimkan detail request ke receiver ?  
Ternyata `command` harus sudah dikonfigurasi dulu sebelumnya dengan data ini, atau mampu mengambil sendiri.

![alt text](https://refactoring.guru/images/patterns/diagrams/command/solution3-en-2x.png)

Kembali ke text editor kita. Setelah kita menerapkan command pattern, kita tidak lagi perlu semua button subclass ini mengimplementasikan berbagai behavior click. Cukup meletakkan singgle field ke dalam base `Button` class yang menyimpan referensi ke sebuah command object dan membuat button mengeksekusi command ketika diclick.

Kamu akan mengimplementasikan banyak class command untuk setiap kemungkinan operasi dan menghubungkan mereka dengan button tertentu, tergantung pada behavior button yang diharapkan.

Element GUI lainnya, seperti menu, shortcuts atau seluruh dialog, dapat di implementasikan dengan cara yang sama. Mereka ditautkan ke sebuah command yang dieksekusi ketika user berinteraski dengan element GUI. Seperti yang kamu duga, element terkait ke operasi yang sama akan ditautkan ke command yang sama, mencegah adanya duplikasi code.

Hasilnya, command menjadi sebuah lapisan tengan yang nyaman yang mengurangi duplikasi antara GUI dan bisnis logic layer. Dan hanya 