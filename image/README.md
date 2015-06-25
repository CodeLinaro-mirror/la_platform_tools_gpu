# image
--
    import "android.googlesource.com/platform/tools/gpu/image"

Package image provides functions for converting between various image formats.

## Usage

#### func  Convert

```go
func Convert(data []byte, width int, height int, srcFmt Format, dstFmt Format) ([]byte, error)
```
Convert uses the registered Converters to convert the image formed from the
parameters data, width and height from srcFmt to dstFmt. If the conversion
succeeds then the converted image data is returned, otherwise an error is
returned. If no direct converter has been registered to convert from srcFmt to
dstFmt, then Convert may try converting via an intermediate format.

#### func  RegisterConverter

```go
func RegisterConverter(src, dst Format, c Converter)
```
RegisterConverter registers the Converter for converting from src to dst
formats. If a converter already exists for converting from src to dst, then this
function panics.

#### type Converter

```go
type Converter func(data []byte, width int, height int) ([]byte, error)
```

Converter is used to convert the the image formed from the parameters data,
width and height into another format. If the conversion succeeds then the
converted image data is returned, otherwise an error is returned.

#### type Format

```go
type Format interface {
	binary.Object

	Check(data []byte, width, height int) error
}
```

Format is the interface for an image and/or pixel format.

Check returns an error if the combination of data, image width and image height
is invalid for the given format, otherwise Check returns nil.

#### func  ATC_RGBA_EXPLICIT_ALPHA_AMD

```go
func ATC_RGBA_EXPLICIT_ALPHA_AMD() Format
```
ATC_RGBA_EXPLICIT_ALPHA_AMD returns a format representing the texture
compression format with the same name.

#### func  ATC_RGB_AMD

```go
func ATC_RGB_AMD() Format
```
ATC_RGB_AMD returns a format representing the texture compression format with
the same name.

#### func  Alpha

```go
func Alpha() Format
```
Alpha returns a format containing a single 8-bit alpha channel per pixel.

#### func  ETC1_RGB8_OES

```go
func ETC1_RGB8_OES() Format
```
ETC1_RGB8_OES returns a format representing the texture compression format with
the same name.

#### func  Luminance

```go
func Luminance() Format
```
Luminance returns a format containing a single 8-bit luminance channel per
pixel.

#### func  LuminanceAlpha

```go
func LuminanceAlpha() Format
```
LuminanceAlpha returns a format containing an 8-bit luminance and alpha channel
per pixel.

#### func  PNG

```go
func PNG() Format
```
PNG returns a format representing the the texture compression format with the
same name.

#### func  RGB

```go
func RGB() Format
```
RGB returns a format containing an 8-bit red, green and blue channel per pixel.

#### func  RGBA

```go
func RGBA() Format
```
RGBA returns a format containing an 8-bit red, green, blue and alpha channel per
pixel.
