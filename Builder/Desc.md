# Builder Design Pattern

Builder design pattern attern is a creational design pattern from GOF. This pattern is used to build a complex object by using a step by step approach.

## Why?

When the object construction of a particular type is complex and requires a lot of parameterization
Say for example a car has some necessary features like Engine, Frame and Seats and some optional features like rooftop 

In these cases we have couple of these options

+ Either create a common constructor which intakes all required and optional values 
+ Or create a costructor for each combination 
+ Or use setters to exclusively set the required properties

### Downside

However the downside to the first two approach is that it makes code cumbersome and difficult to understand while the third leaves it open to modification by client


## Pattern components

### Abstract Builder

Provides contract for all the builders 

### Concrete Builders

Concrete implementations of abstract builders which ultimately holds responsibility for creating product

### Product

The product created by one of the concrete implementations

### Director

Acts as an interface to client and internally calls the passed builder methods to create a product.
This can be an optional component (deviation from GOF) where the client can also choose to call the builders directly with required methods.


## Result

+ With this there's proper separation of concern with regards to creating objects and also make the code cleaner
+ It doesnt leave the Product prone to changes by setters called from client
+ The final build step can hold the responsibility of any validation logic that might be required as part of object creation


## How to run?

To run this project navigate to the current Builder folder and run the module

```cmd
    cd .\Builder
    go run builder
```