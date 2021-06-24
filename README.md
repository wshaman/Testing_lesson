Testing lesson codebase
---
This repo contains code to display some test features in Go.  

Motivation:
---
We are the ACME corp. It's like SHIN-RA, but ACME.  
We want to add our users and create accounts with emails by rules:
* All letters lowercased.
* All that is not letter is replaced with dot `.`
* Multiply dots are shrinked to one (`...` => `.`)
* Same names are allowed, but a 2-degits number is added after name and dot
    * `User` => `user@acme.com`
    * `User` => `user.01@amce.com`

Homework:
--
* Add test to `waitfor.It` func
* Benchmark `app/utils/naming/increment_number.go`. 
  * Run benchmarks
  * With pprof map find slowest part
  * Update slowest function
  * Run benchmark again
  * Result should be: 
    * modifed code
    * pprof files(original and modified)
* Add integration test:
  * make sure each test runs separately
    * this DOESN'T mean to fix the failing code.