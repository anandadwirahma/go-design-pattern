#Builder Pattern

<p align="justify">
<b>Builder pattern</b> merupakan creational design pattern yang memungkinkan kamu untuk membuat object yang complex secara setp by step. Builder pattern memungkinkan kamu untuk membuat tipe dan representasi object yang berbeda menggunakan konstruksi code yang sama. 
</p>

![alt text](https://refactoring.guru/images/patterns/content/builder/builder-en-2x.png)

<p align="justify">
<b>Analoginya</b> misalkan kita ingin membuat sebuah robot dengan bahan playdough berbentuk bola dan menghasilkan robot yang berbadan kotak, mempunyai kepala, tangan dan kaki, maka kita akan mambuatnya step by step.  

Misal:
- Step 1: membentuk bahan menjadi kotak untuk badannya
- Step 2: memasang tangan
- Step 3: menambahkan kepala dan kaki

</p>

_**Problem**_

<p align="justify">
Seandainya kamu mempunyai object yang kompleks dan berisi banyak field. Hal ini akan membuatmu merepotkan ketika inisialisasi constructor, kamu akan terjebak dalam deklarasi object bersarang dan bahkan banyak parameter yang kamu butuhkan.
</p>

![alt text](https://refactoring.guru/images/patterns/diagrams/builder/problem1-2x.png)

Contoh, bayangkan kita membuat sebuah object `House`. Untuk membuat rumah yang simple, kamu membutuhkan menyusun tembok, lantai, pintu, jendela, dan atap. Tapi bagaimana jika kamu ingin membangun rumah yang lebih besar dengan halaman belakang dan tambahan fasilitas seperti sistem pemanas, pipa air, dan kabel listrik ?

Solusi sederhana yaitu meng extend `House` class dan membuat satu set subclass untuk mencakup semua kombinasi dari parameternya. Tapi pada akhirnya kamu akan membuat banyak subclass.

Ada pendekatan lain yang tidak memaksa kita membuat banyak subclass. Kamu bisa membuat sebuah constructor yang tepat berdasarkan `House` class dengan semua parameter yang dibutuhkan untuk house object. Namun pendekatan ini memiliki kekurangan yaitu kamu dipaksa mendeklarasikan parameter yang kemungkinan cukup banyak dan belum tentu di butuhkan semuanya.

_**Solution**_

Builder pattern mensarankanmu dengan meng ekstrak object construction code dari class itu sendiri dan memindahkan ke object terpisah yang dinamakan `builders`.

![alt text](https://refactoring.guru/images/patterns/diagrams/builder/solution1-2x.png)

Builder pattern mengorganisasikan pembuatan object kedalam satu set step2 (`buildWalls`, `buildDoor`, dsb). Untuk membuat sebuah object, kamu dapat menjalankan rangkaian step dari builder object. Bagian terpenting dari hal ini kamu tidak perlu memanggil semua step. Kamu hanya memanggil step yang diperlukan untuk membuat konfigurasi object tertentu.

Beberapa step pembuatn mungkin membutuhkan implementasi yang berbeda ketika kamu membutuhkan mebangun beberapa variasi product. Contoh, rumah A dibangun dengan tembok kayu, tapi rumah B dibangin dengan tembok batu.

Pada case ini, kamu bisa membuat beberapa class builder berbeda yang megimplementasikan set building step yang sama, tetapi dengan cara yang berbeda. Kemudian kamu bisa menggunakan builders pada construction process untuk menyediakan jenis object yang berbeda.

![alt text](https://refactoring.guru/images/patterns/diagrams/builder/solution1-2x.png)

Sebagai contoh, sebuah builder harus membuat beberapa bangunan diantaranya: 
- sebuah rumah yang berbahan kayu dan kaca
- sebuah kastil yang berbahan batu dan besi
- sebuah istana yang menggunakan emas dan berlian

Makan dengan memanggil rangkain step yang sama, kamu bisa membuat ketiga hal diatas dengan membuat builder pada masing masing bangunan. Namun hal ini akan bekerja jika client code memanggil step building yang mampu berinteraksi dengan masing masing builder menggunakan sebuah interface umum.