# Proxy Pattern

**Proxy Pattern** merupakan design pattern yang memungkinkan kamu memberikan pengganti atau placeholder untuk objek lain. Sebuah proxy yang mengcontrol access ke original objek, memungkinkan kamu untuk melakukan sesuatu antara sebelum atau setelah request sampai ke original objek.

_**Problem**_

Mengapa kamu ingin mengcontrol access ke suatu objek ? 
Contohnya:  
Kamu memiliki sebuah objek besar yang menghabiskan banyak sekali system resources. Kamu memerlukannya dari waktu ke waktu, tapi tidak selalu.

![alt text](https://refactoring.guru/images/patterns/diagrams/proxy/problem-en-2x.png)

Kamu dapat mengimplementasikan lazy initialization: contohnya membuat object ini hanya ketika benar benar dibutuhkan. Semua client object akan perlu mengeksekusi beberapa inisialisasi code yang berbeda. Sayangnya, ini mungkin akan menyebabkan banyak duplikasi code.

Idealnya, kita ingin meletakkan code ini langsung di dalam class object kita, tapi hal ini tidak selalu bisa. Contohnya, classnya mungkin bagian dari sebuah 3rd-party library yang tertutup.


_**Solution**_

Proxy pattern memyarankanmu membuat sebuah class proxy baru dengan interface yang sama seperti original service objek. Kemudian update applikasimu sehingga proxy class tsb mengirimkan proxy object ke semua original object client. Setelah menerima request dari sebuah client, proxy membuat sebua service object yang asli dan mendelgasikan semua pekerjaanya.

![alt text](https://refactoring.guru/images/patterns/diagrams/proxy/solution-en-2x.png)

Proxy menyamar seperti objek database itu sendiri. Dia dapat menghandle lazy inisialisasi dan menghasilkan caching tanpa client atau database aslinya mengetahuinya.

Apa benefitnya ? Jika kamu perlu mengeksekusi sesuatu antara sebelum atau sesudah primary logic dari class, proxy memungkinkan kamu melakukan tanpa merubah classnya. Karena proxy mengimplement interface yang sama seperti original classnya, dia dapat diteruskan ke client manapun yang memngharapkan/membutuhkan objek service aslinya.

