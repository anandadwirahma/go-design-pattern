# Facade Pattern

**Facade Pattern** merupakan sebuah structural design pattern yang menyediakan sebuah interface yang disederhanakan ke sebuah library, framework, atau kumpulan class yang complex.

![alt text](https://refactoring.guru/images/patterns/content/facade/facade-2x.png)

_**Problem**_

Bayangkan kamu harus membuat kodemu bekerja dengan kumpulan object yang besar yang memiliki library atau framework yang canggih.  
Biasanya, kamu membutuhkan inisialisasi semua objeknya, melacak dependencynya, mengeksekusi methid dengan urutan yang benar, dsb.

Hasilnya, bisnis logic dari classmu akan menjadi erat digabungkan ke detail implementasi dari class 3rd-party, membuatnya sulit untuk dipahami dan dipelihara.

_**Solution**_

Facade adalah class yang menyediakan sebuah interface sederhana ke subsistem yang komplek yang berisi banyak bagian yang bergerak. Facade mungkin menyediakan fungsionalitas yang terbatas dibandingkan dengan menggunakan subsystem secara langsung. Namun, hanya mencakup fitur2 yang benar benar diperhatikan oleh client.

Facade berguna ketika kamu butuh mengintegrasikan aplikasimu dengan library yang canggih yang memiliki banyak fitur, tapi kamu hanya membutuhkan sebagian kecil fungsionalitasnya.

Contohnya, sebuah aplikasi yang mengupload video pendek ke sosial media berpotensi menggunakan konversi video profesional. Namun, semuanya benar2 membutuhkan sebuah class dengan single method `encode(filename, format)`. Setelah membuat sebuah class dan menghubungkannya dengan library video conversion, kamu akan memiliki facade pertama.