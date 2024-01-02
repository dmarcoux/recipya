---
title: Dependencies
weight: 2
---

The following software is required to build the project.

| Software                                              | Version  |
|-------------------------------------------------------|----------|
| [Go](https://go.dev/dl)                               | 1.21+    |
| [Node.js](https://nodejs.org/en/download)             | 20.10.0+ |
| [Make](https://en.wikipedia.org/wiki/Make_(software)) | any      |
 | [Hugo](https://gohugo.io/installation/)               | latest   |

## Windows

The `make` program is required to build the project but Windows users must install it. 
To verify whether it's installed on your machine, execute `make` in a command prompt or PowerShell.

Follow these steps if not installed:

1. Open either the Command Prompt or Powershell in administrator mode.
2. Execute `winget install GnuWin32.Make`
3. Add `C:\Program Files (x86)\GnuWin32\bin` to the Windows PATH environment variable.


## Recommended CLI Programs 

The following lists CLI programs you should install to help you develop the project.

- The [Goose](https://github.com/pressly/goose?tab=readme-ov-file#install) database migration tool