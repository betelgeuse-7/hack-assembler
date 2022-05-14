## Hack machine language assembler

This is an assembler implemented in Go, for the assembly language used in Nand2Tetris course part 1.

You can read more <strong>about the language</strong> here: [Github | aalhour/Assembler.hack](https://github.com/aalhour/Assembler.hack/blob/master/README.md)

### How to run

You need ```go``` version 1.16+, to compile and run this program.

To install Go on Debian based distros (e.g. Ubuntu/Pardus):
```bash
sudo apt install golang-go
```

If you want to download Go binaries for MacOS, Linux, or Windows, visit [Official Go Website](https://go.dev/dl/)

```bash
git clone https://github.com/betelgeuse-7/hack-assembler
```
```bash
cd hack-assembler
```
If you want an executable:
```bash
go build -o assembler .
```
To run that executable:
```bash
./assembler <file.asm>
```

If you just want to run the program:
```bash 
go run . <file.asm>
```

```<file.asm>``` is an assembly file containing error-free Hack assembly language instructions.

The program will create a ```file.hack``` file containing 16-bit Hack CPU instructions, right next to where ```file.asm``` is located on the file system.