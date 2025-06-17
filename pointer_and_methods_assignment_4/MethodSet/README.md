# 📌 Go Pointers vs Values – Important Notes

This document explains how **value types** and **pointer types** interact with **methods** and **interfaces** in Go.

---

## ✅ Method Receivers

In Go, methods can have:
- **Value receivers**: `func (r Type) Method()`
- **Pointer receivers**: `func (r *Type) Method()`

---

## 🧠 Access Rules Summary

| Method Receiver      | Variable Type | Call Allowed? | Reason                                                |
|----------------------|---------------|---------------|--------------------------------------------------------|
| `func (t Type)`      | `Type`        | ✅ Yes         | Matches receiver exactly                               |
| `func (t Type)`      | `*Type`       | ✅ Yes         | Go auto-dereferences pointer to call value receiver    |
| `func (t *Type)`     | `*Type`       | ✅ Yes         | Matches receiver exactly                               |
| `func (t *Type)`     | `Type`        | ❌ No          | Go does **not** auto-convert value to pointer          |

---

## 🧩 Interface Satisfaction

If an interface requires method `M()`:

- If method `M()` is defined on a **value receiver**, then **both `Type` and `*Type` satisfy** the interface.
- If method `M()` is defined on a **pointer receiver**, then **only `*Type` satisfies** the interface.

---

## 🔧 Examples

### 1. Method on Value Receiver

```go
type File struct{}

func (f File) Log() {
	fmt.Println("Log from value receiver")
}

f1 := File{}
f2 := &File{}

f1.Log() // ✅ OK
f2.Log() // ✅ OK
