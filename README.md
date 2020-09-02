# UniTable

## Install

```
go get -u github.com/unidoc/unipdf

go get -u github.com/JohannesKaufmann/unitable
```

## Usage

This library is only a convenient helper library for [unipdf](https://github.com/unidoc/unipdf). All the heavy work is done by the amazing unipdf library.

---

The approach is similar to [styled-components](https://styled-components.com/) (for css). You start by creating a new base using `NewCellBase` and pass the style object.

```go
header := unitable.NewCellBase(defaultStyle, headerStyle)
```

Then in your 2d-slice, you can call this function. The variable name should reflect how it is supposed to look _(header, item, price, ...)_.

```go
// and now assemble the data with our new helper functions
data := [][]unitable.Renderer{
    {
        header("Name"),
        header("Count"),

        // override the header style later on
        header(
            "Price",
            unitable.Style{
                BackgroundColor: lightBlue,
            },
        ),
    },
    ...
}
```

The different styles will be merged. The later styles will override the previous styles.

Finally pass the data 2d-slice into the `NewTable` function and then draw it.

```go
fonts := map[string]*model.PdfFont{
    "regular": fontRegular,
    "bold":    fontBold,
}

table := unitable.NewTable(creatr, fonts, data)

if err := creatr.Draw(table); err != nil {
    return err
}
```

For a complete example look at the [examples](/examples).

## Extend

Since the `RenderTable` function accepts a 2d slice of the type `Renderer`, you can implement that interface yourself.

## Contributions

This was made as a proof of concept. If you use it and you like it, shoot me a message. Any contributions are welcome.

I developed this with the old version and it works really well.

After I switched to `v3` I couldn't run the example anymore without having a license. It compiles but I don't know whether it still works 🤷‍♂️.
