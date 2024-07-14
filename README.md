# HERatio and BFV Benchmark

This package implements the cryptographic functions and benchmarking tests for the short-paper "HERatio: Homomorphic Encryption of Rationals using Laurent Polynomials" originaly submitted for the "3rd Annual FHE.org Conference on Fully Homomorphic Encryption".

To run the performance benchmarks, you'll need to install Golang (i.e., Go). Both actions are covered in the next subsection.

## Performance Benchmarks

To install the latest version of Go refer to the [Golang Download Page](https://go.dev/learn/) and follow the instructions related to your Operating System (e.g., MacOS, Windows, Linux).

Clone the Github repository into a folder:

```bash
git clone git@github.com:Algemetric/HERatio.git
```

> [!TIP]
> Open a terminal and navigate to a folder where you can organize github projects. Issue the previous command from there.

After installing Go and cloning the github project, move inside the benchmark folder by executing:

```bash
cd ./Implementation/Golang/benchmark
```

Now, execute the set of benchmark functions by issuing:

```bash
go test -bench=. -benchmem
```

The system will list the performance and memory use of functions related to key generation, encryption, decryption, encoding, decoding, and basic mathematical operations over ciphertexts.
