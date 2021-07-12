#Prototype (Clone)

**Prototype** merupakan creational design pattern yang memungkinkan kamu mengcopy object exisiting tanpa membuat dependency code di class tersebut.

_**Problem**_

Jika kamu mempunyai sebuah object dan kamu ingin membuat exact copynya. Bagiamana caranya ?  
Pertama kamu membuat object baru yang sama dengan class tersebut. Kemudian ambil semua field pada object aslinya dan copy valuenya ke dalam object yang baru.

Tapi masalahnya tidak semua object dapat di copy karena beberapa field object tipenya private atau tidak bisa di baca di luar object itu sendiri.

![alt text](https://refactoring.guru/images/patterns/content/prototype/prototype-comic-1-en-2x.png)

Karena kamu harus tahu class objectnya untuk membuat duplikasinya, codemu jadi depend terhadap classnya. Contoh masalah lainnya, terkadang kamu hanya tahu interface yang diikuti object, bukan concrete classnya. Contoh: sebuah parameter di sebuah method menerima object apapun yang mengikuti interfacenya.

_**Solution**_

Prototype pattern mendelegasikan proses cloning ke actual object yang di cloning. Prototype pattern mendeklarasikan common interface untuk semua object yang support cloning. Interface ini memungkinkan kamu mengcloning sebuah object tanpa menduplikasi kodemu ke class object.  
Biasanya interfacenya berisi satu method `clone`.

Impementasi method `clone` sangat mirip di semua class. Method ini membuat sebuah object dari current class dan membawa semua value field dari object lama ke object baru.

Kuncinya: dia akan punya sebuah function yang bisa di delegate untuk melakukan proses clone