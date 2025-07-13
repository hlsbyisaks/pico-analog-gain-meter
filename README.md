# Pico Analog Gain Meter

Measure and display signal amplification using the Raspberry Pi Pico 2 (RP2350-based).

This project reads two analog voltage signals via ADC, calculates gain in decibels (dB), and outputs the results via serial. It's useful for testing and verifying analog amplifier performance.

---

## 🔧 Features

- Reads two analog inputs using built-in ADC
- Calculates gain (dB)
- Displays results over serial

---

## 📦 Hardware

[Raspberry Pi Pico 2 (RP2350)](https://www.raspberrypi.com/products/raspberry-pi-pico-2/)

---

## 📐 Wiring Example

| Pin       | Description                  |
|-----------|------------------------------|
| GP26 / ADC0 | Signal input (pre-amplification) |
| GP27 / ADC1 | Signal output (post-amplification) |
| GND       | Ground reference             |

---

## 🧪 Output Example

```plaintext
Vin (V): 0.24
Vout (V): 1.88
Gain (dB): 17.84
```

---

## 🚀 Building & Flashing

This project is written in TinyGo.

### Prerequisites

- [TinyGo installed](https://tinygo.org/getting-started/)

### Flash to Pico

```bash
tinygo flash -target=pico2 main.go
```

### Read serial
```bash
tinygo monitor
```