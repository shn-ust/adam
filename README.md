# ADAM ![](https://github.com/shn-ust/adam/actions/workflows/go.yml/badge.svg)

Adam is used to find the dependencies of an application in a distributed environment. 

It observes the network traffic in order to find the dependencies. 



## ADAM Architecture
![adam-architecture (1)](https://github.com/shn-ust/adam/assets/142196840/3e6d58cb-9ae3-47fe-91a2-cf11ce685d41)

ADAM needs to be run on the machine where the dependencies are to be mapped.
After finding the dependencies, ADAM sends the data back to the collector using [zeromq](https://zeromq.org/). The collector stores the data (dependencies) to a Redis database.
Multiple "ADAM" instances can be run at a time. 


## Getting started

### Building the image

```
$ docker build -t adam .
```

### Running the image
```
$ docker run --network host -it adam
```

## Papers

- NATARAJAN, ARUN. NSDMiner: Automated Discovery of Network Service Dependencies. (Under the direction of Dr. Peng Ning.) [[paper](https://repository.lib.ncsu.edu/server/api/core/bitstreams/77a238e6-01b9-4e56-8861-0e863393854c/content)]
