# templ-playground

A simple project using [TEMPL](https://templ.guide) as template engine.

You can use [Air](https://github.com/air-verse/air) for live reloading with this project.

## Air Installation

With go 1.22 or higher:

```
go install github.com/air-verse/air@latest
```

The included **.air.toml** file reloads changes to templ templates as well, so install templ if required.

```
go install github.com/a-h/templ/cmd/templ@latest
```

## Running the project
With above installed, at the root of project run:

```
air
```

This will start the web app and auto reload any changes as they are saved.