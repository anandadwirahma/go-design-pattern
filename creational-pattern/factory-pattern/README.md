## Factory Pattern ([Creational Design Pattern](https://refactoring.guru/design-patterns/factory-method)))
Factory pattern is a creational design pattern which solves the problem of creating product objects without specifying their concrete classes.

### Problem
Imagine that you’re creating a logistics management application. The first version of your app can only handle transportation by trucks, so the bulk of your code lives inside the Truck class.

After a while, your app becomes pretty popular. Each day you receive dozens of requests from sea transportation companies to incorporate sea logistics into the app.

At present, most of your code is coupled to the Truck class. Adding Ships into the app would require making changes to the entire codebase. Moreover, if you decide later to add other types of transport to your app, you may need to make all of these changes again.
if you decide later to add other types of transport to your app, you may need to make all of these changes again.
jika nanti Anda memutuskan untuk menambahkan jenis transportasi lain ke aplikasi Anda, Anda mungkin perlu melakukan semua perubahan ini lagi.
If you later decide to add other types of transport to the application, you may need to make all these changes again.
Jika nanti Anda memutuskan untuk menambahkan jenis transportasi lain ke aplikasi, Anda mungkin perlu membuat semua perubahan ini lagi.

As a result, you will end up with pretty nasty code, riddled with conditionals that switch the app’s behavior depending on the class of transportation objects.

<p align="center">
    <img src="https://refactoring.guru/images/patterns/content/factory-method/factory-method-en.png" width="300" height="300">
</p>

### Structure Pattern

<p align="center">
    <img src="https://refactoring.guru/images/patterns/diagrams/factory-method/structure.png" width="300" height="300">
</p>

### Conceptual Of Code
This code contains:

* iGun.go             : as a <i>Product</i>, are declared as interface. which defines all methods a gun should have.
* gun.go              : contains implement of iGun interface. In structure picture in above, the position as <i>ConcreateProduct</i>.
* ak47.go & musket.go : also as <i>ConcreateProduct</i>. Both embed gun struct, so indirectly implement all `iGun` methods.
* gunFactory.go       : serves as a factory. which creates a gun of the desired type based on an incoming arguments.
* main.go             : acts as a client