### 1. What does MVC stands for?

MVC stands for Model-View-Controller.

### 2. What is each layer of MVC responsible for?

Each layer of MVC has its own responsibilities.
* Model: This layer is responsible for handling data and business logic. It interacts with the database or data source and carries out necessary data operations.
* View: This layer is responsible for the presentation logic. It displays the data provided by the Model in a format suitable for interaction, typically a user interface.
* Controller: This layer is responsible for handling user input. It interacts with both the Model and the View. It receives user input, optionally validates it, and then performs the business operation that the input requests.

### 3. What are some benefits and disadvantages to uning MVC?
Benefits of using MVC:

* Separation of Concerns: MVC separates the application into three components: Model, View, and Controller. This makes it easier to manage complexity, as each component can be developed and tested independently.

* Reusability and Modularity: Components defined in MVC can often be reused across different parts of an application or even across different applications. This can lead to more modular and maintainable code.

* Parallel Development: Because of the separation of responsibilities, one developer can work on the View while another works on the Controller logic or the Model, speeding up the development process.

* Efficient Code Organization: MVC forces a certain level of organization of your code, making it easier to locate and fix bugs or add new features.

Disadvantages of using MVC:

* Overhead: MVC can be overkill for simple, small-scale applications where a simple procedural approach might be more efficient.

* Learning Curve: For beginners, understanding and correctly utilizing the MVC pattern can be challenging.

* Complexity: While MVC can help manage complexity, it introduces its own complexity. You have to handle the data flow between Model, View, and Controller, and manage how they interconnect.

* Tight Coupling: While ideally the Model, View, and Controller should be independent, in some cases changes in the Model can require changes in the View and Controller, leading to a higher degree of coupling than desired.


## Read about othe ways to structure code
[gobeyond.dev](https://gobeyond.dev)
[Kat Zien at Gophercon](https://youtube.com/watch?v=oL6JBUk6tj0)