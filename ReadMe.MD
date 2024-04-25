 # Forge
 

This project is inspired by the likes of t3 create-react-app, ng generate, etc but with a focus on go.

There are a few different options for Forge.
- The first is a generic blank project complete with the words of the ancestors of 'Hello, world'.
- Second is a blank rest api service that generates main.go with net/http, this might be extended in the future for other routing options such as gorilla/mux, echo, or chi, but I have found that the latest version of net/http is good enough
- Third is the bases of a cli project with cobra installed as well.
- There are some helper utilities as well, to create a models folder and file, a service folder and file that can be named with `forge service <YourService>`
<br><br>

# That's cool and all but why though?
I honestly just wanted to explore cobra more thoroughly, and this can be useful for me in the future as I create random Go projects. I know it's not hard to create new projects in go, but maybe this can evolve over time as become more familiar with cobra. I've found the best way to learn something is to put it in action.
<br><br>

# Available Commands
- blank: Generic empty go project.
- cli: Cobra focused CLI project
- git: just another way to init a git repo (maybe look in to chaining commands? or adding a flag to set git with -g)
- help: The inbuilt help command
- models: Create a models folder and file
- rest: Create a rest API with net/http (assuming go 1.22) with middleware, helpers, models, etc.
- service: Create a new service structure for apis
<br><br>

# Future stuff
Ideally I would love to add gRPC as a template option, but that seems a bit trickier to do at this time and will play around with it.

<br><br>

# Lastly
If you have any feedback or ways to improve this this I am happy to connect and discuss.