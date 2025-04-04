# pigment 🧪🎨

**Pigment** is a fast, clean, and idiomatic Go library for color management across multiple color spaces. It supports RGB, HSL, HSV, and CMYK models — each as its own subpackage for modular clarity and speed.

> Built for performance, readability, and robots who love colors.

---

## Features

- 📦 Modular design: import only the color spaces you need
- 🔁 Convert between RGB, HSL, HSV, and CMYK
- 🎯 Precise color component control and clamping
- 🧪 Built for extensibility and testing

---

## Installation

```bash
go get github.com/monkeysfoot/pigment
```
## Example

```go
package main

import (
    "fmt"

    "github.com/monkeysfoot/pigment/rgb"
    "github.com/monkeysfoot/pigment/hsl"
)

func main() {
    r := rgb.New(255, 128, 64)
    h := hsl.FromRGB(r)

    fmt.Printf("RGB: %v → HSL: %.2f°, %.2f%%, %.2f%%\n", r, h.H, h.S*100, h.L*100)
}
```

Package Layout

    rgb – Red-Green-Blue color space

    hsl – Hue-Saturation-Lightness

    hsv – Hue-Saturation-Value

    cmyk – Cyan-Magenta-Yellow-Key (Black)

    color.intf.go – Common interface for all color types

    floatclamp.go – Utility for safe float math