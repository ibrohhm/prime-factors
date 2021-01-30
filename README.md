# Prime Factors
Benchmark checking prime factors whether using goroutine or not

## Setup
1. running the project
```
make run
```

2. run all unit test
```
make test
```

## About API
The API is about to retrieve prime factor numbers with given:

url: `/prime-factors`

method: `POST`

payload:
- is_goroutine:
  - `true` -> will find prime factors using go-routine
  - `false` -> will find prime factors without using go-routine
- example_type:
  - `small` -> automatically using small number as params ( number < 1.000 )
  - `medium` -> automatically using medium number as params ( 1.000 < number < 1.000.000 )
  - `high` -> automatically using small big as params ( 1.000.000 < number < 10.000.000 )
  - `custom` -> using `numbers` params as input
- numbers:
  array of number, will find prime factors each given number

## Note
Checking Prime Number using [sieve method](https://en.wikipedia.org/wiki/Sieve_of_Eratosthenes#:~:text=In%20mathematics%2C%20the%20sieve%20of,the%20first%20prime%20number%2C%202.)
