# iceye-home-work

Homework of ICEYE for the job application process

## How to run the app

if you have make tools installed, you can run with `make` - command. If not, try to run this: `docker build -t larvis . && docker run larvis:latest`.

## Decisions / Comments

### Project structure

The app is trying to follow this code structure:
https://github.com/golang-standards/project-layout

The project is also trying to respect domain driven design.

### Used tools
[Air](https://github.com/cosmtrek/air) - to improve development speed. 
It is used to reload files, and only configured locally. Therefore it's not inside the docker, but there are configuration files `.zshrc`and `air.toml`.

### Justifications
#### Why DDD?
 - It keeps code clean. It makes app scale easier. It's not perfect. It's a bit overkill for this small application, but the purpose is to demonstrate my abilities to use DDD.
 
 #### Why no database? 
  - I did not see any justification to add a DB. The configurations could have stored in DB. However they are needed to store in the repo in order to run the app. Therefore, for this version DB is not needed. It could be added on version 2.
  
 #### Testing
  - There are very few unit tests. I aimed to test atleast the most "business" critical code lines. Having much higher test coverage would increase time spent, and sadly I do not have that much time available.
  - In terms of testing, I would add tests for most methods inside inside `internal` - package, they contain practically all the business logic.
  - I would also consider adding e2e or integration test(s) to make sure everything works together. I think having atleast one is more important than having many unit tests.

### Domains

- [root] (main)
- game
- actor|dealer
- deck|hand
- card

Notes:
- Dealer is a special type of Actor
- Hand is a special type of Deck

Domain relationships:

```plantuml
@startuml

Game "1" -- "m" Actor
Actor "1" -- "1" Deck
Deck "m" -- "m" Card

@enduml
```

Having the relationships means, that only respective references are allowed done in the code:

- Game can refer Actor (or Dealer)
- Actor, Non-Dealer can refer Hand 
    - Dealer can refer Deck
- Deck/Hand can refer Card

