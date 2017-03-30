# Go Viper - SaveConfig feature

For now, the Viper SaveConfig function is not embedded. This snippet will take care of this.

## Author

In this snippet, parts of the original  code from Steve Francia has been used.

## Filetypes

It accepts `JSON` and `YAML` files.

## Example

```go
testFile := fmt.Sprintf("%s/%s", workDir, "demo.yaml")

v := viper.New()
v.SetConfigFile(testFile)
v.Set("version", "6.0.0")

if err := SaveConfig(*v); err != nil {
  t.Error("Error ", err)
}
```
