# generator-go-module

Autogenerate the boilerplate for a viam go module using the
[Yeoman](https://github.com/yeoman/yo) scaffolding library.


# Usage

Make sure you have npm installed on your machine.

Install Yo:
``` bash
npm install -g yo
```

Install the generator:
``` bash
npm install -g generator-viam-go-module
```

To create your module, run:

``` bash
yo viam-go-module
```
and follow the prompts to create a module.

Once the module is created, run
``` bash
go mod tidy
```
Then, add defintions for the function stubs prior to building the module.

Note that Board, Input, and Camera modules are currently unsupported.
