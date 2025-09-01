# 🖼️ IFD Image Format — Lossless BMP Compression in Go

Welcome to the **IFD Image Format** project — a custom, lossless image compression format built in Go, designed to convert `.bmp` images into `.ifd` files and render them using the [Ebiten](https://ebiten.org/) game engine.

---

## 🚀 Overview

This project introduces a new image format: `.ifd` (Image Format Dhruv). It compresses raw `.bmp` images into a compact, lossless binary format and provides decoding support for real-time rendering via Ebiten.

- 🔒 **Lossless Compression**: No pixel data is lost — perfect fidelity preserved.
- ⚡ **Fast Encoding/Decoding**: Optimized for performance in Go.
- 🎮 **Ebiten Integration**: Decode and display `.ifd` images seamlessly in games or apps.

---

## 📦 Features

- Convert `.bmp` images to `.ifd` format
- Decode `.ifd` files back to raw image data
- Render decoded images using Ebiten
- CLI tool for batch conversion
- Modular architecture for easy extension

---

## 🛠️ Installation

```bash
git clone https://github.com/JNGCoding/ifd-image-format.git
cd ifd-image-format
go build
